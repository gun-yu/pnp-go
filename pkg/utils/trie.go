package utils

import (
	"pnp-go/pkg/utils/trie_internal"
	"strings"
)

type Trie[T any] struct {
	inner *trie_internal.TrieInternal[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{inner: trie_internal.New[T]()}
}

func (t *Trie[T]) key(key string) string {
	p := NormalizePath(key)

	if !strings.HasSuffix(p, "/") {
		return p + "/"
	}

	return p
}

func (t *Trie[T]) GetAncestorValue(p string) (*T, bool) {
	v, ok := t.inner.GetAncestorValue(t.key(p))
	return &v, ok
}

func (t *Trie[T]) Insert(p string, v T) {
	t.inner.Set(t.key(p), v)
}
