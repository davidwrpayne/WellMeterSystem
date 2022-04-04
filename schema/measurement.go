package schema

import (
	"time"
)

type Measurement struct {
	Verision          int     `json:"version"`
	UUID              string  `json:"uuid"'`
	Distance          float64 `json:"water_level"`
	MeasuredAtRFC3339 string  `json:"measured_at_rfc_3339"`
}

func NewMeasurement(uuid string, distanceCM float64, measuredAt time.Time) *Measurement {
	return &Measurement{
		Verision:          2,
		UUID:              uuid,
		Distance:          distanceCM,
		MeasuredAtRFC3339: measuredAt.Format(time.RFC3339),
	}
}
