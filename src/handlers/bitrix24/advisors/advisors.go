package advisors

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/91diego/backend-guardias/src/database"
	"github.com/91diego/backend-guardias/src/models"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Advisor{})
	return &UserRepo{Db: db}
}

func GetAdvisors(c *gin.Context) {
	api := os.Getenv("BITRIX_SITE")
	token := os.Getenv("BITRIX_TOKEN")
	url := api + token + "/user.get?USER_TYPE=employee&WORK_POSITION=ASESOR%20INMOBILIARIO&UF_DEPARTMENT=59" // + perPage + "&include_totals=" + includeTotals + "&include_fields=" + includeFields + "&page=" + pageNumber + "&sort=" + sortField
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", token)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	c.Data(http.StatusOK, gin.MIMEJSON, body)
}
