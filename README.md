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
	defer resp.Body.Close()
	
	buffer := &bytes.Buffer{}
	if _, err := buffer.ReadFrom(resp.Body); err != nil {
		panic(err)
	}
	
	json_object, err := jpkg.LoadJObject(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	
	// author is of type string
	author, err := json_object.QueryString("['slideshow']['author']")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The author of this book is %s\n", author)
}
```
