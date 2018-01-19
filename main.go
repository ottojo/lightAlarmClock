package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"time"
)

var c mqtt.Client

type Color struct {
	R float64
	G float64
	B float64
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://odroid.lan:1883")
	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer c.Disconnect(200)

	fade(Color{0, 0, 0}, Color{1023, 563, 0}, 30*time.Second, 30*20)
	fade(Color{1023, 563, 0}, Color{1023, 1023, 1023}, 30*time.Second, 30*20)
}

func fade(current, target Color, duration time.Duration, steps int) {
	rateR := target.R - current.R
	rateG := target.G - current.G
	rateB := target.B - current.B
	for current.R < target.R || current.G < target.G || current.B < target.B {
		current.R += rateR / float64(steps)
		current.G += rateG / float64(steps)
		current.B += rateB / float64(steps)
		go setLight(current)
		//go log.Println(current)
		time.Sleep(time.Duration(duration.Nanoseconds()/int64(steps)) * time.Nanosecond)
	}
	go log.Println(current)
}

func setLight(color Color) {
	token := c.Publish("/lights/ledStripWindow/red", 0, false, strconv.Itoa(int(color.R)))
	token.Wait()
	token = c.Publish("/lights/ledStripWindow/green", 0, false, strconv.Itoa(int(color.G)))
	token.Wait()
	token = c.Publish("/lights/ledStripWindow/blue", 0, false, strconv.Itoa(int(color.B)))
	token.Wait()
}
