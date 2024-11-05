package scenes

import (
	"concurrent-parking/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"sync"
	"time"
)

type MainScene struct {
	window fyne.Window
}

const NumSpaces = 20 // Número total de espacios de estacionamiento

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

var carsContainer = container.NewWithoutLayout()

func (s *MainScene) Show() {

	rectangle := canvas.NewRectangle(color.Transparent)
	rectangle.StrokeWidth = 2
	rectangle.StrokeColor = color.White
	rectangle.Resize(fyne.NewSize(140, 620))
	rectangle.Move(fyne.NewPos(220, 10))

	gate := canvas.NewRectangle(color.White)
	gate.Resize(fyne.NewSize(10, 100))
	gate.Move(fyne.NewPos(195, 300))

	for i := 0; i < NumSpaces; i++ {
		spaceRect := canvas.NewRectangle(color.Transparent)
		spaceRect.StrokeWidth = 2
		spaceRect.StrokeColor = color.White

		spaceRect.Resize(fyne.NewSize(60, 25))
		spaceRect.Move(fyne.NewPos(280, 15+float32(i*30)))

		carsContainer.Add(spaceRect)
	}

	carsContainer.Add(rectangle)
	carsContainer.Add(gate)

	s.window.SetContent(carsContainer)
}

func (s *MainScene) Run() {
	p := models.NewParking(make(chan int, 20), &sync.Mutex{})
	poissonDist := models.NewPoissonDist()

	var wg sync.WaitGroup

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			car := models.NewCar(id)
			carImage := car.GetCarImage()
			carImage.Resize(fyne.NewSize(50, 30))
			carImage.Move(fyne.NewPos(-20, 310))

			carsContainer.Add(carImage)
			carsContainer.Refresh()

			car.Park(p, carsContainer, &wg)
		}(i)
		var randPoissonNumber = poissonDist.Generate(float64(2))
		time.Sleep(time.Second * time.Duration(randPoissonNumber))
	}

	wg.Wait()
	fmt.Println("terminado simulacion ...")
}
