# Overview
针对redis的订阅发布功能学习
redis的订阅就是我们经常看到用户下单成功之后收到的推送消息，文章推送等。
发送者--频道--订阅者
这是阅读https://github.com/HDT3213/godis.git 的时候做的笔记


## redis基础
* subscribe channe1 channel2 channel3 … ：订阅一个或者多个频道
* unsubscribe channe1 channel2 channel3 … ：退订订阅的指定频道(关闭客户端终端没用，需要命令退订)
* publish channe1 message：对指定频道发送消息
* pubsub numsub channel1 channel2：查看指定频道的订阅数

## 问题before 阅读代码
看代码之前，我猜测数据结构的设计应该是保存pub sub channel的类型应该是map 
map[channel名字]= 客户端的id的list

然后订阅频道，应该是贤检查字段内部是否存在，不存在那么需要先在map里面创建一个对应的键值对，来保存客户端ID

取消频道订阅，应该是将客户端从链表里面删除

# 目标
允许自己先抄，嘻嘻
自己尝试写一个单机版本的redis，实现tcp服务器， 支持简单pub sub功能即可
不需要支持集群模式




