package worker

import (
	"context"
	"testing"
)

func Benchmark_Push1JobPop1Job(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
			if j, err := q.Pop(ctx); err != nil {
				b.Error(err)
			} else if j == nil {
				b.Error("job is nil")
			}
		}
	})
}
func Benchmark_PrePush100JobPush1JobPop1Job(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	for i := 0; i < 100; i++ {
		if err := q.Push(ctx, job); err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
			if j, err := q.Pop(ctx); err != nil {
				b.Error(err)
			} else if j == nil {
				b.Error("job is nil")
			}
		}
	})
}

func Benchmark_Buffer100PrePush100JobPush1JobPop1Job(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	for i := 0; i < 100; i++ {
		if err := q.Push(ctx, job); err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := q.Push(ctx, job); err != nil {
				b.Error(err)
			}
			if j, err := q.Pop(ctx); err != nil {
				b.Error(err)
			} else if j == nil {
				b.Error("job is nil")
			}
		}
	})
}

func Benchmark_Buffer100PrePush100JobPop1Job(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
			}
		}
	}()

	job := func(context.Context) error {
		return nil
	}

	for i := 0; i < 10000; i++ {
		if err := q.Push(ctx, job); err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if j, err := q.Pop(ctx); err != nil {
				b.Error(err)
			} else if j == nil {
				b.Error("job is nil")
			}
		}
	})
}

/*
func Benchmark_1Buffer100PrePush100JobPop1Job(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	q := queue{buffer: 1,
		eg:    errgroup.Get(),
		qcdur: 100 * time.Microsecond,
		inCh:  make(chan JobFunc, 10000),
		outCh: make(chan JobFunc, 1),
		qLen: func() (v atomic.Value) {
			v.Store(uint64(0))
			return v
		}(),
		running: func() (v atomic.Value) {
			v.Store(false)
			return v
		}(),
	}
	fmt.Println("push start")

	job := func(context.Context) error {
		return nil
	}

	for i := 0; i < 10000; i++ {
		q.inCh <- job
	}
	fmt.Println("push finish")
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
			}
		}
	}()
	var i int64
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Printf("%d¥n", atomic.AddInt64(&i, 1))
			if j, err := q.Pop(ctx); err != nil {
				b.Error(err)
			} else if j == nil {
				b.Error("job is nil")
			}
		}
	})
}
*/
/*
func Benchmark_Parallel_Push(b *testing.B) {
	type fields struct {
		buffer  int
		eg      errgroup.Group
		qcdur   time.Duration
		inCh    chan JobFunc
		outCh   chan JobFunc
		qLen    atomic.Value
		running atomic.Value
	}
	type args struct {
		ctx context.Context
		job JobFunc
	}
	type result struct {
		err error
	}
	type paralleTest struct {
		name      string
		fields    fields
		args      args
		times     int
		result    result
		initFunc  func(paralleTest) *queue
		afterFunc func(paralleTest)
		parallel  []int
	}
	defaultInitPFunc := func(test paralleTest) *queue {
		q := &queue{
			buffer:  test.fields.buffer,
			eg:      test.fields.eg,
			qcdur:   test.fields.qcdur,
			inCh:    test.fields.inCh,
			outCh:   test.fields.outCh,
			qLen:    test.fields.qLen,
			running: test.fields.running,
		}
		q.Start(context.Background())
		return q
	}
	ptests := []paralleTest{
		{
			name: "test push 100 element",
			fields: fields{
				buffer: 1,
				eg:     errgroup.Get(),
				qcdur:  100 * time.Microsecond,
				inCh:   make(chan JobFunc, 10),
				outCh:  make(chan JobFunc, 1),
				qLen: func() (v atomic.Value) {
					v.Store(uint64(0))
					return v
				}(),
				running: func() (v atomic.Value) {
					v.Store(false)
					return v
				}(),
			},
			args: args{
				ctx: context.Background(),
				job: func(context.Context) error {
					return nil
				},
			},
			times:    100,
			parallel: []int{1, 2, 4, 6, 8, 16},
		},
	}
	b.ResetTimer()

	for _, ptest := range ptests {
		test := ptest
		for _, p := range test.parallel {
			name := test.name + "-" + strconv.Itoa(p)
			b.Run(name, func(b *testing.B) {
				b.SetParallelism(p)

				if test.initFunc == nil {
					test.initFunc = defaultInitPFunc
				}
				if test.afterFunc != nil {
					defer test.afterFunc(test)
				}
				if test.times == 0 {
					test.times = 1
				}

				q := test.initFunc(test)
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						for j := 0; j < test.times; j++ {
							test.result.err = q.Push(test.args.ctx, test.args.job)
						}
					}
				})
			})
		}
	}
}
*/
