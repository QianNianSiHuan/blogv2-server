package enum

type LogLevelType int8

const (
	LofInfoLevel LogLevelType = 1
	LofWarnLevel LogLevelType = 2
	LofErrLevel  LogLevelType = 3
)

func (l LogLevelType) String() string {
	switch l {
	case LofInfoLevel:
		return "info"
	case LofWarnLevel:
		return "warn"
	case LofErrLevel:
		return "err"
	}
	return ""
}
