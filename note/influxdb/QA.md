---
Q:设置了删除策略后，为什么能查询到策略范围外的数据？  
A:influxdb的保留策略触发机制默认为30分钟一次，这个数值可以自己在配置文件中设置：  
    [retention]  
    enabled = true // 设置为false防止influxdb强制执行保留策略
    check-interval = "30m0s"   // 多久一次检查 
--- 
Q:按照时间范围备份数据，备份到的内容中有超出时间范围内的数据？  
A:指定时间间隔的备份是在数据块上执行的，而不是逐点执行。由于大多数块都是高度压缩的，因此提取每个块以检查每个点都会给正在运行的系统造成计算和磁盘空间负担。每个数据块都以该块中包含的时间间隔的开始和结束时间戳进行标记。当指定-start或-end时间戳时，将备份所有指定的数据，但是也会备份同一块中的其他数据点。  
  因此其预期行为有：
  + 还原数据时，您可能会看到指定时间段之外的数据。
  + 如果备份文件中包含重复的数据点，则将再次写入这些点，从而覆盖所有现有数据。
---
Q:内存占用过大的问题怎么优化？
A:
+ 数据库设计上避免一个数据库太多series。百万级别的很容易OOM（SHOW SERIES CARDINALITY）   
+ 采用tsl索引，通过修改配置文件中的index-version = “inmem” > index-version = “tsi1”   
参考链接：   
[记一次influx内存调优](https://www.jianshu.com/p/dbbb73b537e1)   
[how-to-overcome-memory-usage-challenges](https://www.influxdata.com/blog/how-to-overcome-memory-usage-challenges-with-the-time-series-index/)

    