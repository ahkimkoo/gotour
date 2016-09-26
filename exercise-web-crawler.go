package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeUrlStore struct {
	links map[string]bool
	mux   sync.Mutex
	wg    sync.WaitGroup
}

func (s *SafeUrlStore) Set(link string) {
	s.mux.Lock()
	s.links[link] = true
	s.mux.Unlock()
}

func (s *SafeUrlStore) Get(link string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.links[link]
	return ok
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, s SafeUrlStore) {
	s.wg.Add(1)
	defer s.wg.Done()
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if ok := s.Get(u); !ok {
			s.Set(u)
			go Crawl(u, depth-1, fetcher, s)
		} else {
			fmt.Printf("\n%s exists\n", u)
		}
	}
	return
}

func main() {
	s := SafeUrlStore{links: make(map[string]bool)}
	Crawl("http://golang.org/", 4, fetcher, s)
	s.wg.Wait()
	time.Sleep(time.Millisecond)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://gloang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
