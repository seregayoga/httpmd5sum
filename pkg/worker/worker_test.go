package worker

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunWorkersInParallel(t *testing.T) {
	body := "Cool body of awesome site!"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	argUrls := []string{ts.URL}
	urls := make(chan string, 1)
	for _, url := range argUrls {
		urls <- url
	}
	close(urls)

	results := make(chan *Result, 1)
	Run(10, 1, urls, results)
	result := <-results

	if result.URL != ts.URL {
		t.Fatalf("Expected URL %s, got %s", ts.URL, result.URL)
	}

	hash := md5.New()
	io.WriteString(hash, body+"\n")

	expectedHash := fmt.Sprintf("%x", hash.Sum(nil))
	actualHash := fmt.Sprintf("%x", result.Hash.Sum(nil))
	if expectedHash != actualHash {
		t.Fatalf("Expected hash %s, got %s", expectedHash, actualHash)
	}
}
