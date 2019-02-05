package worker

import (
	"crypto/md5"
	"fmt"
	"hash"
)

// Result result of fetching url and hashing
type Result struct {
	URL  string
	Hash hash.Hash
	Err  error
}

// NewResult creates new result
func NewResult(url string) *Result {
	return &Result{
		URL:  url,
		Hash: md5.New(),
	}
}

// String implements Stringer interface
func (r *Result) String() string {
	if r.Err != nil {
		return fmt.Sprintf("%s Error while fetching url: %s", r.URL, r.Err)
	}
	return fmt.Sprintf("%s %x", r.URL, r.Hash.Sum(nil))
}
