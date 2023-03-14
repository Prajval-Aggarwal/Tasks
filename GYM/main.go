package main

import (
	"fmt"
	cont "gym-api/controllers"
	"log"
	"net/http"
)

//Controller
// func main() {
// 	fmt.Println("Listening on port:8080")
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/enroll", cont.EnrollHandler)
// 	mux.HandleFunc("/member", cont.MemberHandler)
// 	mux.HandleFunc("/PriceGet", cont.PriceGetHandler)
// 	mux.HandleFunc("/PriceUpdate", cont.PriceUpdate)
// 	mux.HandleFunc("/end-memberShip", cont.EndMemberShipHandler)
// 	log.Fatal(http.ListenAndServe(":8080", mux))

// }

//cont1

func main() {
	fmt.Println("Listening on port:8000")
	mux := http.NewServeMux()
	mux.HandleFunc("/createuser", cont.CreateUserHandler)
	mux.HandleFunc("/makepayent", cont.MakepaymentHandler)
	mux.HandleFunc("/createsubs", cont.CreateSubsHandler)
	mux.HandleFunc("/createEmp", cont.CreateEmphandler)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)

	mux.HandleFunc("/getUsers", cont.GetUsers)
	mux.HandleFunc("/end-memberShip", cont.EndSubscription)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
