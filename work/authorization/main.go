package main

// ======================= 授权服务器 =======================
//+--------+                               +---------------+
//|        |--(1)- Authorization Request ->|   Resource    |
//|        |                               |     Owner     |
//|        |<-(2)-- Authorization Grant ---|               |
//|        |                               +---------------+
//|        |
//|        |                               +---------------+
//|        |--(3)-- Authorization Grant -->| Authorization |
//| Client |                               |     Server    |
//|        |<-(4)----- Access Token -------|               |
//|        |                               +---------------+
//|        |
//|        |                               +---------------+
//|        |--(5)----- Access Token ------>|    Resource   |
//|        |                               |     Server    |
//|        |<-(6)--- Protected Resource ---|               |
//+--------+                               +---------------+

import (
	"authorization/config"
	"authorization/endpoint"
	"authorization/model"
	"authorization/service"
	"authorization/transport"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// 刷新令牌如何生成？ 类似token

func main() {
	var (
		servicePort = flag.Int("service.port", 10098, "service port")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var tokenService service.TokenService
	var tokenGranter service.TokenGranter
	var tokenEnhancer service.TokenEnhancer
	var tokenStore service.TokenStore
	var userDetailsService service.UserDetailsService
	var clientDetailsService service.ClientDetailsService
	var srv service.Service

	tokenEnhancer = service.NewJwtTokenEnhancer("secret")
	tokenStore = service.NewJwtTokenStore(tokenEnhancer.(*service.JwtTokenEnhancer))
	tokenService = service.NewTokenService(tokenStore, tokenEnhancer)

	userDetailsService = service.NewInMemoryUserDetailsService([]*model.UserDetails{
		{
			Username:    "aoho",
			Password:    "123456",
			UserId:      1,
			Authorities: []string{"Simple"},
		},
		{
			Username:    "admin",
			Password:    "123456",
			UserId:      1,
			Authorities: []string{"Admin"},
		},
		{
			Username:    "aoho2",
			Password:    "123456",
			UserId:      1,
			Authorities: []string{"Admin"},
		}})
	clientDetailsService = service.NewInMemoryClientDetailService([]*model.ClientDetails{{
		"clientId",
		"clientSecret",
		1800,
		18000,
		"http://127.0.0.1",
		[]string{"password", "refresh_token"},
	}})

	tokenGranter = service.NewComposeTokenGranter(map[string]service.TokenGranter{
		"password":      service.NewUsernamePasswordTokenGranter("password", userDetailsService, tokenService),
		"refresh_token": service.NewRefreshGranter("refresh_token", userDetailsService, tokenService),
	})

	tokenEndpoint := endpoint.MakeTokenEndpoint(tokenGranter, clientDetailsService)
	tokenEndpoint = endpoint.MakeClientAuthorizationMiddleware(config.KitLogger)(tokenEndpoint)
	checkTokenEndpoint := endpoint.MakeCheckTokenEndpoint(tokenService)
	checkTokenEndpoint = endpoint.MakeClientAuthorizationMiddleware(config.KitLogger)(checkTokenEndpoint)

	srv = service.NewCommonService()

	//创建健康检查的Endpoint
	healthEndpoint := endpoint.MakeHealthCheckEndpoint(srv)

	endpts := endpoint.OAuth2Endpoints{
		TokenEndpoint:       tokenEndpoint,
		CheckTokenEndpoint:  checkTokenEndpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, tokenService, clientDetailsService, config.KitLogger)

	go func() {
		config.Logger.Println("Http Server start at port:" + strconv.Itoa(*servicePort))
		handler := r
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	config.Logger.Println(error)
}

/*
curl -X POST \
'http://localhost:10098/oauth/token?grant_type=password' \
-H 'Authorization: Basic Y2xpZW50SWQ6Y2xpZW50U2VjcmV0' \
-H 'Content-Type: multipart/form-data' \
-H 'Host: localhost:10098' \
-F username=aoho \
-F password=123456
*/

/*
curl -X POST \
'http://localhost:10098/oauth/check_token?token=...' \
-H 'Authorization: Basic Y2xpZW50SWQ6Y2xpZW50U2VjcmV0' \
-H 'Host: localhost:10098' */
/*
{"access_token":{"RefreshToken":{"RefreshToken":null,"TokenType":"jwt","TokenValue":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyRGV0YWlscyI6eyJVc2VySWQiOjEsIlVzZXJuYW1lIjoiYW9obyIsIlBhc3N3b3JkIjoiIiwiQXV0aG9yaXRpZXMiOlsiU2ltcGxlIl19LCJDbGllbnREZXRhaWxzIjp7IkNsaWVudElkIjoiY2xpZW50SWQiLCJDbGllbnRTZWNyZXQiOiIiLCJBY2Nlc3NUb2tlblZhbGlkaXR5U2Vjb25kcyI6MTgwMCwiUmVmcmVzaFRva2VuVmFsaWRpdHlTZWNvbmRzIjoxODAwMCwiUmVnaXN0ZXJlZFJlZGlyZWN0VXJpIjoiaHR0cDovLzEyNy4wLjAuMSIsIkF1dGhvcml6ZWRHcmFudFR5cGVzIjpbInBhc3N3b3JkIiwicmVmcmVzaF90b2tlbiJdfSwiUmVmcmVzaFRva2VuIjp7IlJlZnJlc2hUb2tlbiI6bnVsbCwiVG9rZW5UeXBlIjoiIiwiVG9rZW5WYWx1ZSI6IiIsIkV4cGlyZXNUaW1lIjpudWxsfSwiZXhwIjoxNjM4OTY2MjE4LCJpc3MiOiJTeXN0ZW0ifQ.Or3EVjBulWuiQ5KxZ-GSQL9b8XZxgL9VBpvXhrjZ3WY","ExpiresTime":"2021-12-08T20:23:38.254805+08:00"},"TokenType":"jwt","TokenValue":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyRGV0YWlscyI6eyJVc2VySWQiOjEsIlVzZXJuYW1lIjoiYW9obyIsIlBhc3N3b3JkIjoiIiwiQXV0aG9yaXRpZXMiOlsiU2ltcGxlIl19LCJDbGllbnREZXRhaWxzIjp7IkNsaWVudElkIjoiY2xpZW50SWQiLCJDbGllbnRTZWNyZXQiOiIiLCJBY2Nlc3NUb2tlblZhbGlkaXR5U2Vjb25kcyI6MTgwMCwiUmVmcmVzaFRva2VuVmFsaWRpdHlTZWNvbmRzIjoxODAwMCwiUmVnaXN0ZXJlZFJlZGlyZWN0VXJpIjoiaHR0cDovLzEyNy4wLjAuMSIsIkF1dGhvcml6ZWRHcmFudFR5cGVzIjpbInBhc3N3b3JkIiwicmVmcmVzaF90b2tlbiJdfSwiUmVmcmVzaFRva2VuIjp7IlJlZnJlc2hUb2tlbiI6bnVsbCwiVG9rZW5UeXBlIjoiand0IiwiVG9rZW5WYWx1ZSI6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpWYzJWeVJHVjBZV2xzY3lJNmV5SlZjMlZ5U1dRaU9qRXNJbFZ6WlhKdVlXMWxJam9pWVc5b2J5SXNJbEJoYzNOM2IzSmtJam9pSWl3aVFYVjBhRzl5YVhScFpYTWlPbHNpVTJsdGNHeGxJbDE5TENKRGJHbGxiblJFWlhSaGFXeHpJanA3SWtOc2FXVnVkRWxrSWpvaVkyeHBaVzUwU1dRaUxDSkRiR2xsYm5SVFpXTnlaWFFpT2lJaUxDSkJZMk5sYzNOVWIydGxibFpoYkdsa2FYUjVVMlZqYjI1a2N5STZNVGd3TUN3aVVtVm1jbVZ6YUZSdmEyVnVWbUZzYVdScGRIbFRaV052Ym1Seklqb3hPREF3TUN3aVVtVm5hWE4wWlhKbFpGSmxaR2x5WldOMFZYSnBJam9pYUhSMGNEb3ZMekV5Tnk0d0xqQXVNU0lzSWtGMWRHaHZjbWw2WldSSGNtRnVkRlI1Y0dWeklqcGJJbkJoYzNOM2IzSmtJaXdpY21WbWNtVnphRjkwYjJ0bGJpSmRmU3dpVW1WbWNtVnphRlJ2YTJWdUlqcDdJbEpsWm5KbGMyaFViMnRsYmlJNmJuVnNiQ3dpVkc5clpXNVVlWEJsSWpvaUlpd2lWRzlyWlc1V1lXeDFaU0k2SWlJc0lrVjRjR2x5WlhOVWFXMWxJanB1ZFd4c2ZTd2laWGh3SWpveE5qTTRPVFkyTWpFNExDSnBjM01pT2lKVGVYTjBaVzBpZlEuT3IzRVZqQnVsV3VpUTVLeFotR1NRTDliOFhaeGdMOVZCcHZYaHJqWjNXWSIsIkV4cGlyZXNUaW1lIjoiMjAyMS0xMi0wOFQyMDoyMzozOC4yNTQ4MDUrMDg6MDAifSwiZXhwIjoxNjM4OTUwMDE4LCJpc3MiOiJTeXN0ZW0ifQ.fXOf4CiZc_7Kb3qCu1SOUV9KDPHY0LytwhJzYLooPCg","ExpiresTime":"2021-12-08T15:53:38.255093+08:00"},"error":""}
*/

// 微信授权给拉勾教育
// https://open.weixin.qq.com/connect/qrconnect?appid=wx9d8d3686b76baff8&redirect_uri=https%3A%2F%2Fpassport.lagou.com%2Foauth20%2Fcallback_weixinEduProvider.html&response_type=code&scope=snsapi_login#wechat_redirect

// https://passport.lagou.com/oauth20/callback_weixinEduProvider.html?code=071YpVkl2NTJg84pbmol2Bgyko1YpVka&state=

// heartbeat
// https://kaiwu.lagou.com/web/token/heartbeat
