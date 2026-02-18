package controller

import (
	"context"
	"strconv"

	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/path"
	"github.com/NARUBROWN/spine/pkg/query"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/utils"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/dto"
	"github.com/ppopgi-pang/ppopgipang-spine/stores/service"
)

type StoreController struct {
	service *service.StoreService
}

func NewStoreController(storesService *service.StoreService) *StoreController {
	return &StoreController{service: storesService}
}

// @Summary (공통) 주변 매장 조회
// @Description 위도/경도 기준으로 주변 매장 목록을 페이지네이션과 키워드로 조회합니다.
// @Tags Stores
// @Param latitude query number true "기준 위도"
// @Param longitude query number true "기준 경도"
// @Param radius query int false "검색 반경(미터)"
// @Param keyword query string false "검색 키워드"
// @Param filter query string false "필터(all|scrapped|popular|recent_cert)"
// @Param page query int true "요청 페이지 번호"
// @Param size query int true "한번에 받을 페이지의 사이즈"
// @Success 200 {object} dto.FindNearByDto
// @Router /stores/nearby [GET]
func (s *StoreController) FindNearByStores(ctx context.Context, query query.Values, meta query.Pagination, spineCtx spine.Ctx) (httpx.Response[dto.FindNearByDto], error) {
	latitudeRaw := query.Get("latitude")
	longitudeRaw := query.Get("longitude")

	latitude, _ := strconv.ParseFloat(latitudeRaw, 64)
	longitude, _ := strconv.ParseFloat(longitudeRaw, 64)

	radius := query.Int("radius", 0)
	keyword := query.String("keyword")
	filter := query.String("filter")

	userID, err := utils.GetAuthUserID(spineCtx)
	if err != nil {
		return httpx.Response[dto.FindNearByDto]{}, err
	}

	result, err := s.service.FindNearByStores(ctx, latitude, longitude, radius, meta.Page, meta.Size, keyword, filter, userID)
	if err != nil {
		return httpx.Response[dto.FindNearByDto]{}, err
	}
	return httpx.Response[dto.FindNearByDto]{
		Body: result,
	}, nil
}

// @Summary (공통) 영역 내 매장 조회
// @Description 지도 바운즈(북/남/동/서) 좌표로 매장 목록을 키워드로 조회합니다.
// @Tags Stores
// @Param north query number true "북쪽 위도"
// @Param south query number true "남쪽 위도"
// @Param east query number true "동쪽 경도"
// @Param west query number true "서쪽 경도"
// @Param keyword query string false "검색 키워드"
// @Param filter query string false "필터(all|scrapped|popular|recent_cert)"
// @Success 200 {object} dto.FindInBoundsDto
// @Router /stores/in-bounds [GET]
func (s *StoreController) FindStoresInBounds(ctx context.Context, query query.Values, spineCtx spine.Ctx) (httpx.Response[dto.FindInBoundsDto], error) {
	northRaw := query.Get("north")
	southRaw := query.Get("south")
	eastRaw := query.Get("east")
	westRaw := query.Get("west")

	north, _ := strconv.ParseFloat(northRaw, 64)
	south, _ := strconv.ParseFloat(southRaw, 64)
	east, _ := strconv.ParseFloat(eastRaw, 64)
	west, _ := strconv.ParseFloat(westRaw, 64)

	keyword := query.String("keyword")
	filter := query.String("filter")

	userID, err := utils.GetAuthUserID(spineCtx)
	if err != nil {
		return httpx.Response[dto.FindInBoundsDto]{}, err
	}

	result, err := s.service.FindStoresInBounds(ctx, north, south, east, west, keyword, filter, userID)
	if err != nil {
		return httpx.Response[dto.FindInBoundsDto]{}, err
	}
	return httpx.Response[dto.FindInBoundsDto]{
		Body: result,
	}, nil
}

// @Summary (사용자) 가게 검색
// @Description 사용자 가게 검색 API입니다.
// @Tags Stores
// @Param latitude query float64 false "위도"
// @Param longitude query float64 false "경도"
// @Param keyword query string true "검색 키워드"
// @Param page query int true "요청 페이지 번호"
// @Param size query int true "한번에 받을 페이지의 사이즈"
// @Success 200 {object} dto.StoreSearchResponse
// @Router /stores/search [GET]
func (s *StoreController) SearchStore(ctx context.Context, query query.Values, meta query.Pagination) (httpx.Response[dto.StoreSearchResponse], error) {
	latitudeRaw := query.Get("latitude")
	longitudeRaw := query.Get("longitude")

	latitude, _ := strconv.ParseFloat(latitudeRaw, 64)
	longitude, _ := strconv.ParseFloat(longitudeRaw, 64)

	keyword := query.String("keyword")

	result, err := s.service.SearchStore(ctx, keyword, latitude, longitude, meta.Page, meta.Size)

	if err != nil {
		return httpx.Response[dto.StoreSearchResponse]{}, err
	}

	return httpx.Response[dto.StoreSearchResponse]{
		Body: result,
	}, nil
}

