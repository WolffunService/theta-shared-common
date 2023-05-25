package rivalitemenum

type ExclusiveFeature string

const (
	ETNormalOffer ExclusiveFeature = "NormalOffer"
	ETOnboard     ExclusiveFeature = "Onboard"
	ETRankReward  ExclusiveFeature = "RankReward"
	ETNewArrival  ExclusiveFeature = "NewArrival"
	ETSeasonOffer ExclusiveFeature = "SeasonOffer"
	ETSeasonPass  ExclusiveFeature = "SeasonPass"
	ETEventOffer  ExclusiveFeature = "EventOffer"
	ETLobbyReward ExclusiveFeature = "LobbyReward"
)
