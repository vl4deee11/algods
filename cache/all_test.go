package cache

import (
	"math/rand"
	"testing"
	"time"
)

func TestAllCacheHits(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	hitMap := make(map[int]int)
	maxArg := 100
	numTries := 1000
	for i := 0; i < maxArg; i++ {
		hitMap[i] = 0
	}
	f := func(arg int) int {
		hitMap[arg]++
		return arg * 2 + 1
	}
	decorated := AllCached(f)
	for i := 0; i < numTries; i++ {
		arg := rand.Intn(maxArg)
		decorated(arg)
	}
	for i := 0; i < maxArg; i++ {
		if val, _ := hitMap[i]; val > 1 {
			t.Errorf("Expected less than 1 call to func, got %d", val)
		}
	}
}

func TestAllCacheValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	resultMap := make(map[int]int)
	maxArg := 100
	numTries := 10000
	f := func(arg int) int {
		return arg * 2 + 1
	}
	for i := 0; i < maxArg; i++ {
		resultMap[i] = f(i)
	}
	decorated := AllCached(f)
	for i := 0; i < numTries; i++ {
		arg := rand.Intn(maxArg)
		val := decorated(arg)
		if val != resultMap[arg] {
			t.Errorf("Expected value %d, got %d (from call f(%d))", resultMap[arg], val, arg)
		}
	}
}

func TestAllDecoratorIndependence(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	maxArg := 100
	numTries := 10000

	f := func(arg int) int {
		return arg * 2 + 1
	}
	g := func(arg int) int {
		return arg * arg * 3 - 20
	}
	decoratedF := AllCached(f)
	decoratedG := AllCached(g)
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
