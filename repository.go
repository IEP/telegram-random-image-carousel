package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Record of data
type Record struct {
	PhotoURL    string `json:"photo_url"`
	Description string `json:"description"`
}

// Bucket to contain the records with the bucket name as the key
type Bucket map[string][]Record

// Repository to store bucket information
type Repository struct {
	Bucket     Bucket
	BucketSize map[string]int
}

// LoadData that will be used as the repository
func LoadData(filename string) (*Repository, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &Repository{}, err
	}
	defer f.Close()

	content, _ := ioutil.ReadAll(f)
	var b Bucket
	_ = json.Unmarshal(content, &b)

	repo := Repository{Bucket: b, BucketSize: make(map[string]int)}
	for k, v := range repo.Bucket {
		repo.BucketSize[k] = len(v)
	}

	return &repo, nil
}
