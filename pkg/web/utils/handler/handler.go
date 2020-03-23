package handler

import (
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func LogError(f func(*fiber.Ctx) error) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {
		err := f(ctx)
		if err != nil {
			logrus.Errorln(err)
		}

	}
}
