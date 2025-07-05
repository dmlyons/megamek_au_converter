package main

type PsEvent struct {
	Date         string `yaml:"date"`
	NadirCharge  string `yaml:"nadirCharge"`
	ZenithCharge string `yaml:"zenithCharge"`
}

type PsPlanet struct {
	Name        string         `yaml:"name"`
	Type        string         `yaml:"type"`
	OrbitalDist float64        `yaml:"orbitalDist"`
	SysPos      int            `yaml:"sysPos"`
	Icon        string         `yaml:"icon"`
	Pressure    string         `yaml:"pressure"`
	Atmosphere  string         `yaml:"atmosphere"`
	Gravity     float64        `yaml:"gravity"`
	Diameter    float64        `yaml:"diameter"`
	Density     float64        `yaml:"density"`
	DayLength   float64        `yaml:"dayLength"`
	YearLength  float64        `yaml:"yearLength"`
	Temperature int            `yaml:"temperature,omitempty"`
	Water       int            `yaml:"water,omitempty"`
	Composition string         `yaml:"composition,omitempty"`
	LifeForm    string         `yaml:"lifeForm,omitempty"`
	Landmasses  []PspLandMass  `yaml:"landmass,omitempty"`
	Satellites  []PspSatellite `yaml:"satellite,omitempty"`
	Event       []PspEvent     `yaml:"event,omitempty"`
	SmallMoons  int            `yaml:"smallMoons,omitempty"`
	Ring        string         `yaml:"ring,omitempty"`
}

type PspSatellite struct {
	Name string `yaml:"name"`
	Size string `yaml:"size"`
	Icon string `yaml:"icon"`
}

type PspLandMass struct {
	Name    string `yaml:"name"`
	Capital string `yaml:"capital,omitempty"`
}

type PspEvent struct {
	Date    string `yaml:"date"`
	Faction struct {
		Source string   `yaml:"source"`
		Value  []string `yaml:"value"`
	} `yaml:"faction,omitempty"`
	Population      float64 `yaml:"population,omitempty"`
	SocioIndustrial string  `yaml:"socioIndustrial,omitempty"`
	Hpg             string  `yaml:"hpg,omitempty"`
}

type PlanetarySystem struct {
	ID           string     `yaml:"id"`
	SucsID       int        `yaml:"sucsId"`
	Xcood        float64    `yaml:"xcood"`
	Ycood        float64    `yaml:"ycood"`
	SpectralType string     `yaml:"spectralType"`
	PrimarySlot  int        `yaml:"primarySlot"`
	Events       []PsEvent  `yaml:"event"`
	Planets      []PsPlanet `yaml:"planet"`
}
