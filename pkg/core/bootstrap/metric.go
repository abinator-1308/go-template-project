package bootstrap

import (
	"github.com/devlibx/gox-base/metrics"
	"net/http"
)
import "github.com/devlibx/gox-metrics/provider/multi"
import "github.com/devlibx/gox-metrics/provider/prometheus"
import "github.com/devlibx/gox-metrics/provider/statsd"

type MetricHandler struct {
	MetricsReporter metrics.Reporter
}

func NewMetricService(metricConfig *metrics.Config) (metrics.Scope, *MetricHandler, error) {
	var toRet metrics.Scope
	var err error
	if metricConfig.Enabled {

		// Setup default if missing & build metric
		metricConfig.SetupDefaults()

		// Use correct type of metric exporter
		if metricConfig.EnablePrometheus && metricConfig.EnableStatsd {
			toRet, err = multi.NewRootScope(*metricConfig)
		} else if metricConfig.EnableStatsd {
			toRet, err = statsd.NewRootScope(*metricConfig)
		} else {
			toRet, err = prometheus.NewRootScope(*metricConfig)
		}
	} else {
		toRet, err = metrics.NoOpMetric(), nil
	}

	// Setup a reporter
	var mh metrics.Reporter
	if reporter, ok := toRet.(metrics.Reporter); ok {
		mh = reporter
	} else {
		mh = &MetricHandler{}
	}

	return toRet, &MetricHandler{MetricsReporter: mh}, err
}

func (mh *MetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mh *MetricHandler) HTTPHandler() http.Handler {
	return mh
}
