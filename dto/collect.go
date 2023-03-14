package dto

type CollectDto struct {
	Vid uint `json:"vid"`
}

type CollectMovieDto struct {
	Title       string  `json:"title"`
	Cover       string  `json:"cover"`
	ReleaseTime string  `json:"releaseTime"`
	Score       float64 `json:"score"`
}

type CollectResDto struct {
	Vid       uint            `json:"vid"`
	MovieInfo CollectMovieDto `json:"movieInfo"`
}
