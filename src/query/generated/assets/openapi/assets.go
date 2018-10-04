// Code generated by "esc -modtime 12345 -prefix openapi/ -pkg openapi -ignore .go -o openapi/assets.go ."; DO NOT EDIT.

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

package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
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
	return nil, nil
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
		local:   "openapi/spec.yml",
		size:    11694,
		modtime: 12345,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "openapi",
	},
}
