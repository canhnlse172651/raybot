package battery

import "time"

//nolint:revive
type BatteryState struct {
	Current      uint16 // unit: mA
	Temp         uint8
	Voltage      uint16   // unit: mV
	CellVoltages []uint16 // unit: mV
	Percent      uint8
	Fault        uint8
	Health       uint8
	UpdatedAt    time.Time
}

type ChargeSetting struct {
	CurrentLimit uint16
	Enabled      bool
	UpdatedAt    time.Time
}

type DischargeSetting struct {
	CurrentLimit uint16
	Enabled      bool
	UpdatedAt    time.Time
}
