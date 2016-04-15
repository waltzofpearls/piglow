package main

import (
	"log"
	"time"

	"github.com/wjessop/go-piglow"
)

func main() {
	p, err := piglow.NewPiglow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	// p.SetLED(0, 0)
	// p.SetLED(1, 0)
	p.SetTentacle(0, 255)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
	time.Sleep(time.Second)
	p.SetTentacle(0, 0)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
	time.Sleep(time.Second)
	p.SetTentacle(1, 255)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
	time.Sleep(time.Second)
	p.SetTentacle(1, 0)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
	time.Sleep(time.Second)
	p.SetTentacle(2, 255)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
	time.Sleep(time.Second)
	p.SetTentacle(2, 0)
	err = p.Apply()
	if err != nil { // Apply the changes
		log.Fatal("Couldn't apply changes: ", err)
	}
}
