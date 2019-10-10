// Mandlebrot generates a png image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"log"
	"math/cmplx"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
)

func main() {
	go func() {
		//ops := new(op.Ops)
		w := app.NewWindow()
		gtx := &layout.Context{
			Queue: w.Queue(),
		}
		var button = new(int)
		var b = new(int)

		for e := range w.Events() {

			switch e := e.(type) {
			case app.UpdateEvent:
				gtx.Reset(&e.Config, e.Size)
				key.InputOp{Key: button}.Add(gtx.Ops)
				pointer.RectAreaOp{ // HLevent
					Rect: image.Rectangle{Max: image.Point{X: 500, Y: 500}}, // HLevent
				}.Add(gtx.Ops) // HLevent
				pointer.InputOp{Key: b}.Add(gtx.Ops) // HLevent
				var inc uint8
				for _, bev := range gtx.Events(b) {
					log.Printf("%T\n", bev)
					if bev, ok := bev.(pointer.Event); ok { // HLevent
						switch bev.Type { // HLevent
						case pointer.Press: // HLevent
							//b.pressed = true // HLevent
							log.Printf("Pressed: ")
							log.Println(bev)
						case pointer.Release: // HLevent
							log.Printf("Released: ")
							log.Println(bev)
							//b.pressed = false // HLevent
						}
					}
				}

				for _, ev := range gtx.Events(button) {
					log.Printf("%T\n", ev)
					log.Println(ev)
					switch ev := ev.(type) {
					case key.Event:
						log.Println(ev.Name)
						log.Println(ev.Modifiers)
						if ev.Name == rune(67) && ev.Modifiers == 1 {
							os.Exit(1)

						}
					case key.EditEvent:
						log.Println(ev.Text)

					}
				}
				//c := &e.Config

				img := createFractal(inc, e.Size)
				//cs := layout.RigidConstraints(e.Size)
				widget.Image{Src: img, Rect: img.Bounds()}.Layout(gtx)
				op.InvalidateOp{At: gtx.Now().Add(1 * time.Second)}.Add(gtx.Ops)
				log.Println("bla")
				w.Update(gtx.Ops)
				inc += 10

			}

		}
	}()
	app.Main()

}
func createFractal(inc uint8, size image.Point) image.Image {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	width := size.X
	height := size.Y
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(width)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// image point (px,py) rpresents complex value z
			img.Set(px, py, mandelbrot(z, inc))
		}
	}
	return img
	/*
		file, err := os.Create("Mandelbrot.png")
		if err != nil {
			fmt.Println("error in file")
			os.Exit(0)
		}

		png.Encode(file, img)
		file.Close()
	*/
}

func mandelbrot(z complex128, inc uint8) color.Color {
	const iterations = 200
	const contrast = 25

	var v complex128
	for n := uint8(8); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//rgba := image.NewRGBA(image.Rect(255, 233, int(n), contrast))
			return color.RGBA{(128 + inc) % 255, (233 + inc) % 255, (uint8(n) + inc) % 255, 255 - contrast*n}
			//return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
