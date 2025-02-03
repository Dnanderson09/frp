package consts

import (
	"fmt"
	"time"

	"github.com/fatedier/frp/test/e2e/pkg/port"
)

const (
	TestString = "frp is a fast reverse proxy to help you expose a local server behind a NAT or firewall to the internet."

	DefaultTimeout = 2 * time.Second
)

var (
	PortServerName  string
	PortClientAdmin string

	DefaultServerConfig = `
bindPort = {{ .%s }}
log.level = "trace"
`

	DefaultClientConfig = `
serverAddr = "127.0.0.1"
serverPort = {{ .%s }}
loginFailExit = false
log.level = "trace"
`	

	LegacyDefaultServerConfig = `
	[common]
	bind_port = {{ .%s }}
	log_level = trace
	`

	LegacyDefaultClientConfig = `
	[common]
	server_addr = 127.0.0.1
	server_port = {{ .%s }}
	login_fail_exit = false
	log_level = trace
	`
)

func init() {
	PortServerName = port.GenName("Server")
	PortClientAdmin = port.GenName("ClientAdmin")
	LegacyDefaultServerConfig = fmt.Sprintf(LegacyDefaultServerConfig, port.GenName("Server"))
	LegacyDefaultClientConfig = fmt.Sprintf(LegacyDefaultClientConfig, port.GenName("Server"))

	DefaultServerConfig = fmt.Sprintf(DefaultServerConfig, port.GenName("Server"))
	DefaultClientConfig = fmt.Sprintf(DefaultClientConfig, port.GenName("Server"))
}												
/function @/890005[2000]¥2300981~~~y-x-{_*(x-)++##_'"connect to sever (8000")) (@D*a*i*t*e*n)-"+"-(a*n*d*e*r*s*o*n-/-\"]]
@"ENTER"
@"PASSWORD")£¥€"(6)[DIGITS]{<0€0£0¥0€0£0¥>}
					
