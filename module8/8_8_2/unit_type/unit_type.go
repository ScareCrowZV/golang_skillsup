package unittype

type UnitType string

const (
	Inch              UnitType = "inch"
	CM                UnitType = "cm"
	InchToCentimeters float64  = 0.393701
	CentimetersToInch float64  = 2.54
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {

	value := u.Value

	if t != u.T {
		switch u.T {
		case "inch":
			value = u.Value * CentimetersToInch
		case "cm":
			value = u.Value * InchToCentimeters
		default:
			return 0
		}

	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}
