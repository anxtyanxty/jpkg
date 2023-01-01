package main

import (
	"fmt"
	"io/ioutil"
	"github.com/anxtyanxty/jpkg"
)

func main() {
	body, err := ioutil.ReadFile("example.json")
	if err != nil {
		panic(err)
	}
	slides, err := jpkg.QueryJSONString(string(body), "['slideshow']['slides']")
	if err != nil {
		panic(err)
	}
	for index, slide := range slides.([]interface{}) {
		title, err := jpkg.QueryJSONInterface(slide, "['title']")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Title for Slide #%d: %s\n", index, title)
	}
}
