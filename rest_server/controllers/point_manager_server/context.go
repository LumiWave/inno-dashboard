package point_manager_server

type Common struct {
	Return  int64  `json:"return"`
	Message string `json:"message"`
}

type Point struct {
	PointID          int64  `json:"point_id"`
	Quantity         int64  `json:"quantity"`
	TodayQuantity    int64  `json:"today_quantity"`
	ResetDate        string `json:"reset_date"`
	PreviousQuantity int64  `json:"-"`
}

type MePointValue struct {
	MyUUID     string   `json:"-"`
	DatabaseID int64    `json:"database_id"`
	MUID       int64    `json:"mu_id"`
	Points     []*Point `json:"points"`
}

type MePointInfo struct {
	Common
	MePointValue `json:"value"`
}
