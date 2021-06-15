// Code generated by go-bindata.
// sources:
// pkg/backend/pg/migrations/0001_create.up.sql
// DO NOT EDIT!

package migrations

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.CopyN(&buf,  gz, 1024*1024*256)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __0001_createUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x96\xcf\x6f\xe2\x3a\x10\xc7\xef\xfc\x15\x73\x03\xa4\x56\x7a\x7a\x4f\x6f\x2f\x3d\xb1\xd4\x2b\x55\xdb\xd2\x2e\xa5\xd2\xf6\x14\x4d\xe2\x81\xb8\x38\x36\xf5\x0f\x28\xff\xfd\x2a\x71\x92\x86\x00\x29\xa8\xab\xde\xb0\x67\xc6\x9e\xf9\x78\xbe\x43\xc6\x53\x36\x9a\x31\x98\x8d\xbe\xdf\x32\x78\xf5\xe4\xc9\xc2\xa0\x07\x00\x20\x38\xc4\x62\x61\xc9\x08\x94\xf0\x30\xbd\xb9\x1b\x4d\x9f\xe1\x27\x7b\xbe\x28\xac\x0a\x33\x82\x35\x9a\x24\x45\x33\xf8\xf7\xff\xff\x86\xa0\xb4\x03\xe5\xa5\x0c\xf6\x35\x08\xe5\x68\x41\xa6\xb5\x6f\xc8\x19\x41\x16\x6c\x86\x52\x0a\xe5\x5a\x66\xaf\xc4\xab\xa7\x08\xcd\xc2\x42\xac\xb5\x24\x54\x2d\x0f\xee\x0d\x3a\xa1\x55\x9e\xdc\x7e\x7c\x92\x52\xb2\x14\x2a\xfa\xc0\x4b\xa2\xc8\x3e\xf0\x79\xd1\x71\x64\x93\x94\x32\x84\x17\xab\x55\x1c\x76\x63\x4c\x96\x7a\x3e\x8f\x84\x12\x4e\xa0\x3c\x7a\x06\x70\x9a\xa3\x97\x0e\xfe\xd9\x8d\xcb\xf0\xed\xec\x98\x39\x26\x4e\x1b\x98\x4b\x8d\xc7\x9d\x57\xe8\x2d\xf1\x3d\x66\xb5\xdb\x1c\xa5\xa5\xe0\x7a\x79\x09\x5e\x15\xfe\x20\x14\x17\x09\x3a\xb2\xb0\x49\x49\x01\x86\x0e\x00\xa3\x37\x30\xd8\xa4\x22\x49\x41\x58\xd0\x2e\x25\xb3\x11\xb9\x7b\x96\x79\x87\xb1\xa4\x21\xb8\x14\x5d\x75\xda\x06\x2d\xac\x0c\xad\x85\xf6\x56\x6e\xab\x54\x84\x05\xa5\x37\xd5\x55\xbc\x7c\xde\x33\xf2\x8c\xa5\x4e\x96\x27\xfa\x26\x86\xd0\x11\x8f\xd0\x81\x13\x19\x59\x87\xd9\x6a\x3f\x40\xe9\xcd\x60\x58\x36\xda\x8a\x9f\x17\xc0\x49\x52\x3b\x20\x58\x9e\x26\x37\xbf\x9e\x18\x0c\x72\x49\x5c\xc0\x7a\xd8\x1b\x5e\xf5\x7a\xa5\xae\x6e\x26\xd7\xec\x77\xa9\xab\xa8\xd0\xcc\xfd\xa4\x96\x59\xbe\x6e\xf8\x36\x34\x18\x49\x8c\x49\x56\x4a\x0c\x8f\xd2\x21\xb6\x1d\x31\x7e\xdb\xd7\x22\x4a\xdf\x65\x6f\x88\x1b\x06\xc5\x65\x17\x50\x16\x93\x47\xee\x14\x34\x7b\x7e\x60\x41\x1b\x0e\x9d\xb7\x30\x7a\x04\x36\x79\xba\x2b\x33\xed\x17\xd1\xbc\x1f\xce\xed\x1b\xaf\x94\x50\x8b\x6a\x99\xe8\x6c\x95\x43\xac\xd6\x73\x14\xf2\xdd\x99\x13\xd6\xbf\x13\x54\x09\xc9\x86\x51\xa8\x35\x4a\xc1\xfb\xbd\x3d\x5e\x2f\x3a\x3e\x6d\x62\x1d\x9b\x48\x01\x78\x88\xdd\x91\xe3\x94\xfd\x60\x53\x36\x19\xb3\xc7\xfa\xc5\x04\x1f\xe6\x0f\x78\xcd\x6e\xd9\x8c\xc1\x78\xf4\x38\x1e\x5d\xb3\x70\x0c\x3a\x47\xd9\xca\xd5\x83\x2d\xec\x96\x94\x1a\xc0\x76\x2f\x2f\x06\x5d\x63\xbe\x70\x74\xf5\xbc\x29\x36\x48\x05\xa4\xa7\x37\xaa\x75\x68\x0e\x37\x6a\xc5\x7f\xd7\xb6\xdf\xae\x39\xd2\x2a\xdd\xfb\x49\x49\x38\xac\x0f\xe1\x8f\x8a\x71\x7a\xda\x23\xe4\xee\xdd\xa8\xc3\x6d\x1d\xa0\x3f\xd9\xec\x95\x5a\x43\x26\xc7\x1b\xbd\x46\x51\x96\x17\x95\xa9\x07\x20\x75\xcd\x61\xf7\x30\x96\xf0\x5f\xf4\x65\x60\x9a\xbd\x73\xe6\x4c\x3c\x52\x78\x59\x40\xbb\xf4\xba\xae\x8e\xe2\x0d\x59\x2f\xdd\x97\xd5\xfe\x91\xce\xda\x6c\xc8\x18\x6d\xc0\xd1\x9b\x3b\x2e\x9a\xf6\x47\xc3\x41\xf5\xd4\x4e\x87\x11\x96\x18\x5a\x04\x6b\x38\x1d\x00\xc3\x97\x90\x22\x5b\x31\xfc\x0b\x94\x96\xb4\x85\x78\xeb\x08\xf7\xdf\xe0\xd3\xcd\xf2\x9e\x6f\x94\x5f\x53\x56\xda\xac\x62\x49\xdb\xe1\x55\x67\xe0\x2e\xa5\x66\x6c\x0d\xea\x4f\x00\x00\x00\xff\xff\x85\x49\xc9\xe1\xb0\x0a\x00\x00")

func _0001_createUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_createUpSql,
		"0001_create.up.sql",
	)
}

func _0001_createUpSql() (*asset, error) {
	bytes, err := _0001_createUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_create.up.sql", size: 2736, mode: os.FileMode(436), modTime: time.Unix(1622868146, 0)}
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
	"0001_create.up.sql": _0001_createUpSql,
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
	"0001_create.up.sql": &bintree{_0001_createUpSql, map[string]*bintree{}},
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

