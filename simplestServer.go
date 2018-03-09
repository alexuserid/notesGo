package main 

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {//by writing data in ResponseWriter value we send the data(HTTP server's response) to the HTTP client
	//http.Request is a data structure that represents the client HTTP request
	fmt.Fprintf(w, "Hi, there, I love %s!", r.URL.Path[1:]) //r.URL.Path is component of the request URL. In our case it will be /cats 
	//Fprintf takes r.URL.Path[1:] and puts it instead %s, then the result is writing in w
}

func main() {
	http.HandleFunc("/", handler) //this func tells the http package to handle all requests to the web root ("/") with handle func
	log.Fatal(http.ListenAndServe(":8080", nil)) //ListenAndServe listen on port 8080 on any interface (":8080")
}


//when we run the program and access the URL: http://localhost:8080/cats, the program would present a page contain: Hi there, I love cats!
