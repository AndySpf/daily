Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序，最近使用了一下，非常好用，这里备忘一下。
---
Mac 下编译 Linux 和 Windows 64位可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
---
Linux 下编译 Mac 和 Windows 64位可执行程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
---
Windows 下编译 Mac 和 Linux 64位可执行程序
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
---
GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
GOARCH：目标平台的体系架构（386、amd64、arm）
交叉编译不支持 CGO 所以要禁用它

上面的命令编译 64 位可执行程序，你当然应该也会使用 386 编译 32 位可执行程序
很多博客都提到要先增加对其它平台的支持，但是我跳过那一步，上面所列的命令也都能成功，且得到我想要的结果，可见那一步应该是非必须的，或是我所使用的 Go 版本已默认支持所有平台。

golang 支持的GOOS/GOARCH组合如下
 $GOOS		$GOARCH
    android     arm
    darwin      386
    darwin      amd64
    darwin      arm
    darwin      arm64
    dragonfly   amd64
    freebsd     386
    freebsd     amd64
    freebsd     arm
    linux       386
    linux       amd64
    linux       arm
    linux       arm64
    linux       ppc64
    linux       ppc64le
    linux       mips
    linux       mipsle
    linux       mips64
    linux       mips64le
    netbsd      386
    netbsd      amd64
    netbsd      arm
    openbsd     386
    openbsd     amd64
    openbsd     arm
    plan9       386
    plan9       amd64
    solaris     amd64
    windows     386
    windows     amd64