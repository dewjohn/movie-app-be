package vo

import (
	"movie-app/model"
	"time"
)

type SearchMovieVo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	ReleaseTime string `json:"releaseTime"`
	Score       string `json:"score"`
}

type MovieVo struct {
	ID           uint         `json:"vid"`
	Title        string       `json:"title"`        // 标题
	Cover        string       `json:"cover"`        // 封面
	ReleaseTime  time.Time    `json:"releaseTime"`  // 上映时间
	SheetLength  int          `json:"sheetLength"`  // 片长
	Origin       string       `json:"origin"`       // 地区
	Type         string       `json:"type"`         // 类型
	Director     string       `json:"director"`     // 导演
	Screenwriter string       `json:"screenwriter"` // 编剧
	Actors       string       `json:"actors"`       // 演员
	Language     string       `json:"language"`     // 语言
	Introduction string       `json:"introduction"` // 简介
	Score        float64      `json:"score"`        // 评分
	Resource     []ResourceVo `json:"resource"`     // 视频
	CreatedAt    string       `json:"created_at"`
	UpdatedAt    string       `json:"updated_at"`
}

func ToVideo(movie model.Movie, resource []model.Resource) MovieVo {
	length := len(resource)
	newResource := make([]ResourceVo, length)
	for i := 0; i < length; i++ {
		newResource[i].ID = resource[i].UUID
		newResource[i].Title = resource[i].Title
		newResource[i].Res360 = resource[i].Res360
		newResource[i].Res720 = resource[i].Res720
		newResource[i].Res1080 = resource[i].Res1080
		newResource[i].Original = resource[i].Original
	}
	return MovieVo{
		ID:           movie.ID,
		Title:        movie.Title,
		Cover:        movie.Cover,
		ReleaseTime:  movie.ReleaseTime,
		SheetLength:  movie.SheetLength,
		Origin:       movie.Origin,
		Type:         movie.Type,
		Director:     movie.Director,
		Screenwriter: movie.Screenwriter,
		Actors:       movie.Actors,
		Language:     movie.Language,
		Introduction: movie.Introduction,
		Score:        movie.Score,
		Resource:     newResource,
	}
}

func ToResource(resource []model.Resource) []ResourceVo {
	length := len(resource)
	newResource := make([]ResourceVo, length)
	for i := 0; i < length; i++ {
		newResource[i].ID = resource[i].UUID
		newResource[i].Title = resource[i].Title
		newResource[i].Res360 = resource[i].Res360
		newResource[i].Res720 = resource[i].Res720
		newResource[i].Res1080 = resource[i].Res1080
		newResource[i].Original = resource[i].Original
	}
	return newResource
}
