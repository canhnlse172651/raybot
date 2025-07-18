package command

import (
	"context"
	"time"

	"github.com/tbe-team/raybot/pkg/paging"
	"github.com/tbe-team/raybot/pkg/sort"
	"github.com/tbe-team/raybot/pkg/xerror"
)

var (
	ErrCommandNotFound                    = xerror.NotFound(nil, "command.notFound", "command not found")
	ErrNoNextExecutableCommand            = xerror.NotFound(nil, "command.noNextExecutable", "no next executable command")
	ErrNoCommandBeingProcessed            = xerror.BadRequest(nil, "command.noCommandBeingProcessed", "no command being processed")
	ErrCommandInProcessingCanNotBeDeleted = xerror.BadRequest(nil, "command.inProcessingCanNotBeDeleted", "command in processing can not be deleted")
	ErrCommandAlreadyExists               = xerror.Conflict(nil, "command.alreadyExists", "command already exists")

	ErrRunningCommandNotFound = xerror.NotFound(nil, "command.runningCommandNotFound", "running command not found")
	ErrRunningCommandExists   = xerror.BadRequest(nil, "command.runningCommandExists", "running command already exists")
)

type CreateCommandParams struct {
	Source    Source  `validate:"enum"`
	Inputs    Inputs  `validate:"required"`
	RequestID *string `validate:"omitempty,max=64"` // Optional request ID for idempotency
}

type GetCommandByIDParams struct {
	CommandID int64 `validate:"required,min=1"`
}

type ListCommandsParams struct {
	PagingParams paging.Params `validate:"required"`
	Sorts        []sort.Sort   `validate:"sort=type status source created_at updated_at completed_at"`
	Statuses     []Status      `validate:"dive,enum"`
}

type DeleteCommandByIDParams struct {
	CommandID int64 `validate:"required,min=1"`
}

type Service interface {
	GetCommandByID(ctx context.Context, params GetCommandByIDParams) (Command, error)
	GetCurrentProcessingCommand(ctx context.Context) (Command, error)
	ListCommands(ctx context.Context, params ListCommandsParams) (paging.List[Command], error)
	CreateCommand(ctx context.Context, params CreateCommandParams) (Command, error)
	CancelCurrentProcessingCommand(ctx context.Context) error

	// CancelActiveCloudCommands cancels all QUEUED and PROCESSING commands created by the cloud.
	CancelActiveCloudCommands(ctx context.Context) error

	RunNextExecutableCommand(ctx context.Context) error

	// CancelAllRunningCommands cancels all running commands including the current processing command
	// and the commands in the queue.
	CancelAllRunningCommands(ctx context.Context) error

	DeleteCommandByID(ctx context.Context, params DeleteCommandByIDParams) error
	DeleteOldCommands(ctx context.Context) error
}

type ExecutorService interface {
	Execute(ctx context.Context, cmd Command) error
}

type UpdateCommandParams struct {
	ID             int64
	Status         Status
	SetStatus      bool
	Outputs        Outputs
	SetOutputs     bool
	Error          *string
	SetError       bool
	StartedAt      *time.Time
	SetStartedAt   bool
	CompletedAt    *time.Time
	SetCompletedAt bool
	UpdatedAt      time.Time
}

type Repository interface {
	ListCommands(ctx context.Context, params ListCommandsParams) (paging.List[Command], error)
	GetNextExecutableCommand(ctx context.Context) (Command, error)
	GetCurrentProcessingCommand(ctx context.Context) (Command, error)
	GetCommandByID(ctx context.Context, id int64) (Command, error)
	CreateCommand(ctx context.Context, command Command) (Command, error)
	UpdateCommand(ctx context.Context, params UpdateCommandParams) (Command, error)

	// CancelPendingCommands cancels all pending commands by status QUEUED, PROCESSING, and CANCELING.
	CancelPendingCommands(ctx context.Context) error
	CancelQueuedAndProcessingCommandsCreatedByCloud(ctx context.Context) error

	// DeleteCommandByID deletes a command by id.
	// It does not delete the command if the status is PROCESSING, CANCELING.
	DeleteCommandByID(ctx context.Context, id int64) error
	DeleteOldCommands(ctx context.Context, cutoffTime time.Time) error
}

type RunningCommandRepository interface {
	Get(ctx context.Context) (CancelableCommand, error)
	Add(ctx context.Context, cmd CancelableCommand) error
	Update(ctx context.Context, cmd CancelableCommand) error
	Remove(ctx context.Context) error
}

// ProcessingLock is responsible for controlling a lock mechanism
// that prevents the system from automatically picking up and processing the next command in the queue.
//
// Note: This lock does not handle stopping or canceling a command that is already being executed.
// Its sole purpose is to block the transition to the next command,
// ensuring that no new command is started while the lock is held.
type ProcessingLock interface {
	// WithLock acquires the lock and executes the function.
	// The lock is released when the function returns.
	WithLock(fn func() error) error

	// WaitUntilUnlocked blocks the execution until the lock is released.
	// If the context is canceled, the function returns immediately.
	WaitUntilUnlocked(ctx context.Context) error
}
