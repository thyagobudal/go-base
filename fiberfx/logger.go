package fiberfx

import (
	"context"
	"os"

	"go.elastic.co/apm"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(serviceName, environment string) (*zap.Logger, error) {
	// Criar encoder config com ECS
	encoderConfig := ecszap.NewDefaultEncoderConfig()

	// Criar core do ECS
	core := ecszap.NewCore(
		encoderConfig,
		zapcore.AddSync(zapcore.Lock(zapcore.AddSync(os.Stdout))),
		zap.InfoLevel,
	)

	// Criar logger com campos ECS padrão
	logger := zap.New(core,
		zap.AddCaller(),
		zap.Fields(
			zap.String("service.name", serviceName),
			zap.String("service.environment", environment),
		),
	)

	return logger, nil
}

// LogError registra um erro tanto no logger quanto no APM
func LogError(ctx context.Context, logger *zap.Logger, err error) {
	// Registra no logger
	logger.Error("error occurred", zap.Error(err))

	// Registra no APM
	if e := apm.CaptureError(ctx, err); e != nil {
		e.Send()
	}
}

// LogTrace registra um log de trace no APM e no logger com correlation IDs
func LogTrace(ctx context.Context, logger *zap.Logger, message string, fields ...zap.Field) {
	// Prepare correlation IDs
	correlationFields := []zap.Field{}

	// Get transaction from context
	if tx := apm.TransactionFromContext(ctx); tx != nil {
		traceContext := tx.TraceContext()
		correlationFields = append(correlationFields,
			zap.String("trace.id", traceContext.Trace.String()),
			zap.String("transaction.id", traceContext.Span.String()),
		)

		// Get span from context if exists
		if span := apm.SpanFromContext(ctx); span != nil {
			correlationFields = append(correlationFields,
				zap.String("span.id", span.TraceContext().Span.String()),
			)
		}

		// Set correlation IDs in APM transaction
		tx.Context.SetLabel("trace.id", traceContext.Trace.String())
		tx.Context.SetLabel("transaction.id", traceContext.Span.String())
	}

	// Combine correlation fields with user fields
	allFields := append(correlationFields, fields...)

	// Log with all fields
	logger.Info(message, allFields...)
}
