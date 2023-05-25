package leaderboardmodel

type (
	RequestGetLBFullData struct {
		UserID      string `json:"userID" form:"userID"`
		LBKey       string `json:"LBKey" form:"LBKey"`
		ExposeIsBot bool   `json:"exposeIsBot" form:"exposeIsBot"`
		PageSize    int32  `json:"pageSize" form:"pageSize"`
		PageNumber  int32  `json:"pageNumber" form:"pageNumber"`
	}
	PlayerInfo struct {
		UserId      string `json:"userId"`
		Username    string `json:"username"`
		Trophy      int    `json:"trophy"`
		TrophyLevel int    `json:"trophyLevel"`
		AvatarId    int    `json:"avatarId"`
		GuildName   string `json:"guildName"`
		FrameId     int    `json:"frameId"`
		Country     string `json:"country"`
	}

	LeaderboardDataModel struct {
		ProfileInfo PlayerInfo `json:"profileInfo"`
		Score       float64    `json:"score"`
		Rank        int32      `json:"rank"`
	}

	LeaderboardRequest struct {
		LeaderboardType string `form:"lbType"`
		PageNumber      int32  `form:"pageNumber"`
		PageSize        int32  `form:"pageSize"`
		Country         string `form:"country"`
		GuildId         string `form:"guildId"`
	}
)
