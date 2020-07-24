package viber

// Keyboard struct
type Keyboard struct {
	Type          string   `json:"Type"`
	DefaultHeight bool     `json:"DefaultHeight,omitempty"`
	BgColor       string   `json:"BgColor,omitempty"`
	Buttons       []Button `json:"Buttons"`

	// api level 4
	InputFieldState InputFieldState `json:"InputFieldState,omitempty"`
}

type InputFieldState string

const (
	RegularInputField   InputFieldState = "regular"
	MinimizedInputField InputFieldState = "minimized"
	HiddenInputField    InputFieldState = "hidden"
)

// AddButton to keyboard
func (k *Keyboard) AddButtons(b ...Button) {
	k.Buttons = append(k.Buttons, b...)
}

// NewKeyboard struct with attribs init
func (v *Viber) NewKeyboard(bgcolor string, defaultHeight bool) *Keyboard {
	return &Keyboard{
		Type:          "keyboard",
		DefaultHeight: defaultHeight,
		BgColor:       bgcolor,
	}
}
