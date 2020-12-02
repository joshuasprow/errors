package server

import (
	"fmt"
	"net/http"
)

func borders() (bottom func()) {
	border := func() {
		fmt.Println("-----------")
	}
	border()
	return border
}

func Start(port int, quit chan struct{}) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		b := borders()
		defer b()

		fmt.Println()
	})

	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	<-quit
}
