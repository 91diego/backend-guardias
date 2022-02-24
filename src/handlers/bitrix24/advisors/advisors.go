package advisors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/91diego/backend-guardias/src/models"
	"github.com/gin-gonic/gin"
)

// GetAdvisors retrieve advisors from bitrix24
func GetAdvisors(c *gin.Context) {

	var response models.ResponseAdvisors
	api := "https://intranet.idex.cc/rest/1/" // os.Getenv("BITRIX_SITE")
	token := "evcwp69f5yg7gkwc"               // os.Getenv("BITRIX_TOKEN")
	url := api + token + "/user.get?USER_TYPE=employee&WORK_POSITION=ASESOR%20INMOBILIARIO&UF_DEPARTMENT=59&ACTIVE=true"

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
	c.JSON(http.StatusOK, gin.H{
		"message": "List of advisors",
		"code":    http.StatusOK,
		"items":   response,
	})
}
