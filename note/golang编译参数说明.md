go build -gcflags='-m=2' 可以通过查看编译器的报告来了解是否发生了内存逃逸,-m共有四个级别，超过2的内容太多一般用-m=2  
go run -gcflags '-m -l' -l为关闭inline  
go build -ldflags "-X 'main.VERSION=$(VERSION)'" 可向main文件中注入参数值  
go build -v 编译时显示包名  
go build -p n 开启并发编译，n默认值为cpu逻辑核数    
go build -a 强制重新构建  
go build -n 打印编译时会用到的命令，但不真正执行    
go build -x 打印编译时会用到的命令
go build -race 开启静态检测，如果有竞态监，则在运行程序时会有WARN 
go tool compile -S main.go 获取汇编代码