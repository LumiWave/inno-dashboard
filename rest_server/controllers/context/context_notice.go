package context

// /////// Notice Info
type NoticeInfo struct {
	Id       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Desc     string `json:"desc,omitempty"`
	Url      string `json:"url,omitempty"`
	CreateDt int64  `json:"create_dt,omitempty"`
}

////////////////////////////////////////

type News struct {
	NewsID           int64  `json:"news_id"`
	Title            string `json:"title"`
	StartSDT         string `json:"start_sdt"`
	EndSDT           string `json:"end_sdt"`
	BannerURL        string `json:"banner_url"`
	NewsURL          string `json:"news_url"`
	IsAlwaysVisibled bool   `json:"is_always_visibled"`
}

type ResNewsList struct {
	PageInfo
	TotalCount int64   `json:"total_count"`
	List       []*News `json:"list"`
}
