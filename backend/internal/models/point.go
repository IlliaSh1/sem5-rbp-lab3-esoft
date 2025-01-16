package models

import (
	"errors"

	"github.com/gofiber/fiber/v3/log"
)

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	ErrFailedToScanPoint = errors.New("failed to scan point")
)

func (p *Point) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return ErrFailedToScanPoint
	}

	log.Info(string(bytes))

	result := Point{
		Latitude:  1,
		Longitude: 1,
	}
	*p = result
	return nil
}

// type JSON json.RawMessage

// // Scan scan value into Jsonb, implements sql.Scanner interface
// func (j *JSON) Scan(value interface{}) error {
//   bytes, ok := value.([]byte)
//   if !ok {
//     return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
//   }

//   result := json.RawMessage{}
//   err := json.Unmarshal(bytes, &result)
//   *j = JSON(result)
//   return err
// }

// // Value return json value, implement driver.Valuer interface
// func (j JSON) Value() (driver.Value, error) {
//   if len(j) == 0 {
//     return nil, nil
//   }
//   return json.RawMessage(j).MarshalJSON()
// }

//
//
//
// func (p *Point) String() string {
// 	return fmt.Sprintf("Point(%f, %f)", p.Latitude, p.Longitude)
// }

// func (p *Point) Validate() (err error) {
// 	if p.Latitude < -90 || p.Latitude > 90 {
// 		err = errors.Join(err, fmt.Errorf("latitude must be between -90 and 90"))
// 	}
// 	if p.Longitude < -180 || p.Longitude > 180 {
// 		err = errors.Join(err, fmt.Errorf("longitude must be between -180 and 180"))
// 	}

// 	return err
// }
