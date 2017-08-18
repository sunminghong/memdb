package memdb

import "github.com/nedscode/memdb/persist"

// Storer provides the functionality of a memdb store.
type Storer interface {
	Indexer
	SetIndexer(indexer Indexer)
	SetComparator(comparator Comparator)
	SetExpirer(expirer Expirer)
	SetFielder(fielder Fielder)

	PrimaryKey(fields ...string) *Store
	CreateIndex(fields ...string) *Store
	Unique() *Store
	Reversed(order ...bool) *Store

	Persistent(persister persist.Persister) error

	Get(search interface{}) interface{}
	Put(item interface{}) (interface{}, error)
	PutAll(items []interface{}) error
	Delete(search interface{}) (interface{}, error)

	In(fields ...string) IndexSearcher
	Info(cb InfoIterator)
	Ascend(cb Iterator)
	AscendStarting(at interface{}, cb Iterator)
	Descend(cb Iterator)
	DescendStarting(at interface{}, cb Iterator)

	Expire() int

	Len() int
	Indexes() [][]string
	Keys(fields ...string) []string

	On(event Event, notify NotifyFunc)
}
