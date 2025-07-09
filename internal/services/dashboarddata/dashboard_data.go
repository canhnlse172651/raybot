package dashboarddata

import (
	"context"

	"github.com/tbe-team/raybot/internal/services/appstate"
	"github.com/tbe-team/raybot/internal/services/battery"
	"github.com/tbe-team/raybot/internal/services/cargo"
	"github.com/tbe-team/raybot/internal/services/distancesensor"
	"github.com/tbe-team/raybot/internal/services/drivemotor"
	"github.com/tbe-team/raybot/internal/services/led"
	"github.com/tbe-team/raybot/internal/services/liftmotor"
	"github.com/tbe-team/raybot/internal/services/location"
)

type RobotState struct {
	Battery          battery.BatteryState
	BatteryCharge    battery.ChargeSetting
	BatteryDischarge battery.DischargeSetting
	DistanceSensor   distancesensor.DistanceSensorState
	LiftMotor        liftmotor.LiftMotorState
	DriveMotor       drivemotor.DriveMotorState
	Location         location.Location
	Cargo            cargo.Cargo
	CargoDoorMotor   cargo.DoorMotorState
	AppState         appstate.AppState
	Leds             led.LedsOutput
}

type Service interface {
	GetRobotState(ctx context.Context) (RobotState, error)
}
