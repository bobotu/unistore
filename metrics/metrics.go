package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "unistore"
	raft      = "raft"
)

var (
	RaftWriterWait = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "writer_wait",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})
	WriteWaiteStepOne = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "writer_wait_step_1",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})
	WriteWaiteStepTwo = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "writer_wait_step_2",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})
	WriteWaiteStepThree = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "writer_wait_step_3",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})
	WriteWaiteStepFour = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "writer_wait_step_4",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})

	RaftDBUpdate = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "raft_db_update",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})
	KVDBUpdate = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "kv_db_update",
			Buckets:   prometheus.ExponentialBuckets(0.001, 1.5, 20),
		})
	LockUpdate = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "lock_update",
			Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 15),
		})
	LatchWait = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "latch_wait",
			Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 15),
		})
	RaftBatchSize = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: raft,
			Name:      "batch_size",
			Buckets:   prometheus.ExponentialBuckets(1, 1.5, 20),
		})
)

func init() {
	prometheus.MustRegister(RaftWriterWait)
	prometheus.MustRegister(WriteWaiteStepOne)
	prometheus.MustRegister(WriteWaiteStepTwo)
	prometheus.MustRegister(WriteWaiteStepThree)
	prometheus.MustRegister(WriteWaiteStepFour)
	prometheus.MustRegister(RaftDBUpdate)
	prometheus.MustRegister(KVDBUpdate)
	prometheus.MustRegister(LockUpdate)
	prometheus.MustRegister(RaftBatchSize)
	prometheus.MustRegister(LatchWait)
	http.Handle("/metrics", prometheus.Handler())
}
