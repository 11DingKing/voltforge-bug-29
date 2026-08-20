package cable

import (
	"testing"
	"time"
)

func TestVoltForge29(t *testing.T) {
	loc := time.FixedZone("CST+8", 8*60*60)
	local := time.Date(2026, 8, 20, 0, 5, 0, 0, loc)
	utc := local.UTC()
	if CertificateWindowDay(utc, loc) != "2026-08-20" {
		t.Fatalf("wrong business day: %s", CertificateWindowDay(utc, loc))
	}
}
