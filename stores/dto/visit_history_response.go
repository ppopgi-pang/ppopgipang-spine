package dto

import reviewDtos "github.com/ppopgi-pang/ppopgipang-spine/reviews/dto"

type MyStat struct {
	VisitCount int `json:"visit_count"`
}

type OtherUserStat struct {
	AverageVisitCount   int `json:"average_visit_count"`
	MaxVisitCount       int `json:"max_visit_count"`
	MonthlyVisitorCount int `json:"monthly_visitor_count"`
}

type VisitHistoryResponse struct {
	MyStat          *MyStat                     `json:"my_stat"`
	OtherUserStat   OtherUserStat               `json:"other_user_stat"`
	ReviewImages    []string                    `json:"review_images"`
	ReviewResponses []reviewDtos.ReviewResponse `json:"reviews_responses"`
}
