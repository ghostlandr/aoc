// Copyright © 2018 Graham Holtslander <menello@gmail.com>
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
	"errors"
	"fmt"
	"github.com/gholtslander-va/aoc/cmd/start"
	"github.com/spf13/cobra"
	"os/exec"
	"strconv"
)

var language string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new day of Advent of Code",
	Long: `Generate the following files in the language of your choice:

A new folder called d<date> will be created, i.e. for start 3 a folder
d3 will be created. Example file structure follows:

d3/
	input.txt
	part1.<language extension>
	part2.<language extension>
`,
	Run: func(cmd *cobra.Command, args []string) {
		day, _ := strconv.Atoi(args[0])
		dayFolder := fmt.Sprintf("d%d", day)
		err := exec.Command("mkdir", dayFolder).Run()
		if err != nil {
			// Probably fails because they already created the directory?
			//fmt.Printf("failed to create new directory: %s\n", err)
			//return
		}
		err = exec.Command("touch", fmt.Sprintf("./%s/input.txt", dayFolder)).Run()

		if language == "go" {
			start.GenerateGoFiles(dayFolder)
		}
		if language == "python" {
			start.GeneratePythonFiles(dayFolder)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("day must be an integer value")
		}
		if i > 25 || i < 1 {
			return errors.New("day must be between 1 and 25 (inclusive)")
		}
		if language != "go" && language != "python" {
			return errors.New("only go and python are supported right now")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringVarP(&language,"language", "l", "go", "Language you're starting your day in")
}
