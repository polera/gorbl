package gorbl

import (
	"net"
	"testing"
)

func TestReverseIP(t *testing.T) {
	t.Parallel()
	ip := net.IP{192, 168, 1, 1}

	r := Reverse(ip)

	if r != "1.1.168.192" {
		t.Errorf("Expected ip to equal 1.1.168.192, actual ", r)
	}

}

func TestLookupParams(t *testing.T) {
	t.Parallel()

	res := Lookup("b.barracudacentral.org", "smtp.gmail.com")

	if res.List != "b.barracudacentral.org" {
		t.Errorf("Expected b.barracudacentral.org, actual ", res.List)
	}

	if res.Host != "smtp.gmail.com" {
		t.Errorf("Expected smtp.gmail.com, actual ", res.Host)
	}
}
