# algods

# optimization tricks
  1. use sync.Pool https://pkg.go.dev/sync#Pool
  2. use b2s or s2b for converting bytes to string or string to bytes
  3. https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#bounds_check_elimination
  4. build with - `go build -gcflags=-m example.go` and check malloc and other
