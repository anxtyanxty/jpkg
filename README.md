# jpkg
simple json parser package written in go

### example code
```
package main

import (
	"fmt"
	"net/http"

	"github.com/anxtyanxty/jpkg"
)

func main() {
	resp, err := http.Get("https://httpbin.org/json")
	if err != nil {
		panic(err)
	}
	author, err := jpkg.QueryJSONReader(resp.Body, "['slideshow']['author']")
	if err != nil {
		panic(err)
	}
	fmt.Println("The author of this book is", author)
}
```
