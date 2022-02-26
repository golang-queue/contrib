# contrib

[![Run Tests](https://github.com/golang-queue/contrib/actions/workflows/go.yml/badge.svg)](https://github.com/golang-queue/contrib/actions/workflows/go.yml)

## Logger Interface

```go
type Logger interface {
  Infof(format string, args ...interface{})
  Errorf(format string, args ...interface{})
  Fatalf(format string, args ...interface{})
  Info(args ...interface{})
  Error(args ...interface{})
  Fatal(args ...interface{})
}
```

* [x] zerolog

## Example

zerolog

```go
package main

import (
  "context"
  "fmt"
  "time"

  "github.com/golang-queue/contrib/zerolog"
  "github.com/golang-queue/queue"
)

func main() {
  taskN := 100
  rets := make(chan string, taskN)

  // initial queue pool
  q := queue.NewPool(5, queue.WithLogger(zerolog.New()))
  // shutdown the service and notify all the worker
  // wait all jobs are complete.
  defer q.Release()

  // assign tasks in queue
  for i := 0; i < taskN; i++ {
    go func(i int) {
      if err := q.QueueTask(func(ctx context.Context) error {
        rets <- fmt.Sprintf("Hi Gopher, handle the job: %02d", +i)
        return nil
      }); err != nil {
        panic(err)
      }
    }(i)
  }

  // wait until all tasks done
  for i := 0; i < taskN; i++ {
    fmt.Println("message:", <-rets)
    time.Sleep(20 * time.Millisecond)
  }
}
```
