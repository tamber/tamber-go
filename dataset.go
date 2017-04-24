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
