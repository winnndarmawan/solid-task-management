package middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// RequestID sets or generates a unique request ID
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get("X-Request-ID")
		if id == "" {
			id = uuid.New().String()
		}
		c.Response().Header().Set("X-Request-ID", id)
		c.Set("request_id", id)
		return next(c)
	}
}

// ZapLogger is a custom structured logging middleware
func ZapLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			req := c.Request()
			res := c.Response()

			logger.Info("request",
				zap.String("request_id", c.Get("request_id").(string)),
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", res.Status),
				zap.Duration("latency", stop.Sub(start)),
				zap.String("ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
			)

			return err
		}
	}
}

// RecoverWithZap handles panics and logs them with stack trace
func RecoverWithZap(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					logger.Error("panic recovered",
						zap.Any("error", r),
						zap.String("request_id", c.Get("request_id").(string)),
						zap.Stack("stacktrace"),
					)
					c.Error(echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error"))
				}
			}()
			return next(c)
		}
	}
}
