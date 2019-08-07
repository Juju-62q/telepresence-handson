package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Secret struct {
	Secret string `json:"secret"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		res, err := http.Get(os.Getenv("APP_B_ENDPOINT") + "/secret")
		if err != nil {
			log.Fatal(err)
			return
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatal("StatusCode=%d", res.StatusCode)
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		var secretResponse Secret
		err = json.Unmarshal(body, &secretResponse)
		if err != nil {
			log.Fatal(err)
			return
		}

		c.JSON(200, gin.H{
			"message": "this is added!",
      "secret": secretResponse.Secret,
		})
	})
  r.Run(":80") // listen and serve on 0.0.0.0:8080
}
