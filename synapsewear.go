package synapsewear

// sample raw data:
//
// data={
//   "deviceuuid" : "6C5DF9F1-E559-7EAF-77B2-19C2D7BC5197",
//   "data" : [
//     {
//       "voltage" : 3.9862499237060547,
//       "CO2" : 481,
//       "airpressure" : 1000.9099731445312,
//       "illumination" : 2666,
//       "humidity" : 48,
//       "date" : "2018\/08\/22 19:00:00 -0400",
//       "temperature" : 34.200000762939453,
//       "envsound" : 195,
//       "dateunix" : 1534978800
//     }
// 	]
// }

import (
	"encoding/json"
	"regexp/syntax"
	"strings"
	"time"

	"github.com/google/uuid"
)

const dataPrefix = "data="
const escapedTimeFormat = "2006\\/01\\/02 15:04:05 -0700"
const escapedTimeRegexpStr = "[[:digit:]]{4}\\/[[:digit:]]{2}\\/[[:digit:]]{2}[[:space:]]+[[:digit:]]{2}:[[:digit:]]{2}:[[:digit:]]{2}[[:space:]]+-[[:digit:]]{4}"

type Upload struct {
	DeviceUUID uuid.UUID   `json:"deviceuuid"`
	Data       []Datapoint `json:"data"`
}

type Datapoint struct {
	Voltage      float64   `json:"voltage"`
	CO2          int       `json:"CO2"`
	AirPressure  float64   `json:"airpressure"`
	Illumination int       `json:"illumination"`
	Humidity     int       `json:"humidity"`
	Date         time.Time `json:"date"`
	Temperature  float64   `json:"temperature"`
	EnvSound     int       `json:"envsound"`
	DateUnix     int64     `json:"dateunix"`
}

func ParseString(s string) (Upload, error) {
	var u Upload

	s = strings.TrimPrefix(dataPrefix)

	s, err := scrubEscapedTime(s)

	if err != nil {
		return u, err
	}

	err = json.Unmarshal([]byte(s), &u)

	return u, err
}

func scrubEscapedTime(s string) (string, error) {
	escapedTimeRegexp, err := syntax.Parse(escapedTimeRegexpStr, syntax.PerlX)
	toScrub := escapedTimeRegexp.FindAllString(s, -1)

	for i := range toScrub {
		escapedTimeStr := toScrub[i]
		escapedTime, err := time.Parse(escapedTimeFormat, escapedTimeStr)

		if err != nil {
			return s, err
		}

		rfc3339TimeStr := escapedTime.Format(time.RFC3339)

		s = escapedTimeRegexp.ReplaceAllString(escapedTimeStr, rfc3339TimeStr)
	}

	return s, nil
}
