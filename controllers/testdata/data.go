package testdata

import "github.com/tgm-tmy/go-api/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "1st",
		Contents:    "The first article will be published",
		UserName:    "tgm",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "The second article will be about computer science",
		UserName: "tgm",
		NiceNum:  4,
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st hello world",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "awesome",
	},
}
