package item

import (
	"encoding/json"
	tamber "github.com/tamber/tamber-go"
	"io/ioutil"
)

func LoadItemsFromJSONFile(filepath string) ([]*tamber.Item, error) {
	var items []*tamber.Item
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		return items, err
	}
	err = json.Unmarshal(input, &items)
	return items, err
}

func SaveItemsToJSONFile(items []*tamber.Item, filepath string) error {
	output, err := json.MarshalIndent(&items, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, output, 0644)
}
