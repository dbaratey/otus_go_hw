package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntp, err := ntp.Time("ntp1.stratum2.ru")
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	fmt.Printf("current time: %s\nexact time: %s\n", now.Round(0), ntp.Round(0))
}
