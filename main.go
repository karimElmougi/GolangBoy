package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/karimElmougi/GolangBoy/gameboy"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	romName := args[len(args)-1]

	gameboy.Boot(romName)
	err := ebiten.Run(gameboy.Run, 160, 144, 2, "GolangBoy")
	if err != nil {
		panic(err)
	}
}
