package main

import (
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/karimElmougi/GolangBoy/gameboy"
	"github.com/karimElmougi/GolangBoy/gpu"
	"github.com/karimElmougi/GolangBoy/joypad"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	romName := args[len(args)-1]

	gameboy.Boot(romName)
	go func() {
		ticker := time.NewTicker(time.Second / 60)
		for range ticker.C {
			gameboy.Run()
		}
	}()
	go func() {
		ticker := time.NewTicker(time.Millisecond)
		for range ticker.C {
			joypad.UpdateInputs()
		}
	}()
	err := ebiten.Run(run, 160, 144, 2, "GolangBoy")
	if err != nil {
		panic(err)
	}
}

func run(screen *ebiten.Image) error {
	img, _ := ebiten.NewImageFromImage(gpu.Frame, ebiten.FilterDefault)
	screen.DrawImage(img, &ebiten.DrawImageOptions{})
	return nil
}
