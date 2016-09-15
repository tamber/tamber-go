/*
	Dataset Uploader BETA

	Authors:
	Alexi Robbins - alexi@tamber.com
*/

package tamber

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	go_path "path"
)

func (s *SessionConfig) CallUpload(method, path, key, ext, object, command, dpath, dtype string, resp interface{}) error {
	file, err := os.Open(dpath)
	if err != nil {
		return err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fname := fi.Name()
	if dtype == EventsDatasetName {
		ext := go_path.Ext(fname)
		fname = fname[0:len(fname)-len(ext)] + ".csv"
	}

	fmt.Println("dpath:", dpath)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fname)
	// part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	_ = writer.WriteField("type", dtype)
	// for key, val := range params {
	// 	_ = writer.WriteField(key, val)
	// }
	err = writer.Close()
	if err != nil {
		return err
	}

	path += object + "/" + command
	req, err := s.NewRequest("POST", path, key, ext, writer.FormDataContentType(), body)
	if err != nil {
		return err
	}

	if err := s.Do(req, resp); err != nil {
		return err
	}
	return nil
}
