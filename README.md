# algods
Основные материалы собранные в этом репозитории взяты с сайта [e-maxx.ru](http://e-maxx.ru), переписаны на golang.
Данный репозиторий служит шпаргалкой на тренировках по CP, а так же неким наглядным пособием для лекций по алгоритмам
# optimization tricks
  1. use sync.Pool https://pkg.go.dev/sync#Pool
  2. use b2s or s2b for converting bytes to string or string to bytes
  3. https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#bounds_check_elimination
  4. build with - `go build -gcflags=-m example.go` and check malloc and other or `go build -gcflags "-m -m" example.go` (max 4 -m flags)
  5. build with - `go build -gcflags='-l -l -l' example.go` - to make aggressive inlines (also can use `-l -l` and dont use `-l`)
  6. https://habr.com/ru/company/badoo/blog/301990/
# more
+ add ds AVL tree
+ add algo prim http://e-maxx.ru/algo/mst_prim
+ add algo kruskal http://e-maxx.ru/algo/mst_kruskal_with_dsu
