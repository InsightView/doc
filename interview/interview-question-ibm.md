ibm五面面经

infra

1. 写一个广度优先遍历的扩展
2. 自我介绍 讲一下用到的东西
3. 深度优先遍历

就是写完广度优先需要举一反三, 还用recursive方法写了一遍

DFS: 

    public int maxDepth(TreeNode root) {
        if(root == null) {
            return 0;
        }
        
        Stack<TreeNode> stack = new Stack<>();
        Stack<Integer> value = new Stack<>();
        stack.push(root);
        value.push(1);
        int max = 0;
        while(!stack.isEmpty()) {
            TreeNode node = stack.pop();
            int temp = value.pop();
            max = Math.max(temp, max);
            if(node.left != null) {
                stack.push(node.left);
                value.push(temp+1);
            }
            if(node.right != null) {
                stack.push(node.right);
                value.push(temp+1);
            }
        }
        return max;
    }

BFS:

    public int maxDepth(TreeNode root) {
        if(root == null) {
            return 0;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        int count = 0;
        while(!queue.isEmpty()) {
            int size = queue.size();
            while(size-- > 0) {
                TreeNode node = queue.poll();
                if(node.left != null) {
                    queue.offer(node.left);
                }
                if(node.right != null) {
                    queue.offer(node.right);
                }
            }
            count++;
        }
        return count;
    }
4. mq的一些知识

用的是rocketMQ, 通过轮询所有队列的方法来确定消息被发送到哪一个队列. 根据不同的业务, 可以将ID作为计算队列, 让业务ID相同的消息先后发送到同一个队列中, 在获取到路由信息之后, 根据MessageQueueSelector实现额算法来选择一个队列, 同一个OrderID获取到的肯定是同一个队列

在网络不稳定的情况下, RocketMQ会出现消息重复的问题, 比如发送消息、消费消息、负载均衡的时候,所以需要:

    1. 消费端处理消息的业务逻辑保持幂等性(idempotence)
        多次请求或操作的结果是一致的. 

    2. 保证每条消息都有唯一编号且保证消息处理成功与去重表的日志同时出现




5. 画一个之前做过的项目的架构图
6. linux中top的一些内容 各个参数的意义
7. free中 各个参数的意思 什么swaq之类的
8. 虚拟化内存的一些东西
9. tcp的三次握手和四次挥手 区别
10. awk的命令使用 sed的使用
11. docker的网络
12. k8s和docker的区别 
13. k8s的网络
14. k8s的常用api
15. pod之间是怎样进行通信的 还是网络
16. pod a 和pod b a对外发布 b与a链接对内 怎样设计
17. k8s中暴露服务的方式
18. redis用过什么
19. 用过哪些持久化数据库 我说couchdb然后加上一些内容和对比
20. k8s的service怎么通信的 是一个什么样的概念
21. container是怎样进行资源隔离的 从文件系统和网络
22. etcd的一些概念和用法
23. 多线程 进程和协程
24. 写一个多重二分查找
25. 写一个动态规划
26. k8s怎样实现auto scaling的
27. 一个项目中你怎么设计auto scaling
28. k8s的持久化
29. prometheus 的delay 
30. etcd的调优 和prometheus 的调优
31. stateful和deployment 的区别 
    + https://www.cnblogs.com/weifeng1463/p/10284122.html
    + Deployment用于部署无状态服务，StatefulSet用来部署有状态服务

StatefulSet 稳定的, 唯一的网络标识, 持久的储存; 有序的, 优雅的伸缩和部署, 删除和自动滚动更新,

32. jenkins 的一些内容 做一个全流程怎么做
33. k8s的是怎样负载均衡的
34. ecr和esk和k8s的区别
35. git 怎么做cherry pick
36. socket 编程

是一个TCP/IP协议族的封装, 是应用层和TCP/IP的中间抽象层. 起源于Unix, 就是一切都是文件, 可以通过open write/read close来操作. 所以一个Webserver就是

    1. 建立链接
    2. 接受请求, 从网络中读取一条HTTP请求报文
    3. 处理请求, 访问资源
        1. socket() 创建套接字
        2. bind() 分配套接字地址
        3. listen() 等待连接请求
        4. accept() 循序连接请求
        5. read/wriate 数据交换
        6. close() 关闭连接
    4. 构建响应, 创建带有header的HTTP响应豹纹
    5. 发送响应, 传给客户端

37. kafka 在hyperldger中怎么用的

kafka是一个消息分发的系统, 使用消费者订阅Topic来接受和消费Procuder publish的消息. 数据量大的时候构建kafka集群来处理所有消息. kafka并不是不会记录那些已经消费的消息, 而是会储存一定的时间. 在leader-follower的系统中, leader拥有一个partition, 同时follower有replication, 他们使用投票来跟换leader进行crash tolerance. 这里就引入了zookeeper服务. 

投票的时候还有一个max.pool.interval.ms 想要足够大, 以便使listener处理max.poll.records. 

zookeeper是一个分布式的KV储存库,一般用来储存元数据和处理集群. 在kafka的设计中,使用zookeeper来管理broker, zk上会有一个专门用来记录broker的位置/brokers/ids; 每个broker在启动的时候都在zk上注册, topic也有类似于这种方式, producer需要将消息合理分发到broker上, kafka支持传统4层lbs, zk耶支持zk的lbs; 

传统4层, 生产者ip和端口定相关联的broker, 通常一个生产者对应单个broker, 只需要单个tcp链接. 这样导致不同的broker接受到的消息不均匀,也无法感知broker的新增和删除

使用zk监听zk的broker感知broker和topic的状态来动态lbs. 

kafka in hyperldger

对于每一个chain, 有隔离的partition; 每一个chain都映射到单独的partition topic; orderer node在确定了chain的tx的权限的时候(通过rpc广播), 接替了属于那个chain的通信partition; orderer可以消费这个partition然后返回一个排好序的txlist 然后在所有orderers中同步; 在链上的tx被blocked了, 每一个新的batch来的话都会设置一个计时器; 当block已经了的时候活着计时器超时的时候, 被cut; 每一个orderer维护本地log在每一个chain, 然后block的结果储存在local ledger中; 