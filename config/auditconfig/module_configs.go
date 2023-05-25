package auditconfig

type ModuleConfigs struct {
	BattleEnd struct {
		BattleEndEvent      string `json:"battleEndEvent"`
		AuditCollectionName string `json:"auditCollectionName"`
	} `json:"battleEnd"`

	SpecialEvent struct {
		RewardEvent string `json:"rewardEvent"`
	} `json:"specialEvent"`

	Match struct {
		StartMatch               string `json:"startMatch"`
		StartMatchCollectionName string `json:"startMatchCollectionName"`
	} `json:"match"`

	UpgradeHero struct {
		UpgradeHeroCollectionName string `json:"upgradeHeroCollectionName"`
	} `json:"upgradeHero"`

	TrophyChange struct {
		TrophyAuditCollectionName string `json:"trophyAuditCollectionName"`
	} `json:"trophyChange"`

	IngameShop struct {
		IngameShop               string `json:"ingameShop"`
		IngameShopCollectionName string `json:"ingameShopCollectionName"`
	} `json:"ingameShop"`

	Login struct {
		LoginAuditCollectionName string `json:"loginAuditCollectionName"`
	} `json:"login"`

	Admin struct {
		AdminAuditCollectionName string `json:"adminAuditCollectionName"`
	} `json:"admin"`

	Hero struct {
		HeroEventCollectionName string `json:"heroEventCollectionName"`
	} `json:"hero"`

	CurrencyChange struct {
		CurrencyChangeCollectionName string `json:"currencyChangeCollectionName"`
	} `json:"currencyChange"`

	ThetanBox struct {
		BoxAuditCollectionName string `json:"boxAuditCollectionName"`
	} `json:"thetanBox"`

	BanUnban struct {
		BanUnban               string `json:"banUnban"`
		BanUnbanCollectionName string `json:"banUnbanCollectionName"`
	} `json:"banUnban"`

	Item struct {
		ItemReceived string `json:"itemReceived"`
	}
}
