package date

import (
	"testing"
)

func TestIsParseable(t *testing.T) {
	BasicUTIsParseable(t, "11 des 0001", true)
	BasicUTIsParseable(t, "11 12 0001", false)
	BasicUTIsParseable(t, "11/des/0001", true)
	BasicUTIsParseable(t, "11-des-0001", true)
	BasicUTIsParseable(t, "02 11 0001", false)
	BasicUTIsParseable(t, "1-jan-0001", true)

	BasicUTIsParseable(t, "0001-11 des", true)
	BasicUTIsParseable(t, "0001-11 12", false)
	BasicUTIsParseable(t, "0001-11/des", true)
	BasicUTIsParseable(t, "0001-11-des", true)
	BasicUTIsParseable(t, "0001-02 11", false)
	BasicUTIsParseable(t, "0001-1-jan", true)

	BasicUTIsParseable(t, "des 11 0001", true)
	BasicUTIsParseable(t, "11/des/0001", true)
	BasicUTIsParseable(t, "des-des-0001", false)
}

func BasicUTIsParseable(t *testing.T, dateString string, expected bool) {
	if val := IsParseable(dateString); val != expected {
		t.Errorf("Fail with input %s\nExpecting %v but got %v", dateString, expected, val)
	}
}
