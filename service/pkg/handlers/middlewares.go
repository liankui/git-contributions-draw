package handlers

import (
	"github.com/chaos-io/gokit/middleware"
	"github.com/chaos-io/gokit/tracing"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"go.opentelemetry.io/otel/trace"

	"github.com/liankui/git-contributions-draw/service/pkg/svc"

	// this service api
	pb "github.com/liankui/git-contributions-draw/go/git-contributions-draw/v1"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	var tracer trace.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(trace.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	// var validator *middleware.Validator
	// if value, ok := options["validator"]; ok && value != nil {
	// 	validator = value.(*middleware.Validator)
	// }

	{ // GetDraw
		if tracer != nil {
			in.GetDrawEndpoint = tracing.TraceServer(tracer, "GetDraw")(in.GetDrawEndpoint)
		}
		if count != nil && latency != nil {
			in.GetDrawEndpoint = middleware.Instrumenting(latency.With("method", "GetDraw"), count.With("method", "GetDraw"))(in.GetDrawEndpoint)
		}
		// if validator != nil {
		// 	in.GetDrawEndpoint = validator.Validate()(in.GetDrawEndpoint)
		// }
	}

	return in
}

func WrapService(in pb.GitContributionsDrawServer, options map[string]interface{}) pb.GitContributionsDrawServer {
	return in
}
