# Language

## Concurrency in Go

* Concurrency in Go is done via Goroutines
* Goroutines are simple functions running as lightweight
  processes in a thread pool
* The runtime allows a high number of those Goroutines and
  provides very fast context switches
* Communication via memory is possible but intended to be
  done via typed channels
* Multiples channels can be accessed at the same time via
  `select` statements
* A Goroutine is startet via the `go` statement
* No reference, ID, or instance is returned

## Quote

_In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations._

_Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once._

-- Rob Pike
