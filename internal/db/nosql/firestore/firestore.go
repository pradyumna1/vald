package firestore

import (
	"context"
	"reflect"

	"cloud.google.com/go/firestore"
	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/log"
	"google.golang.org/api/option"
)

type Data interface {
	ID() string
	Data() interface{}
}

type (
	WriteResult *firestore.WriteResult

	UploadResult struct {
		Res WriteResult
	}
)

type Firestore interface {
	Open(context.Context) error
	Close() error
	Get(ctx context.Context, id string, d interface{}) error
	Upload(context.Context, Data) (*UploadResult, error)
	UploadMulti(ctx context.Context, datas ...Data) error
	Delete(ctx context.Context, id string) error
	DeleteMulti(ctx context.Context, ids ...string) error
}

// firestoreClient represent a firestore client for a single collection
type firestoreClient struct {
	projectID      string
	clientOpts     []option.ClientOption
	collectionPath string

	client     *firestore.Client
	collection *firestore.CollectionRef
}

func New(opts ...Option) (Firestore, error) {
	c := &firestoreClient{}
	for _, opt := range append(defaultOptions, opts...) {
		if err := opt(c); err != nil {
			werr := errors.ErrOptionFailed(err, reflect.ValueOf(opt))

			e := new(errors.ErrCriticalOption)
			if errors.As(err, &e) {
				log.Error(werr)
				return nil, werr
			}
			log.Warn(werr)
		}
	}
	return c, nil
}

func (f *firestoreClient) Open(ctx context.Context) error {
	var err error
	f.client, err = firestore.NewClient(ctx, f.projectID, f.clientOpts...)
	if err != nil {
		return err
	}
	f.collection = f.client.Collection(f.collectionPath)
	return nil
}

func (f *firestoreClient) Close() error {
	return f.client.Close()
}

func (f *firestoreClient) Get(ctx context.Context, id string, d interface{}) error {
	dnap, err := f.collection.Doc(id).Get(ctx)
	if err != nil {
		return err
	}
	return dnap.DataTo(d)
}

func (f *firestoreClient) Upload(ctx context.Context, data Data) (*UploadResult, error) {
	r, err := f.collection.Doc(data.ID()).Set(ctx, data)
	if err != nil {
		return nil, err
	}
	return &UploadResult{
		Res: r,
	}, nil
}

func (f *firestoreClient) UploadMulti(ctx context.Context, datas ...Data) error {
	return nil
}

func (f *firestoreClient) Delete(ctx context.Context, id string) error {
	_, err := f.collection.Doc(id).Delete(ctx)
	return err
}

func (f *firestoreClient) DeleteMulti(ctx context.Context, ids ...string) error {
	return nil
}
