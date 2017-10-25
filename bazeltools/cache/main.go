package main

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/runtimes-common/bazeltools/cache/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
