package sensor

type FakeSensor struct {
	fakeMeasurements []float64
	index int
}

func (f FakeSensor) MeasureCM() (float64, error) {
	measurement := f.fakeMeasurements[f.index]
	f.index += 1
	f.index = f.index % len(f.fakeMeasurements)
	return measurement, nil
}

func NewFakeSensor(measurementsToReport []float64) *FakeSensor {
	return &FakeSensor{
		fakeMeasurements: measurementsToReport,
		index: 0,
	}
}

var _ DistanceSensor = (*FakeSensor)(nil)




