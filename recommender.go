package main

import (
	"sort"
	"sync"
)

func scoreItem(item Item, history []string) float64 {
	score := item.Score

	for _, category := range history {
		if category == item.Category {
			score += 0.2
		}
	}

	return score
}

func scoreItems(items []Item, history []string) []Item {
	for i := range items {
		items[i].Score = scoreItem(items[i], history)
	}

	return items
}

func scoreItemsConcurrently(items []Item, history []string) []Item {
	var wg sync.WaitGroup

	for i := range items {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			items[i].Score = scoreItem(items[i], history)
		}(i)
	}

	wg.Wait()

	return items
}

func scoreItemsWithWorkerPool(items []Item, history []string, workers int) []Item {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := range jobs {
				items[i].Score = scoreItem(items[i], history)
			}
		}()
	}

	for i := range items {
		jobs <- i
	}

	close(jobs)
	wg.Wait()

	return items
}

func recommend(items []Item, history []string) []Item {
	result := append([]Item(nil), items...)

	for i := range result {
		result[i].Score = scoreItem(result[i], history)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Score > result[j].Score
	})

	return result
}
