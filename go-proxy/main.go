
/* 
Describe the application and what it is supposed to be doing
Build a simple nginx-ish http site
Host info:
	- port 8080 
	- without reverse proxy
	- host from a single dir on box
	
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}


// package main

// import (
//     "fmt"
//     "net/http"
// )

// func hello(w http.ResponseWriter, req *http.Request) {

//     fmt.Fprintf(w, "hello\n")
// }

// func headers(w http.ResponseWriter, req *http.Request) {

//     for name, headers := range req.Header {
//         for _, h := range headers {
//             fmt.Fprintf(w, "%v: %v\n", name, h)
//         }
//     }
// }

// func main() {

//     http.HandleFunc("/hello", hello)
//     http.HandleFunc("/headers", headers)

//     http.ListenAndServe(":8080", nil)
// }