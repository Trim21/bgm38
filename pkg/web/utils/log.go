package utils

import (
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func HeaderFields(ctx *fiber.Ctx) zap.Field {
	return zap.Object("headers", &status{h: &ctx.Fasthttp.Request.Header})
}

type status struct {
	h *fasthttp.RequestHeader
}

func (s *status) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	s.h.VisitAll(func(key, value []byte) {
		enc.AddByteString(string(key), value)
	})
	return nil
}
