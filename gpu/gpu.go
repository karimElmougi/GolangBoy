package gpu

import (
	"image"
	"image/color"

	"github.com/karimElmougi/GolangBoy/interrupts"
	"github.com/karimElmougi/GolangBoy/mmu"
)

type gpuMode uint8

const (
	HBLANK gpuMode = 0x00
	VBLANK gpuMode = 0x01
	OAM    gpuMode = 0x02
	VRAM   gpuMode = 0x03
)

const (
	LCD_STATUS_ADDR  = 0xff41
	SCAN_LINE_ADDR   = 0xff44
	GPU_CONTROL_ADDR = 0xff40
)

var (
	palette = [4]color.RGBA{
		color.RGBA{255, 255, 255, 255},
		color.RGBA{192, 192, 192, 255},
		color.RGBA{96, 96, 96, 255},
		color.RGBA{0, 0, 0, 255}}
	clock       = 456
	Frame       = image.NewRGBA(image.Rect(0, 0, 160, 144))
	framebuffer = image.NewRGBA(image.Rect(0, 0, 160, 144))
)

func Step(cyclesEllapsed uint64) {
	if !isLcdOn() {
		clock = 456
		mmu.Write(SCAN_LINE_ADDR, 0)
		setMode(HBLANK)
		return
	}

	scanLine := mmu.Read(SCAN_LINE_ADDR)
	previousMode := getMode()
	requestInterrupt := false

	if scanLine >= 144 {
		setMode(VBLANK)
		requestInterrupt = isVblankModeInterruptEnabled()
	} else if clock >= 456-80 {
		setMode(OAM)
		requestInterrupt = isOamModeInterruptEnabled()
	} else if clock >= 456-80-172 {
		setMode(VRAM)
	} else {
		setMode(HBLANK)
		requestInterrupt = isHblankModeInterruptEnabled()
	}

	if requestInterrupt && getMode() != previousMode {
		interrupts.WriteLcdInterrupt()
	}

	if scanLine == mmu.Read(0xff45) {
		setCoincidenceStatus()
		if isCoincidenceInterruptEnabled() {
			interrupts.WriteLcdInterrupt()
		}
	} else {
		resetCoincidenceStatus()
	}

	clock -= int(cyclesEllapsed)
	if clock <= 0 {
		clock += 456
		mmu.RAM[SCAN_LINE_ADDR]++
		scanLine = mmu.Read(SCAN_LINE_ADDR)
		if scanLine == 144 {
			interrupts.WriteVblankInterrupt()
		} else if scanLine > 153 {
			Frame = framebuffer
			framebuffer = image.NewRGBA(image.Rect(0, 0, 160, 144))
			mmu.Write(SCAN_LINE_ADDR, 0)
			renderScanLine(0)
		} else if scanLine < 144 {
			renderScanLine(scanLine)
		}
	}
}

func isHblankModeInterruptEnabled() bool {
	return mmu.Read(LCD_STATUS_ADDR)&0x08 == 0x08
}

func isVblankModeInterruptEnabled() bool {
	return mmu.Read(LCD_STATUS_ADDR)&0x10 == 0x10
}

func isOamModeInterruptEnabled() bool {
	return mmu.Read(LCD_STATUS_ADDR)&0x20 == 0x20
}

func isCoincidenceInterruptEnabled() bool {
	return mmu.Read(LCD_STATUS_ADDR)&0x40 == 0x40
}

func getMode() gpuMode {
	status := mmu.Read(LCD_STATUS_ADDR)
	mode := status & 0x03
	return gpuMode(mode)
}

func setMode(mode gpuMode) {
	status := mmu.Read(LCD_STATUS_ADDR)
	status &= 0xfc
	mmu.Write(LCD_STATUS_ADDR, status|uint8(mode))
}

func setCoincidenceStatus() {
	status := mmu.Read(LCD_STATUS_ADDR)
	mmu.Write(LCD_STATUS_ADDR, status|0x04)
}

func resetCoincidenceStatus() {
	status := mmu.Read(LCD_STATUS_ADDR)
	mmu.Write(LCD_STATUS_ADDR, status&0xfb)
}

func isLcdOn() bool {
	return mmu.Read(GPU_CONTROL_ADDR)&0x80 == 0x80
}

func renderScanLine(currentLine uint8) {
	control := mmu.Read(GPU_CONTROL_ADDR)
	if control&0x01 == 0x01 {
		renderTiles(control, currentLine)
	}
	if control&0x02 == 0x02 {
		renderSprites(control, int(currentLine))
	}
}

