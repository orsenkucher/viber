package viber

// Contact struct
type Location struct {
	Lat float64 `json:"lat"`
	Iot float64 `json:"iot"`
}

// NewKeyboard struct with attribs init
func (v *Viber) NewLocation(lat float64, iot float64) (Location, error) {
	// TODO: Validate lat & iot
	return Location{
		Lat: lat,
		Iot: iot,
	}, nil
}
