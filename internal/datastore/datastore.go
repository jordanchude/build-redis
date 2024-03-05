package datastore

type datastore struct {
	store map[string]string
}

func newDatastore() *Datastore {
	return &Datastore {
		store: make(make[string]string)
	}
}