package main

var user_history = map[string][]string{
	"42": {"electronics", "electronics"},
	"7":  {"home"},
}

var items = []Item{
	{
		ID:       1,
		Name:     "Mechanical Keyboard",
		Category: "electronics",
		Score:    0.95,
	},
	{
		ID:       2,
		Name:     "Gaming Mouse",
		Category: "electronics",
		Score:    0.87,
	},
	{
		ID:       3,
		Name:     "Coffee Mug",
		Category: "home",
		Score:    0.42,
	},
}
