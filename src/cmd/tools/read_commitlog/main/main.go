// Copyright (c) 2021 Uber Technologies, Inc.
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
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pborman/getopt"
	"go.uber.org/zap"

	"github.com/m3db/m3/src/cmd/tools"
	"github.com/m3db/m3/src/dbnode/persist/fs/commitlog"
)

func main() {
	var (
		path     = getopt.StringLong("path", 'p', "", "file path [e.g. /var/lib/m3db/commitlogs/commitlog-0-161023.db]")
		idFilter = getopt.StringLong("id-filter", 'f', "", "ID Contains Filter (optional)")
	)
	getopt.Parse()

	rawLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("unable to create logger: %+v", err)
	}
	logger := rawLogger.Sugar()

	if *path == "" {
		getopt.Usage()
		os.Exit(1)
	}

	bytesPool := tools.NewCheckedBytesPool()
	bytesPool.Init()

	entryCount := 0

	opts := commitlog.NewReaderOptions(commitlog.NewOptions(), false)
	reader := commitlog.NewReader(opts)

	_, err = reader.Open(*path)
	if err != nil {
		logger.Fatalf("unable to open reader: %v", err)
	}

	for {
		entry, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			logger.Fatalf("err reading commitlog: %v", err)
		}

		id := entry.Series.ID
		if *idFilter != "" && !strings.Contains(id.String(), *idFilter) {
			continue
		}

		fmt.Printf("{id: %s, dp: %+v", id, entry.Datapoint) // nolint: forbidigo
		if len(entry.Annotation) > 0 {
			fmt.Printf(", annotation: %s", base64.StdEncoding.EncodeToString(entry.Annotation)) // nolint: forbidigo
		}
		fmt.Println("}") // nolint: forbidigo

		entryCount++
	}

	fmt.Printf("\n%d entries read\n", entryCount) // nolint: forbidigo
}
