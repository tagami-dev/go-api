package services

import (
	"github.com/tgm-tmy/go-api/apperrors"
	"github.com/tgm-tmy/go-api/models"
	"github.com/tgm-tmy/go-api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
