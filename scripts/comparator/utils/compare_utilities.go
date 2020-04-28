// Copyright (c) 2019 Uber Technologies, Inc.
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

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"

	"go.uber.org/zap"
)

// InputQueries is a slice of InputQuery.
type InputQueries []InputQuery

// InputQuery is the JSON representation of a query to be compared.
type InputQuery struct {
	// QueryGroup is the general category for these queries.
	QueryGroup string `json:"queryGroup"`
	// Queries is the list of raw queries.
	Queries []string `json:"queries"`
	// Steps is the list of step sizes for these queries.
	Steps []string `json:"steps"`
}

// PromQLQueryGroup is a list of constructed PromQL query groups.
type PromQLQueryGroup struct {
	// QueryGroup is the general category for these queries.
	QueryGroup string
	// Queries is a list of PromQL compatible queries.
	Queries []string
}

func (q InputQueries) constructPromQL(
	start int64,
	end int64,
) []PromQLQueryGroup {
	queries := make([]PromQLQueryGroup, 0, len(q))
	for _, inQuery := range q {
		queries = append(queries, inQuery.constructPromQL(start, end))
	}

	return queries
}

func (q InputQuery) constructPromQL(start int64, end int64) PromQLQueryGroup {
	queries := make([]string, 0, len(q.Queries)*len(q.Steps))
	for _, inQuery := range q.Queries {
		for _, inStep := range q.Steps {
			values := make(url.Values)
			values.Add("query", inQuery)
			values.Add("step", inStep)
			values.Add("start", strconv.Itoa(int(start)))
			values.Add("end", strconv.Itoa(int(end)))
			query := "/api/v1/query_range?" + values.Encode()

			queries = append(queries, query)
		}
	}

	return PromQLQueryGroup{
		QueryGroup: q.QueryGroup,
		Queries:    queries,
	}
}

func parseFileToQueries(
	fileName string,
	log *zap.Logger,
) (InputQueries, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Error("could not open file", zap.Error(err))
		return nil, err
	}

	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("could not read file", zap.Error(err))
		return nil, err
	}

	queries := make(InputQueries, 0, 10)
	if err := json.Unmarshal(buf, &queries); err != nil {
		log.Error("could not unmarshal queries", zap.Error(err))
		return nil, err
	}

	return queries, err
}

// ParseFileToPromQLQueryGroup parses a JSON queries file
// into PromQL query groups.
func ParseFileToPromQLQueryGroup(
	fileName string,
	start int64,
	end int64,
	log *zap.Logger,
) ([]PromQLQueryGroup, error) {
	queries, err := parseFileToQueries(fileName, log)
	if err != nil {
		return nil, err
	}

	return queries.constructPromQL(start, end), nil
}
