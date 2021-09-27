##简介
    基于TarsBenchmark框架的压力测试工具tb进行接口压力测试。
    
##使用说明
```
    1、设置go test run 测试函数（生成http的post body data文件）；
    2、设置tb -H 参数，填token就可以；
    3、执行benchmark.sh。
```

##tb的常用option说明
```
    -c         连接数量
    -s(可选)    最大速率限制，为空或为0表示自动适配服务最佳速率      
    -n         最大压测进程限制，默认CPU核心数
    -I(可选)    压测持续时间（单位秒），默认1H 
    -p         接口通信协议（http|tars)
    -P         网络传输端口
    -u         url地址
    -F         post body data文件 
    -H         http header
```
详细使用参考[TarsBenchmark](https://github.com/TarsCloud/TarsBenchmark/blob/master/README.zh.md)

    
   
