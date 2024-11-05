package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
)

type Parking struct {
	spaces      chan int
	entrance    *sync.Mutex
	spacesArray [20]bool
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
	return &Parking{
		spaces:      spaces,
		entrance:    entrance,
		spacesArray: [20]bool{},
	}
}

func (p *Parking) GetSpaces() chan int {
	return p.spaces
}

func (p *Parking) GetEntrance() *sync.Mutex {
	return p.entrance
}

func (p *Parking) GetSpacesArray() [20]bool {
	return p.spacesArray
}

func (p *Parking) SetSpacesArray(spacesArray [20]bool) {
	p.spacesArray = spacesArray
}

func (p *Parking) ExitQueue(carsContainer *fyne.Container, carImage *canvas.Image) {
	carImage.Move(fyne.NewPos(205, 350))
	carsContainer.Add(carImage)
	carsContainer.Refresh()
}
