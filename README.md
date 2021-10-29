# algods

# optimization tricks
  1. use sync.Pool https://pkg.go.dev/sync#Pool
  2. use b2s or s2b for converting bytes to string or string to bytes
  3. https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#bounds_check_elimination
  4. build with - `go build -gcflags=-m example.go` and check malloc and other or `go build -gcflags "-m -m" example.go` (max 4 -m flags)
  5. build with - `go build -gcflags='-l -l -l' example.go` - to make aggressive inlines (also can use `-l -l` and dont use `-l`)
  6. https://habr.com/ru/company/badoo/blog/301990/
# more
+ add algo Mo (https://www.geeksforgeeks.org/mos-algorithm-query-square-root-decomposition-set-1-introduction/)
+ add ds treap (https://e-maxx.ru/algo/treap)
