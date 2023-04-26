// package main

// import (
// 	"fmt"
// 	"log"

// 	geo "github.com/vvikash157/go-microservices/geometry"

// 	"net/http"

// 	"rsc.io/quote"
// )

// func rectProps(length, width int) (area, perimeter int) {
// 	area = length * width
// 	perimeter = 2 * (length + width)
// 	return
// }

// func main() {
// 	// fmt.Println("Hello World")
// 	fmt.Println(quote.Go())
// 	A, P := rectProps(5, 9)
// 	fmt.Printf("Area is %v and perimeter is %v", A, P)
// 	fmt.Println(geo.Area(6, 8))
// 	fmt.Println(geo.Diagonal(6, 8))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, you've requested: %s, with token no: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// 	})

// 	fs := http.FileServer(http.Dir("static/"))
//     http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("server has started on your request")

// 	http.ListenAndServe(":80", nil)
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/vvikash157/go-microservices/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking Application Health")
	mapResponse := map[string]string{
		"Health":    "up",
		"timestamp": time.Now().String(),
	}

	json.NewEncoder(w).Encode(mapResponse)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving Home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is Up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching details")
	hostname, err := details.Hostname()

	if err != nil {
		panic(err)
	}

	ip := details.GetIpAddress()
	fmt.Println(hostname, ip)

	mapResponse := map[string]string{
		"hostname": hostname,
		"ip":       ip.String(),
	}

	json.NewEncoder(w).Encode(mapResponse)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}
