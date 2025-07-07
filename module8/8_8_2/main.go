package main

import (
	"fmt"
	u "unittype"
)

func main() {

	autosDimensions := map[string]map[string]float64{
		"BMW": { //Z8
			"length": 440.0,
			"width":  183.0,
			"heigth": 131.7,
		},
		"Mercedes": { //s-class
			"length": 546.9,
			"width":  192.1,
			"heigth": 151.0,
		},
		"Dodge": { //charger
			"length": 524.8,
			"width":  202.7,
			"heigth": 149.7,
		},
	}

	var bmw BMW
	var mercedes Mercedes
	var dodge Dodge

	var unitLength u.Unit
	var unitWidth u.Unit
	var unitHeigth u.Unit

	for auto, dimensions := range autosDimensions {
		if auto == "BMW" {
			for name, value := range dimensions {
				if name == "length" {
					unitLength = u.NewUnit(value, u.CM) //u.Unit{Value: value, T: u.CM}
				}
				if name == "width" {
					unitWidth = u.Unit{Value: value, T: u.CM}
				}
				if name == "heigth" {
					unitHeigth = u.Unit{Value: value, T: u.CM}
				}

			}
			centimetersDimension := newCentimetersDimension(unitLength, unitWidth, unitHeigth)
			bmw = *newBMW("Z8", centimetersDimension, 250, 400)
		}

		if auto == "Mercedes" {
			for name, value := range dimensions {
				if name == "length" {
					unitLength = u.Unit{Value: value, T: u.CM}
				}
				if name == "width" {
					unitWidth = u.Unit{Value: value, T: u.CM}
				}
				if name == "heigth" {
					unitHeigth = u.Unit{Value: value, T: u.CM}
				}

			}
			centimetersDimension := newCentimetersDimension(unitLength, unitWidth, unitHeigth)
			mercedes = *newMercedes("S-Class", centimetersDimension, 250, 612)
		}

		if auto == "Dodge" {
			for name, value := range dimensions {
				if name == "length" {
					unitLength = u.Unit{Value: value, T: u.CM}
				}
				if name == "width" {
					unitWidth = u.Unit{Value: value, T: u.CM}
				}
				if name == "heigth" {
					unitHeigth = u.Unit{Value: value, T: u.CM}
				}

			}
			inchDimension := newInchDimension(unitLength, unitWidth, unitHeigth)
			dodge = *newDodge("Charger", inchDimension, 290, 492)
		}
	}

	fmt.Printf(
		"\n\n===%s===\nModel: %s\nMaxSpeed: %d\nHorsePower: %d\nDimensions(L x W x H): %.1f x %.1f x %.1f %s",
		bmw.Brand(),
		bmw.Model(),
		bmw.MaxSpeed(),
		bmw.EnginePower(),
		bmw.dimensions.Length().Get(u.CM),
		bmw.dimensions.Width().Get(u.CM),
		bmw.dimensions.Height().Get(u.CM),
		u.CM,
	)
	fmt.Printf(
		"\n\n===%s===\nModel: %s\nMaxSpeed: %d\nHorsePower: %d\nDimensions(L x W x H): %.1f x %.1f x %.1f %s",
		mercedes.Brand(),
		mercedes.Model(),
		mercedes.MaxSpeed(),
		mercedes.EnginePower(),
		bmw.dimensions.Length().Get(u.CM),
		bmw.dimensions.Width().Get(u.CM),
		bmw.dimensions.Height().Get(u.CM),
		u.CM,
	)
	fmt.Printf(
		"\n\n===%s===\nModel: %s\nMaxSpeed: %d\nHorsePower: %d\nDimensions(L x W x H): %.1f x %.1f x %.1f %s",
		dodge.Brand(),
		dodge.Model(),
		dodge.MaxSpeed(),
		dodge.EnginePower(),
		bmw.dimensions.Length().Get(u.Inch),
		bmw.dimensions.Width().Get(u.Inch),
		bmw.dimensions.Height().Get(u.Inch),
		u.Inch,
	)

}

type centimetersDimension struct {
	length u.Unit
	width  u.Unit
	height u.Unit
}

func newCentimetersDimension(length, width, height u.Unit) *centimetersDimension {
	d := new(centimetersDimension)
	d.length = length
	d.width = width
	d.height = height

	return d
}

func (d centimetersDimension) Length() u.Unit {
	return d.length
}

func (d centimetersDimension) Width() u.Unit {
	return d.width
}

func (d centimetersDimension) Height() u.Unit {
	return d.height
}

type inchDimension struct {
	length u.Unit
	width  u.Unit
	height u.Unit
}

func newInchDimension(length, width, height u.Unit) *inchDimension {
	d := new(inchDimension)
	d.length = length
	d.width = width
	d.height = height

	return d
}

func (d inchDimension) Length() u.Unit {
	return d.length
}

func (d inchDimension) Width() u.Unit {
	return d.width
}

func (d inchDimension) Height() u.Unit {
	return d.height
}

type BMW struct {
	model       string
	dimensions  u.Dimensions
	maxspeed    int
	enginepower int
}

func newBMW(model string, dimensions u.Dimensions, maxspeed, enginepower int) *BMW {
	car := new(BMW)
	car.model = model
	car.dimensions = dimensions
	car.maxspeed = maxspeed
	car.enginepower = enginepower

	return car

}

func (car BMW) Brand() string {
	return "BMW"
}

func (car BMW) Model() string {
	return car.model
}

func (car BMW) Dimensions() u.Dimensions {
	return car.dimensions
}

func (car BMW) MaxSpeed() int {
	return car.maxspeed
}

func (car BMW) EnginePower() int {
	return car.enginepower
}

type Mercedes struct {
	model       string
	dimensions  u.Dimensions
	maxspeed    int
	enginepower int
}

func newMercedes(model string, dimensions u.Dimensions, maxspeed, enginepower int) *Mercedes {
	car := new(Mercedes)
	car.model = model
	car.dimensions = dimensions
	car.maxspeed = maxspeed
	car.enginepower = enginepower

	return car

}

func (car Mercedes) Brand() string {
	return "Mercedes"
}

func (car Mercedes) Model() string {
	return car.model
}

func (car Mercedes) Dimensions() u.Dimensions {
	return car.dimensions
}

func (car Mercedes) MaxSpeed() int {
	return car.maxspeed
}

func (car Mercedes) EnginePower() int {
	return car.enginepower
}

type Dodge struct {
	model       string
	dimensions  u.Dimensions
	maxspeed    int
	enginepower int
}

func newDodge(model string, dimensions u.Dimensions, maxspeed, enginepower int) *Dodge {
	car := new(Dodge)
	car.model = model
	car.dimensions = dimensions
	car.maxspeed = maxspeed
	car.enginepower = enginepower

	return car

}

func (car Dodge) Brand() string {
	return "Dodge"
}

func (car Dodge) Model() string {
	return car.model
}

func (car Dodge) Dimensions() u.Dimensions {
	return car.dimensions
}

func (car Dodge) MaxSpeed() int {
	return car.maxspeed
}

func (car Dodge) EnginePower() int {
	return car.enginepower
}
