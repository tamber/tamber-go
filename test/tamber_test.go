package test

import (
	"encoding/csv"
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/behavior"
	"github.com/tamber/tamber-go/client"
	"github.com/tamber/tamber-go/discover"
	"github.com/tamber/tamber-go/event"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/user"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	TestProjectKey = "Mu6DUPXdDYe98cv5JIfX"
	TestEngineKey  = "SbWYPBNdARfIDa0IIO9L"
)

func errFunc(exp string, err interface{}) {
	fmt.Printf("\n%s: %v\n", exp, err)
}

func streamProperties(t *testing.T, filepath string, e *client.API) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		t.Error(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		t.Error(err)
		os.Exit(1)
	}
	items := make(map[string]struct{})
	// sanity check, display to standard output
	for _, row := range rawCSVdata {
		id_params := strings.Split(row[1], "-")
		itype := id_params[0]
		id := id_params[1]
		if itype == "artwork" {
			continue
		}

		if _, ok := items[id]; !ok {
			items[id] = struct{}{}

			i, _, err := e.Item.Update(&tamber.ItemUpdateParams{
				Id: row[1],
				Updates: tamber.ItemUpdates{
					Add: tamber.ItemFeatures{
						Properties: map[string]interface{}{
							"type": itype,
						},
					},
				},
			})
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("%+v", i)
			}
		}
	}
}

func TempTest(t *testing.T) {
	c := &client.API{}
	sc := &tamber.SessionConfig{URL: tamber.ApiUrl, HTTPClient: &http.Client{Timeout: 80 * time.Second}}
	sc.SetErrFunc(errFunc)
	c.Init(TestProjectKey, TestEngineKey, sc)
	filepath := "./data.csv"
	streamProperties(t, filepath, c)
}

func EngieVarTest(t *testing.T) {
	c := client.New(TestProjectKey, TestEngineKey, nil)
	_, _, err := c.Event.Track(&tamber.EventParams{
		User:     "user_rlox8k927z7p",
		Behavior: "click",
		Item:     "item_wmt4fn6o4zlk",
		Hit:      tamber.Bool(true),
	})

	if err != nil {
		panic(err)
	}
}

func BasicTest(t *testing.T) {
	t.Log("\nRecs w/ Test Events")
	e, info, err := event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_i5gq90scc1",
		Behavior: "mention",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Event: %+v\n", *e)
	}
	d, info, err := discover.Recommended(&tamber.DiscoverParams{
		User:   "user_jctzgisbru",
		Number: tamber.Int(100),
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			t.Logf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}
}

