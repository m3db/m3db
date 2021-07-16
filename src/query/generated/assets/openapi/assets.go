// Code generated by "esc -modtime 12345 -prefix openapi/ -pkg openapi -ignore .go -o openapi/assets.go ."; DO NOT EDIT.

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

package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/asset-gen.sh": {
		name:    "asset-gen.sh",
		local:   "asset-gen.sh",
		size:    238,
		modtime: 12345,
		compressed: `
H4sIAAAAAAAC/0zKz0rEMBDH8Xue4rfTnBbSsP45LR5EfAHrTUTWdpIO0hlJIgjiu0sr6M5l4PP9dbv4
Khrr7NztMNw/vgwPdzf+4JIVCERB/s8p7o+YzAGAJOzwhDDBC56PaDPrFtYbTZvoB2+QxG2fx9lAmZXL
qYlmpGILvNBvrSPCYlOThUGHi8ura0J4L5zkE+S/pOv28Tuu9pb/gRAkqxVGnw3BzqanWrnVPhuBenKT
KbufAAAA//9BiTev7gAAAA==
`,
	},

	"/index.html": {
		name:    "index.html",
		local:   "openapi/index.html",
		size:    636,
		modtime: 12345,
		compressed: `
H4sIAAAAAAAC/0ySQW/bMAyF7/kVjC+9RJaHDtiQyd6wpceuQ9DLblUk2lYrS55IpzC2/ffBUdLlRr4n
fXwgqNa7h2+PP3/cQc+Db1ZqLcTq+8Pj3Rb2U4CnQb8gaCJk0WEQvyZM8xO4FuY4QTbDDKbXoUMCjsC9
I2idx/VKiGalMhZA9ajtUgAoduyxub/dfYU97qJRMivZHZD1QkyEXBcTt+JjIa+9oAesi6PD1zEmLsDE
wBi4Ll6d5b62eHQGxanZgAuOnfaCjPZYvyvOIO/CC/QJ27romUfaStnGwFR2MXYe9eioNHGQhuhzqwfn
5/p+8TElzdvbqtq8r6rNh6r6s4+HyPFaKiChrwvi2SP1iHwZelJyDXCIdobf5wZg0KlzYQvVpzdp1Na6
0F1pfzNHvoGUvKxVLbzznIQ2GqARjZiSr2/iiEGPThJrdkYuRjkP/qZR8vT0Es8kNzJQMv+XYmwon8mi
d8dUBmQZxiF/+uI1I7E8TMF6pCyWxDpY7WPA8pmKZsl6ouawOaOS+Sj+BQAA//8by2IcfAIAAA==
`,
	},

	"/spec.yml": {
		name:    "spec.yml",
		local:   "openapi/spec.yml",
		size:    26055,
		modtime: 12345,
		compressed: `
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
jRjeuCiq+38AAAD//3KwMhXHZQAA
`,
	},

	"/": {
		name:  "/",
		local: `.`,
		isDir: true,
	},

	"/openapi": {
		name:  "openapi",
		local: `openapi`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	".": {
		_escData["/asset-gen.sh"],
		_escData["/openapi"],
	},

	"openapi": {
		_escData["/index.html"],
		_escData["/spec.yml"],
	},
}
