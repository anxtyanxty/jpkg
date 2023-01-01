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
	slideshow, err := jpkg.QueryJSONString(string(body), "['slideshow']")
	if err != nil {
		panic(err)
	}
	json_string, err := jpkg.ParseToRawString(slideshow)
	if err != nil {
		panic(err)
	}
	fmt.Println("slideshow json:", json_string)
}
