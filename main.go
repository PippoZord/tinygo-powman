package main

import (
	"machine"
	"powman/powman"
	"time"
	"unsafe"
)

func reg32(addr uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(addr))
}

const PADS_BANK0_BASE uintptr = 0x40038000

/*
Unlock GPIOs
*/
func unlockGPIOs() {
	for i := uintptr(0); i < 48; i++ {
		regAddr := PADS_BANK0_BASE + 0x04 + (i * 4)
		*reg32(regAddr) &^= (1 << 8) // Pulisce bit 8 (ISO)
	}
}

func blink() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for i := 0; i < 6; i++ {
		led.High()
		time.Sleep(100 * time.Millisecond)
		led.Low()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	unlockGPIOs()

	powman.PowmanInit(1704067200000)

	blink()

	powman.PowmanOffForMs(10000)
}
