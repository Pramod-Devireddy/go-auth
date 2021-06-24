/*
 * File         : main.go
 * Project      : Go-Auth
 * Created Date : Thursday, Jun 24th 2021, 4:40:32 AM
 * Author       : Pramod Devireddy
 *
 * Last Modified: Thursday, 24th June 2021 4:43:30 am
 * Modified By  : Pramod Devireddy
 *
 * Copyright (c)2021 - Pramod Devireddy
 * ***************************************************************
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/home", home)
	http.HandleFunc("/refresh", refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
