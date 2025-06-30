package standalone

import (
	"fmt"
	"sync"

	"github.com/tbe-team/raybot/internal/application"
	"github.com/tbe-team/raybot/internal/handlers/picserial"
)

func startPICSerial(app *application.Application, interruptChan <-chan any, readyWg *sync.WaitGroup) error {
	service := picserial.New(
		app.Cfg.Hardware.PIC,
		app.Log,
		app.PICSerialClient,
		app.EventBus,
		app.BatteryService,
		app.DistanceSensorService,
		app.LiftMotorService,
		app.DriveMotorService,
		app.LimitSwitchService,
		app.CargoService,
		app.AppStateService,
	)

	cleanup, err := service.Run(app.Context)
	if err != nil {
		return fmt.Errorf("error running PIC serial service: %w", err)
	}

	app.Log.Info("pic serial service started")

	readyWg.Done()
	<-interruptChan

	app.Log.Debug("pic serial service is shutting down")

	if err := cleanup(app.Context); err != nil {
		return fmt.Errorf("error cleaning up PIC serial service: %w", err)
	}

	app.Log.Debug("pic serial service stopped")

	return nil
}
