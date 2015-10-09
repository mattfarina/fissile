// Code generated by go-bindata.
// sources:
// scripts/dockerfiles/Dockerfile-base
// scripts/dockerfiles/Dockerfile-role
// scripts/dockerfiles/monitrc.erb
// scripts/dockerfiles/run.sh
// DO NOT EDIT!

package dockerfiles

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _scriptsDockerfilesDockerfileBase = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x41\x0a\x83\x30\x10\x45\xf7\x73\x8a\x81\x82\x3b\xcd\x15\x6a\xb1\x05\x17\x5a\x90\xf6\x00\x69\x3a\xc6\x80\x26\x69\x32\x2e\x8a\x78\xf7\x4a\x0b\x36\x8b\x59\xfc\x37\xff\xf3\x2e\xdd\xb5\xc1\x65\xc1\xe2\x24\x23\xd5\x93\xd4\x84\xeb\x0a\xd0\x94\x75\x7b\xdb\xee\xdc\xe1\xa0\xfa\xe3\xe0\xa9\x50\x6e\x02\x38\x60\x6d\x23\xcb\x71\x44\x1f\x28\xd0\x6b\x36\xd1\x30\xc5\xf4\x31\x39\x6b\x18\xa0\xbb\xb7\x28\x3d\xe7\x9a\x18\x67\xff\x94\x4c\x98\x65\x3b\x31\x69\x19\xf3\x37\x94\x55\xf5\x0b\x41\x15\x14\x1e\x28\x9c\x67\xb1\xa9\x45\x02\x53\x8b\x72\xb6\x37\x5a\x1b\x0b\xdf\xe9\x1e\xff\xc3\x1d\x09\xf8\x04\x00\x00\xff\xff\x1d\x34\xe1\x4f\xe5\x00\x00\x00")

func scriptsDockerfilesDockerfileBaseBytes() ([]byte, error) {
	return bindataRead(
		_scriptsDockerfilesDockerfileBase,
		"scripts/dockerfiles/Dockerfile-base",
	)
}

func scriptsDockerfilesDockerfileBase() (*asset, error) {
	bytes, err := scriptsDockerfilesDockerfileBaseBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/dockerfiles/Dockerfile-base", size: 229, mode: os.FileMode(420), modTime: time.Unix(1444350497, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsDockerfilesDockerfileRole = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\x51\x4b\xc3\x30\x10\xc7\xdf\xf3\x29\x8e\xe8\xeb\x12\x9f\x85\x81\x95\x55\x50\xb6\x4e\xca\x10\x44\x64\xa4\xe9\x69\x3b\xd7\xa6\x34\x5d\x1d\x8c\x7d\x77\x2f\x4d\xda\x31\x7c\x28\x24\xf7\xff\xfd\xae\x77\xe4\x29\x5d\xaf\xe0\x74\x82\xb2\xce\xf1\x08\x02\x78\xa6\x2c\x6e\xcb\x4a\x7d\x23\x87\xf3\x99\xb1\x55\xf4\x9c\x6c\xe8\x8b\x53\x28\xf4\xd7\x43\xd1\xa0\xd0\xa6\x62\x8c\x9c\xdf\xb2\x2b\xe0\x56\x9b\xba\xc3\x63\x07\xf7\x73\xd2\xc9\x98\x82\xd6\xec\xd1\x55\xa7\xd6\xae\xc0\xaf\x90\x9d\xc9\x2e\xc4\x20\x88\x17\x93\x59\xb8\x23\x68\x19\x3d\xc6\xcb\x20\xcd\x39\x29\x3e\x4f\x54\x85\xd4\x83\x53\x82\x7b\xa4\x61\x67\x3d\xb6\xb6\x34\x75\x80\xa8\xa5\x48\x7d\x22\xde\x7c\xe2\xf1\x2b\x2c\xfc\x71\x9c\x9d\x0f\x0b\x6f\x47\xc4\x09\x6e\x48\xac\xf3\x30\xee\xbf\x13\xbb\x81\x28\xcf\xa1\x51\xfa\x87\x4c\xcb\xa2\xc5\x62\xba\x80\xec\x55\x2b\x7b\xad\x1a\x39\x96\xe4\x28\xb8\x8d\xad\x39\xb4\x3a\x38\x3b\xb7\xee\x85\xa7\xeb\x2c\xc4\x93\xd2\x1e\x6a\x61\x8b\x81\xf6\x47\x90\xa6\xe9\x24\xbd\x06\x21\x71\xb2\x49\xdf\x5f\xd7\xf4\x46\xf0\xc1\xa7\xba\xe7\xf8\x27\xfb\x0b\x00\x00\xff\xff\x3d\x51\x32\x8e\xdf\x01\x00\x00")

