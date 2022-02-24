package developments

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/91diego/backend-guardias/config"
	"github.com/91diego/backend-guardias/src/models"
	"github.com/gin-gonic/gin"
)

// GetDevelopments retrieve developments from bitrix24
func GetDevelopments(c *gin.Context) {

	var response models.ResponseDevelopments
	bitrix, _ := config.SetUp()

	api := bitrix.BitrixSite
	token := bitrix.BitrixToken
	url := api + token + "/crm.deal.fields"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"items":   "",
		})
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"items":   "",
		})
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"code":    http.StatusInternalServerError,
			"items":   "",
		})
	}
	defer res.Body.Close()

	json.Unmarshal(body, &response)
	c.JSON(http.StatusOK, gin.H{
		"message": "List of developments",
		"code":    http.StatusOK,
		"items":   response,
	})
}
