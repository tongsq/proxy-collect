package servers

import "proxy-collect/service/tunnel"

func StartTunnelServer() {
	tunnel.StartTunnels()
}
