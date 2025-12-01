package main

import("fmt";"log";"net/http")
func hello(w http.ResponseWriter,r *http.Request){ fmt.Fprintln(w,"Hello from Go Web Server at RTU!")}
func main(){ http.HandleFunc("/",hello); fmt.Println("http://localhost:8080"); log.Fatal(http.ListenAndServe(":8080",nil)) }
