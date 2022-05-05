package sensor

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

func (R RPIDistanceSensor) MeasureCM() (float64, error) {
	
	panic("implement me")
}
