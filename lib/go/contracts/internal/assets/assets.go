// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/ExampleNFT.cdc (6.724kB)
// ../../../contracts/MetadataViews.cdc (3.291kB)
// ../../../contracts/NonFungibleToken.cdc (4.832kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _examplenftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5f\x8f\xdb\xb8\x11\x7f\xf7\xa7\x98\xdb\x87\xd4\x46\x77\xad\x16\x28\xfa\x60\xac\x6f\x13\x64\xcf\xe8\x3e\xdc\x22\x48\xdc\xf6\x21\x58\xf4\x68\x71\xbc\x26\x56\x22\x05\x92\xb2\xab\xee\xf9\xbb\x17\x43\xea\x0f\x25\x51\xeb\xa4\x01\x6a\x04\x89\x24\xce\x0c\xe7\xff\xfc\xc8\x24\x09\x6c\x0f\xc2\x80\x30\xc0\x24\xe0\xbf\x59\x5e\x64\x08\x82\xfe\xce\x51\x5a\x66\x85\x92\xa0\xf6\xc0\x60\x93\xa9\x13\x3c\x2a\x79\xb3\x29\xe5\xb3\xd8\x65\x08\x5b\xf5\x82\x72\x96\x24\xf0\x60\x89\x5f\x2a\x0b\x05\xd3\x96\xc8\xed\x01\x41\xed\xf7\x22\x15\x2c\x03\x63\x99\xe4\x4c\x73\xd8\x95\x16\x84\x05\x66\x4c\x99\x23\x07\xab\x60\x87\xc4\x7f\x44\x5d\x81\x11\xb9\xc8\x98\xa6\xaf\x07\x75\x82\x9c\xc9\x0a\x1e\x37\x5b\x03\x27\x55\x66\xbc\x53\xc9\xc9\x4e\x95\x46\xd8\x97\x32\x25\xfd\x58\x26\x6c\xb5\x9c\xcd\x44\x5e\x28\x6d\x49\xc7\x46\x45\xa7\x21\xec\xb5\xca\xe1\x6a\x99\x0c\x17\x96\x29\x4f\xaf\x1a\xae\x5f\xd1\x32\xce\x2c\xfb\x87\xc0\x93\x69\x59\x7a\x5f\x3d\xfd\xac\x28\x77\x90\x2a\x69\x35\x4b\x2d\xfc\xe2\x3d\xf6\xb8\xd9\xae\xc6\x1b\xbf\xce\x66\x00\x00\xc4\x70\x74\x96\x59\x96\x7d\x29\x8b\x22\xab\x56\xf0\xf7\x07\x69\xff\xfa\x97\x8e\x00\x8f\x64\xdb\xc7\x5a\xee\x83\x14\x56\xb0\x4c\xfc\x07\xf9\x7c\x31\xa0\xf9\xa7\xb0\x07\xae\xd9\x69\x2e\x78\x23\xe6\xda\x29\xbc\x82\x0f\x9c\x6b\x34\xe6\x6e\xc8\x72\x8f\x85\x32\xc2\xf6\x38\xac\x0a\xe9\x5b\x86\x0c\x49\x8b\x2c\x43\xe7\xda\x2f\x56\x69\xf6\x8c\x9f\x98\x3d\xac\x20\x78\x99\x20\xff\x54\xee\x32\x91\x7a\xea\xee\xb9\x47\xfc\xab\x90\x16\xf5\xa4\xdc\x96\x56\xa3\x51\xa5\x4e\x11\xa2\xae\x5d\x3e\x3c\x6e\xb6\xd7\xfd\xa0\x2d\x3f\xa3\x51\xd9\x11\x35\xbc\x3a\x29\xe1\xae\x9d\xe1\xb3\xd1\x9a\x64\x39\x92\x12\x5a\xc8\xe7\xd1\x22\x47\x93\x6a\x51\x90\x71\x93\x34\xf6\x50\xe6\x3b\xc9\x44\xd6\x52\xb4\x24\x42\x0a\x3b\x6f\xdf\xdc\x97\x2e\x04\xbd\xef\xa1\x16\xfd\x95\x88\x0a\x7d\x82\xd1\xfe\xdd\xf2\x22\xf0\x05\xfd\x0c\x66\xfb\xa5\xe0\xb0\x06\xc1\xc7\x0b\xa4\x03\xac\x9d\x2a\xe3\xc5\x40\x0d\x58\x87\x4a\x8d\x49\x5b\x85\x60\xdd\x29\xd7\x92\x9d\xdd\x53\xcf\x8b\xfb\x52\xc2\x33\x5a\x17\xc6\xf9\x62\x05\x5f\xb7\x55\x81\x4f\x03\xdd\x35\xda\x52\x4b\xf8\xda\xfb\x48\x3f\x22\xbe\xed\xa7\xc2\xbd\x30\x45\xc6\xaa\x9f\xeb\xe2\x69\x7e\x4f\x81\x12\x23\x05\xb4\x4f\x1f\x12\x30\xff\x17\x1c\x05\x9e\x56\x4e\xf4\x62\x05\x1f\x64\xf5\xc5\xea\x32\xb5\x77\x43\x7f\x9e\x84\x4d\x0f\x8e\x78\xb0\x42\xbf\x94\x19\x7c\x5b\xbb\xd5\x88\x27\xb0\x34\xca\x34\x8f\x72\x40\x9b\x41\x6d\x20\xaf\x27\x09\x7b\x09\x35\x8c\xed\x34\x5b\x90\x66\x7d\xcd\xfe\xb6\xdd\x7e\xda\x88\x0c\xa7\x55\xa3\x5f\xa9\xb3\xd5\x20\x3d\x26\xe9\x17\xd1\x95\xfe\xd7\x20\x82\xd0\xf9\x4c\x8e\x12\xed\x1c\xe9\x29\xae\x07\xed\x59\x8a\x41\xff\x1e\x76\xb1\x41\x17\xa1\x0c\xe1\x75\x1b\xb5\xd4\x82\x56\xf0\x7e\xd4\x95\x1e\x37\xdb\x45\x2c\xb1\x1f\xee\x7d\x5a\xfb\xc2\x7f\x1a\x91\xec\x94\xd6\xea\xf4\xb8\xd9\x06\x1d\x7a\xb1\x82\x77\xb1\x0d\x26\x98\x3b\x43\x06\x32\xba\x05\xe2\x1e\x26\x70\xa1\x8c\x8d\x64\xee\x5c\xa3\x29\x33\x0b\xeb\x35\x79\x74\x01\xbf\xff\xde\x7c\xba\x73\xed\x83\xfa\xc7\x44\xf2\x5e\x7d\x64\x92\xa0\x80\x57\x2b\x70\x30\x68\xdc\xa3\x46\x99\xe2\xca\xcd\xf0\x87\xfb\x06\x29\xf8\xd8\x21\xef\x28\x08\x4f\x08\x99\x2a\xad\x31\xb5\x57\x83\xb8\x5f\x8e\x6f\x17\xcb\xd5\x1b\x11\xbe\x1e\x8f\x95\x4f\x5a\x1d\x05\x47\x1d\x59\xfa\x8c\x29\x8a\x63\x74\x69\x2c\x38\x3e\x98\x3a\xba\xc0\xe5\x49\x02\x5c\x78\x1c\xa3\x2b\xf2\x08\xb9\x2a\x55\x72\xaf\x74\x2e\xe4\x33\xb8\x64\x33\x21\x39\x11\x10\x5e\xeb\xec\xb5\x55\x81\x70\x12\xf6\x40\x20\xee\x37\x1f\xfb\xdf\xc8\xc1\x7b\x81\x19\xef\x65\x0c\x01\x11\x75\x92\xc8\x09\x5b\xad\xe0\xfd\xab\xa7\x8e\x8c\xd8\xc7\xcd\xf6\xdc\x9f\x64\x30\x8f\xce\x93\x56\x1c\xdc\xde\xc0\xeb\x39\xd6\x63\x93\xc4\xa9\x47\xb0\x05\x34\xe6\xea\x88\x0e\x6f\x92\x25\x0e\x6a\x79\x4c\xd7\x7a\x87\x49\x0e\x9e\x48\x58\x02\x84\x6e\x99\x65\x19\xea\x51\xf6\x37\x62\xe7\xcd\xc3\xc3\x7d\x90\xfd\xd1\x12\x1d\xd8\xe0\x46\xb8\xc3\x6b\xb7\x37\x03\x83\x96\x5e\xd7\xf9\x0b\x56\x2b\xe8\x36\x58\xc0\xdd\x1d\x14\x4c\x8a\x74\x7e\x95\x0b\x63\x28\x4c\x8f\x9b\xed\xd5\xa2\xdf\x92\x30\x17\x03\xb4\xe6\xb6\x59\x0a\xde\xe0\xb5\x76\x37\x7d\xb7\x64\x1e\x8b\x2d\xa2\x6d\xed\xf6\xc6\xb1\x4e\xb8\xb6\xee\x4b\x60\xd9\x0b\xf9\xd5\xb9\x95\x5c\xc8\x38\xef\x79\xb0\x75\xb0\x09\x52\x2e\x14\xd4\x32\xd5\xf5\x59\x33\x0a\x0e\x4c\x6b\x56\xfd\x6f\x0d\xf1\x2d\x77\xfb\x07\x66\x7e\x82\xf7\xfd\x3e\x35\x1b\xf1\x74\x5d\x8d\x50\x45\xed\xc8\x3e\x19\x59\xc0\xb9\x53\x59\xe2\xa9\x16\x5e\xdb\x10\xd4\xd8\xe9\x20\xd2\x43\x9b\x86\xee\xa8\x92\x71\x50\x12\x47\x7b\xaa\x8c\x6f\xe3\x99\xf1\x55\xf0\xa7\xd6\x80\x48\xd8\x43\xc4\x4d\xf1\x26\xb4\x7d\x39\xda\x1c\x8d\xd5\xaa\x6a\xf7\x9d\x88\xb7\x9f\x28\x75\x6e\xb8\x42\x72\xe1\x69\xda\x29\xad\xd9\x03\xb3\xc0\x34\x8d\xba\x41\xec\xbf\x61\x3e\xc5\x81\xd7\xa0\x34\x5e\xb0\x32\x13\xfa\xb5\xe3\x8c\x64\xfb\x46\xd5\xf4\x75\xab\x9a\xba\x9f\x56\x2c\x49\xc0\x28\x6f\x41\x57\xf8\x90\x32\x82\x68\x8c\x83\xb0\x06\xf2\xba\xbf\xba\x8c\x25\x82\xe6\xeb\x41\x71\xf3\x43\xe3\x35\x6e\xfb\xbb\x48\xf4\x99\xb9\x30\xa0\xcf\xb3\x31\xca\xfd\xa1\x61\x2d\xf6\xb1\x2c\xfc\xc9\xcd\xe8\xc8\x10\x4f\x12\xf8\xa8\x91\x59\x74\x19\x52\xda\x83\xd2\x74\xa0\x1c\x44\x23\xa3\x73\x3d\x57\x27\x99\x32\x63\xc3\xe3\x4d\x58\x08\x1a\xf7\xb0\x9e\xf2\x02\x89\xbe\xe0\x8a\x81\x3b\x49\x1c\x15\xfd\xc0\xde\xef\xc6\x77\x13\xee\xa5\xa9\xdb\xcc\xdc\x81\x83\x3f\xc8\xea\x73\x3d\x35\x5f\xe3\x43\xfa\x1c\xe9\x57\x72\x6f\x7f\xd8\x7c\x92\x83\x1d\x18\x5a\x3b\xa1\x97\x9c\x50\x5b\x1d\xf0\x51\xd6\x7d\x83\x11\x31\x27\xd5\xdd\x65\x34\xc4\x9b\xae\xd3\x37\x2f\x0e\xb3\x92\x84\x7c\x4d\xf0\xb8\xb9\x7d\xa9\x5b\x8d\xac\x94\x44\x57\xa4\xae\x1c\xad\x82\xb4\xce\x3d\xd7\x8b\x31\x2f\x6c\x35\x2c\xf6\x26\x6a\x9e\xf2\x17\x22\xe9\x20\xd2\x3c\x3a\xbe\xa3\x10\xaa\x1d\x92\xcd\x9e\xa1\x94\x81\xf6\x9f\x5b\xcc\xe4\xd5\x06\xc6\x73\x21\x41\x69\x30\x8a\xfa\x07\xcd\xf2\xe6\x2a\xca\xdf\x3c\xa9\x93\xac\xaf\xaa\x6a\x11\x6c\x97\xb9\xd2\xc9\x85\xb4\xce\xb8\xd6\x5d\x49\x12\xbd\xbf\xf0\x77\x1e\xcd\x75\x50\x2d\x85\xb8\x29\xa0\xf4\xaf\xa9\xbd\x44\xef\x1e\xc6\xb9\xd7\x87\xfb\xe1\x70\x6e\x26\x3d\xfd\x91\x35\x76\x4e\x45\x21\x90\x64\x04\x00\xaa\x74\x98\xc4\x1e\x50\xe8\xf0\x73\x5b\xf9\xa3\xc2\xa9\xb5\x99\x0f\xb2\xaf\x96\xbd\x82\x77\xaf\x17\x51\xef\xf9\xff\x75\xa1\x31\x9c\xf9\xbd\x44\x1b\x56\x10\xe1\x5d\x89\x6e\x18\x75\xf9\x31\xb2\x14\xea\x9b\x99\xa0\x14\x83\xdb\xba\xf1\x99\xd8\x1b\x17\x3f\x66\xf7\xcc\x7b\xf3\x64\x1d\xd8\xd9\x3e\xf6\xa9\x16\x23\x63\xdf\xc8\x80\x3f\x18\x60\x69\xaa\x4a\x69\x7b\xf1\x1f\x07\x1d\xc2\xd8\x2e\x07\x20\xee\xf6\xc6\xbb\x6b\xb0\x75\xdc\x33\xb0\x9e\x5a\xf8\x63\xdd\x74\xe7\x7f\x5e\xc4\x3b\x89\xbb\x1a\x5b\xf4\x0f\x42\xdd\xb5\xa7\xb3\xcc\xc9\x03\xe3\x04\xb6\x64\xfe\x0a\xa1\xa7\xc2\x9f\x7a\x85\xf5\x05\x3d\x6e\xa0\xe8\x70\x28\x98\x3d\x98\x3e\x73\xf4\x7a\x13\xd6\x90\x18\xff\x9a\x60\xe4\xc8\x38\x25\xa2\xbb\xe6\x24\x09\xbe\x37\x7e\x83\x80\xd1\x35\x68\x7c\x7f\x4f\xd6\x33\xaf\x19\xe9\x41\x93\xeb\x9a\x0d\xf5\x08\xc3\x8e\x58\xa3\xfe\x5a\x60\xcb\x4e\x33\x28\x68\x07\x6f\x34\xcc\x56\xd1\x3a\xa3\x96\x24\x75\x7e\x7b\xd3\x71\x07\x98\x36\xea\xd0\x45\x4f\xeb\xb6\x46\xeb\xe9\x91\xb2\x82\xed\x44\x26\x6c\x05\x7b\xa5\xa7\x90\x60\x4f\x83\x4c\xc8\x97\xdb\x70\x58\x76\xdb\x5e\x6e\x4e\xd7\x61\x9e\x4e\x5f\x09\x9c\x7f\x9e\x8f\xcf\xb7\xb1\x60\x0f\x1a\x16\xd3\xcf\x68\xdf\xf2\xc6\x2c\x52\xd1\x61\x30\xeb\x11\xf1\x3d\x81\xcc\x3d\x4b\xaf\xab\x79\x31\x17\x62\xe8\x19\x83\xf8\x8d\x92\x31\x50\xd2\x9d\x66\xa6\xff\x5b\xe2\x3c\x3b\xcf\xfe\x1b\x00\x00\xff\xff\xb4\x8f\xc8\x7f\x44\x1a\x00\x00"

