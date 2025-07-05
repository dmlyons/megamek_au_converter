package main

import (
	"encoding/xml"
	"errors"
	"fmt"
)

// Systems was generated 2025-07-05 14:38:27 by dl on framework13 with zek 0.1.28.
type SystemEvents struct {
	XMLName xml.Name `xml:"systems"`
	Text    string   `xml:",chardata"`
	Events  []Event  `xml:"system"`
}

func (se *SystemEvents) FindEvent(sucsID int) (*Event, error) {
	for _, e := range se.Events {
		if e.SucsId == sucsID {
			return &e, nil
		}
	}
	return nil, errors.New("event not found sucsID: " + fmt.Sprintf("%d", sucsID))
}

type Event struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id"`     // A Place, Aalzorg, Abadan,...
	SucsId int    `xml:"sucsId"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
	Events []struct {
		Text        string `xml:",chardata"`
		Date        string `xml:"date"` // 2431-10-03, 2794-03-24, 2...
		NadirCharge struct {
			Text   string `xml:",chardata"` // true, false, true, false,...
			Source string `xml:"source,attr"`
		} `xml:"nadirCharge"`
		ZenithCharge struct {
			Text   string `xml:",chardata"` // false, true, false, true,...
			Source string `xml:"source,attr"`
		} `xml:"zenithCharge"`
	} `xml:"event"`
	Planets []struct {
		Text   string `xml:",chardata"` // a
		SysPos struct {
			Text   string `xml:",chardata"` // 2, 2, 5, 3, 2, 4, 4, 3, 3...
			Source string `xml:"source,attr"`
		} `xml:"sysPos"`
		Events []struct {
			Text    string `xml:",chardata"`
			Date    string `xml:"date"` // 2235-04-30, 2240-01-01, 2...
			Faction struct {
				Text     string `xml:",chardata"` // IND, LA, FC, ARDC, LA, CJ...
				Source   string `xml:"source,attr"`
				Citation string `xml:"citation,attr"`
			} `xml:"faction"`
			Population struct {
				Text   string `xml:",chardata"` // 41458, 78165, 246383, 671...
				Source string `xml:"source,attr"`
			} `xml:"population"`
			SocioIndustrial struct {
				Text   string `xml:",chardata"` // D-C-A-D-C, C-C-A-D-C, B-C...
				Source string `xml:"source,attr"`
			} `xml:"socioIndustrial"`
			Hpg struct {
				Text   string `xml:",chardata"` // A, X, B, X, B, X, B, X, B...
				Source string `xml:"source,attr"`
			} `xml:"hpg"`
			HiringHall struct {
				Text   string `xml:",chardata"` // QUESTIONABLE, STANDARD, Q...
				Source string `xml:"source,attr"`
			} `xml:"hiringHall"`
			Atmosphere struct {
				Text   string `xml:",chardata"` // Breathable, Toxic (Poison...
				Source string `xml:"source,attr"`
			} `xml:"atmosphere"`
			DayLength struct {
				Text   string `xml:",chardata"` // 48
				Source string `xml:"source,attr"`
			} `xml:"dayLength"`
			Pressure struct {
				Text   string `xml:",chardata"` // Standard, Very High, Thin...
				Source string `xml:"source,attr"`
			} `xml:"pressure"`
			Temperature struct {
				Text   string `xml:",chardata"` // 40, 250
				Source string `xml:"source,attr"`
			} `xml:"temperature"`
			Water struct {
				Text   string `xml:",chardata"` // 30, 39
				Source string `xml:"source,attr"`
			} `xml:"water"`
		} `xml:"event"`
	} `xml:"planet"`
}
