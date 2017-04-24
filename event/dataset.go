package event

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"errors"
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"io/ioutil"
	"os"
	"strconv"
)

var lineCols = []string{"user", "item", "behavior", "value", "created"}

func TrackToCSV(writer *csv.Writer, e *tamber.Event) error {
	return writer.Write([]string{
		e.User,
		e.Item,
		e.Behavior,
		strconv.FormatFloat(e.Value, 'f', -1, 64),
		strconv.FormatInt(e.Created, 10),
	})
}

// Returns csv writer for supplied filepath, a defer function, and an error.
// If fileAppend is true and a file already exists for the given filepath, events will be appended.
func GetCSVWriter(filepath string, fileAppend bool) (writer *csv.Writer, df func(), err error) {
	var file *os.File
	if fileAppend { // append if file exists
		if _, err = os.Stat(filepath); os.IsNotExist(err) {
			file, err = os.Create(filepath)
		} else {
			file, err = os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0660)
		}
	} else {
		file, err = os.Create(filepath)
	}

	if err != nil {
		return nil, nil, err
	}
	writer = csv.NewWriter(file)
	df = func() {
		writer.Flush()
		file.Close()
	}
	return
}

// BatchEventsToCSV appends events to the file for the given filepath (if the file does not exist
// it will be created) the filepath provided. BatchEventsToCSV is a convenient wrapper for GetCSVWriter
// and TrackToCsv; there is no performance benefit (it's actually more efficient to call GetCSVWriter once
// before loading multiple event batches - but not noticably faster for typical usage).
func BatchEventsToCSV(events []*tamber.Event, filepath string) error {
	w, df, err := GetCSVWriter(filepath, true)
	defer df()
	if err != nil {
		return err
	}
	for _, e := range events {
		err = TrackToCSV(w, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseLine(line []string) (e *tamber.Event, err error) {
	e = &tamber.Event{}
	for i, name := range lineCols {
		x := string(line[i])
		switch name {
		case "user":
			e.User = x
		case "item":
			e.Item = x
		case "behavior":
			e.Behavior = x
		case "value":
			e.Value, err = strconv.ParseFloat(x, 64)
			if err != nil {
				return nil, err
			}
		case "created":
			e.Created, err = strconv.ParseInt(x, 10, 64)
			if err != nil {
				return nil, err
			}
		}
	}
	return
}

func LoadEventsFromCSV(filepath string) (events []*tamber.Event, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	events = make([]*tamber.Event, len(lines))
	for i, line := range lines {
		if len(line) < len(lineCols) {
			return nil, errors.New(fmt.Sprintf("Incorrect number of values in events CSV file %s at line %d. Found %d, requires 5.", filepath, i, len(line), len(lineCols)))
		}
		events[i], err = parseLine(line)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error at line %d: %s", i, err.Error()))
		}
	}
	return events, nil
}
