package point_manager_server

var gPointManagerServerInfo *PointManagerServerInfo

type HostInfo struct {
	HostUri string
	Ver     string // /v1
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
