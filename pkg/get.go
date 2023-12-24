package pkg

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
)

func GetHarmonicConstituent(station string) (HarmonicConstituents, error) {
	// get data from NOAA API
	baseurl, _ := url.Parse("https://api.tidesandcurrents.noaa.gov/mdapi/prod/webapi/stations/")
	baseurl.Path += station
	baseurl.Path += "/harcon.xml"

	resp, err := http.Get(baseurl.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	harmonicConstituent := HarmonicConstituents{}
	xml.Unmarshal(body, &harmonicConstituent)
	return harmonicConstituent, err
}
