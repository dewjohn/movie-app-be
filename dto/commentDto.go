package dto

type CommentIdDto struct {
	ID uint
}

type ReplyIdDto struct {
	ID uint
}

type CommentDto struct {
	Content string
	Vid     uint
}

type ReplyDto struct {
	Cid       uint
	Content   string
	ReplyUid  uint
	ReplyName string
}
