package vo

import "github.com/google/uuid"

type ResourceVo struct {
	UUID uuid.UUID `json:"uuid"`

	Vid uint `json:"vid"`
	//分P使用的标题
	Title string `json:"title"`
	//不同分辨率
	Res360  string `json:"res360"`
	Res480  string `json:"res480"`
	Res720  string `json:"res720"`
	Res1080 string `json:"res1080"`
	//原始分辨率，适用于早期版本未指定分辨率的视频
	//或者不进行转码处理的情况
	Original string `json:"original"`
}
