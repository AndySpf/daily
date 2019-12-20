Q:init()函数有什么细节？   
A:1.同一个包内，init()函数在所有变量已经初始化好后执行，init()函数执行过后才会执行其他函数   
     2.如果有导入其他包，则导入包的变量初始化以及init()函数会先执行，然后才会执行自己程序的变量初始化和init()   
     3.导入包时使用"_"，可只执行导入包变量初始化以及init()函数而不引入其他功能
---
Q:recover能不能监控整个程序的panic：
A:defer仅服务于当前groutine，因此recover监控panic也只能监控当前groutine内的panic
