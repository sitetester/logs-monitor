package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	tickerProducer := time.NewTicker(1 * time.Second)

	for {
		select {
		case _ = <-tickerProducer.C:

			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 10
			limit := rand.Intn(max-min+1) + min
			log.Println(fmt.Sprintf("Making %d request(s)...", limit))

			for i := 1; i <= limit; i++ {
				_, err := http.Get("http://localhost:8080")
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

}
