package cache

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type Provider interface {
	Set(ctx context.Context, key string, ttl time.Duration, data []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
}

type SerializationType uint8

const (
	SerializationTypeJSON = iota
)

type ObjectCacher[T any] struct {
	prefix            string
	provider          Provider
	serializationType SerializationType
}

func NewObjectCacher[T any](prefix string, p Provider, st SerializationType) *ObjectCacher[T] {
	return &ObjectCacher[T]{
		prefix:            prefix,
		provider:          p,
		serializationType: st,
	}
}

func (c *ObjectCacher[T]) createKey(k string) string {
	return c.prefix + "." + k
}

func NewJsonObjectCacher[T any](prefix string, p Provider) *ObjectCacher[T] {
	return NewObjectCacher[T](prefix, p, SerializationTypeJSON)
}

func (c *ObjectCacher[T]) Get(ctx context.Context, key string) (T, error) {
	t := new(T)
	data, err := c.provider.Get(ctx, c.createKey(key))
	if err != nil {
		if errors.Is(err, ErrCacheMiss) {
			return *t, nil
		}
		return *t, err
	}

	return *t, c.unmarshal(data, &t)
}

func (c *ObjectCacher[T]) Del(ctx context.Context, key string) error {
	return c.provider.Del(ctx, c.createKey(key))
}

func (c *ObjectCacher[T]) Set(ctx context.Context, key string, ttl time.Duration, in T) error {
	data, err := c.Marshal(in)
	if err != nil {
		return err
	}

	return c.provider.Set(ctx, c.createKey(key), ttl, data)
}
