package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-app/common"
	"movie-app/dto"
	"movie-app/model"
	"movie-app/response"
	"movie-app/vo"
	"net/http"
)

func CommentService(comment dto.CommentDto, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	if !IsMovieExit(DB, comment.Vid) {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = "视频不存在"
		return res
	}
	DB.Create(&model.Comment{Vid: comment.Vid, Content: comment.Content, Uid: uid.(uint)})
	return res
}

func GetCommentService(page int, pageSize int, vid int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()

	var count int64
	var comments []vo.CommentVo
	SqlComment := "select comments.id, comments.created_at, content, uid, users.name, users.avatar, reply_count " +
		"from comments,users where comments.deleted_at is null and comments.uid = users.id and vid = ?"
	SqlReplay := "select content, users.name, reply_uid, reply_name from replies, users " +
		"where replies.deleted_at is null and replies.uid = users.id and cid = ? limit 2"

	if !IsMovieExit(DB, uint(vid)) {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = "视频不存在"
		return res
	}
	DB.Model(&model.Comment{}).Where("vid = ?", vid).Count(&count)
	Pagination := DB.Limit(pageSize).Offset((page - 1) * pageSize)
	Pagination.Raw(SqlComment, vid).Scan(&comments)
	for i := 0; i < len(comments); i++ {
		DB.Raw(SqlReplay, comments[i].ID).Scan(&comments[i].Reply)
	}
	res.Data = gin.H{"count": count, "comments": comments}
	return res
}

func ReplyService(reply dto.ReplyDto, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       200,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()

	if !IsCommentExist(DB, reply.Cid) {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = "评论不存在或已被删除"
		return res
	}
	newReply := model.Reply{
		Cid:      reply.Cid,
		Content:  reply.Content,
		Uid:      uid.(uint),
		ReplyUid: reply.ReplyUid,
	}
	DB.Create(&newReply)
	DB.Model(&model.Comment{}).Where("id = ?", reply.Cid).UpdateColumn("reply_count", gorm.Expr("reply_count + ?", 1))
	return res
}

func GetReplyDetailsV2Service(cid int, page int, pageSize int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}

	var replies []vo.ReplyVo
	sql := "select replies.id,replies.created_at,content,uid,users.name,users.avatar,reply_uid,reply_name " +
		"from replies,users where replies.deleted_at is null and replies.uid = users.id and cid = ?"
	DB := common.GetDB()
	if !IsCommentExist(DB, uint(cid)) {
		res.HttpStatus = http.StatusBadRequest
		res.Code = 400
		res.Msg = "评论不存在或已被删除"
		return res
	}
	// DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	// DB.Model(&model.Reply{}).Where("cid = ?", cid)
	DB.Raw(sql, cid).Scan(&replies)
	res.Data = gin.H{"replies": replies}
	return res
}

func DeleteCommentService(id uint, uid interface{}) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("id = ? and uid = ?", id, uid).Delete(&model.Comment{})
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
}

func DeleteReplyService(id uint, uid interface{}) response.ResponseStruct {
	DB := common.GetDB()
	DB.Where("id = ? and uid = ?", id, uid).Delete(&model.Reply{})
	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       http.StatusOK,
		Data:       nil,
		Msg:        response.OK,
	}
}

func IsCommentExist(db *gorm.DB, cid uint) bool {
	var comment model.Comment
	db.Where("id = ?", cid).First(&comment)
	if comment.ID != 0 {
		return true
	}
	return false
}
