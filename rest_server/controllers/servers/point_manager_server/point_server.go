package point_manager_server

var gPointManagerServerInfo *PointManagerServerInfo

type HostInfo struct {
	IntHostUri string
	ExtHostUri string
	IntVer     string // m1.0
	ExtVer     string // v1.0
}

type AuthInfo struct {
	ApiKey string
}

type PointManagerServerInfo struct {
	HostInfo

	AuthInfo
}

func GetInstance() *PointManagerServerInfo {
	return gPointManagerServerInfo
}

func NewPointManagerServerInfo(apiKey string, hostInfo HostInfo) *PointManagerServerInfo {
	if gPointManagerServerInfo == nil {
		gPointManagerServerInfo = &PointManagerServerInfo{
			HostInfo: hostInfo,
			AuthInfo: AuthInfo{
				ApiKey: apiKey,
			},
		}
	}

	return gPointManagerServerInfo
}

func (o *PointManagerServerInfo) SetApiKey(key string) {
	gPointManagerServerInfo.ApiKey = key
}
