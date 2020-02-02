Q:init()函数有什么细节？   
A:1.同一个包内，init()函数在所有变量已经初始化好后执行，init()函数执行过后才会执行其他函数   
     2.如果有导入其他包，则导入包的变量初始化以及init()函数会先执行，然后才会执行自己程序的变量初始化和init()   
     3.导入包时使用"_"，可只执行导入包变量初始化以及init()函数而不引入其他功能
---
Q:recover能不能监控整个程序的panic：
A:defer仅服务于当前groutine，因此recover监控panic也只能监控当前groutine内的panic，且特别注意recover只能显式的写在defer内部，如果是defer内一个函数调用其他函数，recover在另一个函数内，则是不起作用的   
---
Q:在docker容器中，GOMAXPROCS是否能正确判断被限制cpu后的容器实际上有多少cpu资源？   
A:虚拟化容器中golang的GOMAXPROCS不能正确判断当前容器究竟用了多少cpu，而是会用其宿主机的cpu数量作为默认值，这就会导致容器内cpu压力过大，在cpu密集型任务中性能可能有所下降。可使用 uber-go/automaxprocs包，该包可正确判断容器内的cpu个数   
  参考链接:[GOMAXPROCS 与容器的相处之道](http://gaocegege.com/Blog/maxprocs-cpu)  
---