// @Summary (사용자) 가게 기본 정보
// @Description 사용자 가게 기본 정보 API입니다.
// @Tags Stores
// @Param storeId path int64 true "스토어 ID"
// @Success 200 {object} dto.StoreSummaryResponse
// @Router /stores/summary/{storeId} [GET]
func (s *StoreController) FindByStoreSummaryId(ctx context.Context, storeId path.Int) (httpx.Response[dto.StoreSummaryResponse], error) {
	result, err := s.service.FindByStoreSummaryId(ctx, storeId.Value)
	if err != nil {
		return httpx.Response[dto.StoreSummaryResponse]{}, err
	}
	return httpx.Response[dto.StoreSummaryResponse]{
		Body: result,
	}, nil
}

// @Summary (사용자) 가게 상세 페이지 가게정보 탭
// @Description 사용자 가게 정보 API입니다.
// @Tags Stores
// @Param storeId path int64 true "스토어 ID"
// @Success 200 {object} dto.StoreDetailResponse
// @Router /stores/details/{storeId} [GET]
func (s *StoreController) FindByStoreDetailId(ctx context.Context, storeId path.Int, spineCtx spine.Ctx) (httpx.Response[dto.StoreDetailResponse], error) {
	userId, _ := utils.GetAuthUserID(spineCtx)
	result, err := s.service.FindByStoreDetailId(ctx, storeId.Value, userId)
	if err != nil {
		return httpx.Response[dto.StoreDetailResponse]{}, err
	}
	return httpx.Response[dto.StoreDetailResponse]{
		Body: result,
	}, nil
}

// @Summary (사용자) 가게 상세 페이지 방문내역 탭
// @Description 사용자 방문 내역 API입니다.
// @Tags Stores
// @Param storeId path int64 true "스토어 ID"
// @Success 200 {object} dto.VisitHistoryResponse
// @Router /stores/visits/{storeId} [GET]
func (s *StoreController) GetStoreStatById(ctx context.Context, storeId path.Int, spineCtx spine.Ctx) (httpx.Response[dto.VisitHistoryResponse], error) {
	userId, _ := utils.GetAuthUserID(spineCtx)
	result, err := s.service.GetStoreStatById(ctx, storeId.Value, userId)
	if err != nil {
		return httpx.Response[dto.VisitHistoryResponse]{}, err
	}
	return httpx.Response[dto.VisitHistoryResponse]{
		Body: result,
	}, nil
}

// @Summary (사용자) 가게 상세 페이지 리뷰 탭
// @Description 가세 상세 페이지 리뷰 API입니다.
// @Tags Stores
// @Param storeId path int64 true "스토어 ID"
// @Success 200 {object} dto.StoreReviewResponse
// @Router /stores/reviews/{storeId} [GET]
func (s *StoreController) GetStoreReviewsById(ctx context.Context, storeId path.Int) (httpx.Response[dto.StoreReviewResponse], error) {
	result, err := s.service.GetStoreReviewsById(ctx, storeId.Value)
	if err != nil {
		return httpx.Response[dto.StoreReviewResponse]{}, err
	}
	return httpx.Response[dto.StoreReviewResponse]{
		Body: result,
	}, nil
}

// @Summary (사용자) 가장 가까운 가게 1개 조회
// @Description GPS 기반으로 가장 가까운 가게 1개를 반환합니다. 반경 내 가게가 없으면 null 반환.
// @Tags Stores
// @Param latitude query float64 true "현재 위도"
// @Param longitude query float64 true "현재 경도"
// @Param radius query int64 false "검색 반경 (미터, 기본: 100m)"
// @Success 200 {object} dto.FindNearestStoreResponse
// @Router /stores/nearest [GET]
func (s *StoreController) FindNearestStore(ctx context.Context, query query.Values) (httpx.Response[dto.FindNearestStoreResponse], error) {
	latitudeRaw := query.Get("latitude")
	longitudeRaw := query.Get("longitude")

	latitude, _ := strconv.ParseFloat(latitudeRaw, 64)
	longitude, _ := strconv.ParseFloat(longitudeRaw, 64)

	radius := query.Int("radius", 100)

	result, err := s.service.FindNearestStore(ctx, latitude, longitude, radius)
	if err != nil {
		return httpx.Response[dto.FindNearestStoreResponse]{}, err
	}

	return httpx.Response[dto.FindNearestStoreResponse]{
		Body: result,
	}, nil
}
