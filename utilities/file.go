package utilities

import (
	"io/ioutil"
	"mime/multipart"
	"os"
)

func CreateTemporaryFile(file multipart.File, fileHeader *multipart.FileHeader) (*os.File, error) {

	name := "upload-" + fileHeader.Filename
	emptyTempFile, err := ioutil.TempFile(os.TempDir(), name)
	if err != nil {
		return nil, err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Write this byte array to our temporary file
	emptyTempFile.Write(fileBytes)
	tempFile, err := os.Open(emptyTempFile.Name())
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
