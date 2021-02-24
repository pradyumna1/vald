# Benchmark

## Introduction

### What is the problem now in Vald?

In Vald, performance is a critical requirement and factor for users. If Vald does not perform good and stable performance, user may leave Vald.

So it is important to capture the performance statistic for Vald to perform performance tunning and help to find the bottleneck of Vald when there is some performance issue in Vald.

### How can we resolve this problem?

Benchmarking is used to measure the performance of target elements to produce a performance metrics for future performance investigation.

We can measure the performance statistic in 3 different levels to measure different level of performance:

1. Code benchmarking
  - to capture performance metrics by function level (/internal)
  - to perform lowest level of performance investigation/tuning

1. Code level E2E benchmarking
  - to capture performance metrics by use case level (/pkg)

1. Component level E2E benchmarking
  - to capture performance metrics by component level

Unlike unit testing, benchmark test only focus of measure on the performance of the function but not the functionality of the function.

### What can we do with benchmark?

In benchamrking, we can measure performance using specific indicator and resulting in a metrics of performance.

The metrics can be used in different purpose:

- demonstrate whether or not a system meets the criteria set forward
- compare two applications to determine which one works better
- measure a system to find what performs badly

Here is the example of what to measure in benchmarking:

- How much time consumed for each operation (e.g. time/operation)
- How much memory allocated for each operation (e.g. ??? byte/operation)
- How many times the memory is allocated (e.g. ??? times)

After getting the metrics, we can ask ourself the following questions about the metrics:

- How should this software perform under the maximum load?
- How many concurrent users are expected on a daily basis?
- What does “good” performance look like?
- What does “acceptable” performance look like?

### How should we work on benchmark?

- Work on every functions?
- With template
	- Try few important packages first to test if the template is working
- Without template

### Questions/Concerns

1. Missing code benchmark coverage in golang
1. How to detect the changes on benchmark result? 


## Code bench

### Overview

Code bench is the lowest level of benchmarking in above. It performs the benchmark testing on functions and it provide useful performance statistic.

In golang, it supports to get the benchmark metrics by using `go` command.

We can measure the CPU and the memory usage/allocation information of the target function.

e.g. The result is as follow.

```
Benchmark_Uint32/test_rand-4         	17290003	        82.3 ns/op	       0 B/op	       0 allocs/op
```

Reference: https://golang.org/pkg/testing/#BenchmarkResult

### What to test in code bench?

We need to initialize the object in real world scerieno (e.g. from `New()`)

Unlike unit test, we focus on testing the performance of the function execution, we do not check the output of the execution result.

Therefore we do not need to test the error case as it is meaningless to the benchmark, because we can't get the correct benchmark result in error case.

The metrics of the benchmark is: 

- Average time cost of each execution
- Average memory usage / allocation of each execution
- Parallel execution performance

Reference: https://qiita.com/marnie_ms4/items/8706f43591fb23dd4e64

### Testing data

// Unlike stress test, performance test check the system performance under varying loads, while stress testing is check the system behavior under sudden increased load. 

To design testing data, we need to think about the following things:

- How the parameter affects the benchmark result (data volume, etc)
- How the cache affects the result, to cover if the cache affects the performance metrics
- What is the realistic use case
- We need to test the heavy loading cases

For example, given the following implementation to decode json:

```go
func Decode(r io.Reader, data interface{}) (err error) {
	return jsoniter.NewDecoder(r).Decode(data)
}
```

Base on the above criterias, to think about testing the [json decode performance](https://github.com/vdaas/vald/blob/master/internal/encoding/json/json.go#L29), we should test the following data:

- io.Reader with empty data and data is a struct with 1 field
- io.Reader with 10 data and data is a struct with 1 field
- io.Reader with 100 data and data is a struct with 1 field
- io.Reader with empty data and data is a struct with 10 field
- io.Reader with 10 data and data is a struct with 10 field
- io.Reader with 100 data and data is a struct with 10 field
- io.Reader with 100 data and data is a struct with 50 field

These testing data are focus on measuring the performance of how the size of the data affect the performance of the function.


One more example, given the following implementation to [push element to queue](https://github.com/vdaas/vald/blob/master/internal/worker/queue.go#L116):

```go
func (q *queue) Push(ctx context.Context, job JobFunc) error {
	if job == nil {
		return errors.ErrJobFuncIsNil()
	}

	if !q.isRunning() {
		return errors.ErrQueueIsNotRunning()
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case q.inCh <- job:
		return nil
	}
}
```

Since there are cache to store the queue elements, we should test the number of execution to ensure the performance of queue.

Here is the example of testing data:

- push 1 item to queue
- push 10 items to queue
- push 100 items to queue


### Testing result

The testing result may different in different environment, sso it is relative to different environment.

If we found any problems in the testing result, for example the function requires many memory space or the function takes more time to execute than expected, we should report it to authors and ask if the result is expected or not.

### How to implement bench code?

Create / generate a file called `[filename]_bench_test.go`.

For example if we want to test the `internal/rand/rand.go` file, you need to create / generate a file called `internal/rand/rand_bench_test.go` file.

To execute the benchmarking, use the `go test -bench . -benchmem -cpu=1,2,4,6 -benchtime=5s` command.

### Additional things we need to concern

- We should avoid other factors to affect the benchmark result (e.g. the machine is not in idle status, etc)
  - Avoid overhead from initialize value (function value)
  - Avoid possible compiler optimisation: store the return value from the result.
- Different environment have different settings (e.g. hardware & OS settings/version), so the result may different in different environment
- The benchmark result may not accurate if the test is not getting enough sample to calculate the average
- We may need to consider re-run the benchmark test in sometime later


### What we need to think later?

- Consider writing Makefile to execute bench test in different package with different parameter


### Ref

- https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
- https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html#benchmarking
- https://www.castsoftware.com/glossary/software-performance-benchmarking-modeling#:~:text=Software%20performance%20benchmarking%20serves%20different,to%20find%20what%20performs%20badly.
