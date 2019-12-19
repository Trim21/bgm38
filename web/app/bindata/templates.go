// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../templates/index.tmpl
// ../templates/login.tmpl
package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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
	_, err = io.Copy(&buf, gz)
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

// Mode return file modify time
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _indexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x4f\xd1\x0a\x83\x30\x0c\x7c\x17\xfc\x87\x5b\xdf\x9d\x3f\xa0\xfe\x4b\x6d\xc3\x2a\xc4\x46\x34\xe8\x9c\xf8\xef\xa3\x60\x99\x2c\x2f\x77\xe4\xb8\x4b\xae\x79\x78\x71\xba\x4f\x84\xa0\x23\x77\x65\xd1\x24\x04\xdb\xf8\x6a\xcd\x27\x54\x2e\x9a\xae\x2c\x8e\x03\x4a\xe3\xc4\x56\x09\xa6\xb7\x0b\xd5\x81\xac\x7f\xea\x38\xb1\x81\x09\x14\xc8\xe0\x3c\xcb\xa2\xe9\xc5\xef\x29\xc4\x0f\x2b\x1c\xdb\x65\x69\x8d\x93\xa8\x76\x88\x34\xa7\x20\x00\xb8\x8b\xb3\x6c\x79\xfd\x2f\x39\x61\x28\xbd\xb5\x72\x14\xf5\xe7\xce\x13\x88\x59\xb0\xc9\xcc\xfe\xe6\xaf\xfd\xb0\xe6\x33\x17\xcf\x98\x58\x7e\xaf\xbe\xca\x7e\x03\x00\x00\xff\xff\x30\x2d\x03\x48\xfe\x00\x00\x00")

func indexTmplBytes() ([]byte, error) {
	return bindataRead(
		_indexTmpl,
		"index.tmpl",
	)
}

func indexTmpl() (*asset, error) {
	bytes, err := indexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.tmpl", size: 254, mode: os.FileMode(438), modTime: time.Unix(1575461337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loginTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\xbd\x8e\xd3\x40\x10\xee\x23\xe5\x1d\xe6\xb6\xdf\x58\x80\x44\x65\x47\x42\x34\x54\xe8\x24\xb8\x07\x58\xdb\x13\x7b\xc5\xfe\x69\x77\x9c\x10\x22\x4b\xd4\x48\x3c\x00\xd2\x51\xd0\x51\xf2\x52\xe1\x39\xd0\xfa\xec\x3b\x27\xe7\x1c\x10\x5d\x71\x69\xb2\xb6\xbf\xf9\xe6\xfb\xbe\x59\x7b\xd3\x8b\xd2\x16\xb4\x75\x08\x35\x69\xb5\x9c\xcf\xd2\xf8\x0f\x4a\x98\x2a\x63\x9f\x6a\x5e\x18\xb6\x9c\xcf\x76\x3b\x20\xd4\x4e\x09\x42\x60\xb9\x08\x98\xd4\x28\xca\x05\x69\xa7\x18\xb0\x1a\x6b\x64\xd0\xb6\xf3\x59\x9a\xdb\x72\x1b\x49\x4a\xb9\x86\x42\x89\x10\x32\x56\x58\x43\x42\x1a\xf4\x91\x08\x00\x60\xfc\xd0\xdb\x0d\x10\x7e\x24\x5e\xa0\xa1\x3b\xc8\x31\xac\xb0\x8a\x2d\x77\x3b\x58\xbc\x2a\x48\x5a\xf3\x56\x68\x84\xb6\x4d\x93\x52\xae\x07\xd2\xf1\xfa\xb0\xc1\x03\xa4\x5c\x55\xfc\x25\xd8\xd5\x2a\x20\xc5\xf5\x0b\x88\x37\x75\xc9\x9f\x3d\x1f\xcb\x82\x31\x45\x47\x73\xc1\x79\x68\x9c\xf3\x18\x02\xbc\x21\xad\xae\xcc\x07\x63\x37\xe6\xbd\xf0\x15\x12\x70\x7e\x8c\x5f\x59\xaf\x41\x74\xe2\x33\x96\x28\x5b\x49\xc3\x40\x23\xd5\xb6\xcc\x98\xb3\x81\x62\x87\xc3\x9a\x63\xb9\x91\x82\x57\xde\x36\xee\x58\xcd\x2d\x5a\x89\x1c\xd5\xf2\xf7\xf5\xd7\xfd\x97\x1f\xfb\x6f\x3f\xf7\xd7\xbf\xf6\xdf\x3f\x4f\x63\x3b\xbc\x34\xae\x21\x88\xd3\xcf\x58\xb4\xcb\x0e\x7a\xc5\xc1\x79\xab\x18\x18\xa1\x31\x63\x01\xfd\x1a\x3d\x3b\x4d\xd7\xff\xd6\x42\x35\x98\xb1\x38\xad\xd7\xd6\xac\x64\xb5\x78\xd7\x55\x42\xdb\x9e\x14\x9e\xdc\x28\x9f\x08\xa0\x1f\xeb\x23\x45\xd3\x04\xf4\xd1\xcd\x7f\x65\x72\x63\x7f\x28\x9d\xce\xe8\x8c\x54\xae\x7a\xc2\xa7\x90\x8b\x13\x21\x6c\xac\x2f\xff\x31\x97\x01\x3e\x64\x73\x7b\xfd\xd7\x18\x26\xf7\xd7\xfd\x6c\x2e\x7b\xc2\xa7\x90\x4d\xd8\x9a\x02\x9c\xa0\xfa\x8c\x4d\x13\x6b\x2f\x05\xd5\xe7\x05\x73\xce\xab\xd6\x37\x7c\xec\xe0\x4e\x95\xdc\x73\x1f\x9a\x5c\x4b\x7a\x58\x7a\xef\x35\x27\x03\x39\x19\xee\xbc\xd4\xc2\x6f\xe1\xc0\xfb\xa4\xbc\x29\x15\x69\x12\xeb\xc6\x9f\xf8\x89\x53\xe1\xce\x59\x9a\x0c\x47\x54\xd2\x1d\x78\x7f\x02\x00\x00\xff\xff\xf6\xa1\x8c\x48\x00\x07\x00\x00")

func loginTmplBytes() ([]byte, error) {
	return bindataRead(
		_loginTmpl,
		"login.tmpl",
	)
}

func loginTmpl() (*asset, error) {
	bytes, err := loginTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "login.tmpl", size: 1792, mode: os.FileMode(438), modTime: time.Unix(1558294214, 0)}
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
	"index.tmpl": indexTmpl,
	"login.tmpl": loginTmpl,
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
	"index.tmpl": &bintree{indexTmpl, map[string]*bintree{}},
	"login.tmpl": &bintree{loginTmpl, map[string]*bintree{}},
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
