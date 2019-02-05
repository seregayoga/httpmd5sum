# httpmd5sum

## Tool for computing md5 hashes of sites' content

Installation:
`go get github.com/seregayoga/httpmd5sum/cmd/httpmd5sum`

Tests:
`make test` or `go test -race ./...`

Usage:

```
$ httpmd5sum -parallel 15 africa.com https://github.com http://list.ru/ https://wordpress.org https://www.last.fm/
http://list.ru/ 397cba143eabc302ad3628cdebee0c6e
https://www.last.fm/ 8d562f42de351a017734865050e03870
https://github.com b77e616a0f8d0d6bcef415402d0b1fae
https://wordpress.org cd31c07a0383d59b5656bcdd1822f8df
http://africa.com da39b0f1573cd4781ffaf5ed607c0206
```
where `-parallel 15` is number of concurrent requests
