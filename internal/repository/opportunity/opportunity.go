package opportunity

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"open-gsa/internal/api/opportunities"
	"open-gsa/internal/database"
	"strconv"
	"time"
)

type OpportunityFilter struct {
	FromDate string
	ToDate   string
}

func PullLatest() string {
	now := time.Now()

	// only pull data once a day
	count := 1
	lastPull := database.Setting{}
	db := database.GetDbInstance()
	found := db.First(&lastPull, "key = ?", "daily_pull_count")
	if found.RowsAffected > 0 {
		count, _ = strconv.Atoi(lastPull.Value)
	}

	// if the daily count exceeded return
	if count > 5 {
		fmt.Printf("Exceeded daily pull")
	}

	fmt.Printf("Daily Pull Count: %d", count)

	from := now.Format("02/01/2006")
	to := now.AddDate(0, -1, 0).Format("02/01/2006")
	result := opportunities.GetOpportunities(from, to)

	var opportunityItems []opportunities.OpportunityData
	if result.OpportunitiesData != nil {
		opportunityItems = result.OpportunitiesData
		insertToDB(opportunityItems)
	}

	// update daily count
	db.Where(database.Setting{Key: "daily_pull_count"}).Updates(database.Setting{Value: strconv.Itoa(count + 1)})

	err := beeep.Alert("OpenGsaApp", "Pulled latest opportunities.", "static/favicon.png")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Pulled %d items", len(opportunityItems))
}

// insert if not exists or update if exists based on notice ID
func insertToDB(opportunityData []opportunities.OpportunityData) {
	db := database.GetDbInstance()
	db.FirstOrCreate(&database.Setting{Key: "last_pull", Value: time.Now().Format(time.DateTime)})
	db.Where(database.Setting{Key: "last_pull"}).Updates(database.Setting{Value: time.Now().Format(time.DateTime)})
	for _, item := range opportunityData {
		row := &database.Opportunity{
			NoticeID:           item.NoticeID,
			Title:              item.Title,
			SolicitationNumber: item.SolicitationNumber,
			FullParentPathName: item.FullParentPathName,
			FullParentPathCode: item.FullParentPathCode,
			PostedDate:         item.PostedDate,
			Type:               item.Type,
			BaseType:           item.BaseType,
			ArchiveType:        item.ArchiveType,
			ArchiveDate:        item.ArchiveDate,
			ResponseDeadLine:   item.ResponseDeadLine,
			NaicsCode:          item.NaicsCode,
			ClassificationCode: item.ClassificationCode,
			Active:             item.Active,
			Description:        item.Description,
			OrganizationType:   item.OrganizationType,
			UILink:             item.UILink,
		}
		db.Where(database.Opportunity{NoticeID: item.NoticeID}).Assign(row).FirstOrCreate(row)
	}
}

func Search(keyword string, filters OpportunityFilter) []database.Opportunity {
	var result []database.Opportunity
	db := database.GetDbInstance()
	// keywords
	query := db.Where("notice_id LIKE ? OR title LIKE ? OR solicitation_number LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")

	if filters.FromDate != "" {
		query.Where("DATE(posted_date) >= ?", filters.FromDate)
	}
	if filters.ToDate != "" {
		query.Where("DATE(posted_date) <= ?", filters.ToDate)
	}

	query.Find(&result)

	return result
}
