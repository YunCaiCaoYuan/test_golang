package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"resource_server/config"
	"resource_server/endpoint"
	"resource_server/service"
	"resource_server/transport"
	"strconv"
	"syscall"
)

func main() {

	var (
		servicePort = flag.Int("service.port", 10099, "service port")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var tokenService service.ResourceServerTokenService
	var tokenEnhancer service.TokenEnhancer
	var tokenStore service.TokenStore
	var srv service.Service

	tokenEnhancer = service.NewJwtTokenEnhancer("secret")
	tokenStore = service.NewJwtTokenStore(tokenEnhancer.(*service.JwtTokenEnhancer))
	tokenService = service.NewTokenService(tokenStore, tokenEnhancer)

	srv = service.NewCommonService()

	indexEndpoint := endpoint.MakeIndexEndpoint(srv)
	sampleEndpoint := endpoint.MakeSampleEndpoint(srv)
	sampleEndpoint = endpoint.MakeOAuth2AuthorizationMiddleware(config.KitLogger)(sampleEndpoint)
	adminEndpoint := endpoint.MakeAdminEndpoint(srv)
	adminEndpoint = endpoint.MakeOAuth2AuthorizationMiddleware(config.KitLogger)(adminEndpoint)
	adminEndpoint = endpoint.MakeAuthorityAuthorizationMiddleware("Admin", config.KitLogger)(adminEndpoint)

	//创建健康检查的Endpoint
	healthEndpoint := endpoint.MakeHealthCheckEndpoint(srv)

	endpts := endpoint.OAuth2Endpoints{
		HealthCheckEndpoint: healthEndpoint,
		IndexEndpoint:       indexEndpoint,
		SampleEndpoint:      sampleEndpoint,
		AdminEndpoint:       adminEndpoint,
	}

	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, tokenService, config.KitLogger)

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
curl -X GET \
http://localhost:10099/sample \
-H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyRGV0YWlscyI6eyJVc2VySWQiOjEsIlVzZXJuYW1lIjoiYW9obyIsIlBhc3N3b3JkIjoiIiwiQXV0aG9yaXRpZXMiOlsiU2ltcGxlIl19LCJDbGllbnREZXRhaWxzIjp7IkNsaWVudElkIjoiY2xpZW50SWQiLCJDbGllbnRTZWNyZXQiOiIiLCJBY2Nlc3NUb2tlblZhbGlkaXR5U2Vjb25kcyI6MTgwMCwiUmVmcmVzaFRva2VuVmFsaWRpdHlTZWNvbmRzIjoxODAwMCwiUmVnaXN0ZXJlZFJlZGlyZWN0VXJpIjoiaHR0cDovLzEyNy4wLjAuMSIsIkF1dGhvcml6ZWRHcmFudFR5cGVzIjpbInBhc3N3b3JkIiwicmVmcmVzaF90b2tlbiJdfSwiUmVmcmVzaFRva2VuIjp7IlJlZnJlc2hUb2tlbiI6bnVsbCwiVG9rZW5UeXBlIjoiand0IiwiVG9rZW5WYWx1ZSI6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpWYzJWeVJHVjBZV2xzY3lJNmV5SlZjMlZ5U1dRaU9qRXNJbFZ6WlhKdVlXMWxJam9pWVc5b2J5SXNJbEJoYzNOM2IzSmtJam9pSWl3aVFYVjBhRzl5YVhScFpYTWlPbHNpVTJsdGNHeGxJbDE5TENKRGJHbGxiblJFWlhSaGFXeHpJanA3SWtOc2FXVnVkRWxrSWpvaVkyeHBaVzUwU1dRaUxDSkRiR2xsYm5SVFpXTnlaWFFpT2lJaUxDSkJZMk5sYzNOVWIydGxibFpoYkdsa2FYUjVVMlZqYjI1a2N5STZNVGd3TUN3aVVtVm1jbVZ6YUZSdmEyVnVWbUZzYVdScGRIbFRaV052Ym1Seklqb3hPREF3TUN3aVVtVm5hWE4wWlhKbFpGSmxaR2x5WldOMFZYSnBJam9pYUhSMGNEb3ZMekV5Tnk0d0xqQXVNU0lzSWtGMWRHaHZjbWw2WldSSGNtRnVkRlI1Y0dWeklqcGJJbkJoYzNOM2IzSmtJaXdpY21WbWNtVnphRjkwYjJ0bGJpSmRmU3dpVW1WbWNtVnphRlJ2YTJWdUlqcDdJbEpsWm5KbGMyaFViMnRsYmlJNmJuVnNiQ3dpVkc5clpXNVVlWEJsSWpvaUlpd2lWRzlyWlc1V1lXeDFaU0k2SWlJc0lrVjRjR2x5WlhOVWFXMWxJanB1ZFd4c2ZTd2laWGh3SWpveE5qTTVNek0zT1RjeExDSnBjM01pT2lKVGVYTjBaVzBpZlEuaXVkUFNWYWxMU3F6M1puOElraFdMbFRPaDVBM2hfS3pOQnBYMzRwNkJERSIsIkV4cGlyZXNUaW1lIjoiMjAyMS0xMi0xM1QwMzozOTozMS41NzY5ODIxKzA4OjAwIn0sImV4cCI6MTYzOTMyMTc3MSwiaXNzIjoiU3lzdGVtIn0.Ca_ARwjGLnrqPuC6AUf045fFKDXeKOWKRHNkzVcj14s' \
-H 'Host: localhost:10099'
*/
