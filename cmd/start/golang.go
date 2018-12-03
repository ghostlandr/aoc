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

package start

import (
	"fmt"
	"os"
	"text/template"
)

const (
	dayTemplate = `package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("{{.Dir}}/input.txt")
	stringData := strings.Split(string(data), "\n")

	for s := range stringData {
		fmt.Println(s)
	}
}
`
)

type Directory struct {
	Dir string
}

func GenerateGoFiles(dayFolder string) {
	tmpl, err := template.New("code").Parse(dayTemplate)
	if err != nil {
		fmt.Printf("failed to create code template: %s\n", err)
		return
	}
	f, err := os.OpenFile(fmt.Sprintf("./%s/part1.go", dayFolder), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("failed to open file %s\n", err)
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get pwd: %s\n", err)
	}
	d := Directory{Dir:fmt.Sprintf("%s/%s", dir, dayFolder)}
	err = tmpl.Execute(f, d)
	if err != nil {
		fmt.Printf("failed to write go file %s\n", err)
		return
	}
	if err = f.Close(); err != nil {
		fmt.Printf("failed to close file %s\n", err)
		return
	}
	f2, err := os.OpenFile(fmt.Sprintf("./%s/part2.go", dayFolder), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("failed to open file %s\n", err)
		return
	}
	err = tmpl.Execute(f2, d)
	if err != nil {
		fmt.Printf("failed to write go file %s\n", err)
		return
	}
	if err = f2.Close(); err != nil {
		fmt.Printf("failed to close file %s\n", err)
		return
	}
}
