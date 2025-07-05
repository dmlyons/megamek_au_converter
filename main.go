package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	yaml "github.com/goccy/go-yaml"
)

func main() {
	systemFile := flag.String("systems", "systems.xml", "your systems.xml file")
	eventsFile := flag.String("events", "system_events.xml", "your systems.xml file")
	outDir := flag.String("out", "out", "directory to output yml files")

	flag.Parse()

	// read the systems
	var systems Systems

	s, err := os.ReadFile(*systemFile)
	if err != nil {
		log.Fatalf("read systems file %v", err)
	}
	if err := xml.Unmarshal(s, &systems); err != nil {
		log.Fatalf("unmarshal systems %v", err)
	}

	// read the events
	var events SystemEvents

	e, err := os.ReadFile(*eventsFile)
	if err != nil {
		log.Fatalf("read events %v", err)
	}
	if err := xml.Unmarshal(e, &events); err != nil {
		log.Fatalf("unmarshal systems %v", err)
	}

	// create the output directory
	err = os.Mkdir(*outDir, 0777)
	if err != nil {
		log.Printf("create out dir %v", err)
	}

	// create the system yaml files

	// loop through the systems
	for _, system := range systems.Systems {
		e, err := events.FindEvent(system.SucsId)
		if err != nil {
			log.Printf("find event: %v", err)
			continue
		}
		err = process(outDir, &system, e)
		if err != nil {
			log.Printf("process err %v", err)
		}
	}

}

func process(outDir *string, system *System, events *Event) error {
	fn := *outDir + `/` + strings.ReplaceAll(system.ID.Text, " ", "") + `.yml`
	f, err := os.Create(fn)
	if err != nil {
		log.Printf("create yaml file %v", err)
	}
	defer f.Close()

	primarySlot, err := strconv.Atoi(system.PrimarySlot.Text)
	if err != nil {
		return fmt.Errorf("strconv: %w", err)
	}
	ps := PlanetarySystem{
		ID:           system.ID.Text,
		SucsID:       system.SucsId,
		Xcood:        system.Xcood,
		Ycood:        system.Ycood,
		SpectralType: system.SpectralType.Text,
		PrimarySlot:  primarySlot,
		Events:       []PsEvent{},
		Planets:      []PsPlanet{},
	}

	// system events
	evs := []PsEvent{}
	for _, sysEvent := range events.Events {
		evs = append(evs, PsEvent{
			Date:         sysEvent.Date,
			NadirCharge:  sysEvent.NadirCharge.Text,
			ZenithCharge: sysEvent.ZenithCharge.Text,
		})
	}
	ps.Events = evs

	// planets
	planets := []PsPlanet{}
	for _, p := range system.Planets {
		planet := PsPlanet{
			Name:        p.Name.Text,
			Type:        p.Type.Text,
			OrbitalDist: toFloat(p.OrbitalDist.Text),
			SysPos:      toInt(p.SysPos.Text),
			Icon:        p.Icon.Text,
			Pressure:    p.Pressure.Text,
			Atmosphere:  p.Atmosphere.Text,
			Gravity:     toFloat(p.Gravity.Text),
			Diameter:    toFloat(p.Diameter.Text),
			Density:     toFloat(p.Density.Text),
			DayLength:   toFloat(p.DayLength.Text),
			YearLength:  toFloat(p.YearLength.Text),
			Temperature: toInt(p.Temperature.Text),
			Water:       toInt(p.Water.Text),
			Composition: p.Composition.Text,
			LifeForm:    p.LifeForm.Text,
			Landmasses:  []PspLandMass{},
			Satellites:  []PspSatellite{},
			Event:       []PspEvent{},
			SmallMoons:  toInt(p.SmallMoons.Text),
			Ring:        p.Ring.Text,
		}

		// planet landmasses, format is "Name (capital city)" in the xml, not very xml-like I think
		landmasses := []PspLandMass{}
		for _, l := range p.LandMass {
			n, c := nameCapital(l.Text)

			landmasses = append(landmasses, PspLandMass{
				Name:    n,
				Capital: c,
			})
		}
		planet.Landmasses = landmasses

		// planet satellites
		satellites := []PspSatellite{}
		for _, s := range p.Satellites {

			satellites = append(satellites, PspSatellite{
				Name: s.Name.Text,
				Size: s.Size.Text,
				Icon: s.Icon,
			})
		}
		planet.Satellites = satellites

		// planet events
		planet.Event = planetEvents(events, planet.SysPos)

		planets = append(planets, planet)
	}
	ps.Planets = planets

	// marshall the yaml
	b, err := yaml.Marshal(ps)
	if err != nil {
		return fmt.Errorf("yaml marshal: %w", err)
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("file write: %w", err)
	}

	return err
}

func toInt(in string) int {
	if in == "" {
		return 0
	}
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func toFloat(in string) float64 {
	if in == "" {
		return 0
	}
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		panic(err)
	}
	return out
}

func nameCapital(in string) (name, capital string) {
	i := strings.LastIndex(in, `(`)
	var n, c string
	if i < 0 {
		n = strings.TrimSpace(in)
		c = ""
	} else {
		n = strings.TrimSpace(in[:i])
		// trim off the parens
		c = in[i+1 : len(in)-1]
	}

	return n, c
}

func planetEvents(events *Event, syspos int) []PspEvent {
	pEvents := []PspEvent{}
	// find the planet
	for _, pevs := range events.Planets {
		if toInt(pevs.SysPos.Text) == syspos {
			// build up the events
			for _, e := range pevs.Events {
				factions := strings.Split(e.Faction.Text, " ")
				pEvents = append(pEvents, PspEvent{
					Date: e.Date,
					Faction: struct {
						Source string   "yaml:\"source\""
						Value  []string "yaml:\"value\""
					}{
						Source: e.Faction.Source,
						Value:  factions,
					},
					Population:      toFloat(e.Population.Text),
					SocioIndustrial: e.SocioIndustrial.Text,
					Hpg:             e.Hpg.Text,
				})
			}
		}

	}
	return pEvents
}
