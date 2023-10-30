package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeaFood struct {
	Name            string  `json:"name"`
	Origin          string  `json:"origin"`
	OriginalPrice   float32 `json:"original_price"`
	DiscountedPrice float32 `json:"discounted_price"`
}

var fishes = []SeaFood{
	{
		Name:            "Salmon",
		Origin:          "Alaska",
		OriginalPrice:   25.99,
		DiscountedPrice: 20.99,
	},
	{
		Name:            "Tuna",
		Origin:          "Pacific",
		OriginalPrice:   32.99,
		DiscountedPrice: 29.99,
	},
}

var crustaceans = []SeaFood{
	{
		Name:            "Alaskan King Crab",
		Origin:          "Alaska",
		OriginalPrice:   59.99,
		DiscountedPrice: 55.99,
	},
	{
		Name:            "Green Crab",
		Origin:          "Vietnam",
		OriginalPrice:   25.99,
		DiscountedPrice: 22.99,
	},
}

func getFish(c *gin.Context) {
	// j, _ := json.Marshal(fishes)
	c.JSON(http.StatusOK, fishes)
}

func getCrustacean(c *gin.Context) {
	// j, _ := json.Marshal(fishes)
	c.JSON(http.StatusOK, crustaceans)
}

func main() {

	r := gin.Default()
	r.GET("/category/fish", getFish)
	r.GET("/category/crustacean", getCrustacean)
	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080
}
