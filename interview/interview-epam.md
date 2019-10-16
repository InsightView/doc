# JAVA interview 

1. 二分和快排, 需要巩固一下, 二分有各种变种, 快排也有.
2. How to use a callback module in async function
包一层promise就可以了   
3. How nodejs code running in V8 or C++ 
4. the differ between docker and virtual machine

https://www.zhihu.com/question/48174633

虚拟机实现资源的隔离的方式是利用独立的Guest OS，以及利用Hypervisor虚拟化CPU、内存、IO等设备来实现的，对于虚拟机实现资源和环境隔离的方案，Docker显然简单很多。然后Docker并没有和虚拟机一样利用一个独立的Guest OS执行环境的隔离，它利用的是目前当前Linux内核本身支持的容器方式，实现了资源和环境的隔离，简单来说，Docker就是利用Namespace 实现了系统环境的隔离，利用了cgroup实现了资源的限制，利用镜像实例实现跟环境的隔离。

5. the jenkins workflow
    + create a jenkins project
    + add repository
    + build periodically trigger
    + post step

```
    以我之前的工作经验来说：
    先执行单元测试
    将代码编译或者打包
    发送邮件给QA并暂停CI流程
    QA手动的将打包文件部署到测试环境进行验证
    QA回到CI平台触发流程继续，也就是通过了测试
    RD（研发）进入PAAS上线平台，勾选要上线的版本，并发起上线申请
    BOSS/OP（上级）确认上线申请，上线平台创建若干服务容器，将对应版本的打包代码释放到容器内的特定目录
    上线平台切换负载均衡，流量切换到新版本容器，完成上线
```

6. Example for some ES6 feature
    + arrow function; advantage or disadvantage
7. prototype and nodejs constructor
8. destructuring
9. how you use mongodb cluster
    + the master node and slaver node, what the differ
10. how you use redis
11. Do you know Jest for nodejs tdd
12. S3中的一些feature, 比如policy, lifecycle之类的特性
13. IAM的一些基本概念
14. 怎样使自己的代码减少bug
15. typescript 和 javascript的不同和一些用法
16. vue和react的一些基本概念

