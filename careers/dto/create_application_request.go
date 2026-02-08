package dto

type CreateApplicationRequest struct {
	JobPostingId   int64   `json:"job_posting_id" example:"101"`
	Name           string  `json:"name" example:"김민수"`
	Email          string  `json:"email" example:"minsoo.kim@example.com"`
	Phone          *string `json:"phone" example:"010-1234-5678"`
	ResumeFileName *string `json:"resume_file_name" example:"resume.pdf"`
	Memo           *string `json:"memo" example:"백엔드 경력 5년입니다."`
}
