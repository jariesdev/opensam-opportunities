package opportunities

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"open-gsa/internal/database"
	"os"
	"time"
)

type SearchResult struct {
	TotalRecords      int32             `json:"totalRecords"`
	Limit             int32             `json:"limit"`
	Offset            int32             `json:"offset"`
	OpportunitiesData []OpportunityData ` json:"opportunitiesData"`
	Links             []ResultLink      `json:"links"`
}

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

type OpportunityData struct {
	NoticeID                  string             `json:"noticeId"`
	Title                     string             `json:"title"`
	SolicitationNumber        string             `json:"solicitationNumber"`
	FullParentPathName        string             `json:"fullParentPathName"`
	FullParentPathCode        string             `json:"fullParentPathCode"`
	PostedDate                string             `json:"postedDate"`
	Type                      string             `json:"type"`
	BaseType                  string             `json:"baseType"`
	ArchiveType               string             `json:"archiveType"`
	ArchiveDate               string             `json:"archiveDate"`
	TypeOfSetAsideDescription string             `json:"typeOfSetAsideDescription"`
	TypeOfSetAside            string             `json:"typeOfSetAside"`
	ResponseDeadLine          string             `json:"responseDeadLine"`
	NaicsCode                 string             `json:"naicsCode"`
	NaicsCodes                []string           `json:"naicsCodes"`
	ClassificationCode        string             `json:"classificationCode"`
	Active                    string             `json:"active"`
	Award                     Award              `json:"award"`
	PointOfContact            []PointOfContact   `json:"pointOfContact"`
	Description               string             `json:"description"`
	OrganizationType          string             `json:"organizationType"`
	OfficeAddress             OfficeAddress      `json:"officeAddress"`
	PlaceOfPerformance        PlaceOfPerformance `json:"placeOfPerformance"`
	AdditionalInfoLink        string             `json:"additionalInfoLink"`
	UILink                    string             `json:"uiLink"`
	Links                     []Link             `json:"links"`
	ResourceLinks             []string           `json:"resourceLinks"`
}

type PlaceOfPerformance struct {
	City struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"city"`
	State struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"state"`
	Country struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"country"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type OfficeAddress struct {
	Zipcode     string `json:"zipcode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	State       string `json:"state"`
}
type PointOfContact struct {
	Fax      string `json:"fax"`
	Type     string `json:"type"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Title    string `json:"title"`
	FullName string `json:"fullName"`
}

type Award struct {
	Date    string `json:"date"`
	Number  string `json:"number"`
	Amount  string `json:"amount"`
	Awardee struct {
		Name     string `json:"name"`
		Location struct {
			City struct {
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"state"`
			Zip     string `json:"zip"`
			Country struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"location"`
		UeiSAM   string `json:"ueiSAM"`
		CageCode string `json:"cageCode"`
	} `json:"awardee"`
}

type ResultLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode > 299 {
		errorResponse := ErrorResponse{}
		err := json.NewDecoder(r.Body).Decode(&errorResponse)
		if err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("%d, %s", r.StatusCode, errorResponse.Error.Message))
	}

	return json.NewDecoder(r.Body).Decode(target)
}

//func getTestResponse() SearchResult {
//	file, _ := ioutil.ReadFile("./sample-response.json")
//
//	data := SearchResult{}
//
//	_ = json.Unmarshal([]byte(file), &data)
//
//	return data
//}

//func GetOpportunities() {
//	url := "https://api.sam.gov/prod/opportunities/v2/search?api_key={{apiKey}}&limit=15&postedFrom=12/01/2023&postedTo=12/07/2023"
//	resp, err := http.Get(url)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//
//	fmt.Print(body)
//}

func GetOpportunities(dateFrom string, dateTo string, ptype string) (SearchResult, error) {
	//return exampleResponse()

	var url string
	result := SearchResult{}
	apiKey := getApiKey()
	limit := 1000
	url = fmt.Sprintf(getBaseUrl()+"/opportunities/v2/search?api_key=%s&limit=%d", apiKey, limit)
	//url = getBaseUrl() + "/opportunities/v2/search?limit=1000"
	if dateFrom != "" {
		url += "&postedFrom=" + dateFrom
	}
	if dateTo != "" {
		url += "&postedTo=" + dateTo
	}
	if ptype != "" {
		url += "&ptype=" + ptype
	}
	//url = "https://api.sam.gov/prod/opportunities/v2/search?api_key={{apiKey}}&limit=1000&postedFrom=12/01/2023&postedTo=12/07/2023"
	fmt.Printf("Request to URL: %s \n ", url)
	err := getJson(url, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func getBaseUrl() string {
	//envFile, _ := godotenv.Read(".env")
	//return "https://api-alpha.sam.gov"
	return "https://api.sam.gov"
}
func getApiKey() string {
	settingModel := database.Setting{}
	db := database.GetDbInstance()
	result := db.Model(&database.Setting{}).First(&settingModel, "key = ?", "api_key")
	if result.RowsAffected > 0 {
		return settingModel.Value
	}

	return ""
}

func exampleResponse() SearchResult {
	// read file
	data, err := os.ReadFile("./internal/api/opportunities/sample-response2.json")
	if err != nil {
		fmt.Println(err)
	}

	// json data
	var obj SearchResult

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}
