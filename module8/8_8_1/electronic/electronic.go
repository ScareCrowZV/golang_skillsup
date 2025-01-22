package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonCount() int
}

type Smartphone interface {
	OS() string
}

type applePhone struct {
	model string
}

func NewApplePhone(model string) *applePhone {
	ap := new(applePhone)
	ap.model = model

	return ap
}

func (p applePhone) Brand() string {
	return "Apple"
}

func (p applePhone) Model() string {
	return p.model
}

func (p applePhone) Type() string {
	return "smartphone"
}

func (p applePhone) OS() string {
	return "iOS"
}

type androidPhone struct {
	brand string
	model string
}

func NewAndroidPhone(brand string, model string) *androidPhone {
	ap := new(androidPhone)
	ap.brand = brand
	ap.model = model

	return ap
}

func (p androidPhone) Brand() string {
	return p.brand
}

func (p androidPhone) Model() string {
	return p.model
}

func (p androidPhone) Type() string {
	return "smartphone"
}

func (p androidPhone) OS() string {
	return "Android"
}

type radioPhone struct {
	brand       string
	model       string
	buttonCount int
}

func NewRadioPhone(brand string, model string, buttonCount int) *radioPhone {
	ap := new(radioPhone)
	ap.brand = brand
	ap.model = model
	ap.buttonCount = buttonCount

	return ap
}

func (p radioPhone) Brand() string {
	return p.brand
}

func (p radioPhone) Model() string {
	return p.model
}

func (p radioPhone) Type() string {
	return "station"
}

func (p radioPhone) ButtonCount() int {
	return p.buttonCount
}
