package rcfg

import "testing"

func TestGetLatest(t *testing.T) {
	cfg, err := GetLatest(Staging, "test_launcher")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(cfg))
}

func TestGetConfByUser(t *testing.T) {
	cfg, err := GetByUser(Staging, "rivals.lobby", UserContext{
		UserID:     "6199e2a6fe775c92cca39560",
		Attributes: nil,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(cfg))
}
