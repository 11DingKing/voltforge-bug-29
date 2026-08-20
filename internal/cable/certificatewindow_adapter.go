package cable

import "time"

func CertificateWindowSameDay(left, right time.Time, location *time.Location) bool {
	if location == nil {
		location = time.UTC
	}
	leftDay := CertificateWindowDay(left, location)
	rightDay := CertificateWindowDay(right, location)
	return leftDay == rightDay
}
func CertificateWindowExpiry(t time.Time, location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}
	y, m, d := t.In(location).Date()
	return time.Date(y, m, d, 23, 59, 59, 0, location)
}
