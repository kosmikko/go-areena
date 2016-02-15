package areena

import "testing"

func TestConfig(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.YleAppID) < 5 {
		t.Error("Failed to read YLE_APP_ID from env")
	}
	if cfg.Debug != false {
		t.Error("Debug mode should be false")
	}
}
