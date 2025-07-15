package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/tbe-team/raybot/internal/config"
	"github.com/tbe-team/raybot/internal/handlers/http/gen"
	configmocks "github.com/tbe-team/raybot/internal/services/config/mocks"
)

func TestConfigHandler_GetLogConfig(t *testing.T) {
	t.Run("Should get log config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetLogConfig(mock.Anything).Return(config.Log{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/log", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get log config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetLogConfig(mock.Anything).Return(config.Log{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/log", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateLogConfig(t *testing.T) {
	validLogConfig := gen.LogConfig{
		Console: gen.LogConsoleHandler{
			Enable: true,
			Level:  "DEBUG",
			Format: "JSON",
		},
		File: gen.LogFileHandler{
			Enable:        true,
			Path:          "/tmp/raybot.log",
			RotationCount: 10,
			Level:         "DEBUG",
			Format:        "JSON",
		},
	}

	t.Run("Should update log config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateLogConfig(mock.Anything, mock.MatchedBy(func(cfg config.Log) bool {
			return cfg.Console.Enable && cfg.Console.Level == slog.LevelDebug && cfg.Console.Format == config.LogFormatJSON &&
				cfg.File.Enable && cfg.File.Path == "/tmp/raybot.log" && cfg.File.RotationCount == 10 &&
				cfg.File.Level == slog.LevelDebug && cfg.File.Format == config.LogFormatJSON
		})).
			Return(config.Log{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validLogConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/log", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update log config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateLogConfig(mock.Anything, mock.Anything).
			Return(config.Log{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validLogConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/log", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_GetHardwareConfig(t *testing.T) {
	t.Run("Should get hardware config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetHardwareConfig(mock.Anything).Return(config.Hardware{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/hardware", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get hardware config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetHardwareConfig(mock.Anything).Return(config.Hardware{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/hardware", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateHardwareConfig(t *testing.T) {
	validHardwareConfig := gen.HardwareConfig{
		Esp: gen.ESPConfig{
			Serial: gen.SerialConfig{
				Port:        "COM1",
				BaudRate:    115200,
				DataBits:    8,
				StopBits:    1,
				Parity:      "NONE",
				ReadTimeout: 10,
			},
			CommandAckTimeout: 10,
		},
		Pic: gen.PICConfig{
			Serial: gen.SerialConfig{
				Port:        "COM2",
				BaudRate:    115200,
				DataBits:    8,
				StopBits:    1,
				Parity:      "NONE",
				ReadTimeout: 10,
			},
			CommandAckTimeout: 10,
		},
		Leds: gen.LedsConfig{
			System: gen.LedConfig{
				Pin: "GPIO2",
			},
			Alert: gen.LedConfig{
				Pin: "GPIO3",
			},
		},
	}

	t.Run("Should update hardware config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateHardwareConfig(mock.Anything, mock.Anything).
			Return(config.Hardware{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validHardwareConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/hardware", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update hardware config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateHardwareConfig(mock.Anything, mock.Anything).
			Return(config.Hardware{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validHardwareConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/hardware", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_GetCloudConfig(t *testing.T) {
	t.Run("Should get cloud config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetCloudConfig(mock.Anything).Return(config.Cloud{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/cloud", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get cloud config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetCloudConfig(mock.Anything).Return(config.Cloud{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/cloud", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateCloudConfig(t *testing.T) {
	validCloudConfig := gen.CloudConfig{
		Enable:  true,
		Address: "http://localhost:8080",
		Token:   "1234567890",
	}

	t.Run("Should update cloud config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateCloudConfig(mock.Anything, mock.Anything).
			Return(config.Cloud{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validCloudConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/cloud", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update cloud config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateCloudConfig(mock.Anything, mock.Anything).
			Return(config.Cloud{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validCloudConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/cloud", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_GetHTTPConfig(t *testing.T) {
	t.Run("Should get http config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetHTTPConfig(mock.Anything).Return(config.HTTP{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/http", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get http config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetHTTPConfig(mock.Anything).Return(config.HTTP{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/http", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateHTTPConfig(t *testing.T) {
	validHTTPConfig := gen.HTTPConfig{
		Port:    8080,
		Swagger: true,
	}

	t.Run("Should update http config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateHTTPConfig(mock.Anything, mock.Anything).
			Return(config.HTTP{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validHTTPConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/http", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update http config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateHTTPConfig(mock.Anything, mock.Anything).
			Return(config.HTTP{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validHTTPConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/http", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_GetWifiConfig(t *testing.T) {
	t.Run("Should get wifi config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetWifiConfig(mock.Anything).Return(config.Wifi{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/wifi", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get wifi config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetWifiConfig(mock.Anything).Return(config.Wifi{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/wifi", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateWifiConfig(t *testing.T) {
	validWifiConfig := gen.WifiConfig{
		Ap: gen.APConfig{
			Enable:   true,
			Ssid:     "test",
			Password: "1234567890",
			Ip:       "192.168.1.1",
		},
		Sta: gen.STAConfig{
			Enable:   true,
			Ssid:     "test",
			Password: "1234567890",
		},
	}

	t.Run("Should update wifi config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateWifiConfig(mock.Anything, mock.Anything).
			Return(config.Wifi{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validWifiConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/wifi", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update wifi config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateWifiConfig(mock.Anything, mock.Anything).
			Return(config.Wifi{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validWifiConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/wifi", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_GetCommandConfig(t *testing.T) {
	t.Run("Should get command config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetCommandConfig(mock.Anything).Return(config.Command{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/command", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get command config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetCommandConfig(mock.Anything).Return(config.Command{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/command", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateCommandConfig(t *testing.T) {
	validCommandConfig := gen.CommandConfig{
		CargoLift: gen.CargoLiftConfig{
			StableReadCount: 3,
		},
		CargoLower: gen.CargoLowerConfig{
			StableReadCount: 3,
		},
	}

	t.Run("Should update command config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateCommandConfig(mock.Anything, mock.Anything).
			Return(config.Command{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validCommandConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/command", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update command config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateCommandConfig(mock.Anything, mock.Anything).
			Return(config.Command{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validCommandConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/command", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_GetBatteryMonitoringConfig(t *testing.T) {
	t.Run("Should get battery monitoring config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetBatteryMonitoringConfig(mock.Anything).Return(config.BatteryMonitoring{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/monitoring/battery", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to get battery monitoring config if fetching failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().GetBatteryMonitoringConfig(mock.Anything).Return(config.BatteryMonitoring{}, errors.New("fetching failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		req := httptest.NewRequest(http.MethodGet, "/api/v1/configs/monitoring/battery", nil)
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestConfigHandler_UpdateBatteryMonitoringConfig(t *testing.T) {
	validBatteryMonitoringConfig := gen.BatteryMonitoringConfig{
		VoltageLow: gen.BatteryVoltageLowConfig{
			Enable:    true,
			Threshold: 14.0,
		},
		VoltageHigh: gen.BatteryVoltageHighConfig{
			Enable:    true,
			Threshold: 18.0,
		},
		CellVoltageHigh: gen.BatteryCellVoltageHighConfig{
			Enable:    true,
			Threshold: 4.3,
		},
		CellVoltageLow: gen.BatteryCellVoltageLowConfig{
			Enable:    true,
			Threshold: 3.8,
		},
		CellVoltageDiff: gen.BatteryCellVoltageDiffConfig{
			Enable:    true,
			Threshold: 0.5,
		},
		CurrentHigh: gen.BatteryCurrentHighConfig{
			Enable:    true,
			Threshold: 6.0,
		},
		TempHigh: gen.BatteryTempHighConfig{
			Enable:    true,
			Threshold: 60.0,
		},
		PercentLow: gen.BatteryPercentLowConfig{
			Enable:    true,
			Threshold: 20.0,
		},
		HealthLow: gen.BatteryHealthLowConfig{
			Enable:    true,
			Threshold: 60.0,
		},
	}

	t.Run("Should update battery monitoring config successfully", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateBatteryMonitoringConfig(mock.Anything, mock.MatchedBy(func(cfg config.BatteryMonitoring) bool {
			return cfg.BatteryVoltageLow.Enable && cfg.BatteryVoltageLow.Threshold == 14.0 &&
				cfg.BatteryVoltageHigh.Enable && cfg.BatteryVoltageHigh.Threshold == 18.0 &&
				cfg.BatteryCellVoltageHigh.Enable && cfg.BatteryCellVoltageHigh.Threshold == 4.3 &&
				cfg.BatteryCellVoltageLow.Enable && cfg.BatteryCellVoltageLow.Threshold == 3.8 &&
				cfg.BatteryCellVoltageDiff.Enable && cfg.BatteryCellVoltageDiff.Threshold == 0.5 &&
				cfg.BatteryCurrentHigh.Enable && cfg.BatteryCurrentHigh.Threshold == 6.0 &&
				cfg.BatteryTempHigh.Enable && cfg.BatteryTempHigh.Threshold == 60.0 &&
				cfg.BatteryPercentLow.Enable && cfg.BatteryPercentLow.Threshold == 20.0 &&
				cfg.BatteryHealthLow.Enable && cfg.BatteryHealthLow.Threshold == 60.0
		})).
			Return(config.BatteryMonitoring{}, nil)

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validBatteryMonitoringConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/monitoring/battery", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Should not able to update battery monitoring config if updating failed", func(t *testing.T) {
		configService := configmocks.NewFakeService(t)
		configService.EXPECT().UpdateBatteryMonitoringConfig(mock.Anything, mock.Anything).
			Return(config.BatteryMonitoring{}, errors.New("updating failed"))

		h := SetupAPITestHandler(t, func(hs *Service) {
			hs.configService = configService
		})

		body := validBatteryMonitoringConfig
		jsonBody, err := json.Marshal(body)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/v1/configs/monitoring/battery", bytes.NewBuffer(jsonBody))
		rec := httptest.NewRecorder()

		h.ServeHTTP(rec, req)
		require.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
