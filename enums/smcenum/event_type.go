package smcenum

type EventName string

const (
	ET_CosmeticMint  EventName = "CosmeticMinter_mint"
	ET_SellingNFT    EventName = "MarketplaceV3_selling"
	ET_LendingHero   EventName = "HeroRentalV2_renting"
	ET_HeroMint      EventName = "MinterFactoryV2_mint_adminMintTo"
	ET_HeroMultiMint EventName = "MinterFactoryV2_multipleMintTo_adminMultipleMintTo"
	ET_HeroUnlock    EventName = "HeroRentalV2_unlockHero"
	ET_ClaimToken    EventName = "ClaimToken_claimToken"
	ET_Fusion        EventName = "Fusion_fuse"
	ET_Staking       EventName = "Staking_staked"
	ET_BuyBox        EventName = "ThetanBoxHubV2_buyBox(Free)WithSignature"
)
