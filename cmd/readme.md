
# cmd目录

存放结构：
```text
~/cmd/app_name/main.go
~/cmd/appctl/main.go
...
```

`/cmd/app_name/`目录应该只包含main.go，这是app的启动入口，它可以import`/pkg`和`/internal`  
`/cmd/appctl/`目录一般包含main.go在内的多个go文件，它的使用方式是先`go install`，然后运行可执行文件，用作app的辅助工具，
比如Migrate数据表，打开/关闭某个开关

参考[prometheus/cmd][1]

[1]: https://github.com/prometheus/prometheus/tree/master/cmd