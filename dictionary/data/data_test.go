package data

import "testing"

func TestShouldCreateEmptyData(t *testing.T) {
	data := Languages{}
	if data.Size() != 0 {
		t.Errorf("The data should start empty")
	}
	if data.Contains("java") {
		t.Errorf("The data should be empty")
	}

	if data.Size() != 0 {
		t.Errorf("The Should have size 0")
	}
}
