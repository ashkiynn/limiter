package limiter

type InterfaceLimiter interface {
	Check(id int32) bool
	CheckNum(id int32, num uint32) bool
}