func examplenftCdcBytes() ([]byte, error) {
	return bindataRead(
		_examplenftCdc,
		"ExampleNFT.cdc",
	)
}

func examplenftCdc() (*asset, error) {
	bytes, err := examplenftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleNFT.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0xb5, 0x37, 0xb8, 0x62, 0xcf, 0x11, 0x47, 0x9a, 0x9, 0x50, 0xae, 0xc6, 0x8b, 0x96, 0x33, 0x51, 0xeb, 0x33, 0xc5, 0x3d, 0xe5, 0xea, 0x37, 0x66, 0x99, 0x6f, 0x40, 0xa8, 0x38, 0x3b, 0xb0}}
	return a, nil
}

var _metadataviewsCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\xdd\x6e\xdc\x36\x13\xbd\xd7\x53\x0c\xfc\x01\x1f\x76\x03\xaf\x94\xa4\x85\x81\x2e\x60\x04\x41\x1c\xb7\x5b\x24\x85\x61\x3b\xbd\x29\x82\x82\x22\x47\x2b\xb6\x14\x29\x90\x23\xbb\x0b\xc3\xef\x5e\x0c\xf5\xbf\x3f\x4e\xdb\xbd\x58\x48\xe2\xf0\xcc\xcc\x99\xc3\x19\x66\xaf\x5e\x25\xc9\x7d\xa9\x03\x48\x67\xc9\x0b\x49\xa0\xab\xda\x60\x85\x96\x02\x50\x89\x50\x21\x09\x25\x48\x40\x20\x61\x95\xf0\x0a\x6a\xef\x6a\x17\x50\x25\xda\xc2\xf5\xa7\xcd\xcd\xea\xf5\xc5\x77\x17\x69\x92\xdc\x62\xb1\x86\x92\xa8\x0e\xeb\x2c\xdb\x6a\x2a\x9b\x3c\x95\xae\xca\x9c\x2d\x8c\x7b\xcc\xe2\x5f\x6e\x5c\x9e\x55\x22\x10\xfa\xac\x30\xba\x0e\xd9\xdb\xd7\x6f\xdf\xbc\xfe\xe1\xcd\xc5\xca\x16\xb4\xea\x9d\xa5\x95\x4a\x92\x3b\xf2\x8d\xa4\x00\xc2\x2a\xf0\x18\x5c\xe3\x25\x06\x90\xc2\x8e\x21\x82\xb3\x08\xce\x43\xe5\x3c\x26\x43\xa4\xb4\xab\x31\x9c\x83\x14\xc6\xa0\x82\x07\x8d\x8f\x21\x85\x8f\x42\x96\xf1\x39\x2e\x83\xc7\xda\x63\xe0\x2c\x13\x01\x4a\x17\x05\x7a\xc6\xfb\x53\x5b\x05\xae\x18\xb2\x3e\x87\xd0\xc8\x12\x44\x00\x01\xd2\xa3\x20\xe7\x21\xd7\x6e\xeb\x45\x5d\xee\x12\xe7\x41\xc0\xcf\x37\x1f\x7f\x04\x5d\x89\x2d\x42\xa1\x0d\xa6\xc9\xab\x2c\x49\xea\x26\x1f\x19\xfd\xdc\x81\xfd\xca\x91\xc0\x53\x92\x00\x00\x64\x19\xbc\x87\x5b\x0c\xce\x3c\xa0\x67\x4e\x1f\xb4\xc2\x00\x42\x4a\x0c\x01\xc8\x81\x80\x80\x34\x8d\xa5\xcb\xa4\xdb\x3d\x82\x84\xc8\x13\xd3\xd0\xb3\x04\x0b\x4c\xb7\x29\x08\x0b\xbf\x5c\xdf\x2f\xf7\x28\x23\xae\xb6\xb6\x84\xbe\x10\x12\x7b\x18\x72\x7d\x0c\x93\x10\xb8\xfe\xd1\x29\x50\x29\x08\x34\x41\x68\xea\xda\x79\x9a\x47\xc1\xc9\x0e\xae\x07\xe4\x31\xb9\xa7\x68\xd5\x5b\x16\x8d\x85\x2d\x52\x24\x63\xb1\x5c\xc3\x6f\xf7\xbb\x1a\xbf\x1e\x98\xf8\x76\x37\x9b\x2d\x7e\x8f\x51\xac\x81\x2d\x97\x6b\x78\x6f\x77\xad\x36\xde\xc5\x5d\xcf\x47\x08\xfd\xe0\x8c\x41\x49\xda\x59\xd0\x5c\xbb\xad\x77\x4d\xcd\x64\x46\x05\x74\xd8\x9e\x79\x50\xf8\x17\xe4\x3b\xd8\x5c\xfd\x9b\x94\x26\xf0\x87\xc9\xe5\xce\x7b\xf7\xc8\x81\xf7\xe6\x0b\xad\xd6\xf0\x65\x63\xe9\xe2\xfb\xe5\x1a\xfe\xff\xd4\x7f\x7f\x3e\x46\xcc\xe6\xaa\xa5\xa5\xb5\xff\xba\x97\xe2\x95\x0e\xb5\x11\xbb\x36\xab\x5c\x04\x2d\x3b\x51\xc7\x02\x59\x69\x1a\x56\x11\x17\xce\x8a\x0a\xe3\xd9\x51\x18\xa4\xd7\x35\x47\xdb\xa3\xb8\x82\xc5\xe1\xf2\x3f\x50\x52\x0a\x9f\x5d\xa0\xee\x25\x40\x28\x5d\x63\xd4\xbe\x5e\xd8\xc7\x01\x41\x9d\xf0\xfa\x90\xe6\x4c\x18\xa4\x18\xc2\x1a\xee\xc8\x6b\xbb\x3d\x58\x9c\x84\x75\xd2\x86\xca\xa6\xca\xad\xd0\x66\x52\xf4\xa7\x6b\x6d\xb0\x23\x84\x7f\xda\x6a\x5a\x0c\x6f\xfc\x9b\xba\x3d\x9f\xad\x1c\xf1\x39\x37\x78\xc1\x61\x6f\xb2\x9c\x24\xca\xbf\x80\xa6\x48\x23\xd9\x97\xd1\xf3\xe1\xe2\xc4\x2b\x5c\x1e\x94\x63\x66\x3a\xf8\x87\xcb\x31\x96\xc1\xec\x79\x4f\x0c\x1c\x58\xa7\x6f\xb4\xe8\xb5\x9c\x48\x35\x0a\x62\xec\x72\x20\x62\x77\x82\x40\xce\xa3\x02\x67\xb9\x5f\xb8\xa2\x00\x59\x0a\x6d\xf7\x9b\x0a\x03\xb7\xad\x36\x47\x68\x02\x2a\xee\x06\x1e\x63\x93\xe4\x26\x1c\xdb\x5d\x38\x07\xee\x17\xae\xed\xd0\x8e\x4a\xf4\x50\xa1\xd2\xe2\x94\x50\xc6\xe0\x62\xe0\x87\x27\xa7\xf1\x9a\xa5\x3f\x11\xc3\x98\xeb\x4f\xf7\xf7\x37\x63\xbe\x31\x97\x56\xf3\x7d\xcf\xd4\xb9\x41\x10\xc4\xc2\x66\x5b\x58\x38\x1f\x1f\xee\x96\xf0\xe5\xf6\x53\x0a\x27\x82\xea\x71\xd7\xc7\x82\x62\x0d\x36\xde\x0c\x21\xcd\x45\x37\x59\x39\x2a\x8a\xc6\x73\x19\x1b\x3f\x2d\xe0\xcb\x39\xef\xa1\x78\xa4\xc6\xdb\x01\xec\xa4\x0e\x36\x37\xd7\x77\xf7\x83\x72\xda\x5d\xcc\xd2\xa8\xa6\x6e\x3e\xf1\xc0\xea\xcf\x7d\xbf\xb9\x53\x04\x8f\x38\x3b\x99\x63\xa0\x6d\xc4\xdd\x97\x06\x7f\xeb\xea\x0f\xc2\xe3\xa8\x0a\xc5\x6d\x94\x4a\xd4\x3e\x4e\x3e\x6e\x1d\x5a\xa1\x25\x5d\x68\xf4\xb0\xf8\xb0\xb9\x5a\xf6\x18\x5e\x44\xb1\x50\x29\x2c\xf0\xf4\xf5\x28\x09\xbe\xdc\x6e\x52\x78\x0f\xd2\x68\xde\x2a\xea\xda\x68\x29\xe2\x99\x61\x1d\x36\x01\xdb\x4e\xf4\x61\x73\x35\x99\x59\x05\xcf\x6b\x96\x9f\x71\x42\xc5\xb6\xd7\x66\xf0\xa0\x05\xa7\x13\x83\xdd\x0a\xc2\x47\xb1\x3b\xa5\x4a\xb6\x99\x09\x60\xa0\x39\xcb\xd8\x1b\x2b\x8c\x81\x8f\x24\xc5\x7c\xc6\xa0\xa2\x9f\x76\xfa\x8f\x9b\xa7\x38\xb3\xbb\x91\x72\x32\xa4\xba\x2e\x42\xaa\x5d\x26\x9d\x95\x58\x53\xc8\x3a\xfc\x95\x50\xca\xb3\x9c\xed\x36\x3b\x06\xd6\xcb\x52\xf2\x4c\xd9\x97\x65\x96\xc1\x8d\xa0\x32\x1e\x0a\x0b\x2e\x36\x19\x61\xa0\xe6\x6f\xdd\x3c\x8f\xb5\x9d\x8c\xb6\x81\xa6\xb6\x0c\xce\xef\x4e\xe5\x10\xaf\x89\x85\x46\x9e\x0e\x01\x9c\x35\x3b\xb0\x88\x0a\x15\xe8\x62\x84\x8e\x97\x8b\x10\xef\x11\xdf\x86\xfc\x07\xb4\x30\xe8\x2a\xec\x02\x61\x15\x5e\x24\x84\x93\xec\x19\x79\xb7\x77\x52\x27\x64\x9d\xcf\x0d\x8f\x1e\x5c\xa9\x15\x5c\x32\xc3\x87\x4b\x91\xca\xcb\x88\x71\xec\x54\x0f\x34\x35\xb6\xbd\x1e\xf4\x87\x91\xf9\x89\x34\x5b\x41\xfa\x01\xb9\x25\x8d\xf2\xf9\x6f\xca\x29\xdd\xe3\x8a\x5c\xd6\xe9\x65\xc5\x9f\x57\xce\xae\x1e\x31\xcf\xfe\xd7\x7a\x59\x35\xde\x84\x53\x9c\x7d\xab\xf9\xe8\x62\xa0\x15\x2e\x27\xc9\xcf\xad\x60\x6c\x53\x67\x1c\xc1\x3a\xcb\xce\x52\x2e\x9e\xa0\x45\xcf\xe5\xb2\xff\x70\x96\x9d\x0d\xcf\x8c\xb5\x9c\x41\x3d\xcf\xde\x8e\x35\xc2\xd3\x1e\x0e\x5a\xe3\x73\xf2\x77\x00\x00\x00\xff\xff\x11\x66\xde\xf4\xdb\x0c\x00\x00"