func renderTiles(control uint8, currentLine uint8) {
	unsig := false
	tileData := uint16(0x8800)
	scrollY := mmu.Read(0xff42)
	scrollX := mmu.Read(0xff43)
	windowY := mmu.Read(0xff4a)
	windowX := mmu.Read(0xff4b) - 7
	usingWindow := false

	if control&0x20 == 0x20 {
		if windowY <= mmu.Read(SCAN_LINE_ADDR) {
			usingWindow = true
		}
	}

	if control&0x10 == 0x10 {
		tileData = 0x8000
		unsig = true
	}

	testMask := uint8(0x08)
	if usingWindow {
		testMask = 0x40
	}

	backgroundMemory := uint16(0x9800)
	if control&testMask == testMask {
		backgroundMemory = 0x9C00
	}

	yPos := uint8(0)
	if !usingWindow {
		yPos = scrollY + currentLine
	} else {
		yPos = currentLine - windowY
	}

	var tileRow = uint16(yPos/8) * 32

	for pixel := uint8(0); pixel < 160; pixel++ {
		xPos := pixel + scrollX

		if usingWindow && pixel >= windowX {
			xPos = pixel - windowX
		}
		tileCol := uint16(xPos / 8)
		tileAddress := backgroundMemory + tileRow + tileCol
		tileLocation := tileData
		if unsig {
			tileNum := uint16(mmu.RAM[tileAddress])
			tileLocation = tileLocation + uint16(tileNum*16)
		} else {
			tileNum := int16(int8(mmu.RAM[tileAddress]))
			tileLocation = uint16(int32(tileLocation) + int32((tileNum+128)*16))
		}

		var line = (yPos % 8) * 2
		data1 := mmu.RAM[tileLocation+uint16(line)]
		data2 := mmu.RAM[tileLocation+uint16(line)+1]

		colourBit := uint8(int8((xPos%8)-7) * -1)
		colourNum := (((data2 >> colourBit) & 1) << 1) | ((data1 >> colourBit) & 1)

		framebuffer.Set(int(pixel), int(currentLine), getColour(colourNum, 0xff47))
	}
}

func renderSprites(control uint8, currentLine int) {
	ySize := 8
	if control&0x04 == 0x04 {
		ySize = 16
	}

	for sprite := 0; sprite < 40; sprite++ {
		index := sprite * 4
		yPos := int(mmu.Read(uint16(0xfe00+index))) - 16
		xPos := mmu.Read(uint16(0xfe00+index+1)) - 8
		tileLocation := mmu.Read(uint16(0xfe00 + index + 2))
		attributes := mmu.Read(uint16(0xfe00 + index + 3))

		if currentLine < yPos || currentLine >= (yPos+ySize) {
			continue
		}

		line := currentLine - yPos
		if attributes&0x40 == 0x40 {
			line = (line - ySize) * -1
		}

		dataAddress := (uint16(tileLocation) * 16) + uint16(line*2) + 0x8000
		data1 := mmu.Read(dataAddress)
		data2 := mmu.Read(dataAddress + 1)

		for tilePixel := uint8(0); tilePixel < 8; tilePixel++ {
			colourBit := tilePixel
			if attributes&0x20 == 0x20 {
				colourBit = uint8(int8(colourBit-7) * -1)
			}

			colourNum := (((data2 >> colourBit) & 1) << 1) | ((data1 >> colourBit) & 1)

			if colourNum == 0 {
				continue
			}

			pixel := int(xPos) + int(7-tilePixel)
			if pixel >= 0 && pixel < 160 {
				priority := attributes&0x80 != 0x80
				bgTileColour := framebuffer.At(pixel, int(currentLine))
				if priority || bgTileColour == palette[0] {
					paletteAddr := uint16(0xff48)
					if attributes&0x10 == 0x10 {
						paletteAddr = 0xff49
					}
					framebuffer.Set(pixel, int(currentLine), getColour(colourNum, paletteAddr))
				}
			}
		}
	}
}

func getColour(colourNum uint8, addr uint16) color.RGBA {
	customPalette := mmu.Read(addr)

	var i uint8
	switch colourNum {
	case 0:
		i = customPalette&0x02 | customPalette&0x01
	case 1:
		i = customPalette&0x08>>2 | customPalette&0x04>>2
	case 2:
		i = customPalette&0x20>>4 | customPalette&0x10>>4
	case 3:
		i = customPalette&0x80>>6 | customPalette&0x40>>6
	}

	return palette[i]
}
