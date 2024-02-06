package opportunity

import (
	"encoding/json"
	"fmt"
	"github.com/gen2brain/beeep"
	"gorm.io/gorm"
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

type PaginatedResult struct {
	Total int64                  `json:"total"`
	Data  []database.Opportunity `json:"data"`
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

func Search(keyword string, filters OpportunityFilter) PaginatedResult {
	var data []database.Opportunity
	// keywords
	db := database.GetDbInstance()

	tx := db.Debug().Session(&gorm.Session{NewDB: true}).Preload("PointOfContact").Preload("Links").Scopes(PaginatedSearch(keyword, filters))
	if filters.Page > 0 && filters.PerPage > 0 {
		offset := (filters.Page - 1) * filters.PerPage
		tx.Offset(int(offset)).Limit(int(filters.PerPage))
	}
	tx.Order("posted_date desc")
	tx.Find(&data)

	count := int64(0)
	db.Debug().Session(&gorm.Session{NewDB: true}).Model(&database.Opportunity{}).Scopes(PaginatedSearch(keyword, filters)).Count(&count)
	fmt.Printf("Result count: %d %d", count, tx.RowsAffected)

	return PaginatedResult{
		Total: count,
		Data:  data,
	}
}

func PaginatedSearch(keyword string, filter OpportunityFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// keywords
		if keyword != "" {
			db.Where(
				db.Where(
					"notice_id LIKE ?",
					"%"+keyword+"%",
				).Or(
					"title LIKE ?",
					"%"+keyword+"%",
				).Or(
					"solicitation_number LIKE ?",
					"%"+keyword+"%",
				).Or(
					"full_parent_path_name LIKE ?",
					"%"+keyword+"%",
				),
			)
		}

		if filter.FromDate != "" {
			db.Where("DATE(posted_date) >= ?", filter.FromDate)
		}
		if filter.ToDate != "" {
			db.Where("DATE(posted_date) <= ?", filter.ToDate)
		}
		if len(filter.Type) > 0 {
			db.Where("type IN (?)", filter.Type)
		}
		if len(filter.NaicsCode) > 0 {
			db.Where("naics_code IN (?)", filter.NaicsCode)
		}

		return db
	}
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
