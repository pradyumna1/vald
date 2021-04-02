package firestore

import "google.golang.org/api/option"

type Option func(*firestoreClient) error

var defaultOptions = []Option{}

func WithProjectID(pid string) Option {
	return func(client *firestoreClient) error {
		client.projectID = pid
		return nil
	}
}

func WithClientOptions(opts ...option.ClientOption) Option {
	return func(client *firestoreClient) error {
		client.clientOpts = opts
		return nil
	}
}

func WithCollectionPath(c string) Option {
	return func(client *firestoreClient) error {
		client.collectionPath = c
		return nil
	}
}
