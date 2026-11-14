package apperror

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"nodepad-be/internal/common/response"
	"path/filepath"
	"runtime"

	"github.com/zeromicro/go-zero/core/logc"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Error struct {
	Status  int
	Code    string
	Message string
	Source  string
}

func (e *Error) Error() string {
	return e.Message
}

func newError(status int, code string, message string) *Error {
	_, file, line, ok := runtime.Caller(2)

	source := ""
	if ok {
		source = fmt.Sprintf(
			"%s:%d",
			filepath.Base(file),
			line,
		)
	}

	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
		Source:  source,
	}
}

func getTraceInfo(ctx context.Context) (traceID, spanID string) {
	spanContext := oteltrace.SpanContextFromContext(ctx)

	if !spanContext.IsValid() {
		return "", ""
	}

	return spanContext.TraceID().String(),
		spanContext.SpanID().String()
}

func AppErrorHandler() func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		traceID, spanID := getTraceInfo(ctx)

		var appErr *Error

		if errors.As(err, &appErr) {
			logc.Errorf(
				ctx,
				"application error: status=%d code=%s source=%s message=%s",
				appErr.Status,
				appErr.Code,
				appErr.Source,
				appErr.Message,
			)

			return appErr.Status, response.Error{
				Code:    appErr.Code,
				Message: appErr.Message,
				TraceID: traceID,
				SpanID:  spanID,
			}
		}

		logc.Errorf(
			ctx,
			"unhandled error: type=%T error=%v",
			err,
			err,
		)

		internalServerError := InternalServerError()

		return http.StatusInternalServerError, response.Error{
			Code:    internalServerError.(*Error).Code,
			Message: internalServerError.(*Error).Message,
			TraceID: traceID,
			SpanID:  spanID,
		}
	}
}

func BadRequest(message ...string) error {
	msg := http.StatusText(http.StatusBadRequest)
	if len(message) > 0 {
		msg = message[0]
	}

	return newError(http.StatusBadRequest, "BAD_REQUEST", msg)
}

func Unauthorized(message ...string) error {
	msg := http.StatusText(http.StatusUnauthorized)
	if len(message) > 0 {
		msg = message[0]
	}
	return newError(http.StatusUnauthorized, "UNAUTHORIZED", msg)
}

func Conflict(message ...string) error {
	msg := http.StatusText(http.StatusConflict)
	if len(message) > 0 {
		msg = message[0]
	}
	return newError(http.StatusConflict, "CONFLICT", msg)
}

func NotFound(message ...string) error {
	msg := http.StatusText(http.StatusNotFound)
	if len(message) > 0 {
		msg = message[0]
	}
	return newError(http.StatusNotFound, "NOT_FOUND", msg)
}

func InternalServerError(message ...string) error {
	msg := http.StatusText(http.StatusInternalServerError)
	if len(message) > 0 {
		msg = message[0]
	}

	return newError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", msg)
}

func NotFoundMethod(message ...string) error {
	msg := http.StatusText(http.StatusMethodNotAllowed)
	if len(message) > 0 {
		msg = message[0]
	}

	return newError(http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", msg)
}

func NotFoundRoute(message ...string) error {
	msg := http.StatusText(http.StatusNotFound)
	if len(message) > 0 {
		msg = message[0]
	}

	return newError(http.StatusNotFound, "ROUTE_NOT_FOUND", msg)
}
