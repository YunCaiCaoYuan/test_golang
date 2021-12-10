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
