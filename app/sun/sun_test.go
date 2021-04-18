package sun

import (
	"fmt"
	"testing"
	"time"
)

func TestSunriseSunset(t *testing.T) {
	vSunrise, vSunset := SunriseSunset(0, 0, time.Now())
	fmt.Printf("for 0.0 / 0.0 today: rise %s set %s\n", vSunrise.String(), vSunset.String())
}
