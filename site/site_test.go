package site

import "testing"

func TestConstructor(t *testing.T) {
	// test with defaults
	url := "http://foo.bar"
	defaults := New(url)

	if defaults.ExpectedStatusCode != 200 {
		t.Errorf("Defaults: StatusCode expected 200, got %d", defaults.ExpectedStatusCode)
	}

	if defaults.URL != url {
		t.Errorf("Defaults: URL expected %s, got %s", url, defaults.URL)
	}

}
