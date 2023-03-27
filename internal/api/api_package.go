package api

import (
	"context"
	"os"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc/metadata"

	"github.com/hablof/logistic-package-api/internal/model"
	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"
)

const (
	pbMetadataLogLevelKey = "log_level"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "logistic_package_api_not_found_total",
		Help: "Total number of packages that were not found",
	})
	totalCUDevents = promauto.NewCounter(prometheus.CounterOpts{
		Name: "logistic_package_api_cud_event_total",
		Help: "Total number of CUD events",
	})
)

type RepoCRUD interface {
	CreatePackage(ctx context.Context, pack *model.Package, logger zerolog.Logger) (uint64, error)
	DescribePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) (*model.Package, error)
	ListPackages(ctx context.Context, offset uint64, logger zerolog.Logger) ([]model.Package, error)
	RemovePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) error
}

type logisticPackageAPI struct {
	pb.UnimplementedLogisticPackageApiServiceServer
	repo              RepoCRUD
	logger            zerolog.Logger
	allowRiseLogLevel bool
}

// NewLogisticPackageAPI returns api of logistic-package-api service
func NewLogisticPackageAPI(r RepoCRUD, logLevelDebug bool, allowRiseLogLevel bool) pb.LogisticPackageApiServiceServer {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if logLevelDebug {
		l = l.Level(zerolog.DebugLevel)
	} else {
		l = l.Level(zerolog.InfoLevel)
	}

	return &logisticPackageAPI{
		repo:              r,
		logger:            l,
		allowRiseLogLevel: allowRiseLogLevel,
	}
}

func (o *logisticPackageAPI) shouldRiseDebugLevel(ctx context.Context) bool {

	if !o.allowRiseLogLevel {
		return false
	}

	logLevel := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		logLevels := md.Get(pbMetadataLogLevelKey)
		if len(logLevels) > 0 {
			logLevel = logLevels[0]
		}
	}
	if strings.ToLower(logLevel) == "debug" {
		return true
	}
	return false
}

func (o *logisticPackageAPI) setupLogger(ctx context.Context) zerolog.Logger {
	log := o.logger
	if o.shouldRiseDebugLevel(ctx) {
		log = log.Level(zerolog.DebugLevel)
	}

	if sc, ok := opentracing.SpanFromContext(ctx).Context().(jaeger.SpanContext); ok {
		log = log.With().Str("traceID", sc.TraceID().String()).Logger()
	}
	return log
}
