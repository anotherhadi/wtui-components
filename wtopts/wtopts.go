package wtopts

type Opts struct {
	// Styling
	// 256 Colors
	AccentColor    uint8
	TextColor      uint8
	SecondaryColor uint8

	MaxCols int
	MaxRows int
	Style   string
	Loop    bool

	// If you're using Standalone components
	LeftPadding int

	// Per component opts

	// Checkbox
	DefaultCheckbox []bool

	// Confirm
	DefaultConfirm bool
	Affirmative    string
	Negative       string

	// Number Picker
	DefaultNumber float64
	Decimals      int
	Increment     float64
	Maximum       float64
	Minimum       float64

	// RGB/RGBA Picker
	DefaultRed     uint8
	DefaultGreen   uint8
	DefaultBlue    uint8
	DefaultOpacity uint8

	// Combobox
	CaseSensitive bool
}

func GetOpts(customOpts ...Opts) Opts {
	if len(customOpts) > 0 {
		return customOpts[0]
	}
	return DefaultOpts()
}

func DefaultOpts() Opts {
	return Opts{
		AccentColor:    99,
		TextColor:      255,
		SecondaryColor: 237,

		MaxCols: 80,
		MaxRows: 40,
		Style:   "rounded",
		Loop:    true,

		LeftPadding: 2,

		DefaultCheckbox: []bool{},

		DefaultConfirm: true,
		Affirmative:    "Yes",
		Negative:       "No",

		DefaultNumber: 0,
		Decimals:      1,
		Increment:     0.1,
		Maximum:       100,
		Minimum:       -11,

		DefaultRed:     0,
		DefaultGreen:   0,
		DefaultBlue:    0,
		DefaultOpacity: 255,

		CaseSensitive: true,
	}
}
