package context

type NoticeInfo struct {
	Id       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Desc     string `json:"desc,omitempty"`
	Url      string `json:"url,omitempty"`
	CreateDt int64  `json:"create_dt,omitempty"`
}
