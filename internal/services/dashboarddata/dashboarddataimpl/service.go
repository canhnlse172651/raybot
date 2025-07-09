package dashboarddataimpl

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/tbe-team/raybot/internal/services/appstate"
	"github.com/tbe-team/raybot/internal/services/battery"
	"github.com/tbe-team/raybot/internal/services/cargo"
	"github.com/tbe-team/raybot/internal/services/dashboarddata"
	"github.com/tbe-team/raybot/internal/services/distancesensor"
	"github.com/tbe-team/raybot/internal/services/drivemotor"
	"github.com/tbe-team/raybot/internal/services/led"
	"github.com/tbe-team/raybot/internal/services/liftmotor"
	"github.com/tbe-team/raybot/internal/services/location"
)

type service struct {
	batteryStateRepo   battery.BatteryStateRepository
	batterySettingRepo battery.SettingRepository
	distanceSensorRepo distancesensor.DistanceSensorStateRepository
	liftMotorRepo      liftmotor.LiftMotorStateRepository
	driveMotorRepo     drivemotor.DriveMotorStateRepository
	locationRepo       location.Repository
	cargoRepo          cargo.Repository
	appStateRepo       appstate.Repository
	ledRepo            led.Repository
}

func NewService(
	batteryStateRepo battery.BatteryStateRepository,
	batterySettingRepo battery.SettingRepository,
	distanceSensorRepo distancesensor.DistanceSensorStateRepository,
	liftMotorRepo liftmotor.LiftMotorStateRepository,
	driveMotorRepo drivemotor.DriveMotorStateRepository,
	locationRepo location.Repository,
	cargoRepo cargo.Repository,
	appStateRepo appstate.Repository,
	ledRepo led.Repository,
) dashboarddata.Service {
	return &service{
		batteryStateRepo:   batteryStateRepo,
		batterySettingRepo: batterySettingRepo,
		distanceSensorRepo: distanceSensorRepo,
		liftMotorRepo:      liftMotorRepo,
		driveMotorRepo:     driveMotorRepo,
		locationRepo:       locationRepo,
		cargoRepo:          cargoRepo,
		appStateRepo:       appStateRepo,
		ledRepo:            ledRepo,
	}
}

func (s *service) GetRobotState(ctx context.Context) (dashboarddata.RobotState, error) {
	var ret dashboarddata.RobotState
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		ret.Battery, err = s.batteryStateRepo.GetBatteryState(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.BatteryCharge, err = s.batterySettingRepo.GetChargeSetting(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.BatteryDischarge, err = s.batterySettingRepo.GetDischargeSetting(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.DistanceSensor, err = s.distanceSensorRepo.GetDistanceSensorState(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.LiftMotor, err = s.liftMotorRepo.GetLiftMotorState(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.DriveMotor, err = s.driveMotorRepo.GetDriveMotorState(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.Location, err = s.locationRepo.GetLocation(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.Cargo, err = s.cargoRepo.GetCargo(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.CargoDoorMotor, err = s.cargoRepo.GetCargoDoorMotorState(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.AppState, err = s.appStateRepo.GetAppState(ctx)
		return err
	})

	g.Go(func() error {
		var err error
		ret.Leds, err = s.ledRepo.GetLeds(ctx)
		return err
	})

	if err := g.Wait(); err != nil {
		return dashboarddata.RobotState{}, fmt.Errorf("error group wait: %w", err)
	}

	return ret, nil
}
