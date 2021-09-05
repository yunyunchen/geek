package service

import (
	pb "geek/04/api/v1"
	"geek/04/internal/biz"
	"github.com/gin-gonic/gin"
)

func (s *UserService) UpdateArticle(ctx *gin.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	err := s.user.Update(ctx, req.Id, &biz.User{
		Name:  req.Name,
		Email: req.Email,
	})
	return &pb.UpdateUserReply{}, err
}
