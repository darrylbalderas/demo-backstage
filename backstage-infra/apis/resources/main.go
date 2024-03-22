package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api-frameworks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api-frameworks": []string{
				"gin",
				"django",
				"fastapi",
				"spring",
			},
		})
	})
	r.GET("/groups", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"groups": []string{
				"groupA",
				"groupB",
			},
		})
	})
	r.GET("/teams", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"teams": []string{
				"teamA",
				"teamB",
				"teamC",
			},
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": []string{
				"userA",
				"userB",
				"userC",
				"userD",
				"userE",
				"userF",
			},
		})
	})
	r.GET("/systems", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"systems": []string{
				"data-ingestion",
				"ml-training",
				"cloud-platform",
				"analytics",
			},
		})
	})
	r.GET("/resource-types", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"resource-types": []string{
				"aws-s3-bucket",
				"aws-sqs",
				"aws-rds-cluster",
			},
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
