package dto

import "github.com/google/uuid"

type MovieDto struct {
	Title        string  `json:"title"`        // 标题
	Cover        string  `json:"cover"`        // 封面
	ReleaseTime  string  `json:"releaseTime"`  // 上映时间
	SheetLength  int     `json:"sheetLength"`  // 片长
	Origin       string  `json:"origin"`       // 地区
	Type         string  `json:"type"`         // 类型
	Director     string  `json:"director"`     // 导演
	Screenwriter string  `json:"screenwriter"` // 编剧
	Actors       string  `json:"actors"`       // 演员
	Language     string  `json:"language"`     // 语言
	Introduction string  `json:"introduction"` // 简介
	Score        float64 `json:"score"`        // 评分
}

type MovieToAdminDto struct {
	MovieDto
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ResDto struct {
	Res360   string
	Res480   string
	Res720   string
	Res1080  string
	Original string
}

type ModifyMovieDto struct {
	Title        string  `json:"title"`        // 标题
	Cover        string  `json:"cover"`        // 封面
	ReleaseTime  string  `json:"releaseTime"`  // 上映时间
	SheetLength  int     `json:"sheetLength"`  // 片长
	Origin       string  `json:"origin"`       // 地区
	Type         string  `json:"type"`         // 类型
	Director     string  `json:"director"`     // 导演
	Screenwriter string  `json:"screenwriter"` // 编剧
	Actors       string  `json:"actors"`       // 演员
	Language     string  `json:"language"`     // 语言
	Introduction string  `json:"introduction"` // 简介
	Score        float64 `json:"score"`        // 评分
}

type VideoIdDto struct {
	Id uint
}

type UUID struct {
	UUID uuid.UUID
}

type GetMovieListDto struct {
	Page     int
	PageSize int
}

type FilterMovieDto struct {
	Page     int
	PageSize int
	Column   string
	Value    string
}

type ScoreDto struct {
	Vid   uint    `json:"vid"`
	Grade float64 `json:"grade"`
}

type ResourceToAdminDto struct {
	UUID     uuid.UUID `json:"uuid"`
	Original string    `json:"original"`
	Res360   string    `json:"res360"`
	Res480   string    `json:"res480"`
	Res720   string    `json:"res720"`
	Res1080  string    `json:"res1080"`
	Title    string    `json:"title"`
}
type UploadVideoByUrlDto struct {
	Vid uint   `json:"vid"`
	Url string `json:"url"`
}
