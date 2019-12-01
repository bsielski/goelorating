package goelorating

import (
	"testing"
)

func TestDefaultKFactor(t *testing.T) {
	subject := DefaultKFactor
	if subject <= 0  {
		t.Errorf("Subject == %d, should be > 0", subject)
	}
}
func TestDefaultDeviation(t *testing.T) {
	subject := DefaultDeviation
	if subject <= 0  {
		t.Errorf("Subject == %d, should be > 0", subject)
	}
}
func TestDefaultRating(t *testing.T) {
	subject := DefaultRating
	if subject <= 0  {
		t.Errorf("Subject == %d, should be > 0", subject)
	}
}
