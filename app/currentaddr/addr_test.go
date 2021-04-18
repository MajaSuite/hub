package currentaddr

import (
	"fmt"
	"testing"
	"time"
)

func TestGetInternetAddress(t *testing.T) {
	s, a, u, err := GetInternetAddress()
	if err != nil {
		t.Errorf("err %s", a)
	}

	fmt.Printf("FIRST TEST %s://%s updated at %s\n", s, a, u.String())

	time.Sleep(5 * time.Second)

	s, a, u, err = GetInternetAddress()
	if err != nil {
		t.Errorf("err %s", a)
	}

	fmt.Printf("SECOND TEST %s://%s updated at %s\n", s, a, u.String())
}
