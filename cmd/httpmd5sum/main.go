package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/seregayoga/httpmd5sum/pkg/worker"
)

const defaultParallel = 10

func main() {
	parallel, argURLs, err := argsParse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing args: %s", err)
	}

	urls := make(chan string, len(argURLs))
	for _, url := range argURLs {
		urls <- url
	}
	close(urls)

	results := make(chan *worker.Result, len(argURLs))

	worker.Run(parallel, len(argURLs), urls, results)

	for result := range results {
		fmt.Println(result)
	}
}

func argsParse() (int, []string, error) {
	var parallel = flag.Int("parallel", defaultParallel, "number of parallel requests")
	flag.Parse()
	if *parallel <= 0 {
		*parallel = defaultParallel
	}

	argURLs := []string{}
	for _, arg := range flag.Args() {
		u, err := url.Parse(arg)
		if err != nil {
			return 0, nil, err
		}

		url := arg
		if u.Scheme == "" {
			url = "http://" + url
		}

		argURLs = append(argURLs, url)
	}

	return *parallel, argURLs, nil
}
