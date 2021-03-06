# snowflake 

## snowflake 介绍

Twitter Snowflake算法介绍：
https://segmentfault.com/a/1190000011282426

对于分布式的ID生成，以Twitter Snowflake为代表的， Flake 系列算法，属于划分命名空间并行生成的一种算法，生成的数据为64bit的long型数据，在数据库中应该用大于等于64bit的数字类型的字段来保存该值，比如在MySQL中应该使用BIGINT。

![](https://segmentfault.com/img/bVVulC?w=1021&h=346)

* 1位，不用。二进制中最高位为1的都是负数，但是我们生成的id一般都使用整数，所以这个最高位固定是0
* 41位，用来记录时间戳（毫秒）。
  - 41位可以表示$2^{41}-1$个数字，
  - 如果只用来表示正整数（计算机中正数包含0），可以表示的数值范围是：0 至 $2^{41}-1$，减1是因为可表示的数值范围是从0开始算的，而不是1。
  - 也就是说41位可以表示$2^{41}-1$个毫秒的值，转化成单位年则是$(2^{41}-1) / (1000 * 60 * 60 * 24 * 365) = 69$年
* 10位，用来记录工作机器id。
  - 可以部署在$2^{10} = 1024$个节点，包括5位datacenterId和5位workerId
  - 5位（bit）可以表示的最大正整数是$2^{5}-1 = 31$，即可以用0、1、2、3、....31这32个数字，来表示不同的datecenterId或workerId
* 12位，序列号，用来记录同毫秒内产生的不同id。
  - 12位（bit）可以表示的最大正整数是$2^{12}-1 = 4095$，即可以用0、1、2、3、....4094这4095个数字，来表示同一机器同一时间截（毫秒)内产生的4095个ID序号

## snowflake缺点

雪花算法存在的缺点是：

* 依赖机器时钟，如果机器时钟回拨，会导致重复ID生成
* 在单机上是递增的，但是由于设计到分布式环境，每台机器上的时钟不可能完全同步，有时候会出现不是全局递增的情况（此缺点可以认为无所谓，一般分布式ID只要求趋势递增，并不会严格要求递增～90%的需求都只要求趋势递增）

## snowflake 算法异常情况分析

整个ID生成过程中，启动的时候会对外部有依赖，因为需要知道当前的机器ID，之后就独立工作。

几种常见的异常情况分析：
问：如果某个时间戳的所有ID都被用完了，那怎么办？
答：就继续等待下一毫秒然后在生成ID

问：获取当前时间戳时，如果获取到的时间戳比之前一个1已生成的ID的时间戳还要小，怎么办？
答：snowflake的做法是继续获取当前机器的时间戳，直到获取更大的时间戳才继续生成ID（在这个过程中是不分配新的ID）
如果snowfalke运行的服务器上时钟有大量的偏差时，整个snowflake系统就不能正常工作（偏差越多，分配新ID等待的时间越久）


snowflake的官方文档中明确要求必须配置NTP,并且NTP配置成不可向后调整的模式。

## Go语言实现snowflake的第三方库

[sonyflake](https://github.com/sony/sonyflake)

