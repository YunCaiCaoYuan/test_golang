package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"resource_server/model"
	"time"
)

var (
	ErrExpiredToken = errors.New("token is expired")
)

type ResourceServerTokenService interface {
	// 根据访问令牌获取对应的用户信息和客户端信息
	GetOAuth2DetailsByAccessToken(tokenValue string) (*model.OAuth2Details, error)
}

type DefaultTokenService struct {
	tokenStore    TokenStore
	tokenEnhancer TokenEnhancer
}

func NewTokenService(tokenStore TokenStore, tokenEnhancer TokenEnhancer) ResourceServerTokenService {
	return &DefaultTokenService{
		tokenStore:    tokenStore,
		tokenEnhancer: tokenEnhancer,
	}
}

func (tokenService *DefaultTokenService) GetOAuth2DetailsByAccessToken(tokenValue string) (*model.OAuth2Details, error) {
	accessToken, err := tokenService.tokenStore.ReadAccessToken(tokenValue)
	if err != nil {
		return nil, err
	}
	if accessToken.IsExpired() {
		return nil, ErrExpiredToken
	}
	return tokenService.tokenStore.ReadOAuth2Details(tokenValue)
}

type TokenStore interface {

	// 存储访问令牌
	StoreAccessToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) error
	// 根据令牌值获取访问令牌结构体
	ReadAccessToken(tokenValue string) (*model.OAuth2Token, error)
	// 根据令牌值获取令牌对应的客户端和用户信息
	ReadOAuth2Details(tokenValue string) (*model.OAuth2Details, error)
	// 根据客户端信息和用户信息获取访问令牌
	GetAccessToken(oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error)
	// 移除存储的访问令牌
	RemoveAccessToken(tokenValue string) error
	// 存储刷新令牌
	StoreRefreshToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) error
	// 移除存储的刷新令牌
	RemoveRefreshToken(oauth2Token string) error
	// 根据令牌值获取刷新令牌
	ReadRefreshToken(tokenValue string) (*model.OAuth2Token, error)
	// 根据令牌值获取刷新令牌对应的客户端和用户信息
	ReadOAuth2DetailsForRefreshToken(tokenValue string) (*model.OAuth2Details, error)
}

func NewJwtTokenStore(jwtTokenEnhancer *JwtTokenEnhancer) TokenStore {
	return &JwtTokenStore{
		jwtTokenEnhancer: jwtTokenEnhancer,
	}

}

type JwtTokenStore struct {
	jwtTokenEnhancer *JwtTokenEnhancer
}

func (tokenStore *JwtTokenStore) StoreAccessToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) error {
	return nil
}

func (tokenStore *JwtTokenStore) ReadAccessToken(tokenValue string) (*model.OAuth2Token, error) {
	oauth2Token, _, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Token, err

}

// 根据令牌值获取令牌对应的客户端和用户信息
func (tokenStore *JwtTokenStore) ReadOAuth2Details(tokenValue string) (*model.OAuth2Details, error) {
	_, oauth2Details, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Details, err

}

// 根据客户端信息和用户信息获取访问令牌
func (tokenStore *JwtTokenStore) GetAccessToken(oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {
	return nil, nil
}

// 移除存储的访问令牌
func (tokenStore *JwtTokenStore) RemoveAccessToken(tokenValue string) error {
	return nil
}

// 存储刷新令牌
func (tokenStore *JwtTokenStore) StoreRefreshToken(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) error {
	return nil
}

// 移除存储的刷新令牌
func (tokenStore *JwtTokenStore) RemoveRefreshToken(oauth2Token string) error {
	return nil
}

// 根据令牌值获取刷新令牌
func (tokenStore *JwtTokenStore) ReadRefreshToken(tokenValue string) (*model.OAuth2Token, error) {
	oauth2Token, _, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Token, err
}

// 根据令牌值获取刷新令牌对应的客户端和用户信息
func (tokenStore *JwtTokenStore) ReadOAuth2DetailsForRefreshToken(tokenValue string) (*model.OAuth2Details, error) {
	_, oauth2Details, err := tokenStore.jwtTokenEnhancer.Extract(tokenValue)
	return oauth2Details, err
}

type TokenEnhancer interface {
	// 组装 Token 信息
	Enhance(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error)
	// 从 Token 中还原信息
	Extract(tokenValue string) (*model.OAuth2Token, *model.OAuth2Details, error)
}

type OAuth2TokenCustomClaims struct {
	UserDetails   model.UserDetails
	ClientDetails model.ClientDetails
	RefreshToken  model.OAuth2Token
	jwt.StandardClaims
}

type JwtTokenEnhancer struct {
	secretKey []byte
}

func NewJwtTokenEnhancer(secretKey string) TokenEnhancer {
	return &JwtTokenEnhancer{
		secretKey: []byte(secretKey),
	}

}

func (enhancer *JwtTokenEnhancer) Enhance(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {
	return enhancer.sign(oauth2Token, oauth2Details)
}

func (enhancer *JwtTokenEnhancer) sign(oauth2Token *model.OAuth2Token, oauth2Details *model.OAuth2Details) (*model.OAuth2Token, error) {

	expireTime := oauth2Token.ExpiresTime
	clientDetails := oauth2Details.Client
	userDetails := oauth2Details.User
	clientDetails.ClientSecret = ""
	userDetails.Password = ""

	claims := OAuth2TokenCustomClaims{
		UserDetails:   userDetails,
		ClientDetails: clientDetails,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "System",
		},
	}

	if oauth2Token.RefreshToken != nil {
		claims.RefreshToken = *oauth2Token.RefreshToken
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err := token.SignedString(enhancer.secretKey)

	if err != nil {
		return nil, err
	}

	oauth2Token.TokenValue = tokenValue
	oauth2Token.TokenType = "jwt"
	return oauth2Token, nil
}

func (enhancer *JwtTokenEnhancer) Extract(tokenValue string) (*model.OAuth2Token, *model.OAuth2Details, error) {

	token, err := jwt.ParseWithClaims(tokenValue, &OAuth2TokenCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return enhancer.secretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims := token.Claims.(*OAuth2TokenCustomClaims)
	expiresTime := time.Unix(claims.ExpiresAt, 0)

	return &model.OAuth2Token{
			RefreshToken: &claims.RefreshToken,
			TokenValue:   tokenValue,
			ExpiresTime:  &expiresTime,
		}, &model.OAuth2Details{
			User:   claims.UserDetails,
			Client: claims.ClientDetails,
		}, nil
}
