package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/tgm-tmy/go-api/models"
	"github.com/tgm-tmy/go-api/repositories"
	"github.com/tgm-tmy/go-api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)

	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Fatalf("got %d articles, expected %d", num, expectedNum)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title,
					test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents,
					test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {
	const resetIncrementSql = "ALTER TABLE articles AUTO_INCREMENT = 1"
	_, err := testDB.Exec(resetIncrementSql)
	if err != nil {
		t.Fatal(err)
	}
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "tgm",
	}
	expectedArticleNum := 3
	var newArticle models.Article
	newArticle, err = repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum,
			newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `delete from articles  where title = ? and contents = ? and username = ?`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}
