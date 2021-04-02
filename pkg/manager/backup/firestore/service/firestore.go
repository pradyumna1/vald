//
// Copyright (C) 2019-2021 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"
	"reflect"

	"github.com/vdaas/vald/internal/db/nosql/firestore"
	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/pkg/manager/backup/firestore/model"
)

type Firestore interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	GetVector(ctx context.Context, uuid string) (*model.Vector, error)
	GetIPs(ctx context.Context, uuid string) ([]string, error)
	SetVector(ctx context.Context, vec *model.Vector) error
	SetVectors(ctx context.Context, vecs ...*model.Vector) error
	DeleteVector(ctx context.Context, uuid string) error
	DeleteVectors(ctx context.Context, uuids ...string) error
	SetIPs(ctx context.Context, uuid string, ips ...string) error
	RemoveIPs(ctx context.Context, ips ...string) error
}

type client struct {
	firestore firestore.Firestore
}

func New(opts ...Option) (Firestore, error) {
	c := new(client)
	for _, opt := range append(defaultOptions, opts...) {
		if err := opt(c); err != nil {
			return nil, errors.ErrOptionFailed(err, reflect.ValueOf(opt))
		}
	}

	return c, nil
}

func (c *client) Connect(ctx context.Context) error {
	return c.firestore.Open(ctx)
}

func (c *client) Close(ctx context.Context) error {
	return c.firestore.Close()
}

func (c *client) GetVector(ctx context.Context, uuid string) (*model.Vector, error) {
	v := new(model.Vector)
	if err := c.firestore.Get(ctx, uuid, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (c *client) GetIPs(ctx context.Context, uuid string) ([]string, error) {
	v, err := c.GetVector(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return v.IPs, nil
}

func (c *client) SetVector(ctx context.Context, vec *model.Vector) error {
	_, err := c.firestore.Upload(ctx, vec)
	return err
}

func (c *client) SetVectors(ctx context.Context, vecs ...*model.Vector) error {
	return nil
	//return c.firestore.UploadMulti(ctx, vecs...)
}

func (c *client) DeleteVector(ctx context.Context, uuid string) error {
	return c.firestore.Delete(ctx, uuid)
}

func (c *client) DeleteVectors(ctx context.Context, uuids ...string) error {
	return c.firestore.DeleteMulti(ctx, uuids...)
}

func (c *client) SetIPs(ctx context.Context, uuid string, ips ...string) error {
	return nil
}

func (c *client) RemoveIPs(ctx context.Context, ips ...string) error {
	return nil
}
