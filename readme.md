cron 每天定时将基金数据同步到内存，程序启动时也会调用一次
db   数据选型是mongodb存储数据，数据本身没有过多的事务要求，需要高性能的查询
model  项目的逻辑部分
common 可抽取的公共内容
web    接口路由
main.go  入口文件

案例：
    输入  基金类型:all   开始时间:starttime   结束时间:endtime  查询指标：max_retracement(最大回撤)
    输出  所有基金的最大回撤会大到小排名

