package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ppopgi-pang/ppopgipang-spine/auth/dto"
	userEntity "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
	userService "github.com/ppopgi-pang/ppopgipang-spine/users/service"
)

type KakaoOAuthClient struct {
	http        *http.Client
	clientID    string
	redirectURI string
	userService *userService.UserService
}

func NewKakaoOAuthClient(userService *userService.UserService) *KakaoOAuthClient {
	return &KakaoOAuthClient{
		http:        &http.Client{Timeout: 5 * time.Second},
		clientID:    os.Getenv("KAKAO_CLIENT_ID"),
		redirectURI: os.Getenv("KAKAO_REDIRECT_URI"),
		userService: userService,
	}
}

func (c *KakaoOAuthClient) GetUser(token string) (*dto.KakaoUserResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://kapi.kakao.com/v2/user/me",
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kakao oauth failed: status=%d", resp.StatusCode)
	}

	var body dto.KakaoUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &dto.KakaoUserResponse{
		ID: body.ID,
		KakaoAccount: dto.KakaoAccount{
			Email: body.KakaoAccount.Email,
			Profile: dto.KakaoProfile{
				Nickname: body.KakaoAccount.Profile.Nickname,
			},
		},
	}, nil
}

func (c *KakaoOAuthClient) ExchangeCodeForToken(code string) (string, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("client_id", c.clientID)
	form.Set("redirect_uri", c.redirectURI)
	form.Set("code", code)

	req, err := http.NewRequest(
		"POST",
		"https://kauth.kakao.com/oauth/token",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("카카오 토큰 에러: %s", body)
	}

	var tokenResp dto.KakaoTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("카카오 액세스 토큰이 없습니다")
	}

	return tokenResp.AccessToken, nil
}

func (c *KakaoOAuthClient) MapOrCreateUser(kakaoUser *dto.KakaoUserResponse) (*userEntity.User, error) {

	user, err := c.userService.FindByEmail(kakaoUser.KakaoAccount.Email)
	if err == nil {
		return &user, nil
	}

	newUser := &userEntity.User{
		Email:    kakaoUser.KakaoAccount.Email,
		Nickname: kakaoUser.KakaoAccount.Profile.Nickname,
	}
	if err := c.userService.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
