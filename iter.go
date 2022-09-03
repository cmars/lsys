package lsys

import (
	"github.com/frankban/iterate"
)

type Iter[T any] struct {
	iter      iterate.Iterator[T]
	transform func(T) []T
	st        []T
}

func (i *Iter[T]) Next() bool {
	return len(i.st) > 0 || i.iter.Next()
}

func (i *Iter[T]) Value() T {
	var val T
	if len(i.st) > 0 {
		val, i.st = i.st[0], i.st[1:]
		return val
	}
	st := i.transform(i.iter.Value())
	val = st[0]
	i.st = append(i.st, st[1:]...)
	return val
}

func (i *Iter[T]) Err() error {
	return nil
}

func Seed[T any](seed []T, transform func(T) []T) *Iter[T] {
	return &Iter[T]{
		iter:      iterate.FromSlice(seed),
		transform: transform,
	}
}

func Generate[T any](iter *Iter[T]) *Iter[T] {
	return &Iter[T]{
		iter:      iter,
		transform: iter.transform,
	}
}
