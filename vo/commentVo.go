package vo

import "time"

type CommentVo struct {
	ID        uint      `json:"cid"`     //评论ID
	Content   string    `json:"content"` //内容
	Uid       uint      `json:"uid"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Reply     []ReplyVo `json:"reply"`
	CreatedAt time.Time `json:"created_at"`
}

type ReplyVo struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Uid       uint      `json:"uid"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentListVo struct {
	ID         uint      `json:"id"`
	Content    string    `json:"content"`
	Uid        uint      `json:"uid"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	CreatedAt  time.Time `json:"created_at"`
	ReplyCount int       `json:"reply_count"`
}
