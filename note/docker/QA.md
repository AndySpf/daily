***Q:ctr images展示不出来已有镜像，但是docker images可以***   
A:镜像以及容器的下载展示是由dockerd控制的, 如读取镜像列表是dockerd启动时读取：/var/lib/docker/image/{{graphDriver}}(overlay2)/imagedb/content等文件系统(daemon/daemon.go:943 NewDaemon)读取到放在内存的，和containerd无关   
  即containerd本身可以执行拉取镜像等操作，但dockerd没使用，containerd的这部分功能。在containerd没有运行的情况下dockerd也可以正常拉取镜像，查看容器等。
  containerd是使用的bolt键值数据库进行镜像容器信息的管理的，且没有使用联合文件系统，用的是快照（[containerd的graph driver说明](https://blog.mobyproject.org/where-are-containerds-graph-drivers-145fc9b7255)）。可以使用ctr image pull docker.io/library/redis:alpine拉取镜像，然后再ctr images ls可以查看得到
---
Q:通过docker启动一个容器和通过containerd启动一个容器有什么区别？
A:
