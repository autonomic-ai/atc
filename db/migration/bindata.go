// Code generated by go-bindata.
// sources:
// 1510262030_initial_schema.down.sql
// 1510262030_initial_schema.up.sql
// DO NOT EDIT!

package migration

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

var __1510262030_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x5d\x6e\xdb\x30\x0c\x7e\xf7\x29\x74\x80\xde\x20\x4f\x6e\x22\x0c\x06\x9a\xb6\x6b\xb3\x0d\xdb\x0b\xa1\xd8\x8c\xab\x46\x96\x34\x91\xf2\xb6\xdb\x0f\xb5\x15\x44\x76\x95\x37\xf2\xfb\xb1\x49\x7e\xba\x97\x5f\x9a\xc7\x4d\x55\x09\xb1\x7b\x79\x7a\x16\xfb\xfa\x20\x5f\x9a\xfa\xa1\xf9\x25\x77\xe2\x7b\x23\x7f\x54\x42\x08\x61\x14\x23\x31\xb4\x6e\xf0\x06\x19\x3b\x38\x46\x6d\x3a\x02\x8f\x01\xde\xdd\xf1\x6e\x12\x59\xfc\xcb\x45\x82\x83\xb2\xa4\x59\x3b\xbb\xa2\x27\x76\x5b\xbf\x6e\xeb\x9d\xbc\x8e\xf0\x2a\xbf\x7e\x93\x8f\x5b\x39\xb1\xad\xb3\x27\xdd\xc3\x88\x81\x3e\xfc\x84\xbf\xe7\x6f\x3a\x8b\xe0\x4e\x27\xb0\x6a\xc0\xab\xf5\x50\xdf\x3f\xcc\xbe\xa3\x22\x84\x80\xe4\x62\x68\x11\xf8\x9f\x47\x9a\x8d\xd3\x04\x80\x23\x5a\x5e\x20\x7a\x50\x7d\xe6\x68\x55\xfb\xb6\xb4\x68\xeb\xe3\xd2\xe2\x22\xaf\xa0\x54\x4f\x66\xd0\x76\x54\x46\x77\x8a\x5d\xb8\xbb\xec\xc2\x4a\x5b\x0c\x49\xa6\x6d\x87\x1e\x6d\x87\x36\xdd\x6d\xf1\x8f\x77\x77\xcc\x2a\x20\x0c\x5a\x19\xe8\x83\x8b\x9e\xd6\x07\x5f\x18\xbd\xf6\x68\xb4\xc5\xac\x4d\xe5\x72\x3b\x88\x54\x26\x3e\x81\x73\x06\xed\x1b\xb6\x67\x20\xa4\x8f\x24\xca\x9a\x35\x9a\xdd\xfd\x82\xa5\x96\x51\x0d\xa9\x4c\xd9\x62\x07\x2b\xcd\xe8\x4c\x1c\x2e\xcd\x1f\x17\xce\x18\xe0\x66\xae\x89\x2f\x6e\xf2\x89\xbb\xbd\x50\x92\xb2\xa2\x73\xe1\x13\x54\x7e\xb1\x87\x9f\xcf\x32\x7b\x17\xc4\x8a\x23\xad\x32\x9f\x50\x2c\x80\x7d\x09\x64\x04\x67\xba\xfc\x0c\xb9\x3f\x47\xae\xba\x34\xfa\x84\x6e\xaa\x6a\xfb\xb4\xdf\x37\x87\x4d\xf5\x3f\x00\x00\xff\xff\x19\x32\xfa\xb2\xe1\x03\x00\x00")

func _1510262030_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1510262030_initial_schemaDownSql,
		"1510262030_initial_schema.down.sql",
	)
}

