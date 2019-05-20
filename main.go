package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	http.HandleFunc("/api_doc", func(responseWriter http.ResponseWriter, request *http.Request) {
		b,err := ioutil.ReadFile("./en/api.yaml")
		if err != nil {
			panic(err)
		}
		_,_ = responseWriter.Write(b)
	})
	
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
			_,_=responseWriter.Write([]byte(`<!DOCTYPE html>
<html>
  <head>
    <title>Bybit Api Doc</title>
    <!-- needed for adaptive design -->
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!--<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet"> -->

    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <div id="redoc-container"></div>
    <script src="https://cdn.jsdelivr.net/npm/redoc@2.0.0-rc.8-1/bundles/redoc.standalone.js"> </script>
    <script>
        Redoc.init('/api_doc', {
            expandResponses: 'all'
        }, document.getElementById('redoc-container'))
    </script>
  </body>
</html>`))
	})
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
	}
}
