package cable

import "time"

func CertificateWindowSameDay(left, right time.Time, location *time.Location) bool {
	return CertificateWindowDay(left, location) == CertificateWindowDay(right, location)
}
func CertificateWindowExpiry(t time.Time, location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}
	y, m, d := t.In(location).Date()
	return time.Date(y, m, d, 23, 59, 59, 0, location)
}
