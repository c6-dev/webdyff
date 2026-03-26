package ps

// Process is a stub interface matching mitchellh/go-ps
type Process interface {
	Pid() int
	PPid() int
	Executable() string
}

func Processes() ([]Process, error) {
	return nil, nil
}

func FindProcess(pid int) (Process, error) {
	return nil, nil
}
