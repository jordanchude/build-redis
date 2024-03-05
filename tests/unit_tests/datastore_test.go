package datastore

import (
	"build-redis/internal/datastore"
	"testing"
)

func TestSet(t *testing.T) {
	ds := datastore.NewDatastore()
	ds.Set("key", "value")

	if ds.Get("key") != "value" {
		t.Error("Set failed to set value in store")
	}
}
