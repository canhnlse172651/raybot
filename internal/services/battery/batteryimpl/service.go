package batteryimpl

import (
	"context"
	"fmt"

	"github.com/tbe-team/raybot/internal/events"
	"github.com/tbe-team/raybot/internal/services/battery"
	"github.com/tbe-team/raybot/pkg/eventbus"
	"github.com/tbe-team/raybot/pkg/validator"
)

type service struct {
	validator validator.Validator

	publisher        eventbus.Publisher
	batteryStateRepo battery.BatteryStateRepository
	settingRepo      battery.SettingRepository
}

func NewService(
	validator validator.Validator,
	publisher eventbus.Publisher,
	repo battery.BatteryStateRepository,
	settingRepo battery.SettingRepository,
) battery.Service {
	return &service{
		validator:        validator,
		publisher:        publisher,
		batteryStateRepo: repo,
		settingRepo:      settingRepo,
	}
}

func (s service) GetBatteryState(ctx context.Context) (battery.BatteryState, error) {
	return s.batteryStateRepo.GetBatteryState(ctx)
}

func (s service) UpdateBatteryState(ctx context.Context, params battery.UpdateBatteryStateParams) error {
	if err := s.validator.Validate(params); err != nil {
		return fmt.Errorf("validate params: %w", err)
	}

	state, err := s.batteryStateRepo.UpdateBatteryState(ctx, params)
	if err != nil {
		return fmt.Errorf("update battery state: %w", err)
	}

	s.publisher.Publish(events.BatteryUpdatedTopic, eventbus.NewMessage(
		events.BatteryUpdatedEvent{
			BatteryState: state,
		},
	))

	return nil
}

func (s service) UpdateChargeSetting(ctx context.Context, params battery.UpdateChargeSettingParams) error {
	if err := s.validator.Validate(params); err != nil {
		return fmt.Errorf("validate params: %w", err)
	}

	return s.settingRepo.UpdateChargeSetting(ctx, params)
}

func (s service) UpdateDischargeSetting(ctx context.Context, params battery.UpdateDischargeSettingParams) error {
	if err := s.validator.Validate(params); err != nil {
		return fmt.Errorf("validate params: %w", err)
	}

	return s.settingRepo.UpdateDischargeSetting(ctx, params)
}
