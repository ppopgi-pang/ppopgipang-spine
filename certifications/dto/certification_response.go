package dto

import gameDtos "github.com/ppopgi-pang/ppopgipang-spine/gamification/dto"

type CertificationResponse struct {
	ID      int64                   `json:"id" example:"1001"`
	Type    string                  `json:"type" example:"loot"`
	Rewards gameDtos.RewardResponse `json:"rewards"`
}
