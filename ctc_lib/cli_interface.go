/*
Copyright 2018 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ctc_lib

import (
	"fmt"

	"github.com/GoogleCloudPlatform/runtimes-common/ctc_lib/logging"
	"github.com/GoogleCloudPlatform/runtimes-common/ctc_lib/util"
	"github.com/spf13/cobra"
)

type CLIInterface interface {
	printO(c *cobra.Command, args []string) error
	setRun(func(c *cobra.Command, args []string))
	getCommand() *cobra.Command
	ValidateCommand() error
	isRunODefined() bool
	Init()
}

func Execute(ctb CLIInterface) {
	defer errRecover()
	err := ExecuteE(ctb)
	CommandExit(err)
}

func ExecuteE(ctb CLIInterface) (err error) {
	if err := ctb.ValidateCommand(); err != nil {
		return err
	}
	ctb.Init()
	if ctb.isRunODefined() {
		cobraRun := func(c *cobra.Command, args []string) {
			err = ctb.printO(c, args)
			if err != nil {
				Log.Error(err)
			}
		}
		ctb.setRun(cobraRun)
	}

	err = ctb.getCommand().Execute()

	//Add empty line as template.Execute does not print an empty line
	ctb.getCommand().Println()
	if util.IsDebug(Log.Level) {
		logFile, ok := logging.GetCurrentFileName(Log)
		if ok {
			ctb.getCommand().Println("See logs at ", logFile)
		}
	}
	return err
}

// errRecover is the handler that turns panics into returns from the top
// level of Parse.
func errRecover() {
	if e := recover(); e != nil {
		err := fmt.Errorf("%v", e)
		CommandExit(err)
	}
}