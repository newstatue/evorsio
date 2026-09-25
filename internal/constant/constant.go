package constant

const BuildDir = "bin"
const AppName = "evorsio"
const LogDir = "logs"

type Component string

const (
	ComponentApp   Component = "app"
	ComponentFS    Component = "fs"
	ComponentWails Component = "wails"
)

type SubComponent string

const (
	SubComponentServer SubComponent = "server"
	SubComponentMount  SubComponent = "mount"
)

type LogArg string

const (
	LogArgError        LogArg = "error"
	LogArgComponent    LogArg = "component"
	LogArgSubComponent LogArg = "subcomponent"
)

type ErrMsg string
