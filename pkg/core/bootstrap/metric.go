package bootstrap

import (
	"fmt"
	"github.com/devlibx/gox-base/config"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-base/util"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)
import "github.com/devlibx/gox-metrics/provider/multi"
import "github.com/devlibx/gox-metrics/provider/prometheus"
import "github.com/devlibx/gox-metrics/provider/statsd"

type MetricHandler struct {
	MetricsReporter metrics.Reporter
}

func NewMetricService(metricConfig *metrics.Config, appConfig config.App) (metrics.Scope, *MetricHandler, error) {
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
	}

	// Setup DD
	host := metricConfig.Tracing.DD.Host
	port := metricConfig.Tracing.DD.Port
	env := metricConfig.Tracing.DD.Env
	if !metricConfig.Tracing.Enabled {
		fmt.Println("************* datadog tracer is not enabled *************")
	} else if util.IsStringEmpty(host) || port <= 0 {
		fmt.Println("datadog tracer is not enabled - host or port not provided")
	} else {
		agentAddr := fmt.Sprintf("%s:%d", host, port)
		fmt.Println("Setting datadog tracer", "host=", host, "post=", port, "url=", agentAddr)
		if util.IsStringEmpty(env) {
			env = "dev"
		}

		// Set global tracer
		t := opentracer.New(tracer.WithAgentAddr(agentAddr), tracer.WithServiceName(appConfig.AppName), tracer.WithEnv(env))
		opentracing.SetGlobalTracer(t)
	}

	return toRet, &MetricHandler{MetricsReporter: mh}, err
}

func (mh *MetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mh.MetricsReporter == nil {
		w.WriteHeader(200)
	} else {
		mh.MetricsReporter.HTTPHandler().ServeHTTP(w, r)
	}
}

func (mh *MetricHandler) HTTPHandler() http.Handler {
	return mh
}
