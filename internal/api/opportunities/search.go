package opportunities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SearchResult struct {
	TotalRecords      int               `json:"totalRecords"`
	Limit             int               `json:"limit"`
	Offset            int               `json:"offset"`
	OpportunitiesData []OpportunityData ` json:"opportunitiesData"`
	Links             []ResultLink      `json:"links"`
}

type OpportunityData struct {
	NoticeID                  string      `json:"noticeId"`
	Title                     string      `json:"title"`
	SolicitationNumber        string      `json:"solicitationNumber"`
	FullParentPathName        string      `json:"fullParentPathName"`
	FullParentPathCode        string      `json:"fullParentPathCode"`
	PostedDate                string      `json:"postedDate"`
	Type                      string      `json:"type"`
	BaseType                  string      `json:"baseType"`
	ArchiveType               string      `json:"archiveType"`
	ArchiveDate               string      `json:"archiveDate"`
	TypeOfSetAsideDescription interface{} `json:"typeOfSetAsideDescription"`
	TypeOfSetAside            interface{} `json:"typeOfSetAside"`
	ResponseDeadLine          string      `json:"responseDeadLine"`
	NaicsCode                 string      `json:"naicsCode"`
	NaicsCodes                []string    `json:"naicsCodes"`
	ClassificationCode        string      `json:"classificationCode"`
	Active                    string      `json:"active"`
	Award                     interface{} `json:"award"`
	PointOfContact            []struct {
		Fax      interface{} `json:"fax"`
		Type     string      `json:"type"`
		Email    string      `json:"email"`
		Phone    string      `json:"phone"`
		Title    interface{} `json:"title"`
		FullName string      `json:"fullName"`
	} `json:"pointOfContact"`
	Description      string `json:"description"`
	OrganizationType string `json:"organizationType"`
	OfficeAddress    struct {
		Zipcode     string `json:"zipcode"`
		City        string `json:"city"`
		CountryCode string `json:"countryCode"`
		State       string `json:"state"`
	} `json:"officeAddress"`
	PlaceOfPerformance struct {
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
	} `json:"placeOfPerformance"`
	AdditionalInfoLink interface{} `json:"additionalInfoLink"`
	UILink             string      `json:"uiLink"`
	Links              []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	ResourceLinks []string `json:"resourceLinks"`
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

	return json.NewDecoder(r.Body).Decode(target)
}

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

func GetOpportunities2(dateFrom string, dateTo string) *SearchResult {
	var url string
	result := &SearchResult{}
	//apiKey := "PqJ4mFnwiHJwoIUEgKVrmI1Mw2d2yENklw6UJues"
	apiKey := getApiKey()
	limit := 15
	url = fmt.Sprintf(getBaseUrl()+"/opportunities/v2/search?api_key=%s&limit=%d", apiKey, limit)
	if dateFrom != "" {
		url += "&postedFrom=" + dateFrom
	}
	if dateTo != "" {
		url += "&postedTo=" + dateTo
	}
	//url = "https://api.sam.gov/prod/opportunities/v2/search?api_key={{apiKey}}&limit=15&postedFrom=12/01/2023&postedTo=12/07/2023"
	err := getJson(url, result)
	if err != nil {
		return nil
	}

	println(result)

	return result
}

func getBaseUrl() string {
	//envFile, _ := godotenv.Read(".env")
	return "https://api-alpha.sam.gov"
}
func getApiKey() string {
	//envFile, _ := godotenv.Read(".env")
	return "3XASl0fa2aW5gvsMZQdbgqNDp3QPEJ8XnSsPmEFA"
}
