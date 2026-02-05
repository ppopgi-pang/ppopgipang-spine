package utils

import (
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/spine"
)

func GetAuthUserID(spineCtx spine.Ctx) (*int64, error) {
	userIDAny, ok := spineCtx.Get("auth.userId")
	if !ok {
		return nil, nil
	}

	switch value := userIDAny.(type) {
	case int64:
		return &value, nil
	case string:
		if value == "" {
			return nil, nil
		}
		return nil, httperr.Unauthorized("유효하지 않은 사용자 정보입니다.")
	default:
		return nil, httperr.Unauthorized("유효하지 않은 사용자 정보입니다.")
	}
}
