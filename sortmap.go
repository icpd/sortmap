package sortmap

import (
	"bytes"
	"encoding/json"
	"sort"
)

type pair[Key comparable] struct {
	key   Key
	value any
}

func (kv *pair[Key]) Key() Key {
	return kv.key
}

func (kv *pair[Key]) Value() any {
	return kv.value
}

type ByPair[Key comparable] struct {
	Pairs    []*pair[Key]
	LessFunc func(a *pair[Key], j *pair[Key]) bool
}

func (a ByPair[Key]) Len() int           { return len(a.Pairs) }
func (a ByPair[Key]) Swap(i, j int)      { a.Pairs[i], a.Pairs[j] = a.Pairs[j], a.Pairs[i] }
func (a ByPair[Key]) Less(i, j int) bool { return a.LessFunc(a.Pairs[i], a.Pairs[j]) }

type sortMap[Key comparable] struct {
	keys       []Key
	values     map[Key]any
	escapeHTML bool
}

func New[Key comparable]() *sortMap[Key] {
	return &sortMap[Key]{
		keys:       []Key{},
		values:     make(map[Key]any),
		escapeHTML: true,
	}
}

func (o *sortMap[Key]) SetEscapeHTML(on bool) {
	o.escapeHTML = on
}

func (o *sortMap[Key]) Get(key Key) (any, bool) {
	val, exists := o.values[key]
	return val, exists
}

func (o *sortMap[Key]) Set(key Key, value any) {
	_, exists := o.values[key]
	if !exists {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

func (o *sortMap[Key]) Delete(key Key) {
	_, ok := o.values[key]
	if !ok {
		return
	}

	// remove from keys
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			break
		}
	}

	// remove from values
	delete(o.values, key)
}

func (o *sortMap[Key]) Keys() []Key {
	return o.keys
}

// SortKeys Sort the map keys using your sort func
func (o *sortMap[Key]) SortKeys(sortFunc func(keys []Key)) {
	sortFunc(o.keys)
}

// Sort the map using your sort func
func (o *sortMap[Key]) Sort(lessFunc func(a *pair[Key], b *pair[Key]) bool) {
	pairs := make([]*pair[Key], len(o.keys))
	for i, key := range o.keys {
		pairs[i] = &pair[Key]{key, o.values[key]}
	}

	sort.Sort(ByPair[Key]{pairs, lessFunc})

	for i, pair := range pairs {
		o.keys[i] = pair.key
	}
}

func (o sortMap[Key]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(o.escapeHTML)
	for i, k := range o.keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		// add key
		if err := encoder.Encode(k); err != nil {
			return nil, err
		}
		buf.WriteByte(':')
		// add value
		if err := encoder.Encode(o.values[k]); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}