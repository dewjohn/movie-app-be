package dto

type CommentIdDto struct {
	ID uint
}

type CommentDto struct {
	Vid      uint   `json:"vid"`
	Content  string `json:"content"`
	Uid      uint   `json:"uid"`
	ParentID uint   `json:"parentId"`
}
