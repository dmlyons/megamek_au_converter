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

type PlainFloat64 float64

func (f PlainFloat64) MarshalYAML() (any, error) {
	// Format with 'f' to avoid scientific notation
	return strconv.FormatFloat(float64(f), 'f', -1, 64), nil
}

func main() {
	systemFile := flag.String("systems", "systems.xml", "your systems.xml file")
	eventsFile := flag.String("events", "system_events.xml", "your system_events.xml file")
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

	log.Printf("loaded %d systems", len(systems.Systems))

	// read the events
	var events SystemEvents

	e, err := os.ReadFile(*eventsFile)
	if err != nil {
		log.Fatalf("read events %v", err)
	}
	if err := xml.Unmarshal(e, &events); err != nil {
		log.Fatalf("unmarshal systems %v", err)
	}

	log.Println("loaded events")

	// create the output directory
	err = os.MkdirAll(*outDir, 0777)
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

	log.Println("done processing systems")
}

func process(outDir *string, system *System, events *Event) error {
	// clean up the filename
	fn := strings.TrimSpace(system.ID.Text)
	fn = strings.ReplaceAll(fn, " ", "")
	fn = strings.ReplaceAll(fn, "/", "")
	fn = *outDir + "/" + fn + ".yaml"

	f, err := os.Create(fn)
	if err != nil {
		log.Printf("create yaml file %v", err)
	}
	defer f.Close()

	primarySlot, err := strconv.Atoi(system.PrimarySlot.Text)
	if err != nil {
		return fmt.Errorf("strconv: %w", err)
	}
	if system.SucsId == 0 {
		log.Printf("system %s has no sucsid", system.ID.Text)
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
			NadirCharge:  stringPtr(sysEvent.NadirCharge.Text),
			ZenithCharge: stringPtr(sysEvent.ZenithCharge.Text),
		})
	}
	ps.Events = evs

	// planets
	planets := []PsPlanet{}
	for _, p := range system.Planets {
		planet := PsPlanet{
			Name:        p.Name.Text,
			Type:        planetType(p.Type.Text),
			OrbitalDist: toFloat(p.OrbitalDist.Text),
			SysPos:      *toInt(p.SysPos.Text),
			Icon:        p.Icon.Text,
			Pressure:    pressure(p.Pressure.Text),
			Atmosphere:  atmosphere(p.Atmosphere.Text),
			Gravity:     toFloat(p.Gravity.Text),
			Diameter:    toFloat(p.Diameter.Text),
			Density:     toFloat(p.Density.Text),
			DayLength:   toFloat(p.DayLength.Text),
			YearLength:  toFloat(p.YearLength.Text),
			Temperature: toInt(p.Temperature.Text),
			Water:       toInt(p.Water.Text),
			Desc:        stringPtr(p.Desc.Text),
			Composition: stringPtr(p.Composition.Text),
			LifeForm:    lifeForm(p.LifeForm.Text),
			Landmasses:  []PspLandMass{},
			Satellites:  []PspSatellite{},
			Event:       []PspEvent{},
			SmallMoons:  toInt(p.SmallMoons.Text),
			Ring:        stringPtr(p.Ring.Text),
		}

		// planet landmasses, format is "Name (capital city)" in the xml, not very xml-like I think
		landmasses := []PspLandMass{}
		for _, l := range p.LandMass {
			n, c := nameCapital(l.Text)

			landmasses = append(landmasses, PspLandMass{
				Name:    *n,
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

func toInt(in string) *int {
	if in == "" {
		return nil
	}
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return &out
}

func toFloat(in string) PlainFloat64 {
	if in == "" {
		return 0.0
	}
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		panic(err)
	}
	return PlainFloat64(out)
}

func toFloatPtr(in string) *PlainFloat64 {
	if in == "" {
		return nil
	}
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		panic(err)
	}
	pf := PlainFloat64(out)
	return &pf
}

func nameCapital(in string) (name, capital *string) {
	i := strings.LastIndex(in, `(`)
	var n, c string
	if i < 0 {
		n = strings.TrimSpace(in)
		return &n, nil
	} else {
		n = strings.TrimSpace(in[:i])
		// trim off the parens
		c = in[i+1 : len(in)-1]
	}
	return &n, &c
}

func planetEvents(events *Event, syspos int) []PspEvent {
	pEvents := []PspEvent{}
	// find the planet
	for _, pevs := range events.Planets {
		if *toInt(pevs.SysPos.Text) == syspos {
			// build up the events
			for _, e := range pevs.Events {
				var factions *PspFaction
				if e.Faction.Text != "" {
					fcts := strings.Split(e.Faction.Text, " ")
					factions = &PspFaction{
						Source: e.Faction.Source,
						Value:  fcts,
					}
				}
				pEvents = append(pEvents, PspEvent{
					Date:            e.Date,
					Faction:         factions,
					Population:      toFloatPtr(e.Population.Text),
					SocioIndustrial: stringPtr(e.SocioIndustrial.Text),
					Hpg:             stringPtr(e.Hpg.Text),
				})
			}
		}

	}
	return pEvents
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// planetType converts the planet type from the xml to the yaml format
func planetType(in string) string {
	switch in {
	case "Asteroid Belt":
		return "ASTEROID_BELT"
	case "Dwarf Terrestrial":
		return "DWARF_TERRESTRIAL"
	case "Gas Giant":
		return "GAS_GIANT"
	case "Giant Terrestrial":
		return "GIANT_TERRESTRIAL"
	case "Ice Giant":
		return "ICE_GIANT"
	case "Terrestrial":
		return "TERRESTRIAL"
	}
	log.Printf("unknown planet type %s", in)
	return ""
}

// pressure converts the pressure from the xml to the yaml format
func pressure(in string) string {
	switch in {
	case "High":
		return "HIGH"
	case "Low":
		return "THIN"
	case "Normal", "Standard":
		return "STANDARD"
	case "Trace":
		return "TRACE"
	case "Vacuum":
		return "VACUUM"
	case "Very High":
		return "VERY_HIGH"
	case "":
		return ""
	}
	log.Printf(`unknown pressure "%s"`, in)
	return ""
}

// atmosphere converts the atmosphere from the xml to the yaml format
func atmosphere(in string) string {
	switch in {
	case "Breathable":
		return "BREATHABLE"
	case "None":
		return "NONE"
	case "Tainted":
		return "TAINTEDPOISON"
	case "Tainted (Poisonous)":
		return "TAINTEDPOISON"
	case "Toxic (Caustic)":
		return "TOXICCAUSTIC"
	case "Toxic (Poisonous)":
		return "TOXICPOISON"
	case "":
		return ""
	}
	log.Printf("unknown atmosphere %s", in)
	return ""
}

// lifeForm converts the life form from the xml to the yaml format
func lifeForm(in string) *string {
	if in == "" || in == "NONE" {
		return nil
	}

	var lifeForm string

	switch in {
	case "Amphibians", "AMPH":
		lifeForm = "AMPHIBIAN"
	case "Birds", "BIRD":
		lifeForm = "BIRD"
	case "Fish", "FISH":
		lifeForm = "FISH"
	case "Insects", "INSECT":
		lifeForm = "INSECT"
	case "Mammals", "MAMMAL":
		lifeForm = "MAMMAL"
	case "Microbes", "MICROBE":
		lifeForm = "MICROBE"
	case "Plants", "PLANT":
		lifeForm = "PLANT"
	case "Reptiles", "REPTILE":
		lifeForm = "REPTILE"
	default:
		log.Printf("unknown life form %s", in)
	}
	return &lifeForm
}
