// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sql

import (
	"context"
	"time"

	"github.com/TBD54566975/ftl/backend/controller/sql/sqltypes"
	"github.com/TBD54566975/ftl/internal/model"
)

type Querier interface {
	CreateCronJob(ctx context.Context, arg CreateCronJobParams) error
	EndCronJob(ctx context.Context, nextExecution time.Time, key model.CronJobKey, startTime time.Time) (EndCronJobRow, error)
	GetCronJobs(ctx context.Context) ([]GetCronJobsRow, error)
	GetStaleCronJobs(ctx context.Context, dollar_1 sqltypes.Duration) ([]GetStaleCronJobsRow, error)
	StartCronJobs(ctx context.Context, keys []string) ([]StartCronJobsRow, error)
}

var _ Querier = (*Queries)(nil)
