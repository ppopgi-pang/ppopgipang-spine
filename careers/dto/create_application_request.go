package dto

type CreateApplicationRequest struct {
	JobPostingId   int64   `json:"job_posting_id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Phone          *string `json:"phone"`
	ResumeFileName *string `json:"resume_file_name"`
	Memo           *string `json:"memo"`
}
