package datastore

type Datastore struct {
	store map[string]string
}

func NewDatastore() *Datastore {
	return &Datastore{
		store: make(map[string]string),
	}
}

func (ds *Datastore) Set(key, value string) {
	ds.store[key] = value
}

func (ds *Datastore) Get(key string) string {
	return ds.store[key]
}