func PartialTest(t *testing.T) {
	//Create a behavior
	b, info, err := behavior.Create(&tamber.BehaviorParams{
		Name:         "like",
		Type:         tamber.String("exponential"),
		Desirability: 0.5,
		Params: map[string]interface{}{
			"step_auto": true,
		},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Behavior: %+v\n", *b)
	}
	//need to create properties before applying them
	i, info, err := item.Create(&tamber.ItemParams{
		Id: "item_i5gq90scc1",
		Properties: map[string]interface{}{
			"duration": 64.5,
			"genre":    "comedy",
		},
		Tags: []string{"hilarious", "heart warming"},
	})

	//Track an event - user performs the behavior on an item
	e, info, err := event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_i5gq90scc1",
		Behavior: "like",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Event: %+v\n", *e)
	}
	e, info, err = event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_u9nlytt3w5",
		Behavior: "like",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Event: %+v\n", *e)
	}

	//Check User - w/ get recs
	u, info, err := user.Retrieve(&tamber.UserParams{
		Id:      "user_jctzgisbru",
		GetRecs: &tamber.DiscoverNextParams{},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("User: %+v\n", *u)
	}

	//Get User's Recommended Items
	d, info, err := discover.Recommended(&tamber.DiscoverParams{
		User:   "user_jctzgisbru",
		Number: tamber.Int(100),
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			t.Logf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}

	//Check Item - w/ get recs
	i, info, err = item.Retrieve(&tamber.ItemParams{
		Id: "item_i5gq90scc1",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("\nItem: %+v\n", *i)
	}

	//Get Item's Similar Items
	d, info, err = discover.Similar(&tamber.DiscoverParams{
		Item:   "item_i5gq90scc1",
		Number: tamber.Int(100),
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			t.Logf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}
}

func TestTamberGo(t *testing.T) {
	tamber.DefaultProjectKey = TestProjectKey
	tamber.DefaultEngineKey = TestEngineKey

	t.Logf("\n\nBasic Test\n---------\n\n")
	BasicTest(t)

	t.Logf("\n\nPartial Test\n---------\n\n")
	PartialTest(t)

	t.Logf("\n\nExpanded Test\n------------\n\n")

	//User
	t.Log("User - Create")
	u, info, err := user.Create(&tamber.UserParams{
		Id: "user_fwu592pwmo",
		Metadata: map[string]interface{}{
			"city": "San Francisco, CA",
		},
		Events: []tamber.EventParams{
			tamber.EventParams{
				Item:     "item_u9nlytt3w5",
				Behavior: "like",
			},
			tamber.EventParams{
				Item:     "item_i5gq90scc1",
				Behavior: "like",
			},
		},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("User: %+v\n", *u)
	}

	t.Log("User -- Update")
	u, info, err = user.Update(&tamber.UserParams{
		Id: "user_fwu592pwmo",
		Metadata: map[string]interface{}{
			"city": "Mountain View, CA",
			"age":  "55-65",
			"name": "Rob Pike",
		},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("User: %+v\n", *u)
	}

	t.Log("User -- Retrieve")
	u, info, err = user.Retrieve(&tamber.UserParams{
		Id: "user_fwu592pwmo",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("User: %+v\n", *u)
	}

	//Item
	t.Log("Item - Create")
	i, info, err := item.Create(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
		Properties: map[string]interface{}{
			"clothing_type": "pants",
			"stock":         90,
		},
		Tags: []string{"casual", "feminine"},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("\nItem: %+v\n", *i)
	}

	t.Log("Item - Update")
	i, info, err = item.Update(&tamber.ItemUpdateParams{
		Id: "item_nqzd5w00s9",
		Updates: tamber.ItemUpdates{
			Add: tamber.ItemFeatures{
				Properties: map[string]interface{}{
					"stock": 89,
				},
			},
			Remove: tamber.ItemFeatures{
				Tags: []string{"casual"},
			},
		},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("\nItem: %+v\n", *i)
	}

	t.Log("Item - Retrieve")
	i, info, err = item.Retrieve(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("\nItem: %+v\n", *i)
	}

	//event
	//track -- note: repeat behavior
	e, info, err := event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_i5gq90scc1",
		Behavior: "like",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Event: %+v\n", *e)
	}
	//retrieve

	e, info, err = event.Retrieve(&tamber.EventRetrieveParams{
		User: tamber.String("user_jctzgisbru"),
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Event: %+v\n", *e)
	}

	//batch
	batch_resp, info, err := event.Batch(&tamber.EventBatchParams{
		Events: []tamber.EventParams{
			tamber.EventParams{
				User:     "user_y7u9sv6we0",
				Item:     "item_u9nlytt3w5",
				Behavior: "like",
			},
			tamber.EventParams{
				User:     "user_y7u9sv6we0",
				Item:     "item_i5gq90scc1",
				Behavior: "like",
			},
			tamber.EventParams{
				User:     "user_k6q76ohppz",
				Item:     "item_i5gq90scc1",
				Behavior: "like",
			},
			tamber.EventParams{
				User:     "user_y7u9sv6we0",
				Item:     "item_d1zevdf6hl",
				Behavior: "like",
			},
			tamber.EventParams{
				User:     "user_y7u9sv6we0",
				Item:     "item_nqzd5w00s9",
				Behavior: "like",
			},
			tamber.EventParams{
				User:     "user_k6q76ohppz",
				Item:     "item_nqzd5w00s9",
				Behavior: "like",
			},
		},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Batch Response: %+v\n", *batch_resp)
	}

	//discover
	t.Log("Discover - Recommended (w/ TestEvents and Filter)")
	d, info, err := discover.Recommended(&tamber.DiscoverParams{
		User:   "user_jctzgisbru",
		Number: tamber.Int(100),
		TestEvents: []tamber.EventParams{
			tamber.EventParams{
				User:     "user_jctzgisbru",
				Item:     "item_d1zevdf6hl",
				Behavior: "like",
			},
			tamber.EventParams{
				User:     "user_jctzgisbru",
				Item:     "item_nqzd5w00s9",
				Behavior: "like",
			},
		},
		Filter: map[string]interface{}{
			"or": []interface{}{
				map[string]interface{}{
					"gt": []interface{}{
						map[string]interface{}{
							"property": "stock",
						},
						20,
					},
				},
				map[string]interface{}{
					"eq": []interface{}{
						map[string]interface{}{
							"property": "clothing_type",
						},
						"shirt",
					},
				},
			},
		},
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			t.Logf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}
	//similar
	//recommendedSimilar
	//popular
	//hot

	//Behavior
	//create - see BasicTest()
	t.Log("Behavior - Retrieve")
	b, info, err := behavior.Retrieve(&tamber.BehaviorParams{
		Name: "like",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("Behavior: %+v\n", *b)
	}

	//Remove Tests
	t.Log("Item - Remove")
	i, info, err = item.Remove(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
	})
	if err != nil {
		t.Error("err:", err, "info:", info)
	} else {
		t.Logf("\nItem: %+v\n", *i)
	}

}
