package log

type Logger interface {
	Log(msg string, args ...interface{})
}

const (
	FatalLvl logLevel = iota
	ErrLvl
	WarnLvl
	InfoLvl
	TraceLvl
)

type logLevel int
type MultiLevelLogger interface {
	SetLevel(logLevel)

	Fatal(msg string, args ...interface{})
	Err(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Trace(msg string, args ...interface{})
}
