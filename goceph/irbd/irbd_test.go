package irbd_test

import (
	"CephMonitorAPI/goceph/irbd"
	"testing"
)

func TestRandom(t * testing.T) {
	randomInt := irbd.Random()
	t.Error(randomInt)
}