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
	"strconv"
)

func (s *SessionConfig) CallUpload(method, path, key, ext, object, command string, params *UploadParams, resp interface{}) error {
	file, err := os.Open(params.Filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fname := fi.Name()
	if params.Type == EventsDatasetName {
		ext := go_path.Ext(fname)
		fname = fname[0:len(fname)-len(ext)] + ".csv"
	}

	fmt.Println("filepath:", params.Filepath)
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
	_ = writer.WriteField("projectid", strconv.FormatUint(uint64(params.ProjectId), 10))
	_ = writer.WriteField("type", params.Type)
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
