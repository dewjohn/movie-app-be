package service

import (
	"github.com/gin-gonic/gin"
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
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	if !IsMovieExit(DB, comment.Vid) {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.CheckFailCode
		res.Msg = response.MovieNotExit
		return res
	}
	newComment := model.Comment{
		Vid:      comment.Vid,
		Content:  comment.Content,
		ParentId: comment.ParentID,
		Uid:      uid.(uint),
	}
	DB.Create(&newComment)
	res.Data = gin.H{"id": newComment.ID}
	return res
}

func GetCommentService(vid int, replyCount int, page int, pageSize int) response.ResponseStruct {
	var total int64        // 总评论数
	var commentCount int64 // 一级评论数
	var comments []vo.CommentVo

	DB := common.GetDB()
	DB.Model(model.Comment{}).Where("vid = ?", vid).Count(&total)
	DB.Model(model.Comment{}).Where("parent_id != 0 and vid = ?", vid).Count(&commentCount)
	sql := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from users,comments " +
		"where comments.deleted_at is null and comments.uid = users.id and vid = ? and parent_id = 0 limit ? offset ?"
	sqlReply := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from comments,users " +
		"where comments.deleted_at is null and comments.uid = users.id and parent_id = ? limit ?"
	DB.Raw(sql, vid, pageSize, (page-1)*pageSize).Scan(&comments)

	if replyCount > 0 {
		//当需要的回复数大于0时，查询回复
		for i := 0; i < len(comments); i++ {
			DB.Raw(sqlReply, comments[i].ID, replyCount).Scan(&comments[i].Reply)
		}
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "comment_count": commentCount, "comments": comments},
		Msg:        response.OK,
	}
}

func GetCommentListService(vid, page, pageSize int) response.ResponseStruct {
	var total int64 // 记录总数
	var comments []vo.CommentListVo

	DB := common.GetDB()
	DB.Model(model.Comment{}).Where("parent_id = 0 and vid = ?", vid).Count(&total)
	sql := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from users,comments " +
		"where comments.deleted_at is null and comments.uid = users.id and vid = ? and parent_id = 0 limit ? offset ?"
	sqlReplyCount := "select count(*) from comments where comments.deleted_at is null and comments.parent_id <> 0 and comments.parent_id = ?"

	DB.Raw(sql, vid, pageSize, (page-1)*pageSize).Scan(&comments)

	for i := 0; i < len(comments); i++ {
		DB.Raw(sqlReplyCount, comments[i].ID).Scan(&comments[i].ReplyCount)
	}

	return response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       gin.H{"count": total, "comments": comments},
		Msg:        response.OK,
	}
}

func GetReplyDetailsService(cid, offset, page, pageSize int) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	var total int64
	var replies []vo.ReplyVo

	sql := "select comments.id,comments.created_at,content,users.id as uid,name,avatar from comments,users " +
		"where comments.deleted_at is null and comments.uid = users.id and parent_id = ? limit ? offset ?"
	DB := common.GetDB()
	DB.Model(&model.Comment{}).Where("parent_id = ?", cid).Count(&total)
	DB.Raw(sql, cid, pageSize, (page-1)*pageSize+offset).Scan(&replies)
	res.Data = gin.H{
		"count":   total,
		"replies": replies,
	}
	return res
}

func DeleteCommentService(id uint, uid interface{}) response.ResponseStruct {
	res := response.ResponseStruct{
		HttpStatus: http.StatusOK,
		Code:       response.SuccessCode,
		Data:       nil,
		Msg:        response.OK,
	}
	DB := common.GetDB()
	var comment model.Comment
	DB.Where("id = ?", id).First(&comment)
	if comment.ID != 0 && comment.Uid == uid {
		DB.Where("id = ? and uid = ?", id, uid).Delete(&comment)
	}

	return res
}
