package context

import (
	"github.com/gofiber/fiber/v2"
)

type FiberContext struct {
	Ctx *fiber.Ctx
}

func (fc *FiberContext) GetParam(key string) string {
	return fc.Ctx.Params(key)
}

func (fc *FiberContext) GetMiddlewareParam(key string) interface{} {
	return fc.Ctx.Locals(key)
}

func (fc *FiberContext) SetMiddlewareParam(key string, value string) interface{} {
	return fc.Ctx.Locals(key, value)
}

func (fc *FiberContext) SendJson(status int, json interface{}) error {
	return fc.Ctx.Status(status).JSON(json)
}

func (fc *FiberContext) ParseBody(payload interface{}) error {
	return fc.Ctx.BodyParser(payload)
}

func AdaptHandler(handler func(Context) error) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fc := NewFiberContext(ctx)
		return handler(fc)
	}
}

func NewFiberContext(ctx *fiber.Ctx) Context {
	return &FiberContext{Ctx: ctx}
}
