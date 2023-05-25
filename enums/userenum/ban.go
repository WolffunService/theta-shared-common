package userenum

type BanAuditType int

const (
	BAT_BAN              BanAuditType = 1
	BAT_BAN_FIND_MATCH   BanAuditType = 2
	BAT_UNBAN            BanAuditType = 11
	BAT_UNBAN_FIND_MATCH BanAuditType = 12
)
