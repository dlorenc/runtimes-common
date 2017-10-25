package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var bucket = flag.String("bucket", "", "GCS Bucket to use as cache.")

var RootCmd = &cobra.Command{
	Use:   "bazelcache",
	Short: "bazelcache is an HTTP REST cache on top of GCS for Bazel",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if *bucket == "" {
			fmt.Println("Must provide --bucket flag.")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cache = new(sync.Map)
		ctx := context.Background()
		var err error
		client, err = storage.NewClient(ctx)
		if err != nil {
			os.Exit(1)
		}
		bkt = client.Bucket(*bucket)

		http.HandleFunc("/", handler)
		fmt.Println("Starting bazel cache server on localhost:8080 using bucket ", *bucket)
		http.ListenAndServe(":8080", nil)
	},
}

var cache *sync.Map

var client *storage.Client
var bkt *storage.BucketHandle

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()
		StoreInCache(r.URL.String(), b)
	case http.MethodGet:
		b, err := GetFromCache(r.Context(), r.URL.String())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(b)
	default:
		fmt.Println("Unknown method: ", r.Method)
	}
}

func StoreInCache(key string, val []byte) {
	// First store in the memory cache.
	cache.Store(key, val)
	// In the background, store on GCS.
	go StoreInGCS(key, val)
}

func StoreInGCS(key string, val []byte) error {
	obj := bkt.Object(key)
	w := obj.NewWriter(context.Background())
	if _, err := w.Write(val); err != nil {
		return err
	}
	return w.Close()
}

func MaybeServeFromGCS(ctx context.Context, path string) ([]byte, error) {
	// Then check the disk cache
	obj := bkt.Object(path)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}

func GetFromCache(ctx context.Context, path string) []byte, err {
	// First check the memory cache.
	if b, ok := cache.Load(path); ok {
		buf, _ := b.([]byte)
		return buf, nil
	}
	// Then check the disk cache
	b, err := MaybeServeFromGCS(ctx, path)
	if err != nil {
		return nil, err
	}
	// Fill the memory cache then respond.
	StoreInCache(path, b)
	return b, nil
}
