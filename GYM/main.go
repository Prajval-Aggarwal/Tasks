package main

import (
	"fmt"
	cont "gym-api/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/enroll", cont.EnrollHandler)
	mux.HandleFunc("/member", cont.MemberHandler)
	mux.HandleFunc("/PriceGet", cont.PriceGetHandler)
	mux.HandleFunc("/PriceUpdate", cont.PriceUpdateHandler)
	mux.HandleFunc("/end-memberShip", cont.EndMemberShipHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
