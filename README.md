# algods
Данный репозиторий служит шпаргалкой на тренировках по CP, а так же неким наглядным пособием для лекций по алгоритмам

# optimization tricks
  1. https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#bounds_check_elimination
  2. build with - `go build -gcflags=-m example.go` and check malloc and other or `go build -gcflags "-m -m" example.go` (max 4 -m flags)
  3. build with - `go build -gcflags='-l -l -l' example.go` - to make aggressive inlines (also can use `-l -l` and dont use `-l`)
  4. https://habr.com/ru/company/badoo/blog/301990/
# more
+ add ds AVL tree
+ https://medium.com/cheat-sheets/cheat-sheet-for-competitive-programming-with-c-f2e8156d5aa9
+ https://github.com/helloproclub/competitive-programming-cheat-sheet

# run
go
`cat i.txt | go run main.go > o.txt
`

c++
`g++ -std=c++20 -O2 -lm -o x.bin main.cpp && chmod +x ./x.bin | cat i.txt | ./x.bin > o.txt`