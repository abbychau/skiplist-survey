package main

import (
	"time"

	seaSkiplist "github.com/abbychau/jumplist"
)

func abbyInserts(n int) {
	list := seaSkiplist.New()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}
}

func abbyWorstInserts(n int) {
	list := seaSkiplist.New()
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}
}

func abbyAvgSearch(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(float64(i))
	}
}

func abbySearchEnd(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(float64(n))
	}
}

func abbyDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Del(float64(i))
	}
}

func abbyWorstDelete(n int) {
	list := seaSkiplist.New()

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Del(float64(n - i))
	}
}

var abbyFunctions = []func(int){abbyInserts, abbyWorstInserts,
	abbyAvgSearch, abbySearchEnd, abbyDelete, abbyWorstDelete}
