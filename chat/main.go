package main

import (
	"log"
	"net/http"
)

func main() {
	// waiting Request and response html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
        <head>
          <title>チャット</title>
        </head>
        <body>
          チャットをしましょう
        </body>
      </html>
      `))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
