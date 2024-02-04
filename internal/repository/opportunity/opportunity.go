package opportunity

import (
	"encoding/json"
	"fmt"
	"github.com/gen2brain/beeep"
	"open-gsa/internal/api/opportunities"
	"open-gsa/internal/database"
	"strconv"
	"strings"
	"time"
)

type OpportunityFilter struct {
	FromDate  string   `json:"fromDate"`
	ToDate    string   `json:"toDate"`
	Type      []string `json:"type"`
	NaicsCode []string `json:"naicsCode"`
	Page      uint32   `json:"page"`
	PerPage   uint32   `json:"perPage"`
}

func PullLatest() string {
	// only pull data once a day
	count := 1
	lastPull := database.Setting{}
	db := database.GetDbInstance()
	found := db.First(&lastPull, "key = ?", "daily_pull_count")
	if found.RowsAffected > 0 {
		count, _ = strconv.Atoi(lastPull.Value)
	}

	var totalItems int32
	ptypes := []string{"p", "r", "o"}
	for _, ptype := range ptypes {
		// if the daily count exceeded return
		if count > 5 {
			fmt.Printf("Exceeded daily pull")
		}

		fmt.Printf("Daily Pull Count: %d", count)

		from := time.Now().AddDate(0, 0, -7).Format("01/02/2006")
		to := time.Now().Format("01/02/2006")
		result := opportunities.GetOpportunities(from, to, ptype)

		totalItems += result.TotalRecords

		var opportunityItems []opportunities.OpportunityData
		if result.OpportunitiesData != nil {
			opportunityItems = result.OpportunitiesData
			insertToDB(opportunityItems)
		}

		// update daily count
		db.Where(database.Setting{Key: "daily_pull_count"}).Updates(database.Setting{Value: strconv.Itoa(count + 1)})
	}

	err := beeep.Alert("OpenGsaApp", "Pulled latest opportunities.", "static/favicon.png")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Pulled %d items \n", totalItems)
}

// insert if not exists or update if exists based on notice ID
func insertToDB(opportunityData []opportunities.OpportunityData) {
	db := database.GetDbInstance()
	db.FirstOrCreate(&database.Setting{Key: "last_pull", Value: time.Now().Format(time.DateTime)})
	db.Where(database.Setting{Key: "last_pull"}).Updates(database.Setting{Value: time.Now().Format(time.DateTime)})
	for _, item := range opportunityData {
		officeAddress, _ := json.Marshal(item.OfficeAddress)

		row := &database.Opportunity{
			NoticeID:                  item.NoticeID,
			Title:                     item.Title,
			SolicitationNumber:        item.SolicitationNumber,
			FullParentPathName:        item.FullParentPathName,
			FullParentPathCode:        item.FullParentPathCode,
			PostedDate:                item.PostedDate,
			Type:                      item.Type,
			BaseType:                  item.BaseType,
			ArchiveType:               item.ArchiveType,
			ArchiveDate:               item.ArchiveDate,
			ResponseDeadLine:          item.ResponseDeadLine,
			NaicsCode:                 item.NaicsCode,
			ClassificationCode:        item.ClassificationCode,
			Active:                    item.Active,
			Description:               item.Description,
			OrganizationType:          item.OrganizationType,
			UILink:                    item.UILink,
			TypeOfSetAsideDescription: item.TypeOfSetAsideDescription,
			TypeOfSetAside:            item.TypeOfSetAside,
			NaicsCodes:                strings.Join(item.NaicsCodes, ","),
			OfficeAddress:             string(officeAddress),
			AdditionalInfoLink:        item.AdditionalInfoLink,
			ResourceLinks:             strings.Join(item.ResourceLinks, ","),
		}
		db.Where(database.Opportunity{NoticeID: item.NoticeID}).Assign(row).FirstOrCreate(row)

		// store links
		for _, link := range item.Links {
			linkModel := &database.Link{
				Rel:  link.Rel,
				Href: link.Href,
			}
			db.FirstOrCreate(linkModel, &database.Link{
				OpportunityID: row.ID,
				Href:          link.Href,
			})
		}

		// store point of contracts
		for _, poc := range item.PointOfContact {
			linkModel := &database.PointOfContact{
				Fax:      poc.Fax,
				Type:     poc.Type,
				Email:    poc.Email,
				Phone:    poc.Phone,
				Title:    poc.Email,
				FullName: poc.FullName,
			}
			db.FirstOrCreate(linkModel, &database.PointOfContact{
				OpportunityID: row.ID,
				Email:         poc.Email,
			})
		}
	}
}

func Search(keyword string, filters OpportunityFilter) []database.Opportunity {
	var result []database.Opportunity
	db := database.GetDbInstance()
	// keywords
	query := db.Where("notice_id LIKE ? OR title LIKE ? OR solicitation_number LIKE ? OR full_parent_path_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	query.Preload("PointOfContact").Preload("Links")

	if filters.FromDate != "" {
		query.Where("DATE(posted_date) >= ?", filters.FromDate)
	}
	if filters.ToDate != "" {
		query.Where("DATE(posted_date) <= ?", filters.ToDate)
	}
	if len(filters.Type) > 0 {
		query.Where("type IN (?)", strings.Join(filters.Type, ","))
	}
	if len(filters.NaicsCode) > 0 {
		query.Where("naics_code IN (?)", strings.Join(filters.NaicsCode, ","))
	}

	if filters.Page > 0 && filters.PerPage > 0 {
		offset := (filters.Page - 1) * filters.PerPage
		query.Offset(int(offset)).Limit(int(filters.PerPage))
	}

	query.Order("posted_date desc").Find(&result)

	var count int64
	query.Count(&count)

	fmt.Printf("Total Results:%d", count)

	return result
}

func GetTypes() []string {
	var types []string
	db := database.GetDbInstance()
	db.Model(&database.Opportunity{}).Distinct("type").Order("type").Find(&types)
	return types
}

func GetNaicsCodes() []string {
	var types []string
	db := database.GetDbInstance()
	db.Model(&database.Opportunity{}).Distinct("naics_code").Order("naics_code").Find(&types)
	return types
}
