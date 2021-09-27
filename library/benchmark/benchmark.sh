## 生成data文件
go test -v  --tags "local beta" ../../yes/channel/handler/ext_test.go -run TestChannelExtObj_FetchHomeChannels

## 执行tb
./tb -c 10 -n 2 -I 60 -p http -D 47.106.118.73 -P 80 -u https://beta.2tianxin.com/proxymsg -F ../../yes/channel/handler/bin -H "X-Token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzIzODkwOTgsImlzcyI6InllcyIsImFpZCI6OCwidWlkIjoxMDMxMDEyLCJvYmYiOiIxZjNlYWU4MTM0MzFmOWYifQ.wNRGMs3nReGfP7oy_vIx39IJ8QX5SjiedxIF_RvSVhI"
