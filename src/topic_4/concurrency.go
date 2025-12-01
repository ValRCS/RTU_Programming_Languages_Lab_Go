package main

import("fmt";"time")
func worker(id int){ for i:=1;i<=3;i++{ fmt.Printf("Worker %d: step %d\n", id,i); time.Sleep(300*time.Millisecond)}}
func main(){ go worker(1); go worker(2); time.Sleep(2*time.Second)}
