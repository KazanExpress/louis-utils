package archiver

import (
	"archive/zip"
	"bytes"
)

type Archive struct {
	Zip *zip.Reader
}

// OpenZip - opens archive from byte array
func OpenZip(zipBytes []byte) (*Archive, error) {
	var zipR, err = zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		return nil, err
	}

	return &Archive{Zip: zipR}, nil
}

// ListFileNames - returns names of files stored in archive
func (archive *Archive) ListFileNames() []string {

	var names = make([]string, len(archive.Zip.File))

	for i, file := range archive.Zip.File {
		names[i] = file.Name
	}

	return names
}
