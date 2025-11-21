package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	
	fmt.Printf("Serveur lancé--> http://localhost%v\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}