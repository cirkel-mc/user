package constant

type UserStatus int

const (
	NotYetVerified UserStatus = iota
	Active
	InputIdentity
	InputSkill
	Inactive
	Banned
	Deleted
)

func (us UserStatus) Int() int {
	return int(us)
}

func (us UserStatus) String() string {
	s := []string{
		"NOT_YET_VERIFIED",
		"ACTIVE",
		"INPUT_IDENTITY",
		"INPUT_SKILL",
		"INACTIVE",
		"BANNED",
		"DELETED",
	}

	if len(s) < us.Int() {
		return ""
	}

	return s[us]
}
