// Copyright © 2015 Seth Goings
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
  "log"
  "strings"
	"os/exec"

	"github.com/deis/rigger/v2/environments"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mitchellh/cli"
)

type Rigfile struct {
  Workflows []Workflow
}

type Workflow struct {
  Interface WorkflowInterface
  Implementations []WorkflowImplementation
}

type WorkflowInterface struct {
  Url string
}

type WorkflowImplementation struct {
  Name string
  Url string
}

// type WorkflowConfig {
//
// }

// type Rigfile struct {
//   Tools []Tool
//   Commands []Command
// }
//
// type Tool struct {
//   Name string
//   Type string
//   Install InstallMethod
// }
//
// type InstallMethod struct {
//   Method string
//   Url string
//   Env []string
//   Path []string
// }
//
// type Providers struct {
//   Name string
// }

type WorkflowDefinition struct {
  Metadata WorkflowMetadata
  Commands []WorkflowCommands
}

type WorkflowMetadata struct {
  Name string
  Description string
}

type WorkflowCommands struct {
  Name string
  Description string
  Inputs WorkflowInputs
  Outputs WorkflowOutputs
}

type WorkflowOutputs struct {
  Vars []string
}

type WorkflowInputs struct {
  Vars []string
}

type EnvironmentsFile struct {
	Default string
	Rigs map[string]RigInstance
}

type RigInstance struct {
	Rig string
	Url string
}


var Rig Rigfile
var workflow WorkflowDefinition
var Commands map[string]*cobra.Command
var cfgFile string
var Ui *cli.ColoredUi

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rigger [command]",
	Short: "rigger helps you assemble and construct things easily",
	Long: `...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
  loadCommands()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.makeup/makeup.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func loadCommands() {
	Ui = &cli.ColoredUi {
		OutputColor: cli.UiColorNone,
		InfoColor: cli.UiColorBlue,
		ErrorColor: cli.UiColorRed,
		WarnColor: cli.UiColorYellow,
		Ui: &cli.BasicUi {
			Reader: os.Stdin,
			Writer: os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}

  Commands = make(map[string]*cobra.Command)

	environments.Load()

	rigfile, err := os.Open("Rigfile.yaml")
  if err != nil {
    log.Printf("[ERROR] %s", err)
  } else {
    viper.ReadConfig(rigfile)
    viper.Unmarshal(&Rig)
  }

  workflowYaml := Rig.Workflows[0].Interface.Url + "/workflow.yaml"

  // log.Printf("Location of workflow interface: %v", workflowYaml)

  file, err := os.Open(workflowYaml)
  if err != nil {
    log.Printf("[ERROR] %s", err)
  } else {
    viper.ReadConfig(file)
    viper.Unmarshal(&workflow)
  }

  RootCmd.Long = workflow.Metadata.Description

  for _,element := range workflow.Commands {
    // log.Printf("[DEBUG] Adding command %s", element.Name)
    Commands[element.Name] = &cobra.Command{
      Use:   element.Name,
      Short: strings.TrimSpace(element.Description),
      Run: func(cmd *cobra.Command, args []string) {
				var envName string
				if len(args) == 0 {
					envName = environments.Environments.Default
				} else {
					envName = args[0]
				}

				if val, ok := environments.Environments.Rigs[envName]; ok {
					var Rig RigImplementation
					file, err := os.Open(val.Url + "/Rig.yaml")
					if err != nil {
						log.Printf("[ERROR] %s", err)
					} else {
						viper.ReadConfig(file)
						viper.Unmarshal(&Rig)
					}

					execCmd := exec.Command(Rig.Commands[cmd.Name()].Script)
					execCmd.Dir = val.Url
					execCmd.Stdout = os.Stdout
					execCmd.Stderr = os.Stderr
					fmt.Println(execCmd.Dir)
					if err := execCmd.Run(); err != nil {
						log.Fatal(err)
					}
				} else {
					log.Printf("[ERROR] No such environment: %v", envName)
				}
      },
    }
    RootCmd.AddCommand(Commands[element.Name])
  }
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
  log.Print("[DEBUG] Loading runtime config...")
}
