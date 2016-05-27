package webgin

type Article struct{
Id int64 `db:"article_id"`
Created int64
Title string
Content string
}

// Provides a list of articles
func ArticleList(c *gin.Context){
    var articles []Article
    _, err := dbMap.Select(&articles,  "select * from articles order by article_id")
    checkErr(err, "Select failed")
    content := gin.H{}
    for k, v := range articles {
        content[strconv.Itoa(k)] = v
    }
    c.JSON(200, content)
}