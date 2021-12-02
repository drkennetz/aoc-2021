# AoC 2021 in Golang

## Structure and Testing

The repository is structured like this:

```
./inputs
  ./day[xx]
    day[xx].txt
./src
  ./day[xx]
    day[xx].go
    day[xx]_test.go
```

It is designed so that you can just go into the directory and run the tests.

Example from repository root:
```
C:\Users\dennis\aoc-2021>go test ./src/day01/. --bench=. --benchtime=100000x
goos: windows
goarch: amd64
pkg: github.com/drkennetz/aoc-2021/src/day01
BenchmarkPart1-12         100000              2160 ns/op
BenchmarkPart2-12         100000             17140 ns/op
PASS
ok      github.com/drkennetz/aoc-2021/src/day01 1.966s
```

## Day 1 Results
```
C:\Users\drken\aoc-2021\src\day01>go test . -bench=. -benchtime=100000x
goos: windows
goarch: amd64
pkg: github.com/drkennetz/aoc-2021/src/day01
BenchmarkPart1-12                 100000              2101 ns/op
BenchmarkPart1Refactor-12         100000              1040 ns/op
BenchmarkPart2-12                 100000             17080 ns/op
BenchmarkPart2Refactor-12         100000              1084 ns/op
PASS
ok      github.com/drkennetz/aoc-2021/src/day01 2.172s
```

## Day 2 Results
```
C:\Users\drken\aoc-2021\src\day02>go test . -bench=. -benchtime=100000x
goos: windows
goarch: amd64
pkg: github.com/drkennetz/aoc-2021/src/day02
BenchmarkPart1-12         100000              1310 ns/op
BenchmarkPart2-12         100000              1420 ns/op
PASS
ok      github.com/drkennetz/aoc-2021/src/day02 0.309s

```
