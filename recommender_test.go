package main

import "testing"

func generateItems(n int) []Item {
	items := make([]Item, n)

	for i := 0; i < n; i++ {
		items[i] = Item{
			ID:       i,
			Category: "electronics",
			Score:    float64(i) / float64(n),
		}
	}

	return items
}

func TestScoreItem(t *testing.T) {
	item := Item{
		ID:       1,
		Name:     "Keyboard",
		Category: "electronics",
		Score:    0.8,
	}

	history := []string{"electronics"}

	score := scoreItem(item, history)

	if score != 1.0 {
		t.Fatalf("expected score 1.0, got %f", score)
	}
}

func TestRecommend(t *testing.T) {
	items := []Item{
		{
			ID:       1,
			Category: "home",
			Score:    0.8,
		},
		{
			ID:       2,
			Category: "electronics",
			Score:    0.7,
		},
	}

	history := []string{"electronics"}

	result := recommend(items, history)

	if result[0].ID != 2 {
		t.Fatalf("expected item 2 first, got item %d", result[0].ID)
	}
}

func BenchmarkRecommend(b *testing.B) {
	items := generateItems(10000)
	history := []string{"electronics"}

	for i := 0; i < b.N; i++ {
		testItems := append([]Item(nil), items...)
		scoreItems(testItems, history)
	}
}

func BenchmarkConcurrentScoring(b *testing.B) {
	items := generateItems(10000)
	history := []string{"electronics"}

	for i := 0; i < b.N; i++ {
		testItems := append([]Item(nil), items...)
		scoreItemsConcurrently(testItems, history)
	}
}

func BenchmarkWorkerPoolScoring(b *testing.B) {
	items := generateItems(10000)
	history := []string{"electronics"}

	for i := 0; i < b.N; i++ {
		testItems := append([]Item(nil), items...)
		scoreItemsWithWorkerPool(testItems, history, 12)
	}
}
