package common

import "time"

type UserRole int

const (
	DAY  = time.Hour * 24
	WEEK = time.Hour * 24 * 7

	NONE     UserRole = 0 // Role Default
	ADMIN    UserRole = 1 // Role Admin
	ROOT     UserRole = 2 // Role Highest
	SYSADMIN UserRole = 3 // Role System Config
	EDITOR   UserRole = 4 // Role Editor
)
