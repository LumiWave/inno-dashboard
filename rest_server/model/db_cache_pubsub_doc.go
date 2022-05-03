package model

var (
	gIsMaintenance            = false
	gIsExternalTransferEnable = true
	gIsSwapToCoinEnable       = true
	gIsSwapToPointEnable      = false
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

func SetSwapEnable(isToCoinEnable, isToPointEnable bool) {
	gIsSwapToCoinEnable = isToCoinEnable
	gIsSwapToPointEnable = isToPointEnable
}

func GetSwapEnable() (bool, bool) {
	return gIsSwapToCoinEnable, gIsSwapToPointEnable
}

func SetPointUpdateEnable(isEnable bool) {
	gIsPointUpdateEnable = isEnable
}

func GetPointUpdateEnable() bool {
	return gIsPointUpdateEnable
}
