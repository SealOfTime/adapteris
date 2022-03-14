package app

import "fmt"

func (a *App) initLogger() {
	a.Log = &fmtLogger{}
}

type fmtLogger struct{}

func (l *fmtLogger) Log(msg string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s\n", msg), args...)
}
