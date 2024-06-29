package dto

type PaginationInfo struct {
	Page      *string
	SortBy    *string
	Sort      *string
	StartDate *string
	EndDate   *string
	Search    *string
	Limit     *string
}
