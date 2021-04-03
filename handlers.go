package main

import (
	"fmt"
	"net/http"
	"github.com/unrolled/render"
)

type sampleContent struct {
	ID      string 'json:"id"'
	Content string 'json:"content"'
}

func testHandler(formatter * render.Render) http.HandlerFunc  {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			sampleContent{ID: "8675309", Content: "Óla Gabriela"})
	}
}

func main() {

}
