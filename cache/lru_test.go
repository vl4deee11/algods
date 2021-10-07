package cache

import (
	"math/rand"
	"testing"
	"time"
)

func testLruCacheHits(t *testing.T, caseSlice []int, estimatedHits int, cacheSize int) {
	// simple case
	callCount := 0
	f := func(arg int) int {
		callCount++
		return arg*2 + 1
	}
	decorated := LruCached(f, cacheSize)
	for _, c := range caseSlice {
		decorated(c)
	}
	cacheMisses := len(caseSlice) - estimatedHits
	if callCount != cacheMisses {
		t.Errorf("Expected %d cache misses, got %d", cacheMisses, callCount)
	}
}

func TestLruCacheHitsSimple(t *testing.T) {
	// simple case
	testLruCacheHits(t, []int{1, 2, 3, 1}, 1, 3)
	// no hits
	testLruCacheHits(t, []int{1, 2, 3, 4, 5, 6, 7}, 0, 3)
	// front hit
	testLruCacheHits(t, []int{1, 2, 3, 3}, 1, 3)
	// mid hit
	testLruCacheHits(t, []int{1, 2, 3, 2}, 1, 3)
	// little cache
	testLruCacheHits(t, []int{1, 2, 3}, 0, 1)
	// hit without fill
	testLruCacheHits(t, []int{1, 1, 2}, 1, 3)
	testLruCacheHits(t, []int{1, 2, 1, 2}, 2, 3)
	// multiple hits
	testLruCacheHits(t, []int{1, 3, 1, 1, 2}, 2, 3)

}

func TestLruCacheValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	resultMap := make(map[int]int)
	maxArg := 100
	numTries := 100000
	f := func(arg int) int {
		return arg*2 + 1
	}
	for i := 0; i < maxArg; i++ {
		resultMap[i] = f(i)
	}
	decorated := LruCached(f, 5)
	for i := 0; i < numTries; i++ {
		arg := rand.Intn(maxArg)
		val := decorated(arg)
		if val != resultMap[arg] {
			t.Errorf("Expected value %d, got %d (from call f(%d))", resultMap[arg], val, arg)
		}
	}
}

func TestLruDecoratorIndependence(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	maxArg := 100
	numTries := 10000

	f := func(arg int) int {
		return arg*2 + 1
	}
	g := func(arg int) int {
		return arg*arg*3 - 20
	}
	decoratedF := LruCached(f, 5)
	decoratedG := LruCached(g, 3)
	for i := 0; i < numTries; i++ {
		arg := rand.Intn(maxArg)
		valF := decoratedF(arg)
		valG := decoratedG(arg)
		trueF := f(arg)
		trueG := g(arg)
		if valF != trueF {
			t.Errorf("Expected value %d, got %d (from call f(%d))", trueF, valF, arg)
		}
		if valG != trueG {
			t.Errorf("Expected value %d, got %d (from call f(%d))", trueG, valG, arg)
		}
	}
}
