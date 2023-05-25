package rivalitemenum

func GetDropGroup(itemType ItemType) DropGroup { // functask: merge main thì xoá
	switch itemType {
	case ITEnhancer:
		return DGEnhancer
	case ITGold:
		return DGGold
	case ITAvatarFrame, ITAvatar, ITEmoticon, ITNameColor:
		return DGProfile
	case ITGlow, ITVoice, ITDance, ITCurtain, ITVehicle:
		return DGFullEvolve
	case ITFootprint, ITBackBling, ITFlyCraft:
		return DGHalfEvolve
	case ITSkin:
		return DGMinion
	case ITSeasonPoint:
		return DGSPoint
	}
	return DGUnknown
}
