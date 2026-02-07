package powman

import (
	"device/arm"
	"fmt"
)

// This function inizializes the clock setting absTimeMs as starting point
func PowmanInit(absTimeMs uint64) {
	fmt.Printf("Initializing time %d\n", absTimeMs)

	// Stop timer using TIMER register
	*reg32(POWMAN_BASE + TIMER) = PASS | 0

	// Set time
	*reg32(POWMAN_BASE + SET_TIME_15TO0) = PASS | uint32(absTimeMs&0xFFFF)
	*reg32(POWMAN_BASE + SET_TIME_31TO16) = PASS | uint32((absTimeMs>>16)&0xFFFF)
	*reg32(POWMAN_BASE + SET_TIME_47TO32) = PASS | uint32((absTimeMs>>32)&0xFFFF)
	*reg32(POWMAN_BASE + SET_TIME_63TO48) = PASS | uint32((absTimeMs>>48)&0xFFFF)

	// Start timer (RUN bit)
	// Clear timer(CLEAR bit)
	// Clear Alarm (ALARM bit)
	*reg32(POWMAN_BASE + TIMER) = PASS | 0x46

	// Ignore debugger
	*reg32(POWMAN_BASE + DBG_PWRCFG) = PASS | 0x01

}

// Put the system in low power mode setting also the reboot after an alarm is triggered
func powmanPowerOff() {

	// Set the power low mode
	*reg32(POWMAN_BASE + VREG_LP_ENTRY) = PASS | 0x0004

	forceReboot()

	// Switch off system
	// Bit 3: SWCORE, Bit 2: XIP, Bit 1: SRAM0, Bit 0: SRAM1
	*reg32(POWMAN_BASE + STATE) = PASS | 0x00F0

	// Waiting interrupt/alarm
	arm.Asm("wfi")
	for {
	}
}

// Force the low power setting an alarm
// The Alarm triggers system after sleesleepingMs
func PowmanOffForMs(sleepingMs uint64) {

	alarmTime := sleepingMs + getCurrentTime()
	fmt.Printf("Going to sleep for %dms\n", sleepingMs)

	//Enable Interrupt
	*reg32(POWMAN_BASE + INTE) = PASS | uint32(0x02)
	//Stop Timer
	*reg32(POWMAN_BASE + TIMER) = PASS | uint32(0x00)

	//Writing Alarm
	*reg32(POWMAN_BASE + ALARM_TIME_15TO0) = PASS | uint32(alarmTime&0xFFFF)
	*reg32(POWMAN_BASE + ALARM_TIME_31TO16) = PASS | uint32((alarmTime>>16)&0xFFFF)
	*reg32(POWMAN_BASE + ALARM_TIME_47TO32) = PASS | uint32((alarmTime>>32)&0xFFFF)
	*reg32(POWMAN_BASE + ALARM_TIME_63TO48) = PASS | uint32((alarmTime>>48)&0xFFFF)

	// Start Timer and resetting Alarm BIT
	*reg32(POWMAN_BASE + TIMER) = PASS | 0x72

	powmanPowerOff()
}

// Force the system to reboot after
func forceReboot() {
	*reg32(POWMAN_BASE + BOOT0) = 0
	*reg32(POWMAN_BASE + BOOT1) = 0
	*reg32(POWMAN_BASE + BOOT2) = 0
	*reg32(POWMAN_BASE + BOOT3) = 0
}

/*
This function reads current time from READ_TIME_UPPER and READ_TIME_LOWER register
*/
func getCurrentTime() uint64 {
	var now uint64
	for {
		hi1 := uint64(*reg32(POWMAN_BASE + READ_TIME_UPPER))
		lo := uint64(*reg32(POWMAN_BASE + READ_TIME_LOWER))
		hi2 := uint64(*reg32(POWMAN_BASE + READ_TIME_UPPER))
		if hi1 == hi2 {
			now = (hi1 << 32) | lo
			break
		}
	}
	return now
}
