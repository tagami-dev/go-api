package testdata

import "github.com/tgm-tmy/go-api/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "1st",
		Contents: "The first article will be published",
		UserName: "tgm",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "The second article will be about computer science",
		UserName: "tgm",
		NiceNum:  4,
	},
}
