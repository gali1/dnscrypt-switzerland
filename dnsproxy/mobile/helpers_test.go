package mobile

import "testing"

func TestParseDNSStamp(t *testing.T) {
	// AdGuard DNS (DNSCrypt)
	stampStr := "sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20"

	stamp, err := ParseDNSStamp(stampStr)
	if err != nil {
		t.Fatalf("cannot fail: %s", err)
	}

	if stamp.ServerAddr != "176.103.130.130:5443" ||
		stamp.ProviderName != "2.dnscrypt.default.ns1.adguard.com" ||
		stamp.Proto != 0x01 {
		t.Fatalf("wrong stamp data: %v", stamp)
	}
}

func TestTestUpstream(t *testing.T) {
	const timeout = 500 // 500 ms

	err := TestUpstream("123.12.32.1:1493", "", timeout)
	if err == nil {
		t.Fatalf("cannot be successful")
	}

	err = TestUpstream("8.8.8.8:53", "", timeout*10)
	if err != nil {
		t.Fatalf("cannot fail: %s", err)
	}

	// Test for DoT with 2 bootstraps. Only one is valid
	err = TestUpstream("tls://dns.adguard.com", "1.2.3.4\n8.8.8.8", timeout*10)
	if err != nil {
		t.Fatalf("cannot fail: %s", err)
	}
}
