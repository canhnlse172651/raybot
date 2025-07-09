package systemimpl

import (
	"context"
	"sync"

	"github.com/tbe-team/raybot/internal/services/system"
)

type repository struct {
	mu         sync.RWMutex
	systemInfo system.Info
	status     system.Status
}

func NewRepository() system.Repository {
	return &repository{
		systemInfo: system.Info{},
		status:     system.StatusNormal,
	}
}

func (r *repository) GetInfo(_ context.Context) (system.Info, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.systemInfo, nil
}

func (r *repository) UpdateInfo(_ context.Context, params system.UpdateInfoParams) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if params.SetLocalIP {
		r.systemInfo.LocalIP = params.LocalIP
	}
	if params.SetCPUUsage {
		r.systemInfo.CPUUsage = params.CPUUsage
	}
	if params.SetMemoryUsage {
		r.systemInfo.MemoryUsage = params.MemoryUsage
	}
	if params.SetTotalMemory {
		r.systemInfo.TotalMemory = params.TotalMemory
	}
	if params.SetUptime {
		r.systemInfo.Uptime = params.Uptime
	}

	return nil
}

func (r *repository) GetStatus(_ context.Context) (system.Status, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.status, nil
}

func (r *repository) UpdateStatus(_ context.Context, status system.Status) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.status = status

	return nil
}
