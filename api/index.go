package api

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r.URL.Path)

	var res []byte

	var client http.Client
	resp, err := client.Get("https://atulya.me/README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		res = bodyBytes
	}

	fmt.Print(string(res))

	l := blackfriday.Run(res)
	fmt.Print(string(l))
	w.Write([]byte(l))
	w.Header().Set("Content-Type", "text/plain")

	//return string(l)
}