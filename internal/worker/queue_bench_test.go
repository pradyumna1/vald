package worker

import (
	"context"
	"testing"

	"github.com/vdaas/vald/internal/errors"
)

func Benchmark_Push(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue()
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				_, err := q.Pop(ctx)
				if err == context.Canceled {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Buffer100Push(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueBuffer(100))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				_, err := q.Pop(ctx)
				if err == context.Canceled {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Buffer1000Push(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueBuffer(1000))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				_, err := q.Pop(ctx)
				if err == context.Canceled {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_CheckDur1000msPush(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueCheckDuration("1000ms"))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				_, err := q.Pop(ctx)
				if err == context.Canceled {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_CheckDur10msPush(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueCheckDuration("10ms"))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				_, err := q.Pop(ctx)
				if err == context.Canceled {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Pop(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue()
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	job := func(context.Context) error {
		return nil
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				err := q.Push(ctx, job)
				if err == context.Canceled || (err != nil && errors.Is(err, errors.ErrQueueIsNotRunning())) {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := q.Pop(ctx); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Buffer100Pop(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueBuffer(100))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	job := func(context.Context) error {
		return nil
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				err := q.Push(ctx, job)
				if err == context.Canceled || (err != nil && errors.Is(err, errors.ErrQueueIsNotRunning())) {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := q.Pop(ctx); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_Buffer1000Pop(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueBuffer(1000))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	job := func(context.Context) error {
		return nil
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				err := q.Push(ctx, job)
				if err == context.Canceled || (err != nil && errors.Is(err, errors.ErrQueueIsNotRunning())) {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := q.Pop(ctx); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_CheckDur1000msPop(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueCheckDuration("1000ms"))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	job := func(context.Context) error {
		return nil
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				err := q.Push(ctx, job)
				if err == context.Canceled || (err != nil && errors.Is(err, errors.ErrQueueIsNotRunning())) {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := q.Pop(ctx); err != nil {
				b.Error(err)
			}
		}
	})
}

func Benchmark_CheckDur10msPop(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(func() {
		cancel()
	})
	q, err := NewQueue(WithQueueCheckDuration("10ms"))
	if err != nil {
		b.Fatal(err)
	}
	cerr, err := q.Start(ctx)
	if err != nil {
		b.Fatal(err)
	}

	job := func(context.Context) error {
		return nil
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-cerr:
				if err != nil {
					b.Error(err)
				}
			default:
				err := q.Push(ctx, job)
				if err == context.Canceled || (err != nil && errors.Is(err, errors.ErrQueueIsNotRunning())) {
					return
				}
				if err != nil {
					b.Error(err)
				}
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := q.Pop(ctx); err != nil {
				b.Error(err)
			}
		}
	})
}
