package dto

import reviewDtos "github.com/ppopgi-pang/ppopgipang-spine/reviews/dto"

type StoreReviewResponse struct {
	ReviewImages    []string                    `json:"review_images"`
	ReviewResponses []reviewDtos.ReviewResponse `json:"reviews_responses"`
}
