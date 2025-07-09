package system

import (
	"context"
	"time"
)

type Service interface {
	Reboot(ctx context.Context) error

	// StopEmergency stops all motors and cancel all queued and processing commands.
	StopEmergency(ctx context.Context) error

	// GetInfo returns the system information.
	GetInfo(ctx context.Context) (Info, error)

	// GetStatus returns the system status.
	GetStatus(ctx context.Context) (Status, error)

	// SetStatusError sets the system status to error.
	SetStatusError(ctx context.Context) error
}

type SysInfoCollectorService interface {
	Run(ctx context.Context)
	Stop()
}

type UpdateInfoParams struct {
	LocalIP        string
	SetLocalIP     bool
	CPUUsage       float64
	SetCPUUsage    bool
	MemoryUsage    float64
	SetMemoryUsage bool
	TotalMemory    uint64
	SetTotalMemory bool
	Uptime         time.Duration
	SetUptime      bool
}

type Repository interface {
	GetInfo(ctx context.Context) (Info, error)
	UpdateInfo(ctx context.Context, params UpdateInfoParams) error
	GetStatus(ctx context.Context) (Status, error)
	UpdateStatus(ctx context.Context, status Status) error
}
