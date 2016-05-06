package timestamp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var ts Timestamp = 20160229010203

func TestTimestampString(t *testing.T) {
	Convey("Soll einen ts aus einem String erzeugen", t, func() {
		s := "29.02.16 01:02:03"
		tsNeu := FromString(s)
		So(tsNeu, ShouldEqual, 20160229010203)
	})
	Convey("Soll den Zeitpunkt als String ausgeben", t, func() {
		So(ts.String(), ShouldEqual, "29.02.16 01:02:03")
	})

	Convey("Soll das Datum des Timestamps als String ausgeben", t, func() {
		So(ts.DateStr(), ShouldEqual, "29.02.16")
	})

	Convey("Soll die Zeit des Timestamps als String ausgeben", t, func() {
		So(ts.TimeStr(), ShouldEqual, "01:02:03")
	})

}
