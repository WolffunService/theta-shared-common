package authenum

type LoginType int

const (
	LT_UNKNOWN         LoginType = 0
	LT_FIRST_GET_TOKEN LoginType = 1
	LT_TOKEN           LoginType = 2
	LT_EMAIL           LoginType = 3
	LT_AS_GUEST        LoginType = 4
	LT_GUEST_ACCOUNT   LoginType = 5
)

type LoginPlatform int

const (
	LP_UNKNOWN LoginPlatform = 0
	LP_ANDROID LoginPlatform = 1
	LP_IOS     LoginPlatform = 2
	LP_PC      LoginPlatform = 3
)
