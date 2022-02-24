package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"gorm.io/gorm"
)

type AdvisorBitrix struct {
	UserID        string `json:"ID"`
	PersonalSreet string `json:"PERSONAL_STREET"`
}

type Advisor struct {
	gorm.Model
	ID           int
	BitrixID     string `json:"ID"`
	Name         string `json:"NAME"`
	LastName     string `json:"LAST_NAME"`
	Email        string `json:"EMAIL"`
	Photo        string `json:"PERSONAL_PHOTO"`
	WorkPosition string `json:"WORK_POSITION"`
	UserType     string `json:"USER_TYPE"`
	Active       bool   `json:"ACTIVE"`
}

type ResponseAdvisors struct {
	Result []BitrixAdvisors `json:"result"`
}

type BitrixAdvisors struct {
	ID           string `json:"ID"`
	Name         string `json:"NAME"`
	LastName     string `json:"LAST_NAME"`
	Email        string `json:"EMAIL"`
	Photo        string `json:"PERSONAL_PHOTO"`
	WorkPosition string `json:"WORK_POSITION"`
	UserType     string `json:"USER_TYPE"`
	Active       bool   `json:"ACTIVE"`
}

// UpdateBitrixGuardAdvisor update field PERSONAL_STREET by USER_ID on bitrix24
func UpdateBitrixGuardAdvisor(advisorBitrix *AdvisorBitrix) (err error) {

	api := "https://intranet.idex.cc/rest/1/" // os.Getenv("BITRIX_SITE")
	token := "evcwp69f5yg7gkwc"               // os.Getenv("BITRIX_TOKEN")

	params := url.Values{}
	params.Add("ID", advisorBitrix.UserID)
	params.Add("PERSONAL_STREET", advisorBitrix.PersonalSreet)

	resp, err := http.PostForm(api+token+"/user.update", params)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
	// Unmarshal result
	guardAdvisor := AdvisorBitrix{}
	err = json.Unmarshal(body, &guardAdvisor)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	log.Println("Guard assigned to advisor: ", advisorBitrix.UserID)
	return nil
}
