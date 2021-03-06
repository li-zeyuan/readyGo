# 批量导入

### 背景

- 批量导入代码混乱，迭代难道大
- 导入处理逻辑存在爆内存风险
- 导入10W数据时，偶发内存OOM

### 流程

```sequence
前端->后端: get 导入结果获取接口
后端->前端: 上次导入结果或导入进度（列表）

Note over 前端: 
前端->后端: post 开始导入接口[oss_url]
Note over 后端: 任务写入pg，通知seagull导入
后端->消息队列: 
消息队列->消费者: 

Note over 消费者: 读取pg中的task
Note over 消费者: rides加导入任务正在执行标记
Note over 消费者: 查pg，循环执行导入任务
Note over 消费者: 读取excel文件，md去重
Note over 消费者: 数据解析、校验

前端->>后端: get 轮询获取进度
后端->>前端: 进度

消费者->>消费者: 分批(1000)：校验、非法数据标红写入excel、入库、流转

前端->>后端: put 删除待导入文件
后端->>前端: ok

Note over 消费者: 更新pg进度、excel导入状态
Note over 消费者: 清除rides任务正在执行标记

Note over 后端: 多个excel导入完成

前端->后端: get 导入结果获取接口
后端->前端: 导入结果
```



### 优化点

- 用面向对象的方式重新批量导入功能
- 优化处理逻辑，由对一次性sheet操作优化成分片的方式
- 定位内存OOM故障点，修复内存暴涨问题

### 遇到问题

- 消息顺序问题
  - 解决：生产者发送消息并保存到pg，消费者根据pg的id顺序去处理消息
- 消息重复消费问题
  - 解决：excel文件的md5保存到pg，再次导入时，取excel文件md5值并与pg中的md5值去重。
- 超大消息传输问题
  - 解决：生产者先将excel上传到oss，然后将oss链接通过消息体发送给消费者

- 内存OOM问题
  - 解决：
    - 1、优化导入流程，由对一次性sheet操作优化成分片的方式
    - 2、go pprof 内存分析，找出故障点（第三方包excel数据处理方式为普通循环读写，解析出的sheet结构体大，gc不及时）；优化成缓存读写的方式（创建一个缓存buf，buf满后写入excel文件，然后清空buf，并且弃掉一些如样式，字体等不重要的信息）。