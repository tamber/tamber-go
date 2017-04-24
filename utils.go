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

func StringPointer(v string) *string {
	return &v
}

func Int64Pointer(v int64) *int64 {
	return &v
}

func IntPointer(v int) *int {
	return &v
}

func FloatPointer(v float64) *float64 {
	return &v
}
