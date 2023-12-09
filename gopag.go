package gopag

import (
	"fmt"
	"strconv"
)

type PaginationInfo struct {
	CurrentPage int32 `json:"current_page"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
	TotalPages  int32 `json:"total_pages"`
	NextPage    int32 `json:"next_page"`
	PrevPage    int32 `json:"prev_page"`
	TotalPage   int32 `json:"total_page"`
	TotalData   int32 `json:"total_data"`
}

func CalculateLimitOffset(pageSizeStr, pageNumberStr string) (limit, offset int16, err error) {
	// Convert strings to integers
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting pageSize to int: %v", err)
	}

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting pageNumber to int: %v", err)
	}

	// Validate and adjust values
	if pageSize < 1 {
		pageSize = 1
	}

	if pageNumber < 1 {
		pageNumber = 1
	}

	// Calculate limit and offset
	offset = int16((pageNumber - 1) * pageSize)
	limit = int16(pageSize)

	return limit, offset, nil
}

func CalculatePaginationInfo(limit, offset int16, dataSize int) PaginationInfo {
	var pageInfo PaginationInfo

	limit32 := int32(limit)
	offset32 := int32(offset)
	dataSize32 := int32(dataSize)

	// Calculate CurrentPage
	if limit > 0 {
		pageInfo.CurrentPage = (offset32 / limit32) + 1
	}

	// Set Limit and Offset
	pageInfo.Limit = limit32
	pageInfo.Offset = offset32

	// Calculate TotalPages
	if limit > 0 {
		pageInfo.TotalPages = (dataSize32 + limit32 - 1) / limit32
	}

	// Calculate NextPage and PrevPage
	if pageInfo.CurrentPage < pageInfo.TotalPages {
		pageInfo.NextPage = pageInfo.CurrentPage + 1
	}
	if pageInfo.CurrentPage > 1 {
		pageInfo.PrevPage = pageInfo.CurrentPage - 1
	}

	// Set TotalPage and TotalData
	pageInfo.TotalPage = pageInfo.TotalPages
	pageInfo.TotalData = dataSize32

	return pageInfo
}
