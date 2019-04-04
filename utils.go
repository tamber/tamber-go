package tamber

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

func Gzip(filepath string) error {
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath + ".gz")
	if err != nil {
		return err
	}
	w, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		return err
	}
	w.Write(input)
	w.Close()
	return nil
}

func String(v string) *string {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func Int(v int) *int {
	return &v
}

func Float(v float64) *float64 {
	return &v
}

type StringId string

func (s StringId) GetUserParams() *UserParams {
	return &UserParams{Id: string(s)}
}

func (s StringId) GetItemParams() *ItemParams {
	return &ItemParams{Id: string(s)}
}
