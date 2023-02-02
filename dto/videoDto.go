package dto

type VideoDto struct {
	Title        string `json:"title"`        // 标题
	Cover        string `json:"cover"`        // 封面
	ReleaseTime  string `json:"releaseTime"`  // 上映时间
	SheetLength  int    `json:"sheetLength"`  // 片长
	Origin       string `json:"origin"`       // 地区
	Type         string `json:"type"`         // 类型
	Director     string `json:"director"`     // 导演
	Screenwriter string `json:"screenwriter"` // 编剧
	Actors       string `json:"actors"`       // 演员
	Language     string `json:"language"`     // 语言
	Introduction string `json:"introduction"` // 简介
}

type ResDto struct {
	Res360   string
	Res480   string
	Res720   string
	Res1080  string
	Original string
}

type ModifyVideoDto struct {
	Vid          int    `json:"vid"`          // 视频ID
	Title        string `json:"title"`        // 标题
	Cover        string `json:"cover"`        // 封面
	ReleaseTime  string `json:"releaseTime"`  // 上映时间
	SheetLength  int    `json:"sheetLength"`  // 片长
	Origin       string `json:"origin"`       // 地区
	Type         string `json:"type"`         // 类型
	Director     string `json:"director"`     // 导演
	Screenwriter string `json:"screenwriter"` // 编剧
	Actors       string `json:"actors"`       // 演员
	Language     string `json:"language"`     // 语言
	Introduction string `json:"introduction"` // 简介
}
