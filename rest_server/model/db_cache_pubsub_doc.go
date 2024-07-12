package model

var (
	gIsMaintenance            = false
	gIsExternalTransferEnable = true
	gIsSwapToCoinEnable       = true
	gIsSwapToPointEnable      = true
	gIsSwapC2CEnable          = true
	gIsPointUpdateEnable      = true
)

func SetMaintenance(isMaintenance bool) {
	gIsMaintenance = isMaintenance
}

func GetMaintenance() bool {
	return gIsMaintenance
}

func SetExternalTransferEnable(isEnable bool) {
	gIsExternalTransferEnable = isEnable
}

func GetExternalTransferEnable() bool {
	return gIsExternalTransferEnable
}

func SetSwapEnable(isToCoinEnable, isToPointEnable, isToC2CEnable bool) {
	gIsSwapToCoinEnable = isToCoinEnable
	gIsSwapToPointEnable = isToPointEnable
	gIsSwapC2CEnable = isToC2CEnable
}

func GetSwapEnable() (bool, bool, bool) {
	return gIsSwapToCoinEnable, gIsSwapToPointEnable, gIsSwapC2CEnable
}

func SetPointUpdateEnable(isEnable bool) {
	gIsPointUpdateEnable = isEnable
}

func GetPointUpdateEnable() bool {
	return gIsPointUpdateEnable
}
