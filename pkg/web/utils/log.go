package utils

import (
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func HeaderFields(ctx *fiber.Ctx) zap.Field {
	return zap.Object("headers", &status{ctx.Fasthttp.Request.Header})
}

type status struct {
	fasthttp.RequestHeader
}

func (s *status) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	s.VisitAll(func(key, value []byte) {
		enc.AddByteString(string(key), value)
	})
	return nil
}
