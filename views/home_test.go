package views

import (
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	comp, err := componentToString(Home())
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(comp, "Events Example App") {
		t.Errorf("Expected Events Example App', got '%s'", comp)
	}
}
