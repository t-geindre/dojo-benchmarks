# ⏱️ Benchmarks Dojo

Let's write some benchmarks!

## 🔢 Fibonacci

As a code example that can run very slowly, we will use the Fibonacci sequence.

> The Fibonacci sequence is a sequence in which each number is the sum of the two preceding ones.
> It starts with 0 and 1, and the next number is the sum of the previous two.
> The sequence starts like this: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55 ...

An implementation of the Fibonacci sequence is already provided in the `fibonacci.go` file.
It also comes with a test to ensure that the implementation is correct, at least for the first numbers.

## 🧪 Try it!

```shell
$ go build
$ ./dojo
```

The `main.go` file contains a call to the `Fibonacci` function with the number 10.
The 11th number (index `10`) in the Fibonacci sequence is 55, so the output should be `55`.

You can now try to run the program with higher numbers and see how long it takes to compute them.

## ⏱️ Benchmark

With high numbers, the Fibonacci sequence can take a long time to compute.

Before trying to optimize the implementation, we need to measure how long it takes to compute the numbers so we can compare the different implementations.

📝 Write a benchmark for the `Fibonacci` function in the `fibonacci_test.go` file.

```go
func BenchmarkFibonacci(b *testing.B) {
    // Do something
}
```

This benchmark can be done with a fairly high number or with multiple numbers to see how the time grows.

```go
func BenchmarkFibonacci(b *testing.B) {
	for _, idx := range []int{5, 10, 15, 20, 25} {
		b.Run(fmt.Sprintf("Fibonacci(%d)", idx), func(b *testing.B) {
			// Do something
		})
	}
}
```

Check this [Golang documentation](https://pkg.go.dev/testing#hdr-Benchmarks) to know how to write a relevant benchmark.

> 💡 **Tips**: In the given benchmarks examples, there is more to add than just a call to the `Fibonacci` function.


Once you've written the benchmark, you can run it with the following command:

```shell
$ go test -bench=. ./...
```

For a more accurate result, you can run the benchmark multiple times:

```shell
$ go test -bench=. -count=10 ./...
```

📝 Save your benchmark results in a file to compare them later.

```shell
$ go test -bench=. -count=10 ./... > benchmarks-slow.txt
```

## 💪 Improve

Now is the time to try to improve the implementation of the Fibonacci sequence. To do so, we need to understand why the current implementation is slow.

Here is a simple explanation of the problem:
![Fibonacci tree](doc/fib-tree.jpg)

We use a recursive implementation, which means that it will compute the same numbers multiple times.
And this tree grows exponentially with the number we want to compute.

To solve this problem, we can use a technique called "[memoization](https://en.wikipedia.org/wiki/Memoization)".
It consists of storing the results of the computations to avoid recomputing them.

📝 Improve Fibonacci sequence implementation using memoization in the `fibonacci.go` file.

> 💡 **Tips**: You can use a map to store the results of the computations.
> ```go
> var memo = map[int]int{}
> 
> func Fibonacci(index int) int {
> }
> ```

## ⏱️ Benchmark again

Now that you have improved the implementation, you can run the benchmark again to see if it is faster.

```shell
$ go test -bench=. -count=10 ./...
```

📝 Save your new benchmark results in a file to compare them with the previous ones.

```shell
$ go test -bench=. -count=10 ./... > benchmarks-fast.txt
```

## 🆚 Compare

Now is the time to compare the results of the benchmarks. To do so, we'll use the `benchstat` tool.

📝 Install the `benchstat` tool:

```shell
$ go install golang.org/x/perf/cmd/benchstat@latest
```

📝 Compare the two benchmark results files:

```shell
$ benchstat benchmarks-slow.txt benchmarks-fast.txt
```

The `benchstat` tool will show you the difference between the two benchmarks.