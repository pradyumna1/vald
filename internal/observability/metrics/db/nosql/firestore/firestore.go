package firestore

import (
	"github.com/vdaas/vald/internal/db/nosql/cassandra"
	"github.com/vdaas/vald/internal/observability/metrics"
)

type Observer interface {
	metrics.Metric
	cassandra.QueryObserver
}
