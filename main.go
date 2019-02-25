package main

import (
	"fmt"
	"jianghexinxi/pkg/mqtt"
	"jianghexinxi/pkg/setting"
)

func init() {
	setting.Setup()
	mqtt.Setup()
}
func main() {
	go mqtt.Subscribe()
	fmt.Println(<-mqtt.Chan)
	for {

	}

}
