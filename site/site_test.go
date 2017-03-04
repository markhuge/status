package site

import "testing"

func TestConstructor(t *testing.T) {
	// test with defaults
	defaults := New()

	if defaults.ExpectedStatusCode != 200 {
		t.Errorf("Defaults: StatusCode expected 200, got %d", defaults.ExpectedStatusCode)
	}

}
