package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("ntp1.stratum2.ru")
	if err != nil {
		log.Fatal(err)
	}
	nowTime := time.Now()
	fmt.Printf("current time: %s\nexact time: %s\n", nowTime.Round(0), ntpTime.Round(0))
}
