package main

type PsEvent struct {
	Date         string     `yaml:"date"`
	NadirCharge  *YesNoBool `yaml:"nadirCharge,omitempty"`
	ZenithCharge *YesNoBool `yaml:"zenithCharge,omitempty"`
}

type PsSourceWithValue struct {
	Source string `yaml:"source"`
	Value  any    `yaml:"value"`
}

func (p PsSourceWithValue) MarshalYAML() (any, error) {
	if p.Source == "" {
		return p.Value, nil
	}
	// Use an alias to avoid recursion
	type Alias PsSourceWithValue
	return Alias(p), nil
}

// YesNoBool is a custom type to render boolean values as "yes" or "no" in YAML.
type YesNoBool bool

func (y YesNoBool) MarshalYAML() (any, error) {
	if y {
		return "yes", nil
	}
	return "no", nil
}

type PsPlanet struct {
	Name        PsSourceWithValue `yaml:"name"`
	Type        PsSourceWithValue `yaml:"type"`
	OrbitalDist float64           `yaml:"orbitalDist"`
	SysPos      int               `yaml:"sysPos"`
	Icon        PsSourceWithValue `yaml:"icon"`
	Pressure    PsSourceWithValue `yaml:"pressure"`
	Atmosphere  string            `yaml:"atmosphere"`
	Gravity     float64           `yaml:"gravity"`
	Diameter    float64           `yaml:"diameter"`
	Density     float64           `yaml:"density"`
	DayLength   float64           `yaml:"dayLength"`
	YearLength  float64           `yaml:"yearLength"`
	Temperature *int              `yaml:"temperature,omitempty"`
	Water       *int              `yaml:"water,omitempty"`
	Desc        *string           `yaml:"desc,omitempty"`
	Composition *string           `yaml:"composition,omitempty"`
	LifeForm    *string           `yaml:"lifeForm,omitempty"`
	Landmasses  []PspLandMass     `yaml:"landmass,omitempty"`
	Satellites  []PspSatellite    `yaml:"satellite,omitempty"`
	Event       []PspEvent        `yaml:"event,omitempty"`
	SmallMoons  *int              `yaml:"smallMoons,omitempty"`
	Ring        *YesNoBool        `yaml:"ring,omitempty"`
}

type PspSatellite struct {
	Name string `yaml:"name"`
	Size string `yaml:"size"`
	Icon string `yaml:"icon"`
}

type PspLandMass struct {
	Name    string  `yaml:"name"`
	Capital *string `yaml:"capital,omitempty"`
}

type PspEvent struct {
	Date            string      `yaml:"date"`
	Faction         *PspFaction `yaml:"faction,omitempty"`
	Population      *float64    `yaml:"population,omitempty"`
	SocioIndustrial *string     `yaml:"socioIndustrial,omitempty"`
	Hpg             *string     `yaml:"hpg,omitempty"`
}

type PspFaction struct {
	Source string   `yaml:"source"`
	Value  []string `yaml:"value"`
}

type PlanetarySystem struct {
	ID           string            `yaml:"id"`
	SucsID       int               `yaml:"sucsId"`
	Xcood        float64           `yaml:"xcood"`
	Ycood        float64           `yaml:"ycood"`
	SpectralType PsSourceWithValue `yaml:"spectralType"`
	PrimarySlot  PsSourceWithValue `yaml:"primarySlot"`
	Events       []PsEvent         `yaml:"event"`
	Planets      []PsPlanet        `yaml:"planet"`
}
