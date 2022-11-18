package demo

import "testing"

func TestDemo(t *testing.T) {
	expected := "This a demo module!"

	if result := demo(); result != expected {
		t.Errorf("Demo() = %q, not equal expected value: %q\n", result, expected)
	}
}
