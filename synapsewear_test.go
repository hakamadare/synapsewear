package synapsewear

import (
	"fmt"
	"testing"
	"time"

	"github.com/MakeNowJust/heredoc"
)

var sampleData string = heredoc.Doc(`
data={
  "deviceuuid" : "6C5DF9F1-E559-7EAF-77B2-19C2D7BC5197",
  "data" : [
    {
      "voltage" : 3.9862499237060547,
      "CO2" : 481,
      "airpressure" : 1000.9099731445312,
      "illumination" : 2666,
      "humidity" : 48,
      "date" : "2018\/08\/22 19:00:00 -0400",
      "temperature" : 34.200000762939453,
      "envsound" : 195,
      "dateunix" : 1534978800
    }
	]
}
`)

func TestEscapedTimeParsing(t *testing.T) {
	escapedTimeStr := `2018\/08\/22 19:00:00 -0400`
	escapedTime, err := time.Parse(escapedTimeFormat, escapedTimeStr)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(fmt.Sprintf("%v", escapedTime))
}

func TestUploadParsing(t *testing.T) {
	_, err := ParseString(sampleData)

	if err != nil {
		t.Fatal(err)
	}
}
