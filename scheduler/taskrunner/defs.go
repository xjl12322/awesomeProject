package taskrunner
const (
	READY_TO_DISPATCH = "d"

	READY_TO_EXECUTE = "a"

	CLOSE = "c"
)
type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error