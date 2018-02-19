package main

import (
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		data := map[string]interface{}{
			"code":    "ok",
			"error":   false,
			"payload": "Hello World",
		}

		res, err := json.Marshal(data)

		if err != nil {
			panic(err)
		}

		w.Write(res)
	})

	http.ListenAndServe(":1337", nil)
}