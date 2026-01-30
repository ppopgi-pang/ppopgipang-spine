package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/path"
	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/service"
)

type CareerController struct {
	service *service.CareerService
}

func NewCareerController(careerService *service.CareerService) *CareerController {
	return &CareerController{service: careerService}
}

// @Summary (공통) 모집 공고 목록 조회
// @Description 공고 목록을 활성화 상태, 페이지네이션 QueryString 옵션으로 가져옵니다.
// @Tags Careers
// @Param isActive query bool false "공고 활성화 상태"
// @Param page query int true "요청 페이지 번호"
// @Param size query int true "한번에 받을 페이지의 사이즈"
// @Success 200 {object} dto.JobPostingListResponse
// @Router /careers/job-postings [GET]
func (c *CareerController) GetJobPostings(ctx context.Context, query query.Values, page query.Pagination) (httpx.Response[dto.JobPostingListResponse], error) {
	isActive := query.GetBoolByKey("isActive", false)

	result, err := c.service.GetJobPostings(ctx, isActive, page.Page, page.Size)
	if err != nil {
		return httpx.Response[dto.JobPostingListResponse]{}, err
	}

	return httpx.Response[dto.JobPostingListResponse]{
		Body: result,
	}, nil
}

// @Summary (관리자) 모집 공고 생성
// @Description 관리자 모집 공고 생성 API
// @Tags Careers
// @Param req body dto.JobPostingRequest true "요청 Body"
// @Accept json
// @Produce json
// @Router /careers/job-postings [POST]
func (c *CareerController) CreateJobPosting(ctx context.Context, dto *dto.JobPostingRequest) error {
	return c.service.CreateJobPosting(ctx, dto)
}

// @Summary (사용자) 모집 공고 상세 조회
// @Description 사용자 모집 공고 상세 조회 API
// @Tags Careers
// @Param id path int64 true "요청 공고"
// @Produce json
// @Router /careers/job-postings/{id} [GET]
func (c *CareerController) GetJobPosting(ctx context.Context, id path.Int) (httpx.Response[dto.JobPostingResponse], error) {
	result, err := c.service.GetJobPosting(ctx, id.Value)

	if err != nil {
		return httpx.Response[dto.JobPostingResponse]{}, err
	}

	return httpx.Response[dto.JobPostingResponse]{
		Body: result,
	}, nil
}
