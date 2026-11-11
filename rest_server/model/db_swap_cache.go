package model

import (
	"strconv"

	"github.com/LumiWave/inno-dashboard/rest_server/config"
)

// redis member point lock key generate
func MakeMemberSwapLockKey(AUID int64) string {
	return config.GetInstance().DBPrefix + "-MEMBER-SWAP-" + strconv.FormatInt(AUID, 10) + "-lock"
}

// 같은 계정의 지갑 등록이 동시에 처리되지 않게 잡는 락 키
func MakeWalletRegistLockKey(AUID int64) string {
	return config.GetInstance().DBPrefix + "-WALLET-REGIST-" + strconv.FormatInt(AUID, 10) + "-lock"
}
