package monitor

import "log"

type Monitor struct {
	*log.Logger
}

// Write implements the io.Writer interface.
// make monitor non-intrusive, always return error = nil
func (m *Monitor) Write(p []byte) (int, error) {
	err := m.Output(2, string(p))
	if err != nil {
		log.Println(err) // use the log package’s default Logger
	}
	return len(p), nil
}
