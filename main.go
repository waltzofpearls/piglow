package main

import (
	"log"
	"time"

	"github.com/wjessop/go-piglow"
)

const (
	BrtMax   = 155
	BrtMin   = 0
	Interval = 10 * time.Millisecond
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
	for i := BrtMin; i <= BrtMax; i++ {
		fnSetLeds(uint8(i))
		if err := p.Apply(); err != nil {
			p.Error(err)
		}
		time.Sleep(Interval)
	}
}

func (p *PiGlow) Gloom(fnSetLeds func(uint8)) {
	for i := BrtMax; i >= BrtMin; i-- {
		fnSetLeds(uint8(i))
		if err := p.Apply(); err != nil {
			p.Error(err)
		}
		time.Sleep(Interval)
	}
}

func (p *PiGlow) Error(err error) {
	p.SetAll(0)
	p.Apply()
	log.Fatal("Couldn't apply changes: ", err)
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
