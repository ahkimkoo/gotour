package main

import (
	"fmt"
	"github.com/huichen/sego"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")

	text := []byte("我被青春撞了一下腰")
	segments := segmenter.Segment(text)
	fmt.Println(sego.SegmentsToString(segments, false))
}
