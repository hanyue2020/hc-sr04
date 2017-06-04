package main

import "fmt"
import (
	"github.com/stianeikeland/go-rpio"
	"time"
	"os"
)

var (
	pin_send2 rpio.Pin = rpio.Pin(2)
	pin_recv3 rpio.Pin = rpio.Pin(3)
)

func checkDistance() float64 {
	pin_send2.Low()
	time.Sleep(time.Microsecond * 30)
	pin_send2.High()
	time.Sleep(time.Microsecond * 30)
	pin_send2.Low()
	time.Sleep(time.Microsecond * 30)
	for {
		status := pin_recv3.Read()
		if status == rpio.High {
			break;
		}
	}
	begin := time.Now()
	for {
		status := pin_recv3.Read()
		if status == rpio.Low {
			break
		}
	}
	end := time.Now()
	diff := end.Sub(begin)
	//fmt.Println("diff = ",diff.Nanoseconds(),diff.Seconds(),diff.String()) 1496548629.307,501,127
	result_sec := float64(diff.Nanoseconds()) / 1000000000.0
	//fmt.Println("begin = ", begin.UnixNano(), " end = ", end.UnixNano(), "diff = ", result_sec, diff.Nanoseconds())
	return result_sec * 340.0 / 2
}

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	pin_send2.Output()
	pin_recv3.Input()

	time.Sleep(time.Second * 2)
	for {
		distance := checkDistance();
		fmt.Println("distance = ", distance)
		time.Sleep(time.Millisecond * 500)
	}
}
