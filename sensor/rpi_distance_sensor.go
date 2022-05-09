package sensor

import (
	"errors"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

type RPIDistanceSensor struct {
	trigger     int
	echo        int
	mode        int
	warnings    bool
	triggerTime float64
	maxWaitTime float64
}

var _ DistanceSensor = (*RPIDistanceSensor)(nil)

func NewRPIDistanceSensor(triggerGPIOPin, echoGPIOPin, mode int, warnings bool) *RPIDistanceSensor {
	return &RPIDistanceSensor{
		trigger:     triggerGPIOPin,
		echo:        echoGPIOPin,
		mode:        mode,
		warnings:    warnings,
		triggerTime: 0.00001,
		maxWaitTime: 0.015, // max time waiting for response in case something is missed
	}
}

const triggerTimeMicro = 10 * time.Microsecond
const maxTimeSeconds = 0.015
const speedOfSoundCmPerS = 34300.0

//
//func roundTripDistanceCmToS(distance float64) float64 {
//	return (distance * 2.0) / speedOfSoundCmPerS
//}

func (s RPIDistanceSensor) MeasureCM() (float64, error) {
	fmt.Println("measuring on the raspberry pi")
	err := rpio.Open()
	if err != nil {
		return 0.0, err
	}
	defer rpio.Close() // guarantees we unmap the memory for gpio
	fmt.Printf("trigger pin %v ,  echo pin %v \n", s.trigger, s.echo)
	// setup pins
	echoPin := rpio.Pin(s.echo)
	triggerPin := rpio.Pin(s.trigger)
	echoPin.Input()
	echoPin.PullUp() // Online it looks like people use PUllUP? why not pulldown
	triggerPin.Output()

	fmt.Println("configured pins for the correct input vs output")
	// Set trigger pin high for at least 10 Microseconds
	triggerPin.High()
	var echoStartTime time.Time
	var echoStopTime time.Time
	var startTime time.Time = time.Now() // used for timeouts

	// pulse trigger
	time.Sleep(triggerTimeMicro)
	echoStartTime = time.Now() // set startTime in case of super fast response
	triggerPin.Low()

	fmt.Println("triggered trigger pin")

	// wait for start of echo response / wait for pin to go high
	for echoPin.Read() == rpio.Low && time.Since(startTime).Seconds() <= maxTimeSeconds { // wait while echo pin is
		fmt.Println("waiting for pin to go high")
		echoStartTime = time.Now()
	}

	if time.Since(startTime).Seconds() > maxTimeSeconds {
		return 0.0, errors.New("timeout reached while waiting for echo start")
	}

	// Wait for echo end
	echoStopTime = time.Now()
	for echoPin.Read() == rpio.High && time.Since(startTime).Seconds() <= maxTimeSeconds {
		fmt.Println("waiting for pin to go low")
		echoStopTime = time.Now() // continue overwriting stop time
	}

	elapsedSeconds := echoStopTime.Sub(echoStartTime).Seconds()
	if elapsedSeconds <= maxTimeSeconds {
		fmt.Println("calculating distance")
		return elapsedSeconds * speedOfSoundCmPerS / 2.0, nil
	} else {
		return 0.0, errors.New("timeout reached while waiting for echo end")
	}
}
