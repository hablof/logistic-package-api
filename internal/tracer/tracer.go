package tracer

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"

	"github.com/hablof/logistic-package-api/internal/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type loggerJaeger struct {
}

// Error implements jaeger.Logger
func (loggerJaeger) Error(msg string) {
	log.Error().Err(fmt.Errorf("JAEGER ERROR: %s", msg))
}

// Infof implements jaeger.Logger
func (loggerJaeger) Infof(msg string, args ...interface{}) {
	log.Info().Msgf(fmt.Sprintf("JAEGER INFO: %s", msg), args...)
}

func (loggerJaeger) Debugf(msg string, args ...interface{}) {
	log.Debug().Msgf(fmt.Sprintf("JAEGER DEBUG: %s", msg), args...)
}

// NewTracer - returns new tracer.
func NewTracer(cfg *config.Config) (io.Closer, error) {
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}

	if cfg.Jaeger.IsRateLimiting {
		cfgTracer.Sampler.Type = jaeger.SamplerTypeRateLimiting
		cfgTracer.Sampler.Param = cfg.Jaeger.SpansPerSecond
	}

	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(loggerJaeger{}))
	if err != nil {
		log.Err(err).Msgf("failed init jaeger: %v", err)

		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	log.Info().Msgf("Traces started")

	return closer, nil
}
