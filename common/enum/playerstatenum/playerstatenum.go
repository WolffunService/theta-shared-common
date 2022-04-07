package playerstatenum

import "fmt"

type StatName string

const (
	Test                  StatName = "test_%d"
	PlayerBattle          StatName = "player_battle"
	PlayerWinBattle       StatName = "player_win_battle"
	HeroBattle            StatName = "hero_battle_%d"  //heroId
	SkillBattle           StatName = "skill_battle_%d" //skillId
	ModeBattle            StatName = "mode_battle_%d"  //modeId
	RarityGTHCBattle      StatName = "ra_gthc_b_%d"    //rarityid
	RarityGTHCEarn        StatName = "ra_gthc_e_%d"    //rarityid
	FirstOpenDate         StatName = "first_open_dt"
	ConnectWalletDate     StatName = "conn_wallet_dt"
	AccountCreateDate     StatName = "acc_create_dt"
	IAPTotalValue         StatName = "iap_tt"
	RentalTotalValue      StatName = "ren_tt"
	BuyHeroTotalValue     StatName = "buy_hero_tt"
	BuyBoxTotalValue      StatName = "buy_box_tt"
	SellTotalValue        StatName = "sell_tt"
	NFTHeroCount          StatName = "nft_hero_tt"
	NFTHeroSellingCount   StatName = "selling_tt"
	NFTHeroRentedOutCount StatName = "rented_out_tt"
	NFTHeroForRentCount   StatName = "for_rent_tt"
	MKPVisit              StatName = "mkp_visit"
	CreatorViewPoint      StatName = "creator_view"
)

func (s StatName) GetWithCode(code int) string {
	return fmt.Sprintf(string(s), code)
}

func (s StatName) String() string {
	return fmt.Sprintf(string(s))
}
