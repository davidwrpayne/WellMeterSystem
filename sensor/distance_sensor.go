package sensor


type DistanceSensor interface {
	MeasureCM() (float64, error) // reports Distance in Centimeters
}

