package main

import (
	"reflect"
	"testing"
)

func TestRepository(t *testing.T) {
	repo, err := LoadData("data.example.json")
	if err != nil {
		t.Fatal(err)
	}

	referenceRepo := &Repository{
		Bucket: map[string][]Record{
			"bucket_1": {
				{PhotoURL: "https://a.com/photo1", Description: "Photo 1"},
				{PhotoURL: "https://a.com/photo2", Description: "Photo 2"},
			},
		},
		BucketSize: map[string]int{
			"bucket_1": 2,
		},
	}

	if !reflect.DeepEqual(repo, referenceRepo) {
		t.Fatalf("referenceRepo != repo (%+v)", repo)
	}
}

func TestNonExistingRepository(t *testing.T) {
	repo, err := LoadData("non_existing_data.json")
	if err == nil {
		t.Fatal("LoadData supposed to return error")
	}
	if !reflect.DeepEqual(repo, &Repository{}) {
		t.Fatal("loaded repository supposed to be equal to default repository")
	}
}
