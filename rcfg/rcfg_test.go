package rcfg

import "testing"

func TestGetLatest(t *testing.T) {
	cfg, err := GetLatest(Production, "tr_country_whitelist")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(cfg))
}
