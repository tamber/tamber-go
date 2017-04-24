package test

import (
	"fmt"
	tamber "github.com/tamber/tamber-go"
	"github.com/tamber/tamber-go/behavior"
	"github.com/tamber/tamber-go/client"
	"github.com/tamber/tamber-go/discover"
	"github.com/tamber/tamber-go/event"
	"github.com/tamber/tamber-go/item"
	"github.com/tamber/tamber-go/user"
)

func BasicTest() {
	fmt.Println("\nRecs w/ Test Events")
	e, info, err := event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_i5gq90scc1",
		Behavior: "mention",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Event: %+v\n", *e)
	}
	d, info, err := discover.Recommended(&tamber.DiscoverParams{
		User:   "user_jctzgisbru",
		Number: 100,
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			fmt.Printf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}
}

func PartialTest() {
	//Create a behavior
	b, info, err := behavior.Create(&tamber.BehaviorParams{
		Name:         "like",
		Type:         "chi-squared",
		Desirability: 0.5,
		Params: map[string]interface{}{
			"k": 1.0,
		},
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Behavior: %+v\n", *b)
	}
	//need to create properties before applying them
	i, info, err := item.Create(&tamber.ItemParams{
		Id: "item_i5gq90scc1",
		Properties: &map[string]interface{}{
			"duration": 64.5,
			"genre":    "comedy",
		},
		Tags: &[]string{"hilarious", "heart warming"},
	})

	//Track an event - user performs the behavior on an item
	e, info, err := event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_i5gq90scc1",
		Behavior: "like",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Event: %+v\n", *e)
	}
	e, info, err = event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_u9nlytt3w5",
		Behavior: "like",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Event: %+v\n", *e)
	}

	//Check User - w/ get recs
	u, info, err := user.Retrieve(&tamber.UserParams{
		Id:      "user_jctzgisbru",
		GetRecs: &tamber.DiscoverParams{},
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("User: %+v\n", *u)
	}

	//Get User's Recommended Items
	d, info, err := discover.Recommended(&tamber.DiscoverParams{
		User:   "user_jctzgisbru",
		Number: 100,
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			fmt.Printf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}

	//Check Item - w/ get recs
	i, info, err = item.Retrieve(&tamber.ItemParams{
		Id: "item_i5gq90scc1",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("\nItem: %+v\n", *i)
	}

	//Get Item's Similar Items
	d, info, err = discover.Similar(&tamber.DiscoverParams{
		Item:   "item_i5gq90scc1",
		Number: 100,
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			fmt.Printf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}
}

func Test() {
	tamber.DefaultProjectKey = "Mu6DUPXdDYe98cv5JIfX"
	tamber.DefaultEngineKey = "SbWYPBNdARfIDa0IIO9L"

	fmt.Printf("\n\nBasic Test\n---------\n\n")
	BasicTest()

	fmt.Printf("\n\nPartial Test\n---------\n\n")
	PartialTest()

	fmt.Printf("\n\nExpanded Test\n------------\n\n")

	//User
	fmt.Println("User - Create")
	u, info, err := user.Create(&tamber.UserParams{
		Id: "user_fwu592pwmo",
		Metadata: &map[string]interface{}{
			"city": "San Francisco, CA",
		},
		Events: &[]tamber.Event{
			tamber.Event{
				Item:     "item_u9nlytt3w5",
				Behavior: "like",
			},
			tamber.Event{
				Item:     "item_i5gq90scc1",
				Behavior: "like",
			},
		},
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("User: %+v\n", *u)
	}

	fmt.Println("User -- Update")
	u, info, err = user.Update(&tamber.UserParams{
		Id: "user_fwu592pwmo",
		Metadata: &map[string]interface{}{
			"city": "Mountain View, CA",
			"age":  "55-65",
			"name": "Rob Pike",
		},
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("User: %+v\n", *u)
	}

	fmt.Println("User -- Retrieve")
	u, info, err = user.Retrieve(&tamber.UserParams{
		Id: "user_fwu592pwmo",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("User: %+v\n", *u)
	}

	//Item
	fmt.Println("Item - Create")
	i, info, err := item.Create(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
		Properties: &map[string]interface{}{
			"clothing_type": "pants",
			"stock":         90,
		},
		Tags: &[]string{"casual", "feminine"},
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("\nItem: %+v\n", *i)
	}

	fmt.Println("Item - Update")
	i, info, err = item.Update(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
		Updates: &tamber.ItemUpdates{
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
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("\nItem: %+v\n", *i)
	}

	fmt.Println("Item - Retrieve")
	i, info, err = item.Retrieve(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("\nItem: %+v\n", *i)
	}

	//event
	//track -- note: repeat behavior
	e, info, err := event.Track(&tamber.EventParams{
		User:     "user_jctzgisbru",
		Item:     "item_i5gq90scc1",
		Behavior: "like",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Event: %+v\n", *e)
	}
	//retrieve

	e, info, err = event.Retrieve(&tamber.EventParams{
		User: "user_jctzgisbru",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Event: %+v\n", *e)
	}

	//batch
	batch_resp, info, err := event.Batch(&tamber.EventBatchParams{
		Events: []tamber.Event{
			tamber.Event{
				User:     "user_y7u9sv6we0",
				Item:     "item_u9nlytt3w5",
				Behavior: "like",
			},
			tamber.Event{
				User:     "user_y7u9sv6we0",
				Item:     "item_i5gq90scc1",
				Behavior: "like",
			},
			tamber.Event{
				User:     "user_k6q76ohppz",
				Item:     "item_i5gq90scc1",
				Behavior: "like",
			},
			tamber.Event{
				User:     "user_y7u9sv6we0",
				Item:     "item_d1zevdf6hl",
				Behavior: "like",
			},
			tamber.Event{
				User:     "user_y7u9sv6we0",
				Item:     "item_nqzd5w00s9",
				Behavior: "like",
			},
			tamber.Event{
				User:     "user_k6q76ohppz",
				Item:     "item_nqzd5w00s9",
				Behavior: "like",
			},
		},
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Batch Response: %+v\n", *batch_resp)
	}

	//discover
	fmt.Println("Discover - Recommended (w/ TestEvents and Filter)")
	d, info, err := discover.Recommended(&tamber.DiscoverParams{
		User:   "user_jctzgisbru",
		Number: 100,
		TestEvents: &[]tamber.Event{
			tamber.Event{
				User:     "user_jctzgisbru",
				Item:     "item_d1zevdf6hl",
				Behavior: "like",
			},
			tamber.Event{
				User:     "user_jctzgisbru",
				Item:     "item_nqzd5w00s9",
				Behavior: "like",
			},
		},
		Filter: &map[string]interface{}{
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
		fmt.Println("err:", err, "info:", info)
	} else {
		for _, rec := range *d {
			fmt.Printf("Item: %s :: Score: %s", rec.Item, rec.Score)
		}
	}
	//similar
	//recommendedSimilar
	//popular
	//hot

	//Behavior
	//create - see BasicTest()
	fmt.Println("Behavior - Retrieve")
	b, info, err := behavior.Retrieve(&tamber.BehaviorParams{
		Name: "like",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("Behavior: %+v\n", *b)
	}

	//Remove Tests
	fmt.Println("Item - Remove")
	i, info, err = item.Remove(&tamber.ItemParams{
		Id: "item_nqzd5w00s9",
	})
	if err != nil {
		fmt.Println("err:", err, "info:", info)
	} else {
		fmt.Printf("\nItem: %+v\n", *i)
	}

}

func EngieVarTest() {
	c := client.New("Mu6DUPXdDYe98cv5JIfX", "SbWYPBNdARfIDa0IIO9L", nil)
	_, _, err := c.Event.Track(&tamber.EventParams{
		User:     "user_rlox8k927z7p",
		Behavior: "click",
		Item:     "item_wmt4fn6o4zlk",
	})

	if err != nil {
		panic(err)
	}
}