func scriptsDockerfilesDockerfileRoleBytes() ([]byte, error) {
	return bindataRead(
		_scriptsDockerfilesDockerfileRole,
		"scripts/dockerfiles/Dockerfile-role",
	)
}

func scriptsDockerfilesDockerfileRole() (*asset, error) {
	bytes, err := scriptsDockerfilesDockerfileRoleBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/dockerfiles/Dockerfile-role", size: 479, mode: os.FileMode(420), modTime: time.Unix(1444350497, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsDockerfilesMonitrcErb = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\xce\xdd\x6a\xc3\x30\x0c\x05\xe0\x7b\x3f\xc5\xc1\x10\xd8\x06\xcb\x8f\x6f\x36\xc6\xda\x77\x51\x2d\xa7\x49\x71\x6d\x63\x3b\x09\x7d\xfb\xba\xf5\x4d\x21\x14\x81\x40\x9c\x4f\x42\xc9\x64\x30\x99\xab\x77\x18\x7a\xf1\x98\xac\x3f\x8f\xb3\x35\xe8\x56\x8a\xdd\xaa\x29\x74\x25\x9c\x73\xed\x6d\x49\xc5\x93\x4d\x39\x07\x46\xf0\x31\x43\xfd\x2a\x05\x72\x8c\x25\x19\x10\x73\x34\x29\x61\x50\x3f\x6d\x5f\x6a\x10\x00\x59\xeb\x37\xc8\xff\xe6\x80\xf0\x21\x27\x3d\xb6\xf5\x58\xf1\x51\x7e\xa2\x39\xca\xbf\x7d\x18\x28\xa5\xcd\x47\xae\x00\xd1\x10\x7f\x7b\x67\x6f\x42\xcc\x4e\xdb\x85\xf7\x1f\x7e\xd5\xcd\xa8\xdf\x8a\x8b\x3f\xbd\xa8\x7b\x00\x00\x00\xff\xff\x77\x3e\x7e\x87\xfc\x00\x00\x00")

func scriptsDockerfilesMonitrcErbBytes() ([]byte, error) {
	return bindataRead(
		_scriptsDockerfilesMonitrcErb,
		"scripts/dockerfiles/monitrc.erb",
	)
}

