package dto

type RewardResponse struct {
	Exp            int                 `json:"exp" example:"50"`
	TotalExp       int                 `json:"total_exp" example:"1250"`
	CurrentLevel   int                 `json:"current_level" example:"5"`
	LevelUp        bool                `json:"level_up" example:"true"`
	NewLevel       *int                `json:"new_level" example:"6"`
	ExpToNextLevel int                 `json:"exp_to_next_level" example:"150"`
	NewStamp       *StampResponse      `json:"new_stamp"`
	NewBadges      *[]NewBadgeResponse `json:"new_badges"`
}
