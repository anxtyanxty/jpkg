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
	author, err := jpkg.QueryJSONString(string(body), "['slideshow']['author']")
	if err != nil {
		panic(err)
	}
	fmt.Println("The author of this book is", author)
}
