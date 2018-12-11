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
		size:    23209,
		modtime: 12345,
		compressed: `
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
DBs6ttxdaGHJI/Uj2Jv1/hcAAP//0Ji7tqlaAAA=
`,
	},

	"/": {
		isDir: true,
		local: "openapi",
	},
}
