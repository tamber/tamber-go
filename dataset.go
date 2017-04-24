package tamber

import (
	"compress/gzip"
	"io/ioutil"
)

func Gzip(filepath string) error {
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		err
	}
	f, _ = os.Create(filepath + ".gz")
	w, _ = gzip.NewWriterLevel(f, gzip.BestCompression)
	w.Write(input)
	w.Close()
}
