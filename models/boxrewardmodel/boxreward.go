package boxrewardmodel

import "github.com/WolffunService/theta-shared-common/enums/boxrewardenum"

type BoxRewardProbability struct {
	Name        string                      `json:"name"`
	Type        boxrewardenum.BoxRewardType `json:"type"`
	Rate        float64                     `json:"rate"`
	DisplayRate float64                     `json:"displayRate,omitempty"`
}
