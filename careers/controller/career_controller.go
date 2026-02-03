package controller

import (
	"context"

	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/path"
	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/careers/service"
	commonDto "github.com/ppopgi-pang/ppopgipang-spine/commons/dto"
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
func (c *CareerController) CreateJobPosting(ctx context.Context, dto *dto.JobPostingRequest, spineCtx spine.Ctx) error {
	roleAny, ok := spineCtx.Get("auth.role")
	roleString := roleAny.(string)

	if !ok {
		panic("컨텍스트에서 Role을 가져오는데 실패했습니다.")
	}

	if roleString != "admin" {
		httperr.Unauthorized("권한이 Admin이 아닙니다.")
	}

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

// @Summary (관리자) 모집 공고 수정
// @Description 관리자 모집 공고 수정 API
// @Tags Careers
// @Param id path int64 true "수정할 공고의 ID"
// @Param req body dto.JobPostingModifyRequest true "요청 Body"
// @Accept json
// @Produce json
// @Router /careers/job-postings/{id} [PATCH]
func (c *CareerController) UpdateJobPosting(ctx context.Context, id path.Int, dto *dto.JobPostingModifyRequest, spineCtx spine.Ctx) error {
	roleAny, ok := spineCtx.Get("auth.role")
	roleString := roleAny.(string)

	if !ok {
		panic("컨텍스트에서 Role을 가져오는데 실패했습니다.")
	}

	if roleString != "admin" {
		httperr.Unauthorized("권한이 Admin이 아닙니다.")
	}

	return c.service.UpdateJobPosting(ctx, id.Value, dto)
}

// @Summary (관리자) 모집 공고 삭제
// @Description 관리자 모집 공고 삭제 API
// @Tags Careers
// @Param id path int64 true "삭제할 공고의 ID"
// @Router /careers/job-postings/{id} [DELETE]
func (c *CareerController) DeleteJobPosting(ctx context.Context, id path.Int, spineCtx spine.Ctx) error {
	roleAny, ok := spineCtx.Get("auth.role")
	roleString := roleAny.(string)

	if !ok {
		panic("컨텍스트에서 Role을 가져오는데 실패했습니다.")
	}

	if roleString != "admin" {
		httperr.Unauthorized("권한이 Admin이 아닙니다.")
	}
	return c.service.DeleteJobPosting(ctx, id.Value)
}

// @Summary (사용자) 지원서 제출
// @Description 사용자 지원서 제출 API
// @Tags Careers
// @Param req body dto.CreateApplicationRequest true "지원서 요청 Body"
// @Router /careers/applications [POST]
func (c *CareerController) CreateApplication(ctx context.Context, dto *dto.CreateApplicationRequest) (httpx.Response[commonDto.CommonResponse], error) {
	id, err := c.service.CreateApplication(ctx, dto)
	if err != nil {
		return httpx.Response[commonDto.CommonResponse]{}, nil
	}

	return httpx.Response[commonDto.CommonResponse]{
		Body: commonDto.CommonResponse{
			ID:      id,
			Message: "이력서가 정상적으로 제출되었습니다.",
		},
	}, nil
}

// @Summary (관리자) 지원서 목록 조회
// @Description 관리자 지원서 목록 조회 API
// @Tags Careers
// @Param jobPostingId query int64 false "찾을 공고 ID"
// @Param status query string true "지원서 상태"
// @Param page query int true "페이지 번호"
// @Param size query int true "한번에 가져올 목록의 개수"
// @Success 200 {object} dto.ApplicationListResponse
// @Router /careers/applications [GET]
func (c *CareerController) GetApplications(ctx context.Context, query query.Values, meta query.Pagination, spineCtx spine.Ctx) (httpx.Response[dto.ApplicationListResponse], error) {
	jobPostingId := query.Int("jobPostingId", 0)
	status := query.String("status")

	result, err := c.service.GetApplications(ctx, jobPostingId, status, meta.Page, meta.Size)
	if err != nil {
		return httpx.Response[dto.ApplicationListResponse]{}, err
	}

	return httpx.Response[dto.ApplicationListResponse]{
		Body: result,
	}, nil
}

// @Summary (관리자) 지원서 상세 조회
// @Description 관리자 지원서 상세 조회 API
// @Tags Careers
// @Param id path int64 true "지원서 ID"
// @Success 200 {object} dto.ApplicationResponse
// @Router /careers/applications/{id} [GET]
func (c *CareerController) GetApplication(ctx context.Context, id path.Int, spineCtx spine.Ctx) (httpx.Response[dto.ApplicationResponse], error) {
	roleAny, ok := spineCtx.Get("auth.role")
	roleString := roleAny.(string)

	if !ok {
		panic("컨텍스트에서 Role을 가져오는데 실패했습니다.")
	}

	if roleString != "admin" {
		httperr.Unauthorized("권한이 Admin이 아닙니다.")
	}

	result, err := c.service.GetApplication(ctx, id.Value)

	if err != nil {
		return httpx.Response[dto.ApplicationResponse]{}, err
	}

	return httpx.Response[dto.ApplicationResponse]{
		Body: result,
	}, nil
}
