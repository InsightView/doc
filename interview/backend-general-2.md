## interview 
比较简单, 但是一些没接触过的

1. why dockerfile needs to put those apt-update and install in one line?

make sure those cache is not store on the disk, when split into two line, update will cache those package, binary or other stuff

2. redis 分布式事务锁的用法 需要结合业务场景

3. redis 的有序集合使用

4. 一个论坛的在线人数, 使用redis来解决. 可以使用有序集合, 那如何优化. 比如不断更新的人数消耗过大. 这样可以使用前缀的字段表示, 从而实现过滤. 设置ttl字段之类的. 通过查询约束来解决. https://www.v2ex.com/t/298920

5. https的加密方式

6. https的状态码 3 4 5 

7. restful的 put delete

8. 加密算法, 非对称加密; sha256; ecdsa

9. websocket 长链接的问题, 在什么时候用到了
做的一个数据实时可视化的项目, 就是实时更新的两种方式, 一个轮询, 一个websocket. 怎么websocket内部怎么实现的, 是什么样子的一种机制 https://www.jianshu.com/p/ef39cc3085ac

10. 区块链算力排序的一些概念

11. golang 怎么优化代码, 使用的工具pprof

看完这个就好了https://github.com/dgryski/go-perfbook/blob/master/performance-zh.md

12. 有没有做过深度一些的项目