func metadataviewsCdcBytes() ([]byte, error) {
	return bindataRead(
		_metadataviewsCdc,
		"MetadataViews.cdc",
	)
}

func metadataviewsCdc() (*asset, error) {
	bytes, err := metadataviewsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "MetadataViews.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2, 0x49, 0x9e, 0xd6, 0x60, 0x9f, 0x9b, 0x70, 0x80, 0x1d, 0x15, 0x67, 0xc0, 0x6b, 0x19, 0x2, 0x66, 0xbb, 0xd4, 0x75, 0x44, 0xa3, 0xa1, 0x39, 0x69, 0xf1, 0x61, 0x91, 0x88, 0x37, 0x66, 0x5f}}
	return a, nil
}

var _nonfungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x41\x8f\xdb\xba\x11\xbe\xeb\x57\xcc\xcb\x03\x9a\xdd\xc0\x6b\xf7\x50\xf4\x60\x20\x68\xda\xb7\x6f\x01\x5f\xb6\x0f\x5b\x17\x3d\x04\x01\x4c\x8b\x23\x9b\x08\x45\x2a\x24\x65\xc7\x0d\xf6\xbf\x17\x33\x24\x25\xca\xf6\x26\x9b\x5b\x73\x89\x57\x22\xbf\x99\xf9\xe6\x9b\x8f\xd4\xe2\xdd\xbb\xaa\xfa\xf5\x57\x58\xef\x11\x1e\xb4\x3d\xc2\xa3\x35\x77\x0f\xbd\xd9\xa9\xad\x46\x58\xdb\xcf\x68\xc0\x07\x61\xa4\x70\x92\x17\x6e\x1e\xad\xc9\xef\xf9\xf5\x06\x6a\x6b\x82\x13\x75\x00\x65\x02\xba\x46\xd4\x58\x55\x84\x37\xfc\x09\x61\x2f\x02\x08\xad\xc1\x58\x73\xd7\x64\xf4\xc0\xe8\x79\xb7\x87\xda\xf6\x5a\xd2\xdf\x8d\x75\x2d\x04\x3b\xaf\x56\x0d\x08\xe8\x3d\x3a\x38\x0a\x13\x3c\x04\x0b\x12\x3b\x6d\x4f\x20\xc0\xe0\x11\x4c\x13\x86\xfd\x33\x08\x7b\x54\x6e\xcc\xe6\xc8\x70\x06\x51\x56\xc1\x82\x6a\x3b\x8d\x2d\x9a\x40\xcb\xe0\xbc\x88\x31\xd7\x39\xe7\x7e\x89\xb3\x17\x07\xca\x18\x1a\xab\x89\x26\x2a\x86\x80\x5c\xaf\xd1\x83\x30\x12\x8c\x68\x95\xd9\x55\x5c\x6a\x98\x54\xef\x3b\xac\x55\xa3\xd0\xcf\x13\x83\x0f\xeb\x0d\x38\xf4\xb6\x77\x99\xaa\xda\x3a\x1c\x1e\x41\x38\x75\x89\x33\x87\x9d\x43\x8f\x54\xbb\x30\xf0\xf8\xb0\x06\x65\x18\xdd\xb7\xc2\x8d\xb5\x27\xe0\xdf\xac\xd6\x58\x07\x65\xcd\x06\x9e\x26\xf8\x23\x34\xa1\xfa\x60\x1d\x65\xcd\xd4\xbe\xf5\x8c\x5b\x0f\x7b\xe7\xd5\x8a\x5a\x59\xeb\x5e\xf2\xa2\x06\x8f\xd0\xf4\x86\xdf\x71\x0b\x04\x33\x40\x59\xd8\xa3\x41\x47\x8f\x50\x78\xa5\x4f\x55\x6b\x0f\xa9\xad\x9e\x12\x25\x5a\x6c\x1f\xc0\x36\xbc\xba\x0c\xc1\xf9\xfe\xe1\xec\x41\x49\x74\x1b\x5e\xb9\x79\xc2\x1a\xd5\x81\xfe\x1c\xd2\x1d\x48\xf4\x5c\x87\x2f\x9f\x80\xc4\x5a\x0b\x87\x45\x72\x47\x15\xf6\xe0\x6d\x8b\xd0\x39\x64\xd0\xce\x7a\xa6\x49\x2a\x5e\x51\x25\x56\xbf\xf4\xca\x21\x27\x35\x72\x56\x74\xb7\x46\x17\x84\x32\xa9\xa7\x0c\xb4\xc5\xbd\x38\x28\xeb\x86\x69\xf0\x51\x29\x27\xa0\x14\x3c\x76\xc2\x89\x80\xb0\xc5\x5a\xf4\x94\x66\x80\x9d\x3a\xa0\xe7\x18\xac\x60\xfa\x21\xb6\x4a\xab\x70\xa2\x48\x7e\x4f\xfb\x04\x38\x6c\xd0\xa1\xa9\x91\x44\x1a\x15\x5c\xa6\x44\xe9\x5a\xa3\x4f\x80\x5f\x3b\xeb\x13\x5e\xa3\x50\xcb\xa8\xba\xb1\x76\x65\xc0\x1a\x04\xeb\xa0\xb5\x0e\xab\xc4\xf9\x48\xd7\x1c\x56\x34\x83\xde\xa6\xc4\x28\x29\x7f\x9e\x55\x2b\x3e\x23\xd4\xbd\x0f\xb6\x1d\x9a\x90\x48\x9b\x0c\xd0\xb4\x11\x34\x96\x16\x0e\xc2\x29\xdb\x13\xa4\x32\xbb\xd4\x0b\x82\x8f\x7a\x98\x57\xd5\x3f\x4e\xd0\x7b\xe2\x73\x40\xe6\x12\x46\xa0\x59\x4a\xca\x36\x2c\xc9\xa9\xc6\x3d\xd4\xc2\x80\x47\x23\x2b\xda\xe5\xa2\x58\xb2\xda\x3a\x44\x77\x17\xec\x1d\xfd\x3f\xe3\xd8\x24\x3c\x6a\x99\xd9\x51\x7e\x1c\x84\xa7\x99\xd2\x12\x50\x23\xa1\x6a\xd0\x28\x77\xe8\xaa\x8b\x71\x5a\x5b\x0e\x95\xa7\x8e\x54\x6f\x6c\xd8\xa3\xe3\x14\x67\x83\x2d\xb1\x37\x78\xe2\xe6\xc4\xd0\xd2\x89\x38\x1a\x8f\x0f\xeb\xaa\x71\xb6\xbd\xe8\x29\xfb\x94\x81\x3a\x3b\x88\xc4\xce\x7a\x15\x86\x4e\x82\x35\x93\x58\x6f\x7d\x35\xd5\x68\x6d\xa9\x13\x21\xca\x37\x38\x61\x7c\x83\x6e\x5e\x55\xef\x16\x55\xb5\x58\xb0\x93\xb7\x24\xde\x38\xd5\xe7\xd6\x3c\x87\x7f\x32\x74\xf9\x96\x9a\xa5\x35\x6d\x56\x6d\x67\x5d\x88\x6d\x29\xfa\xad\x7c\xe1\xed\x8b\x45\xd5\xf5\xdb\x2b\xd0\x97\xae\xfa\xad\xaa\x00\x00\x52\x56\xc1\x06\xa1\xc1\xf4\xed\x16\x1d\x7b\x42\x6c\x1d\x2b\x55\xf9\xe8\x7a\xca\x00\x7e\x55\x3e\xf0\x44\xd0\x5e\x0a\x75\x10\x2e\x6e\xfe\x57\xdf\x75\xfa\xb4\x84\x7f\xaf\x4c\xf8\xeb\x5f\x06\xf0\xdf\x0f\x31\x4d\x11\x00\x5b\x15\x02\x4a\x38\x12\xc7\xa9\x0f\x45\xaa\x54\x87\x0a\x4a\x68\xf5\x5f\x94\x69\xfb\x10\x06\x19\xe6\xb7\xb4\x78\x35\x2e\xbc\xb9\xbd\x16\x4a\xf9\x69\x34\x91\x0e\x34\xe5\x07\x25\x98\x59\xde\xa7\x8c\x54\xb5\x08\xac\xc6\xc1\x38\x2f\x7c\x31\x01\x07\x38\x8a\x02\x04\x48\x47\xf3\x32\xdb\xc5\x02\x56\x17\x7b\x95\x07\x63\x43\xf4\x5d\x10\x75\x6d\x7b\x13\xde\x7a\x36\x7b\xb1\xc3\x19\x6c\x08\x66\xc3\xad\x86\x2d\xc2\xc6\x28\xbd\x99\x5f\xe7\xe0\x3f\x29\xf4\x8d\x92\x99\xec\x19\x67\xb1\x84\xbf\x4b\xe9\xd0\xfb\xbf\x5d\xa5\xe4\x25\x3e\x92\xc6\x51\xf2\x20\x4d\x0e\x82\xb3\xaa\x42\x66\x2a\x59\xdd\x6b\x88\x2a\xd1\x5f\x28\xe8\x3e\x2e\x99\xd4\x13\xec\xb5\x6a\x56\xd3\x4b\x4b\x92\x90\x1f\xce\xff\xf1\x7a\x72\x1e\xe9\xf2\xd0\x82\x15\xa9\xef\x1b\xaf\x28\xe6\xa0\x37\xea\x4b\x8f\xb0\xba\x4f\xa4\x89\x7a\xcf\x32\xdd\x0b\x3f\x2c\x25\x40\x8d\x01\xc6\x84\xf9\xd5\xf3\x90\xe7\x53\x3c\xc3\xda\x81\x7b\xf2\x93\x94\x1c\xa9\xec\x9a\x81\x52\x0d\x79\x3f\x5f\xa5\x1a\x65\xe2\x19\x94\x32\x27\x53\x42\x19\x1d\x8f\x30\x13\x1e\x3b\x3c\xd5\x72\x59\xeb\xe3\xc3\x7a\x79\x5e\xe6\x0f\x73\x2f\x38\xb6\xd0\xa2\x54\x74\x72\x66\xb9\x7b\xc8\xb6\x59\x98\xe6\x2b\xb8\xce\x97\x89\x29\xdf\x83\x27\x3b\xa4\xcb\xc9\x70\x8d\x1a\x62\x14\x9a\x22\xd7\x8b\x8b\x54\x80\x78\x1a\x47\x46\xdc\xa4\xb4\xa6\x37\x03\xec\x4d\xfe\xb1\xba\xcf\xb5\xde\x2e\xe1\xc3\x94\x0f\xde\x48\xf7\x90\xe9\x23\xfa\xe7\xd0\xf7\x3a\xcc\x95\x84\xf7\xef\xa1\xc4\x7a\x43\x42\x59\xdd\x67\xe5\x8f\x5e\x10\x67\xaa\xed\x7d\xa0\x21\xe6\xab\xa0\x68\x11\x44\x1c\x17\xba\xd9\xa0\xa7\x51\x58\xdd\xbf\x99\x44\x7b\xae\xa6\xbf\x7e\xd0\x8d\x34\x53\x3e\xf3\xf0\x53\xad\xc8\x17\xb9\xec\xff\x29\x50\x3e\xe9\x82\xf8\x3c\x36\x42\xf0\x2f\xe1\x76\x3d\x4b\x99\x7a\x20\xa4\x2c\x5b\x70\x16\xba\x08\x5f\x76\x24\x81\xdf\x30\x3f\xb1\x05\xb7\x2f\x17\xca\x03\x33\xb8\x64\x3a\xc6\x6b\xdb\xb6\x7c\xd7\xca\x1b\xba\x7e\xab\x95\xdf\x43\x63\xdd\xf0\x71\x31\xc9\xe5\x85\xfa\xc7\x8c\xff\x20\x84\xfa\x6c\x36\xbe\x9b\x6e\xb9\x68\x87\x61\x75\xef\x6f\x6e\x97\xf0\x31\x6a\xeb\xd3\xc5\x92\xad\x75\xce\x1e\x1f\x1f\xd6\x85\xb5\xdd\x2e\xe1\x4f\x79\x58\xaf\x1b\x46\x2a\x28\x0d\x80\xa9\x1d\x5d\x27\x26\x9f\x1f\x85\x4d\x6c\x31\xdf\xb4\x65\xfe\xfa\x18\xee\x06\xe4\x34\xd9\x5f\x5e\x14\xc6\x48\xc7\x72\x98\xd2\xd9\x20\x92\xd9\x35\xba\x4a\xd9\xdc\x2b\x7e\x27\x1c\xdf\x50\xf7\x56\xcb\xd1\x95\x53\x3e\x57\x24\x92\xef\x0d\x74\x80\x48\x5a\xbb\x84\x0f\xdf\x22\x3f\x4b\xda\xfb\x5c\xfd\x5f\xd8\xc4\xf7\x06\x24\xce\xc7\xe5\x40\x8c\xb9\x78\x90\x03\x39\x25\xd0\xb0\x29\x44\x17\x49\x1b\x95\x04\xe1\x9c\x38\xbd\x4e\x8d\x25\x60\x54\x22\x38\x0c\xbd\x33\x69\x62\x9d\x38\x65\x7b\xa2\x77\x71\xa6\x1c\xe6\x9e\xd4\xd7\x7b\xf2\x82\xae\xcb\x60\x4f\x39\x4a\x52\x37\xca\xf1\x2b\x29\xde\xc4\xcb\x2f\xe1\x2b\x71\x16\x0b\xf0\x76\x3c\xbf\x63\x73\xf8\xf3\xc1\xa1\x90\x20\x45\x10\x4c\x11\xdf\xc1\x5b\x0c\x7b\x2b\xd3\xa9\xa3\xc2\xcf\x4c\xd8\xb9\xc7\x3b\xbc\x62\xf1\x1e\x75\x33\x1f\x54\xf8\x51\xc9\x4f\xf0\xcb\x7b\x30\x4a\x2f\xe1\x0d\x61\x48\x8b\xf1\xe2\xc6\xf7\xde\xcb\xaa\x7e\x79\xad\x8f\xd7\x0e\x45\xc0\xdf\xdb\x2e\x9c\x8a\x0f\x86\xf8\x94\x5b\x86\xf4\xea\xd2\xc9\x21\x7e\x4e\x45\xce\xcf\x25\x5d\x12\x79\x62\x0a\xed\x91\xe9\xf7\x55\x49\xd2\xd5\xd8\xd4\xe0\x0f\x45\x2a\x85\x0b\x5e\x9e\x86\xe9\x24\xcc\xd2\x98\x6b\x34\xbb\xb0\xa7\x63\xf1\xcf\xe9\x34\x8c\x31\x64\x39\x8a\xf9\x18\xe4\xca\x0a\xa2\x32\x35\xcf\xd5\xff\x02\x00\x00\xff\xff\x33\x4d\x81\x27\xe0\x12\x00\x00"

func nonfungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_nonfungibletokenCdc,
		"NonFungibleToken.cdc",
	)
}

func nonfungibletokenCdc() (*asset, error) {
	bytes, err := nonfungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NonFungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x61, 0xca, 0x9d, 0xaa, 0x66, 0x36, 0xdf, 0xbc, 0x51, 0xdb, 0x7b, 0x51, 0xd8, 0x3d, 0x6f, 0x4e, 0x9c, 0x8e, 0x50, 0x28, 0x7c, 0x18, 0x1d, 0x2, 0xb2, 0xc2, 0x2b, 0x26, 0xa1, 0xfe, 0x2d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"ExampleNFT.cdc":       examplenftCdc,
	"MetadataViews.cdc":    metadataviewsCdc,
	"NonFungibleToken.cdc": nonfungibletokenCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"ExampleNFT.cdc": {examplenftCdc, map[string]*bintree{}},
	"MetadataViews.cdc": {metadataviewsCdc, map[string]*bintree{}},
	"NonFungibleToken.cdc": {nonfungibletokenCdc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
