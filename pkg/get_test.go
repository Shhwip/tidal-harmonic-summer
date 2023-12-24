package pkg

import (
	"fmt"
	"testing"
)

func TestGetHarmonicConstituent(t *testing.T) {
	station := "9452210"
	HarmonicConstituents, err := GetHarmonicConstituent(station)
	if err != nil {
		t.Error(err)
	}
	fmt.Print(HarmonicConstituents)
}
