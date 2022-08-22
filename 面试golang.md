1. make 和 new 的区别

   ```shell
   make:用来初始化map,channel和切片，返回类型本身
   new：返回type指针
   ```

2. 了解过golang的内存管理吗?

   > GC?标记回收?三色回收

3. 调用函数传入结构体时，应该传值还是指针?说出你的理由?

   ```shell
   在struct传参传递指针，进行计算或对原结构体进行修改会传递指针
   传递指针产生逃逸，栈上指针逃逸到堆上
   ```

4. 线程有几种模型?Goroutine的原理了解过吗，讲一下实现和优势?

   > GMP模型
   >
   > 优势:线程系统级，占用资源多。携程轻量级，占用资源少

   Goroutine什么时候会发生阻塞?

   ```shell
   发生阻塞，调度器会进行等待。
   ```

   PMG模型中Goroutine有哪几种状态? 

   > GRuning
   >
   > GWaiting

   每个线程/协程占用多少内存知道吗? 

   ```shell
   线程2M
   携程2k
   ```

   如果Goroutine—直占用资源怎么办,PMG模型怎么解决的这个问题?

   > goroutine正常模式和饥饿模式
   >
   > 信号协作

   如果若干线程中一个线程OOM，会发生什么?如果是Goroutine 呢?

   > 

   项目中出现过OOM吗，怎么解决的?

   ```shell
   链路追踪。抓火焰图。pprof
   ```

   项目中错误处理是怎么做的?

   > 统一的错误处理。不推荐

   如果若干个Goroutine,其中有一个panic，会发生什么?

   > panic捕获

   defer可以捕获到其Goroutine的子Goroutine 的panic吗?

   ```shell
   不行
   ```

   开发用Gin框架吗?Gin怎么做参数校验?

   ```shell
   shouldBingJson,shouldBind 处理完参数后进行参数验证
   ```

   中间件使用过吗?怎么使用的。Gin的错误处理使用过吗?Gin中自定义校验规则知道怎么做吗?自定义校验器的返回值呢?

   ```shell
   日志记录。ssl https
   reflect包。typeof和valeof方法校验参数类型和参数值
   default recover 捕获异常
   ```

   golang中解析tag是怎么实现的？反射原理是什么？通过反射调用函数

   ```shell
   通过反射
   type value field获取原方法，源类型
   ```

   golang的锁机制了解过吗? Mutex的锁有哪几种模式，分别介绍一下? Mutex锁底层如何实现了解过吗?

   ```shell
   正常模式 饥饿模式
   读写锁 map携程不安全 修改map前lock sync.Map携程安全的锁
   ```

   channel、channel使用中需要注意的地方？

   ```shell
   close的channel不能传递信息
   
   ```

   数据库用的什么？数据库锁有了解吗？mysql锁机制讲一下。mysql分库分表。

   ```shell
   行锁表锁
   ```

   讲一下redis分布式锁？redis主从模式和集群模式的区别了解过吗？redis的数据类型有哪些？redis持久化怎么做的？

   ```shell
   setNx方法 redlock
   string hash zset list ...
   aof rdb
   ```

   

   编程题：你了解的负载均衡算法有什么？实现一个负载均衡算法。

```go
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type Client struct {
	Name string
}

type LoadBalance struct {
	ClientArr []Client
	Size      int
}

func getLoadBalance(size int) *LoadBalance {
	var clients []Client
	for i := 1; i <= size; i++ {
		clients = append(clients, Client{Name: "load" + strconv.Itoa(i)})
	}
	return &LoadBalance{ClientArr: clients, Size: size}
}

func (l *LoadBalance) getClients() *Client {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(100)
	return &l.ClientArr[intn%l.Size]
}

func (c *Client) Do() {
	fmt.Println(c.Name)
}

func main() {
	//随机负载均衡算法
	balance := getLoadBalance(5)
	balance.getClients().Do()
	//反射获取
	p := &People{}
	of := reflect.ValueOf(p)
	of.MethodByName("Eat").Call([]reflect.Value{})
}

type People struct {
	Name string
	Age  int
}

func (p *People) Eat() {
	fmt.Println("people eat")
}

```

banchmark 
