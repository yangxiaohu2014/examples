// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package main generated by go-bindata.// sources:
// ../template_blocks_0/views/500.html
// ../template_blocks_0/views/index.html
// ../template_blocks_0/views/layouts/error.html
// ../template_blocks_0/views/layouts/main.html
// ../template_blocks_0/views/partials/footer.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _views500Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4e\xf3\x30\x14\x85\xf7\x48\x79\x87\xfb\x67\xe9\x8f\xd4\xa4\xea\x5a\x42\x37\x06\x06\xa6\x22\x21\x46\xc7\x3e\xad\x2d\x9c\x7b\x23\xdb\x49\x89\xa2\xbc\x3b\x6a\x5a\x04\x2c\xac\xf7\x7e\xe7\x3b\x3a\xf5\xbf\xb2\xa4\x37\xe9\x49\x2b\x26\x83\xa3\x63\x50\x2b\x01\x94\xac\x62\x12\x06\x35\x5e\xf4\x7b\x95\x67\x2f\x16\x17\x40\xf5\x3e\x2d\x77\x17\xa9\xd0\xc2\x09\x9c\x0a\x3a\x5b\xa7\x2d\x45\x2b\xbd\x37\xd4\x5c\xd2\xa0\x56\x39\xa6\x84\xb6\xf3\x2a\x61\x15\xa9\x11\x33\x56\x79\x76\x90\x35\x61\x00\x93\x3b\x92\x4b\xab\x48\xad\x8b\xd1\xf1\x89\xfe\x47\x80\x1c\x1b\x7c\x54\x36\xb5\xfe\x6e\x7d\x7d\x2b\x63\x60\x48\xf5\x49\x5a\x95\x9c\x56\xde\x8f\xd4\x8c\x4b\xc3\xe0\x70\x26\xf0\xc9\x31\xaa\x3c\x7b\xb5\x60\x1a\xa5\x27\x06\x0c\x25\xf9\x63\xce\x7a\xe1\xac\x1a\x70\xe1\x9a\x1b\x13\x3b\x68\x77\x74\x7a\x97\x67\x65\xb9\xcf\xb3\x69\xfa\x52\x7c\x0f\x9d\xe7\x3c\xab\xed\x76\xff\xc4\x09\x81\x95\xa7\x03\xc2\x80\x40\x8f\x21\x48\xa8\x37\x76\x7b\xcd\x81\xcd\x82\xfe\x92\xb4\x88\x51\x9d\x70\x93\x74\x14\xd3\xe8\xf1\x50\x68\xf1\x12\x76\x01\xe6\xbe\xd8\x4f\x53\xf5\x7c\xa5\xe6\xb9\xde\x74\x3f\x65\x9f\x01\x00\x00\xff\xff\xb3\xfa\x91\xbd\xaa\x01\x00\x00")

func views500HtmlBytes() ([]byte, error) {
	return bindataRead(
		_views500Html,
		"views/500.html",
	)
}

func views500Html() (*asset, error) {
	bytes, err := views500HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/500.html", size: 426, mode: os.FileMode(438), modTime: time.Unix(1596515591, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb4\xf3\xcc\x4b\x49\xad\x50\x70\xca\x4f\xa9\xb4\xd1\xcf\x30\xb4\x03\x04\x00\x00\xff\xff\xcc\x4b\x98\x69\x13\x00\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 19, mode: os.FileMode(438), modTime: time.Unix(1596514225, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsLayoutsErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xbf\x4e\xf3\x40\x10\xc4\x7b\x4b\x7e\x87\xfd\xb6\xfe\x6c\x43\x47\x71\xe7\x26\x40\x0b\x45\x28\x28\x37\x77\xa3\xf8\xc4\xfd\x89\xe2\x55\x22\x74\xf2\xbb\xa3\x18\x83\x44\xb5\xda\x99\x9f\x66\x57\x63\xfe\x3d\xbe\xec\xf6\xef\xaf\x4f\x34\x69\x8a\x63\xdb\x98\xdb\xa4\x28\xf9\x68\x19\x99\x57\x05\xe2\xc7\xb6\x21\x22\x32\x09\x2a\xe4\x26\x39\xcf\x50\xcb\x6f\xfb\xe7\xee\x81\xff\x78\x59\x12\x2c\x5f\x02\xae\xa7\x72\x56\x26\x57\xb2\x22\xab\xe5\x6b\xf0\x3a\x59\x8f\x4b\x70\xe8\xd6\xe5\x3f\x85\x1c\x34\x48\xec\x66\x27\x11\xf6\xbe\xbf\xfb\xcd\xd2\xa0\x11\x63\xad\xfd\xae\x78\x2c\x8b\x19\xbe\x85\xb6\x31\xc3\xf6\x8e\x39\x14\xff\xb9\xe1\xb5\x92\x22\x9d\xa2\x28\x88\xb7\x8b\x4c\xfd\xb2\xb4\xcd\x0f\x70\x88\xc5\x7d\x10\x27\xcc\xb3\x1c\xb1\x9a\xb5\x22\xfb\x1b\x63\x86\x2d\xcb\x0c\x6b\x0b\x5f\x01\x00\x00\xff\xff\xbe\xb7\x11\x67\x15\x01\x00\x00")

func viewsLayoutsErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsLayoutsErrorHtml,
		"views/layouts/error.html",
	)
}

func viewsLayoutsErrorHtml() (*asset, error) {
	bytes, err := viewsLayoutsErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/layouts/error.html", size: 277, mode: os.FileMode(438), modTime: time.Unix(1596514340, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsLayoutsMainHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xb1\x4e\xf4\x30\x10\x84\xfb\x48\x79\x87\xf9\x5d\xff\x49\xa0\xa3\xb0\xd3\x70\xd0\x21\x28\x42\x41\xb9\x24\x1b\x62\xc9\x71\xa2\x64\xb9\x13\xb2\xfc\xee\xc8\x39\x0b\xe9\x2a\xcf\xfa\xb3\x66\xc6\xab\xff\x9d\x5e\x1f\xbb\x8f\xb7\x27\x4c\x32\xbb\xb6\x2c\x74\x3a\xe1\xc8\x7f\x19\xc5\x5e\x1d\x37\x4c\x43\x5b\x16\x00\xa0\x67\x16\x42\x3f\xd1\xb6\xb3\x18\xf5\xde\x3d\x57\x0f\xea\x86\x79\x9a\xd9\xa8\xb3\xe5\xcb\xba\x6c\xa2\xd0\x2f\x5e\xd8\x8b\x51\x17\x3b\xc8\x64\x06\x3e\xdb\x9e\xab\x63\xf8\x0f\xeb\xad\x58\x72\xd5\xde\x93\x63\x73\x5f\xdf\xfd\x79\x89\x15\xc7\x6d\x08\xb0\x23\xea\x2e\x0d\x88\x31\x84\x1b\xcd\x6e\x4f\xea\xc4\x23\x7d\x3b\xc1\x0b\x59\x8f\x03\x27\xe6\x07\xc4\xa8\x9b\xab\x4f\x59\xe8\x26\xff\x42\x7f\x2e\xc3\x4f\x4e\x09\x01\xc2\xf3\xea\x48\x18\x2a\x17\x55\xa8\x11\x63\x59\x94\x85\x1e\x97\x45\x78\x4b\x25\x56\xda\x52\x4f\xa8\x2c\xf6\xe6\xca\x14\xea\x14\x92\x1f\xa6\x94\xec\xae\x9b\x63\x9d\xbf\x01\x00\x00\xff\xff\x44\x95\x63\x98\x5e\x01\x00\x00")

func viewsLayoutsMainHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsLayoutsMainHtml,
		"views/layouts/main.html",
	)
}

func viewsLayoutsMainHtml() (*asset, error) {
	bytes, err := viewsLayoutsMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/layouts/main.html", size: 350, mode: os.FileMode(438), modTime: time.Unix(1596514155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsPartialsFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb6\x73\xcb\xcf\x2f\x49\x2d\x52\x08\x48\x2c\x2a\xc9\x4c\xcc\xb1\xd1\xcf\x30\xb6\x03\x04\x00\x00\xff\xff\x08\xe6\xe9\xf8\x17\x00\x00\x00")

func viewsPartialsFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsPartialsFooterHtml,
		"views/partials/footer.html",
	)
}

func viewsPartialsFooterHtml() (*asset, error) {
	bytes, err := viewsPartialsFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/partials/footer.html", size: 23, mode: os.FileMode(438), modTime: time.Unix(1596514093, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"views/500.html":             views500Html,
	"views/index.html":           viewsIndexHtml,
	"views/layouts/error.html":   viewsLayoutsErrorHtml,
	"views/layouts/main.html":    viewsLayoutsMainHtml,
	"views/partials/footer.html": viewsPartialsFooterHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"views": {nil, map[string]*bintree{
		"500.html":   {views500Html, map[string]*bintree{}},
		"index.html": {viewsIndexHtml, map[string]*bintree{}},
		"layouts": {nil, map[string]*bintree{
			"error.html": {viewsLayoutsErrorHtml, map[string]*bintree{}},
			"main.html":  {viewsLayoutsMainHtml, map[string]*bintree{}},
		}},
		"partials": {nil, map[string]*bintree{
			"footer.html": {viewsPartialsFooterHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
