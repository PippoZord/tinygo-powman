package powman

import "unsafe"

const (
	//POWMAN BASE ADDRESS
	POWMAN_BASE uintptr = 0x40100000
	//PASSWORD FOR POWMAN
	PASS uint32 = 0x5AFE0000

	// GENERIC REGISTER OFFSET
	VREG_LP_ENTRY uintptr = 0x10
	STATE         uintptr = 0x38
	TIMER         uintptr = 0x88
	INTE          uintptr = 0xE4

	// OFFSET FOR TIME REGISTER
	SET_TIME_15TO0  uintptr = 0x6c
	SET_TIME_31TO16 uintptr = 0x68
	SET_TIME_47TO32 uintptr = 0x64
	SET_TIME_63TO48 uintptr = 0x60

	// OFFSET FOR ALARM REGISTER
	ALARM_TIME_15TO0  uintptr = 0x84
	ALARM_TIME_31TO16 uintptr = 0x80
	ALARM_TIME_47TO32 uintptr = 0x7C
	ALARM_TIME_63TO48 uintptr = 0x78

	// OFFSET FOR READING TIME
	READ_TIME_UPPER uintptr = 0x70
	READ_TIME_LOWER uintptr = 0x74

	// OFFSET FOR BOOT REGISTER
	BOOT0 uintptr = 0xD0
	BOOT1 uintptr = 0xD4
	BOOT2 uintptr = 0xD8
	BOOT3 uintptr = 0xDC

	// OFFSET POWER CONFIGURATION
	DBG_PWRCFG uintptr = 0xA4
)

func reg32(addr uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(addr))
}
