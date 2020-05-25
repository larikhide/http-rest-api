package main

import (
	"github.com/larikhide/http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
 s := apiserver.New() //starts server
 if err := s.Start(); err != nil {
 	log.Fatal(err)
 }
}