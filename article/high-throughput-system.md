## Building a high throughput system

### 康威定律 conway‘s law

### soa 架构

### Add LRU for you code 

add LRU to preset cache limit, the utility rate of memory can be increase. 

But when the data change freauqnetly, when the cache correct rate raised before in a long time. It may caused invaild of new cache, the service can not update real time. The system needs to clear policy based on real-time. Three points needs to be followed: set cache timer; query cache timer; cache the cache timer. If the timer is out, the cache data will be cleared. 

Hash LRU cache is used when the traffic raised dramatically. The LRU cache will coat lots of time in this condition. We could add hash into our system. HashLRUcache will reduce the time of query. the shard of the hash should be setted by the core of CPU.

zero copy policy: using address to copy the complex data. But due to the real-time clear policy, the address will be invaild when the data is cleared. The solution is atom reference counter. We still needs to ensure the consistency of every thread.

