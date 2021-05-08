package localaddr

import (
	"errors"
	"net"
	"runtime"
	"strings"
)

var ErrCantFindInterfaces = errors.New("Cant find local interfaces")
var ErrCantFindIPv4 = errors.New("Cant find local ipv4")

func Local4Address() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, ErrCantFindInterfaces
	}

	for _, interf := range ifaces {
		if "darwin" == runtime.GOOS && !strings.HasPrefix(interf.Name, "en") {
			continue
		}

		if "linux" == runtime.GOOS && !strings.HasPrefix(interf.Name, "e") {
			continue
		}

		addrs, _ := interf.Addrs()
		for _, addr := range addrs {
			ip, _, _ := net.ParseCIDR(addr.String())
			if ip.To4() != nil {
				return ip, nil
			}
		}
	}

	return nil, ErrCantFindIPv4
}
