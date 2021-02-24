package worker

import (
	"context"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/vdaas/vald/internal/errgroup"
)

func Benchmark_Push(b *testing.B) {
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
	type test struct {
		name      string
		fields    fields
		args      args
		times     int
		result    result
		initFunc  func(test) *queue
		afterFunc func(test)
	}

	defaultInitFunc := func(test test) *queue {
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
	tests := []test{
		{
			name: "test push 1 element",
			fields: fields{
				buffer: 10,
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
			times: 1,
		},
		{
			name: "test push 10 element",
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
			times: 10,
		},
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
			times: 100,
		},
	}
	for _, tc := range tests {
		test := tc

		if test.initFunc == nil {
			test.initFunc = defaultInitFunc
		}
		if test.afterFunc != nil {
			defer test.afterFunc(test)
		}
		if test.times == 0 {
			test.times = 1
		}

		q := test.initFunc(test)
		b.Run(test.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for j := 0; j < test.times; j++ {
					test.result.err = q.Push(test.args.ctx, test.args.job)
				}
			}
		})
	}
	b.ResetTimer()
}

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
