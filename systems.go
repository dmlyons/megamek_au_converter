package main

import "encoding/xml"

// Systems was generated 2025-07-05 14:38:06 by dl on framework13 with zek 0.1.28.
type Systems struct {
	XMLName xml.Name `xml:"systems"`
	Text    string   `xml:",chardata"`
	Systems []System `xml:"system"`
}

type System struct {
	Text string `xml:",chardata"`
	ID   struct {
		Text   string `xml:",chardata"` // A Place, Aalzorg, Abadan,...
		Source string `xml:"source,attr"`
	} `xml:"id"`
	SucsId       int     `xml:"sucsId"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
	Xcood        float64 `xml:"xcood"`  // -123.764, 583.881, -70.31...
	Ycood        float64 `xml:"ycood"`  // 272.182, 256.931, -95.949...
	SpectralType struct {
		Text   string `xml:",chardata"` // G5V, G9V, F1V, G9V, G2V, ...
		Source string `xml:"source,attr"`
	} `xml:"spectralType"`
	PrimarySlot struct {
		Text   string `xml:",chardata"` // 2, 2, 5, 3, 2, 4, 4, 3, 3...
		Source string `xml:"source,attr"`
	} `xml:"primarySlot"`
	Planets []struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Source  string `xml:"source,attr"`
		Name    struct {
			Text   string `xml:",chardata"` // Sullivan, A Place, Sraosa...
			Source string `xml:"source,attr"`
		} `xml:"name"`
		Type struct {
			Text   string `xml:",chardata"` // Giant Terrestrial, Terres...
			Source string `xml:"source,attr"`
		} `xml:"type"`
		OrbitalDist struct {
			Text   string `xml:",chardata"` // 0.36, 0.63, 0.9, 1.44, 2....
			Source string `xml:"source,attr"`
		} `xml:"orbitalDist"`
		SysPos struct {
			Text   string `xml:",chardata"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
			Source string `xml:"source,attr"`
		} `xml:"sysPos"`
		Pressure struct {
			Text   string `xml:",chardata"` // Very High, Low, Very High...
			Source string `xml:"source,attr"`
		} `xml:"pressure"`
		Atmosphere struct {
			Text   string `xml:",chardata"` // Toxic (Poisonous), Breath...
			Source string `xml:"source,attr"`
		} `xml:"atmosphere"`
		Composition struct {
			Text   string `xml:",chardata"` // Hydrogen and Helium, plus...
			Source string `xml:"source,attr"`
		} `xml:"composition"`
		Gravity struct {
			Text   string `xml:",chardata"` // 1.34, 1.02, 2.05, 1.02, 2...
			Source string `xml:"source,attr"`
		} `xml:"gravity"`
		DayLength struct {
			Text   string `xml:",chardata"` // 20, 27, 15, 13, 12, 16, 1...
			Source string `xml:"source,attr"`
		} `xml:"dayLength"`
		YearLength struct {
			Text   string `xml:",chardata"` // 1, 1.1, 1.3, 2, 4.1, 10.2...
			Source string `xml:"source,attr"`
		} `xml:"yearLength"`
		Diameter struct {
			Text   string `xml:",chardata"` // 13500, 10500, 120000, 800...
			Source string `xml:"source,attr"`
		} `xml:"diameter"`
		Density struct {
			Text   string `xml:",chardata"` // 7, 6.8337, 1.2, 0.9, 1, 1...
			Source string `xml:"source,attr"`
		} `xml:"density"`
		Satellites []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
			Name   struct {
				Text   string `xml:",chardata"` // Vertumnus, Feronia, Venus...
				Source string `xml:"source,attr"`
			} `xml:"name"`
			Size struct {
				Text   string `xml:",chardata"` // medium, medium, medium, m...
				Source string `xml:"source,attr"`
			} `xml:"size"`
			Icon string `xml:"icon"` // rock12, oddmoon3, oddmoon...
		} `xml:"satellite"`
		Icon struct {
			Text   string `xml:",chardata"` // gasg6, green47, asteroid2...
			Source string `xml:"source,attr"`
		} `xml:"icon"`
		Temperature struct {
			Text   string `xml:",chardata"` // 31, 6, -151, 152, 31, 10,...
			Source string `xml:"source,attr"`
		} `xml:"temperature"`
		Water struct {
			Text   string `xml:",chardata"` // 32, 0, 37, 71, 0, 0, 0, 0...
			Source string `xml:"source,attr"`
		} `xml:"water"`
		LifeForm struct {
			Text   string `xml:",chardata"` // Birds, Birds, MAMMAL, Amp...
			Source string `xml:"source,attr"`
		} `xml:"lifeForm"`
		Desc struct {
			Text   string `xml:",chardata"` // Despite A Place's poor so...
			Source string `xml:"source,attr"`
		} `xml:"desc"`
		LandMass []struct {
			Text   string `xml:",chardata"` // Campbell (New Keene), Bon...
			Source string `xml:"source,attr"`
		} `xml:"landMass"`
		SmallMoons struct {
			Text   string `xml:",chardata"` // 12, 19, 15, 22, 21, 17, 2...
			Source string `xml:"source,attr"`
		} `xml:"smallMoons"`
		Ring struct {
			Text   string `xml:",chardata"` // true, true, true, true, t...
			Source string `xml:"source,attr"`
		} `xml:"ring"`
	} `xml:"planet"`
}
