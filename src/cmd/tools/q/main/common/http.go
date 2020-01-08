package common

import (
	"flag"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	EndPoint *string
)

// This is used across the commands
func init() {
	EndPoint = flag.String("endpoint", "http://bmcqueen-ld1:7201", "url for endpoint")
}

func DoGet(url string, logger *zap.SugaredLogger, getter func(reader io.Reader, logger *zap.SugaredLogger)) {

	logger.Debugf("url:%s:\n", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func() {
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}()

	getter(resp.Body, logger)

}

func DoPost(url string, data io.Reader, logger *zap.SugaredLogger, getter func(reader io.Reader, logger *zap.SugaredLogger)) {

	logger.Debugf("url:%s:\n", url)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, data)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer func() {
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}()

	logger.Debugf("resp:%d:\n", resp.StatusCode)

	getter(resp.Body, logger)

}

func DoDelete(url string, logger *zap.SugaredLogger, getter func(reader io.Reader, logger *zap.SugaredLogger)) {

	logger.Debugf("url:%s:\n", url)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer func() {
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}()

	getter(resp.Body, logger)

}

