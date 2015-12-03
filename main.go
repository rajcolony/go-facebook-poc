package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"bytes"
	"net/url"
)

func main() {
	println("Init facebook.go > main()")
	http.HandleFunc("/facebook", facebook)
	http.HandleFunc("/", welcome)
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func facebook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	if (request.Method != "POST") {
		writer.Write([]byte("{error:'Unsupported HTTP method - " + request.Method + ". : Supports ONLY - HTTP POST'}"))
	} else {
		urlStr := "https://graph.facebook.com/v2.5/feed"
		data := url.Values{}
		data.Set("access_token", request.FormValue("access_token"))
		data.Add("message", request.FormValue("message"))

		client := &http.Client{}
		r, error := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, _ := client.Do(r)
		if error != nil {
			fmt.Printf("%s", error)
			//os.Exit(1)
		}

		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			//os.Exit(1)
		}
		//fmt.Printf("%s\n", string(contents))
		writer.Write(contents)
	}
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Write([]byte("<h1>Welcome to Go Facebook POC!!!</h1>"))
	writer.Write([]byte("HTTP Post request on /facebook uri with params: access_token, message in order to post on facebook."))
}