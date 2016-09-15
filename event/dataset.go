package event

import (
	"encoding/csv"
	tamber "github.com/tamber/tamber-go"
	"strconv"
)

func TrackToCSV(writer *csv.Writer, e *tamber.Event) error {
	return writer.Write([]string{
		e.User,
		e.Item,
		e.Behavior,
		strconv.FormatFloat(e.Value, 'f', -1, 64),
		strconv.FormatInt(e.Created, 10),
	})
}