func scriptsDockerfilesMonitrcErb() (*asset, error) {
	bytes, err := scriptsDockerfilesMonitrcErbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/dockerfiles/monitrc.erb", size: 252, mode: os.FileMode(420), modTime: time.Unix(1444350497, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsDockerfilesRunSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x53\x4f\x8f\xd3\x3e\x10\xbd\xe7\x53\xcc\xcf\x8d\xb4\x97\x4d\xfd\x03\x6e\x8b\x72\xe3\x02\x07\xb4\x02\x6e\xcb\x2a\x72\x92\x49\xe3\x55\x6b\x47\xb6\x53\x56\x44\xfd\xee\x8c\x9d\x7f\x4d\x5a\x2e\x50\x89\x9c\xdc\xce\xf3\xbc\xf7\x66\x9e\x37\xff\xf1\x5c\x2a\x9e\x0b\x5b\x47\x16\x1d\x24\x18\x45\x85\x56\xb6\xdd\x67\xa2\x2c\x0d\x5a\x9b\xc6\x6f\x22\x59\xc1\xd3\x13\x24\x3f\x21\x5e\xd6\xe0\xf9\xf9\x3d\xb8\x1a\x55\x04\xb0\xba\xc5\xba\x0e\xa4\x2a\xf1\x15\xb6\xc0\x4a\xac\x44\xbb\x77\xd9\x12\xc3\xe0\x74\x62\x51\x25\x03\x63\x25\x77\x99\x75\xda\x60\xd6\x18\xac\xe4\x6b\x1a\xbf\x5d\xd2\xae\x01\x2b\xee\x8b\xfb\xbf\x13\xb0\x06\xce\x2a\x8c\xde\x63\x26\x95\x75\x42\x15\xfe\x40\x77\xd3\xf8\xdd\x99\x8a\x2b\x80\x73\x15\xd7\xee\xff\x1f\x3a\x6f\xe0\xd1\xe8\xc2\x0f\xcc\xe1\xa1\xd9\x0b\x87\x36\x22\x79\x3f\xa4\xab\xfb\xae\xf0\x90\xce\x62\xfd\x1f\x5e\x95\x87\x18\xa1\x76\x08\xb1\xbc\x87\xf8\x45\xe7\x1e\xb6\xfd\xa4\x73\x4b\xc5\x0d\xa4\x37\xfc\xa8\xdd\xf8\x7d\x1b\x25\x42\xa5\x0d\x78\x56\xd2\xe1\xd9\xb7\x9f\xc5\x01\xe1\xf6\xd4\xb3\xcd\x17\xb2\x39\x4e\xc8\x7b\x0d\xac\xb3\x1e\xa2\xe6\xba\x71\xbc\x2e\x2a\xde\x6f\x72\x47\xd9\x9d\x4e\xf0\x9d\x96\x90\x24\xa5\x70\x02\xee\x3a\x46\x77\xd9\x03\x74\xc0\x14\xc9\xa6\x13\x5b\xd9\xa0\x11\xdf\x03\x0b\x53\xa7\x6a\xdc\x5d\x59\x9f\x07\x34\xc2\x10\xde\xa1\xb1\xbe\xdb\x09\x4e\x77\x03\x8f\x6e\x5d\xd3\x3a\x00\xc6\x8f\xc2\xf0\x63\x21\x1a\x4e\xdd\x2d\x5f\xd1\xd0\xef\xc9\xd3\xf6\x03\x5a\x27\x95\x70\x52\xab\x47\xe1\x6a\x2f\xa2\x6f\xd6\xbf\x0c\x6a\x16\x77\xcb\x47\x32\x21\x86\xd4\x0f\x88\x75\x8a\x27\x58\x48\x13\x7d\xe4\x36\x24\x2b\xc8\x98\x79\xfc\x3a\xc7\xf2\x28\x72\xac\xae\x8c\x24\xd6\x14\x17\x66\xa6\xfc\x86\xca\xe4\xeb\xab\x6e\x4d\x81\xde\x52\x78\x4d\x7f\x18\x10\x9f\x04\x54\xe5\x10\xfd\xe1\x74\xf6\x78\x0e\x5a\x49\x67\x8a\x2d\x9a\x7c\x7a\x48\x7f\x95\x08\xba\x96\x84\xa6\xc9\x41\x58\x5a\xf1\xed\x22\x81\xae\xe0\x83\xdc\x7f\xb6\xe2\x4b\x77\xfd\x96\xc7\x81\x9d\x8d\x93\x45\xcb\x89\x7f\x69\x55\x14\xca\x90\x1c\x3f\xfe\x0a\x00\x00\xff\xff\x59\x5d\x63\x1f\x24\x06\x00\x00")

func scriptsDockerfilesRunShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsDockerfilesRunSh,
		"scripts/dockerfiles/run.sh",
	)
}

func scriptsDockerfilesRunSh() (*asset, error) {
	bytes, err := scriptsDockerfilesRunShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/dockerfiles/run.sh", size: 1572, mode: os.FileMode(420), modTime: time.Unix(1444350497, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"scripts/dockerfiles/Dockerfile-base": scriptsDockerfilesDockerfileBase,
	"scripts/dockerfiles/Dockerfile-role": scriptsDockerfilesDockerfileRole,
	"scripts/dockerfiles/monitrc.erb": scriptsDockerfilesMonitrcErb,
	"scripts/dockerfiles/run.sh": scriptsDockerfilesRunSh,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"scripts": &bintree{nil, map[string]*bintree{
		"dockerfiles": &bintree{nil, map[string]*bintree{
			"Dockerfile-base": &bintree{scriptsDockerfilesDockerfileBase, map[string]*bintree{
			}},
			"Dockerfile-role": &bintree{scriptsDockerfilesDockerfileRole, map[string]*bintree{
			}},
			"monitrc.erb": &bintree{scriptsDockerfilesMonitrcErb, map[string]*bintree{
			}},
			"run.sh": &bintree{scriptsDockerfilesRunSh, map[string]*bintree{
			}},
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

