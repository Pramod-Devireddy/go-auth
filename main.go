/*
 * File         : main.go
 * Project      : Go-Auth
 * Created Date : Thursday, Jun 24th 2021, 4:40:32 AM
 * Author       : Pramod Devireddy
 *
 * Last Modified: Thursday, 24th June 2021 7:55:53 am
 * Modified By  : Pramod Devireddy
 *
 * Copyright (c)2021 - Pramod Devireddy
 * ***************************************************************
 */

package main

import (
	"log"
	"net/http"
	"tollysearch/mux"

	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter()

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	router.Use(cors)

	router.HandleFunc("/login", login)
	router.HandleFunc("/home", home)
	router.HandleFunc("/refresh", refresh)

	log.Fatal(http.ListenAndServe(":8080", router))
}
