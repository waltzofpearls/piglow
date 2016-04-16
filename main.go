package main

import (
	"log"
	"time"

	"github.com/wjessop/go-piglow"
)

type PiGlow struct {
	*piglow.Piglow
}

func NewPiGlow() (*PiGlow, error) {
	parent, err := piglow.NewPiglow()
	if err != nil {
		return nil, err
	}
	return &PiGlow{parent}, nil
}

func (p *PiGlow) Glow(fnSetLeds func(uint8)) {
	for i := 0; i <= 255; i++ {
		fnSetLeds(uint8(i))
		p.Apply()
		time.Sleep(10 * time.Millisecond)
	}
}

func (p *PiGlow) Gloom(fnSetLeds func(uint8)) {
	for i := 255; i >= 0; i-- {
		fnSetLeds(uint8(i))
		p.Apply()
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	p, err := NewPiGlow()
	if err != nil {
		log.Fatal("Couldn't create a Piglow: ", err)
	}

	p.Glow(func(brightness uint8) {
		p.SetAll(brightness)
	})
	p.Gloom(func(brightness uint8) {
		p.SetAll(brightness)
	})
}
