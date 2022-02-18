package model

import (
	"strconv"

	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/config"
)

// redis member point lock key generate
func MakeMemberSwapLockKey(AUID int64) string {
	return config.GetInstance().DBPrefix + "-MEMBER-SWAP-" + strconv.FormatInt(AUID, 10) + "-lock"
}
