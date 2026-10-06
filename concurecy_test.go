package main

import (
	"sync"
	"testing"
)

func TestAddRequestConcurrent(t *testing.T) {
	var wg sync.WaitGroup

	results := make([]error, 100)
	u := createUser("Ali")
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = u.addRequest()
		}(i)
	}

	wg.Wait()

	wanted := 0

	for _, err := range results {
		if err == nil {
			wanted++
		}
	}

	if wanted != 60 {
		t.Errorf("прошло %d запросов - ожидали 60", wanted)
	}

}