func _1510262030_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1510262030_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1510262030_initial_schema.down.sql", size: 993, mode: os.FileMode(420), modTime: time.Unix(1510611496, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1510262030_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x5f\x73\xdb\xb8\x73\xef\xfe\x14\x78\xb3\x3c\xa3\x9f\x27\x6e\xa7\xed\xd4\xb9\x64\x46\x27\x33\x89\x7a\xb6\x74\x27\xcb\x97\x4b\x3b\x37\x1c\x48\x84\x25\x3a\x14\xa9\x23\x29\xc7\xbe\x4e\xbf\x7b\x87\x24\x48\x00\xc4\xe2\x1f\x4d\x59\x99\x36\x37\xf7\x10\x13\xc0\xee\x62\x77\xb1\x58\xec\x1f\xfd\xec\x7d\x9c\x4c\xdf\x9e\x9c\x20\x34\x9e\x7b\xa3\x85\x87\x16\x5f\x7e\xf5\xd0\x72\x1f\x46\x81\x9f\xe5\x38\xdf\x67\x68\x74\x8b\xbc\xe9\xdd\x0d\x1a\x9c\x20\xfa\xdf\xe9\x8e\xc4\x41\x18\xaf\x4f\x87\xf4\xd3\x69\x96\xe3\x34\x27\x01\xfb\x80\x97\x89\xf8\x21\xdb\xaf\x56\x84\x04\xfc\xa7\x7b\x1c\x46\xfc\xdf\x24\x4d\x93\x94\x04\xa7\x27\x08\x9d\xbd\x3d\x39\x69\x13\xb5\x4a\xe2\x1c\x87\x31\x49\x0b\xc2\xd6\x44\xa2\xeb\x74\xb5\x21\xab\xaf\x0c\xdc\x9a\xe4\xec\x8f\x74\x1f\x5b\xc1\xcd\x21\xb8\x29\xc1\xb9\xb0\xdd\xf2\x0b\x4f\x7a\x40\xb2\x3c\x4d\x9e\x85\x49\x74\x7b\x76\x58\xfd\x24\x0a\x5e\x8a\x59\x81\xe9\x31\x89\xf6\x5b\xf2\xaa\x9b\xe3\x51\x1e\x70\x67\xdf\x92\xf4\xab\x5a\x6c\xe9\x3e\x8e\xdb\x4a\x1a\x09\x0a\x17\xe1\x96\x1a\x17\x1f\xf8\x09\x29\xc9\xc3\x94\x27\x80\xe1\x1f\xfd\x7c\xed\xa1\x25\xce\x88\x9f\x92\x2c\xd9\xa7\x2b\xe2\xe7\xcf\x3b\x92\x35\xf8\xc3\x00\x85\x71\x4e\xd6\x24\x45\xd3\xd9\x02\x4d\xef\xae\xaf\x6b\xb8\x31\xde\x12\x94\x93\xa7\xbc\x19\x91\x37\x78\xeb\xfd\x76\xe7\x4d\xc7\x20\x0e\x3f\x0c\xfc\x8c\xfc\x45\xa1\xdd\x2e\x46\xf3\x05\xfa\x3c\x59\x7c\x42\x17\xf4\xd3\x64\x3a\x9e\x7b\x37\xde\x74\x81\x7e\xfe\xd2\x7c\x9c\xce\xd0\xcd\x64\xfa\xfb\xe8\xfa\xce\xe3\xbe\x8c\xfe\xe0\xbf\x8c\x47\xe3\x4f\x1e\xba\xa0\xa4\x8c\xae\x17\xde\xdc\x86\x12\x34\xfb\x3c\xf5\xae\x0a\x64\xc0\x9c\xf3\x30\x68\xc9\xae\xe2\x5d\x69\x63\xc8\x23\x89\x73\xc6\xb4\xea\x23\x63\x5d\xcd\xb1\x02\x10\x5a\x6d\x70\x8a\x57\x39\x49\xd1\x23\x4e\x0b\xbd\x18\xfc\xf3\x3f\x9d\x49\xcc\xdd\xe1\xe7\x28\xc1\x81\xc8\xdf\x7a\xb0\x44\xe7\x6b\x44\xf3\x48\xd2\x2c\x4c\x62\x93\x74\xf8\x2d\x84\x5b\xbc\xe6\xf6\xbc\xc2\xab\x0d\xa7\x06\xe2\x77\x60\x6b\xed\x2d\xdb\x62\x8d\x77\x7b\x0b\xc6\xd1\xfd\x90\x80\x11\x28\x4f\x92\xf5\xb1\x1e\xd9\x26\x41\x78\x1f\x92\xc0\xcf\xc3\x62\x4a\xb8\x25\x59\x8e\xb7\x3b\xf4\x2d\xcc\x37\xc9\x3e\x2f\xbf\xa0\xbf\x93\x98\xa0\x2b\xef\xc3\xe8\xee\x7a\x81\xe2\xe4\xdb\xe0\xcc\x72\x0f\xc9\x3e\xef\x6f\x13\xe4\x69\x17\x85\xab\x30\x47\xcb\x24\x89\x08\x8e\x1b\x8a\xee\x71\x94\x91\xd7\xdd\x59\x47\x2b\x50\x8f\xd0\x8b\x57\xb8\x85\xa5\x39\xab\x0d\x09\xf6\x11\x09\x2c\xf7\x5b\xde\xd0\xd0\x66\xd9\x4e\x1b\x4e\xc6\x20\x57\xa0\x89\xeb\x30\x86\xce\xe5\xc5\xbf\x9e\x89\x73\xfc\x2d\xc9\x71\x80\x73\x5c\x6e\xb8\x1e\x5b\x25\xdb\x5d\x44\x72\xeb\x3d\x3c\x24\x4b\x40\xf0\x29\xc1\x3b\x3b\x7a\x73\x82\xb7\xba\xc3\xbf\xc5\xf1\x1e\x47\xd1\xb3\x9f\xa7\xe1\x7a\x4d\x52\x15\x61\xf5\xfc\x02\x4e\xba\x22\xbb\x3c\x5c\x46\x44\x9a\x9a\xa7\xfb\x66\x66\x9c\xc4\x2b\x22\xec\x7d\xb7\x5f\x46\xe1\xca\xdf\x45\x38\x46\x0f\x59\xc2\x96\x9d\xfe\xf7\xff\x9c\x5e\x5e\x16\x9f\x9a\xa9\xe1\x8e\x44\x05\x13\x19\xe9\xba\xfb\xa2\x54\xc0\x63\x5d\x11\x3c\x72\xee\x56\x28\x3f\x2b\x2e\x02\x6a\x18\xe3\x47\x1c\x85\x01\xce\x93\xb4\x39\x3c\x11\xce\x72\x36\x40\x02\xab\x63\x7a\x7a\xf1\xef\xff\xf6\xe6\x1f\x6f\x2e\xfe\xf1\xe6\x02\xbd\x79\x73\x59\xfe\x7f\x7a\x79\xa9\x5b\x6a\x71\x0b\xaf\x92\xf8\x3e\x5c\xfb\xd4\x12\xbd\x0a\x6b\x45\x26\xd5\xee\x22\x33\x2d\x1b\x1c\x07\x91\xc2\x82\xa8\x2c\x69\xa1\x6e\xc5\x67\x41\x13\x65\xf5\x1a\xb6\x2f\x2f\x79\x88\xfa\x5e\x8d\x11\x6b\x30\x93\x42\x66\xf7\xfe\x3e\x23\x81\xbf\x7c\xd6\x89\x6c\x68\xb6\x91\xad\x13\xcb\x1b\x48\x22\x39\xee\x8d\x06\x34\xce\xe5\xe5\x65\x7b\x4e\x1b\xc1\x26\x7c\xc0\xab\xaf\xd6\x16\x28\x08\xb3\x02\x62\x18\xef\xad\x97\x14\x96\xaf\x74\x84\x2a\x49\x35\x44\x16\x3a\x09\x5e\xb9\xc5\xfc\x2c\x27\x3b\xc6\x5b\xcb\x45\x38\xcf\xc9\x76\x97\xbb\x2c\x29\xa4\x18\xc6\x6b\x3f\x08\x53\xb2\xca\x93\xf4\xd9\x65\xf1\x2e\x4d\x56\x24\xcb\x0a\x51\xa7\x4e\xeb\x64\x85\x6b\x56\xbe\x81\x97\x88\x96\xdf\x34\x5b\x72\xa9\x0c\xf3\x1b\x82\x5c\x39\x5e\x90\xe5\xba\xa6\x22\xce\x61\x55\xe5\x5f\x96\x8f\x5b\x9f\xa9\xb3\x7c\x28\xaa\x79\x6b\x92\x6b\x67\xd1\x73\xcb\xdc\xd2\xca\xb0\x55\xe0\x33\x92\x95\xe6\xcd\xea\xa2\x61\x26\xe9\x48\x97\x8d\x44\x00\xbb\x70\xd8\x90\xe2\xd2\x09\xe3\x80\xec\x48\x1c\x14\xcf\x01\xd0\xa1\xd6\x18\xa5\x96\x32\x4a\xf2\x2a\x00\xf9\x6a\xef\xae\xbe\x43\x34\x20\xee\xc3\x34\xcb\xfd\x64\xb5\xda\xa7\x29\x29\x3c\x87\xda\xd4\x58\xdc\x54\xaa\x9d\x1d\x49\x48\x06\x72\x98\xc8\x54\x13\x15\x02\x7c\x48\x96\x4c\x58\x6a\x66\xd3\xd3\xb6\xdf\x2e\x8b\x2b\x80\xfc\x65\x61\x12\x76\xb8\xb8\xbb\x0c\x9e\x9f\x5a\x76\x90\x75\x83\xe5\x1b\x25\xeb\x75\x71\x49\xda\x5b\x2b\xca\xba\x80\xe4\x24\xdd\x86\xb1\xfd\x1d\x84\x9f\xfc\x30\xf6\xef\xa3\x70\xbd\xc9\xfd\x94\x14\x1e\x97\xed\xda\xca\x40\xc0\xdc\xc5\xab\x3c\x7c\x94\x3d\x5f\x18\x50\xe9\x2d\xa7\x7b\x85\xb7\x0c\xaf\x61\x7e\xb3\x46\xe3\x0b\x55\x38\x92\x76\x73\xa8\x99\x26\x17\x1f\x35\x5a\xeb\x67\x24\x0d\x71\xe4\xaf\xd3\x64\xbf\xb3\x32\x38\xfc\x02\x58\x10\xa2\x49\x32\xf1\x4a\x20\xe0\x98\x9c\x83\x08\x11\xf9\x28\x4e\x91\xb8\x7a\x33\x5a\x78\xf3\xc9\xe8\x7a\xf2\x9f\xde\x15\xfa\x7d\xe2\x7d\x46\x11\xce\x0b\x0f\xb4\x79\x5a\xfa\xf4\x49\xb2\x23\x69\x71\x5d\xa3\xd1\x6d\x41\x5d\xb9\x41\x3a\xb5\x3e\x7f\xfc\x1c\x2e\xd8\x5d\x30\xc5\xbb\xf6\xc6\x8b\xe2\x14\x0d\x96\xfe\xc5\x79\x18\x9c\x15\x53\xea\x75\xfc\x4c\x84\x3e\xcc\x67\x37\x68\x40\xa3\x00\x4b\xff\x42\x1c\x45\xe8\x3f\x66\x93\x69\x65\xbc\x1e\xd0\x6c\x8a\x06\x83\x87\xf3\x30\x40\xef\x8a\xa9\xe7\x95\x14\xcf\xce\xce\x84\x45\x9f\x3f\x79\x73\x0f\x95\x98\x69\x2c\xe0\xa7\xf7\x68\x74\x7d\x8d\x06\xa3\xf9\x7c\xf4\xe5\xbf\x9a\x68\xfc\xe5\x25\x1f\x31\x18\xb2\xa8\xbc\x38\xf0\x67\x0b\xfe\xc7\xf9\xec\xee\xd7\xf2\x95\xd6\x90\xc0\x8d\x97\x73\x29\x03\x96\xe7\x61\xd0\x18\xd6\xf3\xc2\xe8\xb2\xbf\x28\x52\xf6\x77\x1d\x9f\x10\xa6\xd0\x10\x04\xfb\x56\xc7\x1a\xf8\x2f\xeb\x30\x96\xfe\x6e\x02\x08\x6c\xa0\x11\x31\xfb\x54\x51\xcf\xfe\x6e\x42\x03\xec\x13\x7d\x51\xb0\x0f\xf2\xa3\x9f\x8d\x09\x0f\x7c\x6e\xe7\x85\x55\x62\x7f\x72\x4f\x79\xee\x23\xbb\x06\x4e\x00\xc5\x38\xe1\xd5\x41\xa9\x88\x51\xa5\x22\xd1\x79\x73\x47\xbc\x2b\x85\x50\xab\x48\xa9\xc6\xd3\x19\xba\x1a\x2d\x46\x65\x12\x67\xee\x7d\x98\x7b\xb7\x9f\xdc\xcf\x05\x10\xdc\x8e\xc9\xd3\xff\x35\xdf\x48\xda\xd2\x91\x8c\x9f\x8a\x0e\x66\xfb\xa4\x19\x16\xa6\x8f\xad\xe9\xc7\xdc\x85\xf1\xf7\x61\xee\xde\xa1\xd1\xf4\xcb\x0f\x6b\xf7\xc3\xda\xd9\x5b\x3b\xe0\x28\xbc\x05\x2d\x42\x12\x13\x3f\xb9\xbf\x2f\x4d\xd2\xc1\xad\x80\x60\x5e\x6b\x96\x31\xb3\x5a\xe7\x9e\x2c\xde\x01\x5d\xb2\x0a\x36\xef\x9a\x24\x0d\x48\x1a\xc6\x6b\x0b\x1a\xca\xe0\x2c\xcb\x42\x1c\x38\x34\x6b\x1b\xc2\xaf\xb4\xd3\xf2\x75\x41\x3d\xf0\x87\x2c\x89\x35\x97\x46\x23\xa8\x23\x5d\x16\x6d\xfc\xec\x92\x68\x46\x14\xaf\x8d\x62\x5c\xb8\xb5\x41\xbd\xd8\xa7\x91\x10\xc1\x55\xb1\x58\x95\xf1\x6a\x65\x38\xf7\xd9\x4b\xb2\x9f\xdc\xab\x53\x0a\x61\xd9\x11\x60\xe5\xa6\xb4\xc3\x5f\xca\x9c\xa3\xea\x28\xa5\x78\x9b\xf9\x1b\x9c\x6d\x14\x79\x53\x3e\xdd\xa5\xd1\xad\x16\xe9\x47\xd2\x30\x98\x0a\xa6\x67\xad\x71\x85\xb6\x69\x43\x8a\x3d\x49\x85\x3c\xed\xc2\x94\x64\x3e\xce\x95\xc9\x3e\x2b\x6e\x43\x14\x1e\x9d\xf9\x1a\xa2\x20\x59\x40\xd3\xed\x44\x63\x25\x0c\xb9\x92\x03\x4a\xc4\x54\xa3\xea\x83\xa0\x34\x00\xf6\x62\xfa\x4e\x04\x63\x14\x85\x91\xf9\xd6\xf5\x41\x9a\x8c\x9c\xfa\x7a\x67\x29\x26\x5d\x4d\x8b\x70\x7b\x97\xba\xf3\x7a\x77\x77\x6f\x91\x44\x39\x9b\xae\xb6\x1c\x36\x7a\x76\xcc\xa2\x2a\x43\x3d\x95\x55\x29\x55\x3d\xc9\x26\x20\x5e\xd9\x8b\xb2\xd4\x52\xcc\x02\x1f\x3e\xea\xfd\x43\xe7\x38\x51\x1d\x59\xdd\xd4\x9a\xa6\x52\xb2\xc2\x2d\x7c\x61\x5d\xd3\x12\x67\xe1\xca\xc7\xfb\x7c\x83\xf8\xfa\x16\x1c\x6c\xc3\x58\xaf\x79\xe5\x1a\x9e\xf9\xb2\x38\xc6\xb3\xe9\xed\x62\x3e\x9a\x4c\x17\x85\xd0\xb3\x3c\xc5\x61\x9c\xfb\x25\xd5\xe5\xeb\xce\x8f\x93\xdc\x27\xdb\x5d\xfe\x8c\xc6\x9f\xbc\xf1\x2f\xc5\x4b\x93\xc4\xeb\x7c\x33\x28\x46\xcf\xd0\x7b\xf4\xa6\x7c\x67\xaa\x24\x57\x41\x3a\x8e\xd4\x78\xdc\x4c\x62\xe5\x57\x8b\xb8\x50\x9e\xe2\x38\x0b\xf3\x30\x89\x35\xd1\x21\x3a\xb2\x24\xf7\x49\x4a\x7c\xb6\x44\x15\x1d\x62\xe1\x93\x61\x2b\xe6\xd3\x0a\x93\x6f\xf1\x13\x14\x32\xd2\xc4\x8c\xae\xbd\x0f\x0b\x39\x70\xc4\x10\xa2\x77\xe8\xe1\x5c\x0e\x1c\x09\x2b\x4d\x69\x80\x0c\x00\x9a\xe9\x23\x52\xad\x08\x7c\x46\xff\x7d\x86\x46\xd3\xab\x7e\xc3\xf3\x5d\x23\x56\x57\x93\xdb\xc5\x64\x3a\x5e\x94\x7b\xab\xe3\x43\x67\x3f\x02\x59\xf5\xce\xfb\x0b\x64\x31\x45\x53\x9e\x9b\x4a\xbf\x98\x76\xa9\x26\xb6\x94\xae\x56\xb6\x81\x72\xfe\x16\x3f\xa1\xc9\x6d\x69\x58\x6b\xdd\xeb\x4f\xf3\xd0\x6c\x5e\xc0\x0b\x03\xf4\x5e\x4d\xf1\x16\x3f\x51\x6a\x67\xf3\x2b\x6f\x5e\xea\x66\x2d\xc4\x52\xdb\xa0\x70\x9d\x36\x5e\xa7\x34\x51\x40\x5e\x42\xae\x06\xb6\xba\x96\xb4\x0f\x7a\xe1\xbd\xee\xe0\x5a\x93\x18\x2f\xa1\x1a\xdc\x3c\xdd\xcb\x2e\x84\xa6\x9a\xee\xa5\xd5\xc8\xa2\x77\x57\x46\xf1\x34\x21\x3c\xcd\x2d\x07\x30\xf7\x48\x77\x9e\x9a\x12\x76\x03\x02\x73\x14\xde\x4b\xd5\x91\x62\xa5\x28\xba\xba\xca\x46\x84\xd0\x9b\xaa\x19\x6c\x9e\xc3\xad\xa2\x2b\xb5\x67\x94\xa4\xe1\x3a\x8c\x71\xe4\xd3\xce\x19\x8e\x86\x66\x4a\x59\x2c\x2f\x17\x5c\xee\x70\xcb\x35\xda\x24\x59\xee\x17\x5f\x15\x44\xee\xa2\x70\x85\x8b\x2b\xf1\x3e\x4d\xb6\xad\x7a\x6c\x75\xdd\x98\xb6\x12\x53\x68\x31\x82\xca\x30\x85\x09\x40\x34\x4d\xec\xca\x68\x0d\xc8\x38\x5a\x7c\xb5\x0a\x55\x48\x85\x6f\x8a\x88\x24\x9d\x97\xe3\xec\xab\x72\x0e\xef\x6b\xe2\xb8\x70\x2c\x59\x95\xb2\x1f\xec\xd3\x30\x5e\xfb\x61\x1c\xe6\x21\x8e\xc2\xbf\x71\x79\x1d\xd4\x3e\xe7\x60\x50\x6d\x47\x4c\x6a\xd5\xcd\x4f\x22\xa7\x86\x42\x13\x94\x34\x46\x7b\xb0\xc4\xef\x7f\x9e\xd1\x4b\x61\xa0\xdc\xb0\x78\x79\x68\x79\x08\x4e\x15\x59\x23\x4e\x11\x54\xa8\x1e\xa2\xd7\x8a\x9e\x22\xaa\x14\xd5\x54\x23\x51\xd0\x6c\x99\x2e\x61\x96\x44\x5a\x3d\x7a\xa6\xf5\xfb\xa9\xe5\x38\x96\x15\x14\xb0\x73\x96\xaf\xfa\xae\xb0\x76\x4a\xf6\x31\x03\xa8\xac\xdf\xb6\x39\x4b\x65\xa9\xa9\x39\xd8\x64\x93\xb7\xd2\x70\x5e\xbd\x8b\x23\x09\xc3\x48\x10\x93\x8f\x7a\xaa\x5e\x64\x1d\x72\x19\x2e\x46\xb0\x4b\x38\x16\xa6\xec\xb8\x22\x30\x65\x2c\xe0\x69\xb6\xac\xef\x9a\xbf\xe8\x26\x09\x7d\x01\xb6\xe2\xfe\x75\x91\xd7\x77\x94\xec\x70\xa1\x4d\x23\x4d\x87\xd4\x87\x74\x37\xb8\x48\x52\xb2\x8e\x70\x07\x5a\xab\x5b\x43\xf6\x6f\xf2\x8d\x75\xbb\xaf\x4c\xef\x71\x45\x25\x13\x22\xc9\x85\x9b\xa2\x15\x02\xe3\x3c\x0e\x02\x31\xfa\x4c\xf3\x7a\xc6\x0e\xbe\x2a\x64\xeb\x73\xdd\x50\xd2\x2b\x47\x3a\x60\xd5\xbd\x27\xb8\xcc\x11\xce\xef\x93\x54\x74\x7e\x73\xbc\xce\x5a\x37\xe1\x7a\x8d\xd7\x64\x15\xe1\x70\xeb\xd7\x49\x72\xa9\x5d\xc3\x1c\xf5\xdc\xe4\xf9\xce\xdf\xa5\xc9\xd3\xb3\xdf\x4e\xb5\x17\x43\x99\x62\x2c\x4e\xaa\x01\xe1\x23\xd7\xc5\xe9\xe2\x9f\x0b\x8d\xf2\xcd\x0e\xea\x06\xf9\xcb\x4b\x61\xdc\x26\x7d\xc4\x39\xc0\x85\x2c\xfd\x6f\x1b\x12\xfb\x14\x9e\xe4\xea\xfe\xf4\x9e\xf5\xde\x8b\xb8\xa8\xd3\xc8\xa6\xd1\x06\x7c\x70\xd6\xa0\xd4\x1a\xc9\xad\x93\xc4\x24\xba\x76\x95\xe7\xd9\x38\xdd\x4a\x42\x8a\x69\xcd\x24\x98\x0c\x4a\x47\x43\x06\x17\x80\x81\x68\x90\x5d\xcb\xea\x84\x55\x47\x62\x36\xbd\x06\xbb\xe5\xe9\xa4\xf1\xec\xfa\xee\x66\x5a\x98\xa8\x5b\x6f\xc1\x1e\xfd\xe4\x29\x7f\xc4\xd1\xe0\x54\xed\x82\x9c\x5e\x5e\xa6\x64\xbd\x8a\x70\x96\xa9\xd1\x56\xd1\x2c\x3b\x4c\x7c\x47\xa7\x0d\x70\xee\x6c\x5a\x21\x90\xba\x78\x6c\x90\x28\xdb\x77\xac\x50\x1a\x7a\x52\x6c\x08\x28\x23\xd3\x56\xc8\xb8\x16\x01\x5b\xc0\xad\x0e\x01\x7b\x34\x50\x3d\xbd\x0d\x52\xb9\xd4\xd7\x0a\xa7\xaa\x8c\xd5\x06\x25\x2b\x7f\xb3\x42\xd5\x2e\x82\xb2\x41\xd1\xf6\xa5\xad\x10\xc1\x9e\xa5\x1b\x3a\xd0\x7f\x74\xc4\xae\x71\x8c\x3a\x10\xd3\x0d\x7d\x27\x84\x0e\x56\xac\xb3\x01\x63\x81\x5f\x27\x3c\x2e\x28\xaa\x74\xa7\x15\x78\x3e\x3b\x67\x03\x1a\x0a\x60\x5b\x21\x52\x87\x44\xad\xd0\xd2\x10\xa8\x1d\x2a\x21\xee\x60\x03\x5e\x13\x73\xb0\xc2\x68\x7c\x5c\x3b\x10\xd1\xe9\xe4\x6b\x9f\x96\x9d\x90\x77\xb7\x03\x2e\xef\x24\x07\xd2\xf8\x57\x90\x0b\x1d\xf2\x23\xc0\xca\xcd\x90\x45\x49\xdd\xc7\xd1\xd5\x15\xef\x42\x42\x32\x2f\x53\xf6\x5f\xc9\x33\xba\x9b\x4e\x7e\xbb\xf3\x50\x95\xa5\x3f\x08\xaa\x5d\x81\xe6\xd7\xf9\xe4\x66\x34\xff\x82\x7e\xf1\xbe\xa0\x41\x18\x18\x7c\x27\x05\x70\x9a\xba\x72\x81\x47\x9f\x46\x30\x40\xae\x90\x81\xce\xab\xf8\xb2\x8f\xc3\xbf\xf6\xc4\x8e\x35\xcc\xbf\x52\xe2\xa8\xfd\xaf\x2a\xeb\x20\x70\xbd\xfa\xd4\x07\x70\x27\xae\xa8\x3c\x34\x18\x8b\xd2\x9f\x3b\x06\xce\x4a\x36\xf4\x17\x0b\x7c\xae\xe1\xa9\x66\x69\x9d\x28\x65\x43\x5a\x67\x10\x46\x5f\xba\x7b\x4e\xdb\x93\x1d\x44\x0d\x64\xd1\x91\x74\xc6\xa3\x81\x4c\xd9\xc3\xe5\xd9\x45\xee\x70\x03\x43\xa4\x65\x8e\xe4\x7d\xc2\x58\x65\x27\xd5\x69\x3b\x9d\xb1\xf4\xa9\x07\x8d\xf7\x0b\x23\x67\xce\x71\x69\x1d\xea\x50\x00\x6f\x1e\x86\x75\x80\xe0\xe5\x38\x9c\xd8\x57\xb6\x20\xa8\x21\x3a\x42\x6b\xdd\xcd\x30\xdc\xf6\x05\x7e\x78\x0c\x72\x25\x60\xf3\x43\x41\x5c\xa7\x80\x60\x57\xe5\x25\xc3\xda\x2f\x1c\xf2\xed\x05\x1d\x5f\x1b\x26\xba\x41\x77\xa2\x23\xa3\xaa\xa7\x82\x15\xc6\xd7\xc1\x21\x25\x39\xe0\x88\xbc\x9f\x25\x0a\x91\xd0\x65\x43\x38\x19\x36\xe4\xab\xde\x2d\xdf\x43\x06\xca\xa9\x1f\xd2\xb2\x89\x02\x75\xf6\x76\xb1\x0b\xe6\x2e\x52\x31\x80\x76\x84\x5a\x3e\xa2\x60\x88\xd5\xfb\xaa\x4f\x1a\xfb\xbb\x84\x80\x27\x19\x8c\x12\x7a\xbb\x39\x6d\x89\xbe\xc8\x14\xd0\xb9\xe7\x5a\xf9\xa3\xab\xbc\xea\x94\x2a\x5b\x46\x30\x5f\x04\xbb\x7f\x6a\xb9\x2c\x0b\xe4\x7b\x72\xc3\x43\x64\x70\x44\xd5\xaf\x47\x98\x04\xcd\x6b\xb3\x83\xfb\xde\x0b\x5e\x9e\x1b\x90\xe1\x51\xf3\x06\x34\x53\x8e\x6f\x64\x2d\xb9\x2f\xba\x52\xad\xde\xb2\x96\xe8\x5f\x7e\x6f\xc9\x2f\x5a\x2d\x6a\xfe\xe5\xdb\xdf\xbb\xae\x7e\xcc\x95\xb9\x13\x5e\xac\xc5\x87\x97\x81\x94\x89\xe4\xed\x17\xcd\xca\x51\x74\x93\xe9\x95\xf7\x87\xf0\x43\xbc\x4d\x57\xb6\xdf\xfc\x52\xee\x6c\x2a\xfe\x54\xef\xdd\xed\x64\xfa\x11\x2d\xf3\x94\x10\x5a\x17\x5b\x5a\xc9\x7a\x7e\x0b\x93\x0e\x45\x18\x3c\xd9\x41\xd7\xc0\x04\x7f\x7c\x97\xfd\xa4\x52\x03\x1e\xfe\x91\xde\x7e\xb0\xc9\x45\x15\x2e\x68\xa5\xd5\x3a\xfc\xd5\xd3\x02\xe6\x20\x8d\xdc\x3b\xef\xa9\x05\x13\xfe\xb1\x5d\x2b\x2c\x43\xf8\x97\x7a\xcd\xc8\xc1\x65\xe6\xdd\xb9\x62\xa3\xbf\x37\xac\xe0\x5f\xfd\x6b\xc4\xae\x0c\x94\xa0\x1a\x38\xa8\xc5\xe3\xce\xc2\x1a\xbd\x05\x0f\x41\xcc\x8e\xf8\x32\x5f\xa8\xae\x67\x1d\x16\x0d\x9a\x16\x7c\xb1\x18\x9f\xfd\xe6\xaf\x0e\x05\x2d\xb3\x50\x40\xa4\x05\xf3\x9a\xf5\x7c\x37\x9a\x02\x08\x37\x45\x07\x89\x16\xd6\x2b\x80\xd0\x0e\x10\xcd\xfa\xfa\x2d\xae\x00\xd0\x7a\x94\x0b\x10\xb8\xd8\x19\x6f\xcf\xb8\xfc\xaa\xbd\xa6\xf2\x49\x56\xd5\xcf\x49\x6a\x60\xab\xd6\xd8\xe1\x92\x7e\x92\xd2\x88\xa9\xbd\xc2\x84\xa7\xfe\x59\x5b\x0d\x60\x3a\xc5\x04\x89\x93\x97\x0a\x92\xa5\xcc\xf8\x12\x22\x0d\x34\x6e\x9a\x25\x44\x63\xe5\x98\x19\x9b\x09\x04\x48\x89\x32\xe2\xc9\x4e\xab\x32\x2d\x6f\x7b\x7e\x95\x38\xb8\x1f\x86\xb2\xc6\xc3\xd6\xe8\xdc\x9f\x02\xd6\x13\xdf\x27\x48\x1f\x87\xab\xc2\xa9\x0e\xe3\x8c\x94\x8d\x29\x8f\xa5\x10\xab\xa4\xa0\x80\x24\x4a\xbe\x91\xb4\xf2\xb1\xc0\x1d\x55\x81\x5a\xd1\x1e\x95\xc5\x03\x4e\xd6\x08\x08\xca\xd6\x81\xc5\xca\xc0\x03\x65\x03\x66\x9e\x0b\x7c\x30\x74\xcc\x51\xd2\x4d\x7d\x75\xe2\x61\x86\x37\x23\xc7\x4a\x19\x7c\xb9\x12\xc1\x56\x75\x64\xa8\xa2\xce\x18\x20\xdb\x29\x0b\xf0\xbb\x42\x12\x78\x37\x66\xb0\xb8\x2a\x67\x7a\x58\x71\x84\xb5\xe5\xa9\xa2\xa9\x2d\x18\x0e\xeb\x5b\x6e\xe8\x3e\x6b\xb9\xd1\xc0\xb8\xc3\xfd\x03\x41\x6f\x5f\x09\x46\x0c\xc6\x1b\x01\xc2\x02\x3a\xe7\x46\x54\x76\x4e\xb9\x39\x0e\x2c\x23\x53\x22\xaa\x97\x18\x30\x81\x2f\x60\x13\x62\x30\x11\xdd\x03\x1d\x99\xa2\xd6\x59\x46\xde\x56\x16\x5d\xc4\x42\x8f\xd1\x20\x50\x08\x9b\xa3\x34\xa5\x68\xac\x80\xa1\x2a\x69\x70\xb2\xde\x2d\xc8\x06\x69\x01\x08\x1c\xa5\xa3\x24\xbe\x23\xdd\x26\x92\x5d\xa9\x15\x4c\xaa\xb2\xf5\x93\x22\x51\x77\xaf\xdb\x98\x57\x28\xec\x0a\xbc\x8f\xa0\x92\x1c\x78\x4f\xfa\xcd\x98\xd0\x95\x9a\x5e\x17\xd4\x3a\xe2\x1d\x96\x7d\xa8\x43\xb4\x0d\xfe\xa5\xbe\xad\x60\x97\xa3\x0e\xaf\xb6\x0d\x6c\x5d\x00\x64\x6b\x54\xc5\x9d\x51\xa0\xb4\x33\x51\x05\x4e\x0c\xd3\x82\x64\xb1\x76\x3f\x15\x90\x66\x86\x16\x0e\x77\xd7\x81\x50\x74\xb7\x5d\x2b\x02\xad\xb4\x62\x20\x60\xdd\x22\x1b\x6c\xf5\x93\x40\x07\x5c\xf9\x1e\x68\xc1\x02\xad\xa1\x0e\xb0\xc9\x16\x82\x42\x87\xd7\xd2\x2a\x94\x17\x22\x04\x31\x89\xad\x7c\x06\x0c\xc2\x64\x10\xba\x26\xea\xae\x14\xbd\xa6\x9e\xad\xe3\x55\x66\x17\xfb\x77\xc1\x6d\x52\x16\x45\xf8\x1e\x54\x1a\x45\xe9\x5c\x87\x8b\x54\x81\xd5\x74\xd6\x6c\x08\x70\x3e\x7a\x76\x19\x04\xe5\x30\x48\x97\xbb\x37\x65\xf5\xbe\xb6\x23\x55\xc9\x81\x8e\xa4\x76\xe5\x28\x9f\x18\x61\x97\x35\x50\x67\x68\xfb\x7e\x03\xe0\xc2\xa7\x42\x09\xdc\xee\x34\x08\x17\x47\xdd\x98\xa4\xbb\x38\xe0\x12\x40\x38\xb4\xaf\xa9\x0b\x34\xa4\x2a\xfc\xfb\xaf\xe4\x19\x7d\x98\xcd\xbd\xc9\xc7\x69\x95\xb8\x69\x5e\x54\x68\xee\x7d\xf0\xe6\xde\x74\xec\xd1\x9f\x61\xce\x0a\x77\xa7\xa0\xff\xca\xbb\xf6\x16\x1e\x1a\x8f\x6e\xc7\xa3\x2b\xef\x55\x08\x96\xab\x2b\x64\xca\x65\x3b\xc1\x6f\xa1\x05\xb1\xb5\x97\xb9\x77\xbb\x98\x4f\xc6\x0b\xc3\x66\x34\xb5\x58\x70\x5e\xe3\xf0\xfc\xb5\x25\x09\x8e\xd4\xcb\xf4\xc1\x91\x79\x9e\x58\xc0\x7d\xec\x42\x39\xcd\x0b\xe8\x48\x97\x92\x1c\x07\x67\xa7\x03\x51\xdf\x1b\x43\x0d\xe5\xc1\x5c\xd5\x89\x4c\x23\xff\x26\xe3\x29\x6b\x42\x42\xbd\xd3\x43\xcd\x1d\x40\x4b\x6d\x08\x79\x3a\xca\x10\xa8\x1b\x0d\x0e\x45\xc2\xbd\xa8\xd7\xad\x57\x75\xff\xf5\x40\x8f\x2a\xf1\x01\xd0\xa7\xcc\x91\xf0\xf4\x32\xd0\x87\xa6\xb9\x9d\x42\x51\x52\x2c\xe5\x5a\x5e\x9b\xde\x97\xeb\x5f\x8f\xc4\xf0\x45\x38\x32\x41\xbc\x87\xc1\x13\x45\xdd\x08\xfa\x4b\x82\x87\x38\x18\x96\x89\x1b\x35\xcd\x46\x97\x54\xde\x8f\xc1\x8f\x74\x14\x82\xb6\x22\xac\x20\x5b\x8c\x57\x08\x7b\xe8\x5f\x3f\x6d\xab\xe1\x4b\xc2\xa8\x7f\x2b\x90\x54\xff\xb8\x1e\x47\x4c\x01\xb3\x3f\xeb\x7c\x58\xcc\xd6\x75\x8b\x25\x19\x7c\x1d\x03\xec\xe8\x29\x5c\x3c\x47\xaa\x7a\x6a\xc1\xa0\xa9\x31\xf9\x2c\xf4\xc2\xba\x9e\x88\x64\x09\x27\xb5\xa7\xd2\xbb\x7b\x62\x6a\x28\x79\x4d\xd7\xa4\x73\x67\xc5\x61\xc5\xdb\x99\xac\xa3\x08\xd4\xba\x6b\xe3\xe0\x5e\x9e\xb1\xd3\xe3\xe0\x14\x00\x59\x3d\x9b\xfe\x0d\x31\xd7\x79\xc8\x87\xcd\x8b\x08\x34\x38\x73\xee\x57\xe4\x21\x89\x3d\x76\x90\xc0\x2a\xca\x61\x4e\xdf\x6a\xa9\x6e\x92\x5b\x30\xd9\x55\x22\xb2\x33\xdd\xbd\x75\xf5\x7c\x7f\xbb\xb2\xed\xe3\x81\xbb\x76\x80\xf3\x09\xc6\x2e\x85\xc3\x2a\x07\xd2\xbb\x1e\x86\x17\x76\x22\x1d\xe9\x18\x74\x6b\x42\x3a\xb8\x0b\xd0\x81\xbe\x57\x50\x67\xd3\xdb\xc1\xe0\x31\x83\x59\xf8\xd7\xe3\xa5\x91\xaa\xef\x80\x83\x76\xfd\x46\x75\x16\x18\xe2\x5d\x3d\x36\x14\x7e\x24\x56\xf4\xad\x2a\x38\x5c\xfb\x93\xc3\xc9\xb1\xa2\xf0\xf0\x91\x0b\x97\x36\x2a\x5b\x73\xa9\xcd\xf8\x00\x51\x00\xb3\xf1\xec\x73\x17\xaf\x13\x79\x71\xa1\xc8\xc6\x88\x2b\x73\xdd\xda\xa8\x0a\x64\xd2\xfb\xe4\xa5\x90\x18\x57\x53\x2d\xe6\xcf\x01\x8a\xb9\x8c\x9b\x23\xb5\x3d\xb6\xc7\xbd\xa6\x37\x60\xb2\x0e\x07\xea\xfa\x3b\xb8\xda\xf7\xd0\xfa\x77\x68\x97\xe6\xf0\x3b\x38\xba\xb9\x74\xdd\x62\x4f\xfd\x92\x5a\x88\x16\x2e\x80\x36\x62\xec\x14\x2a\x3e\x16\x03\xfe\xbf\x4a\xbe\x79\x0b\x1e\x3c\x04\xf3\x82\xee\xda\xc3\x46\xf5\x5e\x40\xd8\x11\x2c\xb4\xa1\xc7\xb7\x57\x31\x8e\x67\x37\x37\x93\xc5\xdb\x93\xff\x0d\x00\x00\xff\xff\xa3\x4b\x73\x92\x0a\x98\x00\x00")

func _1510262030_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1510262030_initial_schemaUpSql,
		"1510262030_initial_schema.up.sql",
	)
}

func _1510262030_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1510262030_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1510262030_initial_schema.up.sql", size: 38922, mode: os.FileMode(420), modTime: time.Unix(1510611490, 0)}
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
	"1510262030_initial_schema.down.sql": _1510262030_initial_schemaDownSql,
	"1510262030_initial_schema.up.sql": _1510262030_initial_schemaUpSql,
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
	"1510262030_initial_schema.down.sql": &bintree{_1510262030_initial_schemaDownSql, map[string]*bintree{}},
	"1510262030_initial_schema.up.sql": &bintree{_1510262030_initial_schemaUpSql, map[string]*bintree{}},
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
