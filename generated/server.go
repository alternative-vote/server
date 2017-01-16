package generated

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//StartServer Start the HTTP server
func StartServer(address string) {
	server := &http.Server{
		Handler: Stack(),
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 120 * time.Second,
		ReadTimeout:  120 * time.Second,
	}

	fmt.Println("listening at", address, "...")
	log.Fatal(server.ListenAndServe())
}
