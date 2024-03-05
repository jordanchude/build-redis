package datastore

import (
	"testing"

	. "github.com/jordanchude/build-redis/internal/datastore"
)

func TestSet(t *testing.T) {
	ds := NewDatastore()
	ds.Set("key", "value")

	if ds.store["key"] != "value" {
		t.Error("Set failed to set value in store")
	}
}
