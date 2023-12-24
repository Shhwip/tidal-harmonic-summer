package pkg

import (
	"encoding/xml"
)

type HarmonicConstituents struct {
	XMLName xml.Name `xml:"HarmonicConstituents"`
	Self    string   `xml:"self,attr"`
	Units   string   `xml:"units"`
	HarCons []HarCon `xml:"HarmonicConstituent"`
}

type HarCon struct {
	XMLName     xml.Name `xml:"HarmonicConstituent"`
	Type        string   `xml:"type,attr"`
	Number      int      `xml:"number"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Amplitude   float64  `xml:"amplitude"`
	PhaseGMT    float64  `xml:"phase_GMT"`
	PhaseLocal  float64  `xml:"phase_local"`
	Speed       float64  `xml:"speed"`
}
