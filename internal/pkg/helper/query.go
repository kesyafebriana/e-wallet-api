package helper

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/kesyafebriana/e-wallet-api/internal/dto"
)

func QueryGetAllTransaction(pagination *dto.PaginationInfo, totalData int) (string, int) {
	var queryBuilder strings.Builder
	var totalPage int
	var limit int

	queryBuilder.WriteString(`SELECT id, sender_wallet_id, recipient_wallet_id, amount, source_of_funds, description, created_at, updated_at FROM transactions 
	WHERE (sender_wallet_id = $1 OR recipient_wallet_id = $1) `)

	if *pagination.StartDate != "" {
		_, err := time.Parse("2006-01-02", *pagination.StartDate)
		if err == nil {
			queryBuilder.WriteString(fmt.Sprintf("AND created_at > '%s' ", *pagination.StartDate))
		}
	}

	if *pagination.EndDate != "" {
		_, err := time.Parse("2006-01-02", *pagination.EndDate)
		if err == nil {
			queryBuilder.WriteString(fmt.Sprintf("AND created_at < '%s' ", *pagination.EndDate))
		}
	}

	if *pagination.Search != "" {
		queryBuilder.WriteString(fmt.Sprintln("AND description ILIKE '%", *pagination.Search, "%' "))
	}

	if *pagination.SortBy != "" {
		queryBuilder.WriteString(fmt.Sprintf("ORDER BY %s ", *pagination.SortBy))
	} else {
		queryBuilder.WriteString(`ORDER BY created_at `)
	}

	if *pagination.Sort == "asc" {
		queryBuilder.WriteString(`ASC `)
	} else {
		queryBuilder.WriteString(`DESC `)
	}

	if *pagination.Limit != "" {
		limit, _ = strconv.Atoi(*pagination.Limit)

		queryBuilder.WriteString(fmt.Sprintf("LIMIT %s ", *pagination.Limit))
	} else {
		limit = 10

		queryBuilder.WriteString("LIMIT 10 ")
	}

	totalPage = int(math.Round(float64(totalData / limit)))

	log.Println(totalData, totalPage, limit)
	if *pagination.Page != "" {
		page, err := strconv.Atoi(*pagination.Page)

		if err == nil && page <= totalPage {
			offset := (limit * page) - limit
			res := strconv.Itoa(offset)
			queryBuilder.WriteString(fmt.Sprintf("OFFSET %s", res))
		}
	}

	log.Println(queryBuilder.String())
	return queryBuilder.String(), totalPage
}
