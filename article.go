package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Id      int64 `db:"articleId"`
	Created int64
	Title   string
	Content string
}

// Provides a list of articles
func ArticlesList(c *gin.Context) {
	var articles []Article
	_, err := dbmap.Select(&articles, "select * from articles order by articleId")
	checkErr(err, "Select failed")
	content := gin.H{}
	for k, v := range articles {
		content[strconv.Itoa(k)] = v
	}

	c.JSON(200, content)
}

// Allows recovering one specific article
func ArticleDetail(c *gin.Context) {
	articleID := c.Params.ByName("id")
	aID, _ := strconv.Atoi(articleID)
	article := getArticle(aID)
	content := gin.H{"title": article.Title, "content": article.Content}
	c.JSON(200, content)
}

// Post article to database
func ArticlePost(c *gin.Context) {
	var json Article

	c.Bind(json)
	article := createArticle(json.Title, json.Content)
	if article.Title == json.Title {
		content := gin.H{
			"result":  "Success",
			"title":   article.Title,
			"content": article.Content,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}

func createArticle(title, body string) Article {
	article := Article{
		Created: time.Now().UnixNano(),
		Title:   title,
		Content: body,
	}

	err := dbmap.Insert(&article)
	checkErr(err, "Insert failed")
	return article
}

func getArticle(articleID int) Article {
	article := Article{}
	err := dbmap.SelectOne(&article, "select * from articles where articleId=?", articleID)
	checkErr(err, "SelectOne failed")
	return article
}
