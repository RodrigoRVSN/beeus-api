package context

type Context interface {
	GetParam(key string) string
	GetMiddlewareParam(key string) interface{}
	SetMiddlewareParam(key string, value string) interface{}
	SendJson(status int, json interface{}) error
	ParseBody(interface{}) error
}
