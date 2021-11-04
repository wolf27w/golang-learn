//RabbitMQ介绍
//1.1.1. MQ简介
//简单释义
//消息总线(Message Queue)，是一种跨进程、异步的通信机制，用于上下游传递消息。由消息系统来确保消息的可靠传递。
//
//背景描述
//当前市面上mq的产品很多，比如RabbitMQ、Kafka、ActiveMQ、ZeroMQ和阿里巴巴捐献给Apache的RocketMQ。甚至连redis这种NoSQL都支持MQ的功能。
//
//适用场景
//上下游逻辑解耦&&物理解耦
//保证数据最终一致性
//广播
//错峰流控等等

//         RabbitMQ的特点