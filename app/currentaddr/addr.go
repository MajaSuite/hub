package currentaddr

import (
	"time"
	"net"
	"net/http"
	"io/ioutil"

	"gohub/app"
)

var (
	Type string
	Addr net.IP
	Update time.Time
)

func convertIP(ip net.IP) (string, string) {
	if Addr.To4() != nil {
		return "ipv4", Addr.String()
	}

	return "ipv6", Addr.String()
}

// return:
// type ip lastUpdate error
func GetInternetAddress() (string, string, time.Time, error) {
	if Addr != nil {
		t, a := convertIP(Addr)
		return t, a, Update, nil
	}

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", "", time.Now(), app.ErrNotConnected
	}
	defer resp.Body.Close()

	buf, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", "", time.Now(), app.ErrNotConnected
	}

	Addr = net.ParseIP(string(buf))
	t, a := convertIP(Addr)
	Update = time.Now()

	return t, a, Update, nil
}
