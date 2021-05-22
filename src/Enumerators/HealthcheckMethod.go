package healthcheck

type Method byte

const (
	Ping   Method = 1
	Http   Method = 2
	Telnet Method = 3
)
