package config

import "testing"

// There's probably a better way to test this, but whatever
func TestRead(t *testing.T) {
	conf := Read("config.yml")
	if len(conf.Sites) != 3 {
		t.Errorf("Expected 3 config items, found %d", len(conf.Sites))
	}

}

func TestMissingFile(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Did not panic on invalid config file")
		}
	}()
	Read("no_such_file.yml")
}
