package worker

import (
	"io"
	"net/http"
	"sync"
)

// DownloadWorker creates worker that reads urls and writes to results
func DownloadWorker(wg *sync.WaitGroup, urls <-chan string, results chan<- *Result) {
	for url := range urls {
		result := NewResult(url)

		var resp *http.Response
		resp, result.Err = http.Get(url)
		if result.Err == nil {
			io.Copy(result.Hash, resp.Body)
		}

		results <- result

		wg.Done()
	}
}

// Run runs workers in goroutines
func Run(parallel int, urlsNum int, urls <-chan string, results chan<- *Result) {
	wg := &sync.WaitGroup{}
	wg.Add(urlsNum)
	for p := 0; p < parallel; p++ {
		go DownloadWorker(wg, urls, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}
