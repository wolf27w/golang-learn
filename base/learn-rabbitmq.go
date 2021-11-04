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

//RabbitMQ是由Erlang语言开发的AMQP的开源实现。
//
//AMQP：Advanced Message Queue，高级消息队列协议。它是应用层协议的一个开放标准，为面向消息的中间件设计，基于此协议的客户端与消息中间件可传递消息，并不受产品、开发语言灯条件的限制。
//
//可靠性(Reliablity)：使用了一些机制来保证可靠性，比如持久化、传输确认、发布确认。
//灵活的路由(Flexible Routing)：在消息进入队列之前，通过Exchange来路由消息。对于典型的路由功能，Rabbit已经提供了一些内置的Exchange来实现。针对更复杂的路由功能，可以将多个Exchange绑定在一起，也通过插件机制实现自己的Exchange。
//消息集群(Clustering)：多个RabbitMQ服务器可以组成一个集群，形成一个逻辑Broker。
//高可用(Highly Avaliable Queues)：队列可以在集群中的机器上进行镜像，使得在部分节点出问题的情况下队列仍然可用。
//多种协议(Multi-protocol)：支持多种消息队列协议，如STOMP、MQTT等。
//多种语言客户端(Many Clients)：几乎支持所有常用语言，比如Java、.NET、Ruby等。
//管理界面(Management UI)：提供了易用的用户界面，使得用户可以监控和管理消息Broker的许多方面。
//跟踪机制(Tracing)：如果消息异常，RabbitMQ提供了消息的跟踪机制，使用者可以找出发生了什么。
//插件机制(Plugin System)：提供了许多插件，来从多方面进行扩展，也可以编辑自己的插件。
//1.1.3. rabbitmq简单使用

//所有MQ产品从模型抽象来说，都是一样的过程：
//
//消费者(consumer)订阅某个队列。
//生产者(product)创建消息，然后发布到队列中(queue)，最终将消息发送到监听的消费者。
//这只是最简单抽象的描述，具体到RabbitMQ则由更详细的概念需要解释。