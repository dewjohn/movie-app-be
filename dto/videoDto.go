package dto

type VideoDto struct {
	Uid   uint   `json:"uid"`
	Title string `gorm:"type:varchar(50);not null;index"`
	Cover string `gorm:"size:255;not null"`
	//Videos       string    `gorm:"size:255;"`    // 先用字符串视频链接，下一版本引入本地视频
	ReleaseTime  string `json:"releaseTime"`  // 上映时间
	SheetLength  int    `json:"sheetLength"`  // 片长
	Origin       string `json:"origin"`       // 地区
	Type         string `json:"type"`         // 类型
	Director     string `json:"director"`     // 导演
	Screenwriter string `json:"screenwriter"` // 编剧
	Actors       string `json:"actors"`       // 演员
	Language     string `json:"language"`     // 语言
	Introduction string `json:"introduction"` // 简介
	//Score        float64   `json:"score"`        // 评分
}
