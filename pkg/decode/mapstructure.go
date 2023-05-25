package decode

import (
	"github.com/mitchellh/mapstructure"
)

func ToTimeHookFn(m *mapstructure.DecoderConfig) {
	const layout = "2006-01-02T15:04:05"
	m.DecodeHook = mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeHookFunc(layout),
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	)
}
