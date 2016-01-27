// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"log"
	"fmt"
	"os"
	// "os/exec"
	"strconv"
	"github.com/deis/rigger/v2/environments"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RigImplementation struct {
  Metadata RigMetadata
	Init RigInit
  Commands map[string]RigCommand
}

type RigMetadata struct {
  Name string
  Description string
}

type RigInit struct {
  Vars []RigVar
}

type RigVar struct {
  Name string
	Question string
	Default string
}

type RigCommand struct {
  Script string
}

var RigImpl RigImplementation

func PromptHeader(question string, variable string, defaultAnswer string) {
	Ui.Info(fmt.Sprintf("%s %s [ %s ]", question, variable, defaultAnswer))
}

func ListPrompt(question string, variable string, defaultAnswer string, list []WorkflowImplementation) string {
	var answer string = defaultAnswer

	PromptHeader(question, variable, defaultAnswer)

	for i, value := range list {
		fmt.Printf("%d %v\n", i+1, value.Name)
	}
	ret, _ := Ui.Ask("#? ")

	if ret != "" {
		i, _ := strconv.Atoi(ret)
		answer = list[i-1].Name
	}

	fmt.Printf("You chose: %v\n", answer)

	return answer
}

func SimplePrompt(question string, variable string, defaultAnswer string) string {
	var answer string

	for answer == "" {
		PromptHeader(question, variable, defaultAnswer)

		answer, _ = Ui.Ask("? ")

		if answer == "" {
			answer = defaultAnswer
		}
	}

	fmt.Printf("You chose: %v\n", answer)

	return answer
}

func chooseProvider(cmd *cobra.Command) string {
	var provider string

	if cmd.Flag("provider").Changed {
		provider = cmd.Flag("provider").Value.String()
	} else {
		provider = ListPrompt("What provider would you like to use?", "PROVIDER", "gcp", Rig.Workflows[0].Implementations)
	}
	return provider
}

func iterateInit(initSection RigInit) {
	for i := range initSection.Vars {
		SimplePrompt(RigImpl.Init.Vars[i].Question, RigImpl.Init.Vars[i].Name, RigImpl.Init.Vars[i].Default)
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <name>",
	Short: "Instantiate a workflow with an actual implementation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			rigName := args[0]

			provider := chooseProvider(cmd)
			for _, impl := range Rig.Workflows[0].Implementations {
				if impl.Name == provider {
					log.Printf("[DEBUG] Instantiated workflow with name: %s via provider %s", rigName, impl.Name)
					RigImpl = loadRigImpl(impl.Url)
					iterateInit(RigImpl.Init)

					environments.Add(rigName, environments.RigInstance {
						Rig: impl.Name,
						Url: impl.Url,
						// Vars: RigImpl.Init.Vars,
					})
					environments.Save()
				}
			}
		} else {
			Ui.Error("Please pass in a name for the instantiated workflow.")
			cmd.Usage()
		}
	},
}

func concretizeCommands() {
	for _, value := range Commands {
		value.Run = func(cmd *cobra.Command, args []string) {
			Ui.Error("Undefined!")
		}
	}
}

func loadRigImpl(url string) RigImplementation {
	rigImplYaml := url + "/rig.yaml"
	var rigYaml RigImplementation

	log.Printf("[DEBUG] Location of rig implementation: %v", rigImplYaml)

	file, err := os.Open(rigImplYaml)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	} else {
		viper.ReadConfig(file)
		viper.Unmarshal(&rigYaml)
	}

	return rigYaml
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().String("provider", "default", "Pick a rig")
}
