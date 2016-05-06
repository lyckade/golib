package timestamp

import (
	"strconv"
	"time"
)

// Timestamp bildet Zeitpunkte als Integer Wert ab
type Timestamp int

const timestampLayout string = "20060102150405"
const stringLayout string = "02.01.06 15:04:05"
const dateStrLayout string = "02.01.06"
const timeStrLayout string = "15:04:05"

// FromString erzeugt einen Timestamp aus einem String
func FromString(s string) Timestamp {
	return FromLayout(stringLayout, s)
}

// FromLayout erezugt einen Timestamp aus einem LayoutString und einem
// string
func FromLayout(layoutString string, s string) Timestamp {
	t, _ := time.Parse(layoutString, s)
	i, _ := strconv.Atoi(t.Format(timestampLayout))
	ts := Timestamp(i)
	return ts
}

func (ts *Timestamp) String() string {
	t := ts.Time()
	return t.Format(stringLayout)
}

//Time gibt ein time.Time() Objekt des Zeitstempels zurück
func (ts *Timestamp) Time() time.Time {
	t, _ := time.Parse(timestampLayout, strconv.Itoa(int(*ts)))
	return t
}

//DateStr liefert das Datum als  string
func (ts *Timestamp) DateStr() string {
	return ts.Time().Format(dateStrLayout)
}

//TimeStr liefert die Zeit als  string
func (ts *Timestamp) TimeStr() string {
	return ts.Time().Format(timeStrLayout)
}
