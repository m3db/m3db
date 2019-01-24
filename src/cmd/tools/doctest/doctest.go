// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	operationRegEx  = `(?s)<operation>.*?</operation>|(?s)<operation_json>.*?</operation_json>`
	validationRegEx = `(?s)<validation>.*?</validation>`

	operationOpen      = `<operation>`
	operationClose     = `</operation>`
	operationJSONOpen  = `<operation_json>`
	operationJSONClose = `</operation_json>`
	validationOpen     = `<validation>`
	validationClose    = `</validation>`

	markdownHeader = "<!-- Code generated by doctest. DO NOT EDIT. -->\n<!-- source: %s -->\n\n"
	bashHeader     = "#!/usr/bin/env bash\n\n# Code generated by doctest. DO NOT EDIT.\n# source: %s\n\n"
)

var (
	file           string
	destDir        string
	markdown, bash bool

	goPath = build.Default.GOPATH + "/src/"
)

func init() {
	flag.StringVar(&file, "file", "", "File to be parsed")
	flag.StringVar(&destDir, "dest", "", "Directory where the output will be saved")

	flag.BoolVar(&markdown, "markdown", true, "generate a markdown file based on the source (default=true)")
	flag.BoolVar(&bash, "bash", true, "generate a bash script based on the source (default=true)")
}

func main() {
	flag.Parse()

	if file == "" {
		log.Fatal("'file' flag is required")
	}

	if destDir == "" {
		destDir = filepath.Dir(file)
	}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read source file: %v\n", err)
	}

	operationRegEx := regexp.MustCompile(operationRegEx)
	validationRegEx := regexp.MustCompile(validationRegEx)

	if bash {
		if err := bashScript(f, destDir, file, operationRegEx, validationRegEx); err != nil {
			log.Fatalf("unable to create bash script: %v", err)
		}
	}

	if markdown {
		if err := markdownFile(f, destDir, file, validationRegEx, operationRegEx); err != nil {
			log.Fatalf("unable to create markdown file: %v", err)
		}
	}
}

func markdownFile(contents []byte, destDir, fileName string, validationRegEx *regexp.Regexp, opRegex *regexp.Regexp) error {
	// create new markdown file based on name of source file
	mdFileName := destDir + "/" + strings.Replace(filepath.Base(fileName), ".source", "", -1)
	mdFile, err := os.Create(mdFileName)
	if err != nil {
		return err
	}
	defer mdFile.Close()
	mdFile.Chmod(0644)

	sourceFile := strings.TrimPrefix(fileName, goPath)
	mdFile.Write([]byte(fmt.Sprintf(markdownHeader, sourceFile)))

	// find all of the operations in the source file
	matches := opRegex.FindAll(contents, 100)
	for _, match := range matches {
		validations := validationRegEx.FindAll(match, 100)
		if len(validations) == 0 {
			continue
		}

		// remove the validations and if the operation block is empty, replace it with nothing
		noValidationBlock := bytes.Replace(match, validations[0], nil, -1)

		if string(noValidationBlock) == fmt.Sprintf(operationOpen+"\n"+operationClose) {
			contents = bytes.Replace(contents, match, []byte(""), -1)
		}
	}

	// convert tags to markdown syntax
	OpOpen := regexp.MustCompile(operationOpen)
	OpClose := regexp.MustCompile(operationClose)
	OpJSONOpen := regexp.MustCompile(operationJSONOpen)
	OpJSONClose := regexp.MustCompile(operationJSONClose)

	contents = OpOpen.ReplaceAll(contents, []byte("```"))
	contents = OpClose.ReplaceAll(contents, []byte("```"))
	contents = OpJSONOpen.ReplaceAll(contents, []byte("```json"))
	contents = OpJSONClose.ReplaceAll(contents, []byte("```"))
	contents = validationRegEx.ReplaceAll(contents, nil)
	mdFile.Write(contents)

	return nil
}

func bashScript(contents []byte, destDir, fileName string, operationRegEx, validationRegEx *regexp.Regexp) error {
	// create bash script file based on name of source file
	scriptFileName := destDir + "/" + strings.Replace(filepath.Base(fileName), ".md.source", ".sh", -1)
	script, err := os.OpenFile(scriptFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer script.Close()

	scriptTemplate, err := ioutil.ReadFile("script_template.txt")
	if err != nil {
		return err
	}

	sourceFile := strings.TrimPrefix(fileName, goPath)
	script.Write([]byte(fmt.Sprintf(bashHeader, sourceFile)))

	script.Write(scriptTemplate)

	// find all of the operations in the source file
	matches := operationRegEx.FindAll(contents, 100)

	for _, match := range matches {
		// find any validations
		validationRaw := validationRegEx.Find(match)

		// write operations to script
		match = validationRegEx.ReplaceAll(match, nil)
		operation := removeTags(string(match), []string{operationOpen, operationClose, operationJSONOpen, operationJSONClose})
		script.WriteString(operation)

		// write validations to script
		if len(validationRaw) > 0 {
			validation := removeTags(string(validationRaw), []string{validationOpen, validationClose})
			script.WriteString(validation)
		}
	}

	return nil
}

func removeTags(text string, tagNames []string) string {
	for _, tag := range tagNames {
		text = strings.Replace(text, tag, "", -1)
	}
	return text
}
