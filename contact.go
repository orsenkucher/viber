package viber

// Contact struct
type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

// NewKeyboard struct with attribs init
func (v *Viber) NewContact(name string, phone string) (Contact, error) {
	// TODO: Validate phone
	return Contact{
		Name:        name,
		PhoneNumber: phone,
	}, nil
}
