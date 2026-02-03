package dto

type JobPostingModifyRequest struct {
	Title        *string `json:"title"`
	Description  *string `json:"description"`
	Department   *string `json:"Department"`
	PositionType *string `json:"position_type"`
	Location     *string `json:"location"`
	IsActive     *bool   `json:"is_active"`
}
