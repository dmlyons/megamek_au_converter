package main

import "encoding/xml"

// Systems was generated 2025-07-05 14:38:06 by dl on framework13 with zek 0.1.28.
type Systems struct {
	XMLName xml.Name `xml:"systems"`
	Text    string   `xml:",chardata"`
	Systems []System `xml:"system"`
}

type TextWithSource struct {
	Text   string `xml:",chardata"`
	Source string `xml:"source,attr"`
}

type System struct {
	Text         string         `xml:",chardata"`
	ID           TextWithSource `xml:"id"`
	SucsId       int            `xml:"sucsId"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
	Xcood        float64        `xml:"xcood"`  // -123.764, 583.881, -70.31...
	Ycood        float64        `xml:"ycood"`  // 272.182, 256.931, -95.949...
	SpectralType TextWithSource `xml:"spectralType"`
	PrimarySlot  TextWithSource `xml:"primarySlot"`
	Planets      []struct {
		Text        string         `xml:",chardata"`
		Primary     string         `xml:"primary,attr"`
		Source      string         `xml:"source,attr"`
		Name        TextWithSource `xml:"name"`
		Type        TextWithSource `xml:"type"`
		OrbitalDist TextWithSource `xml:"orbitalDist"`
		SysPos      TextWithSource `xml:"sysPos"`
		Pressure    TextWithSource `xml:"pressure"`
		Atmosphere  TextWithSource `xml:"atmosphere"`
		Composition TextWithSource `xml:"composition"`
		Gravity     TextWithSource `xml:"gravity"`
		DayLength   TextWithSource `xml:"dayLength"`
		YearLength  TextWithSource `xml:"yearLength"`
		Diameter    TextWithSource `xml:"diameter"`
		Density     TextWithSource `xml:"density"`
		Satellites  []struct {
			Text   string         `xml:",chardata"`
			Source string         `xml:"source,attr"`
			Name   TextWithSource `xml:"name"`
			Size   TextWithSource `xml:"size"`
			Icon   string         `xml:"icon"` // rock12, oddmoon3, oddmoon...
		} `xml:"satellite"`
		Icon        TextWithSource   `xml:"icon"`
		Temperature TextWithSource   `xml:"temperature"`
		Water       TextWithSource   `xml:"water"`
		LifeForm    TextWithSource   `xml:"lifeForm"`
		Desc        TextWithSource   `xml:"desc"`
		LandMass    []TextWithSource `xml:"landMass"`
		SmallMoons  TextWithSource   `xml:"smallMoons"`
		Ring        TextWithSource   `xml:"ring"`
	} `xml:"planet"`
}
