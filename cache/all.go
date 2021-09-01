package cache

func AllCached(f func(int) int) func(int) int {
	var cacheMap = make(map[int]int)
	decorated := func(arg int) int {
		if val, ok := cacheMap[arg]; ok {
			return val
		}
		val := f(arg)
		cacheMap[arg] = val
		return val
	}
	return decorated
}
