package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	minutesCPU   = 5
	numCPU       = 10
	amountMemory = 2048
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	http.HandleFunc("/cpu", loadCPU)
	http.HandleFunc("/memory", loadMemory)
	log.Println("Listen: 8000.")
	http.ListenAndServe(":8000", nil)
}

func loadCPU(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < numCPU; i++ {
		go func() {
			now := time.Now()
			for {
				if time.Since(now) > time.Minute*time.Duration(minutesCPU) {
					return
				}
			}
		}()
	}
	msg := fmt.Sprintf("Allocating %d CPU's!\n", numCPU)
	log.Print(msg)
	w.Write([]byte(msg))
}

func loadMemory(w http.ResponseWriter, r *http.Request) {
	var buffer [amountMemory * 1e6]byte
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 1
	}
	msg := fmt.Sprintf("Allocated memory %d bytes!\n", len(buffer))
	log.Print(msg)
	w.Write([]byte(msg))
}
