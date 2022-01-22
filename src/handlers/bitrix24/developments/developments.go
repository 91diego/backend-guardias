package developments

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/91diego/backend-guardias/src/models"
	"github.com/gin-gonic/gin"
)

// GetDevelopments retrieve developments from bitrix24
func GetDevelopments(c *gin.Context) {

	var response models.ResponseDevelopments
	api := os.Getenv("BITRIX_SITE")
	token := os.Getenv("BITRIX_TOKEN")
	url := api + token + "/crm.deal.fields"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    400,
			"items":   "",
		})
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    400,
			"items":   "",
		})
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"code":    400,
			"items":   "",
		})
	}
	defer res.Body.Close()

	json.Unmarshal(body, &response)
	c.JSON(200, gin.H{
		"message": "List of developments",
		"code":    200,
		"items":   response,
	})
}
