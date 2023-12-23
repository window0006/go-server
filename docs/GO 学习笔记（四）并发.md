# 并发基础
作为一名前端，多少还是知道并发的，比如 react 的 concurrency 模式，基于 fiber 节点链表结构，实现了自己的并发调度算法，可以更高效完成 diff 与渲染更新。<br />但是这和后端开发的时候说的并发是一回事吗？系统层面，硬件层面支持的并发又是怎么回事？
## 并发与并行
可以认为，并发（Concurrency）就是任务排好队，按照调度算法，分时轮流执行，一个 CPU 也是可以并发的。<br />而并行（Parallelism）则是多个 CPU 在同一时刻去做一批任务。所谓**单核只能并发，多核可以并行。**

- 单核并发会有切换上下文的成本，而并行却没有。但是同时，因为所有线程都在一个 cpu 上处理，因此可以共享缓存和内存，可以以相对较低的成本实现线程之间的通信。
- 多核并行会让每个线程都在不同的核心上运行，它的好处是没有上下文切换的成本，但是却有较高的通信成本。
> 操作系统允许开发者指定线程应该在哪个核心上执行，以控制同一个进程内的线程在同一个核心上执行，降低通信成本。这种技术被称为线程亲和性（Thread Affinity）。
> 但是这需要开发者做好取舍，既不浪费多核的性能，也平衡好通信成本。
> go 语言中可以通过 `syscall` 包中提供的 `SchedSetaffinity` 方法设置线程亲和性。但是因为 goroutine 是经过 go 运行时调度到不同操作系统线程的，因此也无法保证设置之后就都在一个 cpu 核心上执行。
> 还可以通过系统的 taskset 命令来设置。

## 进程、线程、协程
并发和并行最终是为了提升程序的执行效率，尽量不浪费硬件资源，程序则表现为一个个进程，进程套着线程，线程可以交给不同的 cpu 核心通识执行，协程则是程序内部的调度逻辑了。
### 进程（process）
进程是操作系统资源分配的最小单位。当一个程序运行时，操作系统会为他创建一个进程，并为其分配必要的资源（内存、cpu 等）。
> 资源申请：比如我们在代码中声明一个长度为 10 的数组，就相当于是申请了一块足以存放 10 个数字的大小的内存。
> 编译器生成可执行文件时，会根据变量的类型和作用域确定需要多少内存空间。当程序运行时，操作系统会根据可执行文件中的信息为程序分配内存，以便程序能够正常运行。

一个进程就代表着一个程序，进程可以区分主进程和子进程，拿我们熟悉的浏览器来讲，打开浏览器就是启动浏览器主进程，而浏览器本身是多进程架构的：<br />![image.png](https://cdn.nlark.com/yuque/0/2023/png/303138/1691379874337-2acd9850-2d50-4d08-81da-87e6e8667c70.png#averageHue=%23faf9f9&clientId=u99660d17-9e63-4&from=paste&height=247&id=hZewJ&originHeight=494&originWidth=1142&originalType=binary&ratio=2&rotation=0&showTitle=false&size=132059&status=done&style=none&taskId=u42894a27-160e-4890-8003-5e67c5a0a4c&title=&width=571)<br />主进程可以负责子进程的管理，不同子进程则负责不同的功能，提供服务供其他子进程或者主进程调用。
### 线程（thread）
线程是 CPU 调度执行的最小单位。<br />进程中至少存在一个线程。可以使用多个线程并发或者并行提升运算效率，但同时也会存在线程之间的通信成本和管理成本。<br />![image.png](https://cdn.nlark.com/yuque/0/2023/png/303138/1691380126284-021cdd44-2c99-4e70-9efb-2e9bbd605575.png#averageHue=%23f8f7f7&clientId=u99660d17-9e63-4&from=paste&height=395&id=LnaQl&originHeight=789&originWidth=1142&originalType=binary&ratio=2&rotation=0&showTitle=false&size=238790&status=done&style=none&taskId=ubfffa37c-56cd-48cd-aa7d-07132f859ee&title=&width=571)<br />多线程就像是火车上的每节车厢，而进程就是火车。车厢最终组成了一个进程。<br />线程又分了多种模型：

- 内核级，就是我们上面说的 CPU 调度执行的最小单位，由操作系统实现 CPU 抢占式调度。
- 用户级（即协程），由应用程序创建及管理，程序内部实现调度逻辑。
- 两级混合，goroutine 的实现就是两级混合的，goroutine 与线程关系是 n:m，可以在 n 个系统线程上多工调度 m 个 goroutine，有 GMP 模型实现调度。

![image.png](https://cdn.nlark.com/yuque/0/2023/png/303138/1691407183865-651b7b9a-2933-45fc-8240-da1169058387.png#averageHue=%23a9cf8a&clientId=u467c301c-3529-4&from=paste&height=466&id=VAwGI&originHeight=932&originWidth=1012&originalType=binary&ratio=2&rotation=0&showTitle=false&size=326961&status=done&style=none&taskId=u1dea37d5-bddd-4912-9d35-b43692a4142&title=&width=506)
## 协程（co-routine）
进程和线程是系统层面的，而协程完全是程序内实现的调度逻辑，系统是不知道有协程这种东西存在的。<br />当然，协程的出现是因为它能更省内存，能更灵活地实现任务的切换，毕竟线程上下文切换的成本太高了，当并发数较大的时候，就不希望很频繁地切换线程了。<br />go 是从语言层面上就支持了协程的，go 语言的 goroutine 本质上就是一种协程。
> 协程并不是由操作系统调度的，而且应用程序也没有能力和权限执行 cpu 调度。怎么解决这个问题？ 
> 答案是，协程是基于线程的。内部实现上，维护了一组数据结构和 n 个线程，真正的执行还是线程，协程执行的代码被扔进一个待执行队列中，由这 n 个线程从队列中拉出来执行。这就解决了协程的执行问题。
> 那么协程是怎么切换的呢？
> 答案是，golang 对各种 io 函数 进行了封装，这些封装的函数提供给应用程序使用，而其内部调用了操作系统的异步 io 函数，当这些异步函数返回 busy 或 bloking 时，golang 利用这个时机将现有的执行序列压栈，让线程去拉另外一个协程的代码来执行，基本原理就是这样，利用并封装了操作系统的异步函数。包括 linux 的 epoll、select 和 windows 的 iocp、event 等。

## goroutine
goroutine 即 go 协程，通过 `go` 关键字即可轻松创建。
```go
func main() {
    // ...
	go func() {
		fmt.Println("hello world")
	}()
    // ...
}
```
go 的 main 函数也是一个特殊的 goroutine，作为 go 程序的入口，它可以创建其他 goroutine，当 main goroutine 执行完了，go 程序也就退出了，其他正在执行的 goroutine 也会被终止。
### 运行机制
我们可以类比一下 js 的时间循环机制。<br />我们知道浏览中的渲染进程有一个主线程可以执行 js。<br />script 中的 js 会被作为一个宏任务放入宏任务队列中，它可能会绑定 web 交互事件，可能会调用 setTimeout 在宏任务队列中创建下一个宏任务，也可能会调用 new Promise 等方法在微任务队列中添加微任务。<br />在主线程的执行完一个宏任务的代码逻辑之后，它会接着清空当前宏任务重的微任务队列，完了之后再结束当前的宏任务。<br />在事件循环机制的作用下，主线程会持续从任务队列中被取出宏任务来执行。
```javascript
setTimeout(() => console.log('hello'))
console.log('world')
```
如果我们将 main goroutine 当成执行 js 的主线程一样，那么它的 goroutine 队列就像宏任务队列一样（没有微任务队列了）。
```go
func main() {
	go func() {
		fmt.Println("hello")
	}()
	fmt.Println("world")
	time.Sleep(time.Second)
}
```
main goroutine 在执行到 `go` 关键字所在的语句时就会新建一个 goroutine，此时它并不会停下来，就像 js 中的 setTimeout 一样，下面的代码会接着执行。<br />但是不一样定的是 setTimeout 的回调会在主线程中的代码执行完之后再由主线程接着执行，而 goroutine 则会与 main goroutine 一起并发执行（如果被调度到其他核的线程上，则可能并行执行）。
> 这么看来 goroutine 类比为 web worker 会更合适一些……

go 程序中可能会启动非常多的 go 协程，那么这些协程实如何被执行的呢？
## GMP 调度模型
> 一切程序都运行在操作系统上，cpu 则是真正干活的人。

**G-M-P** 分别代表：

- **G** - Goroutine，go 协程，是参与调度与执行的最小单位，每一次使用 `go func` 就会生成一个 G
- **M** - Machine，指的是系统级线程，由 go 的 runtime 动态创建合理的数量，它是运行 goroutine 的实体，调度器会把可运行的 goroutine 分配过来
- **P** - Processor，指的是逻辑处理器是 GMP 模型的关键，根据 `GOMAXPROCS` 配置，go 程序启动时会创建对应数量的 P，每一个 P 都会维护一个待执行的 goroutine 队列（这是本地队列，除此之外，还有一个全局 goroutine 队列）。

调度过程如下：<br />![image.png](https://cdn.nlark.com/yuque/0/2023/png/303138/1691580040026-192203dc-d158-469d-bff1-cf96c8b8c774.png#averageHue=%23f6f6f6&clientId=u22a4d355-7199-4&from=paste&height=791&id=S1Rx9&originHeight=1582&originWidth=2484&originalType=binary&ratio=2&rotation=0&showTitle=false&size=764213&status=done&style=stroke&taskId=u0e454398-f4cb-4df9-bfda-3b7ff50a3fa&title=&width=1242)<br />在新建 G 时，会优先选择 P 的本地队列，如果本地队列满了，则将 P 的本地队列的一半的 G 移动到全局队列。<br />当 P 的本地队列为空时，它会尝试从其他 P 的本地队列中窃取一些 G，或者从全局队列中获取 G，这其实可以理解为调度资源的共享和再平衡。这也是 work stealing 调度算法。<br />如果找到了可运行的 G，P 会唤醒一个空闲的 M 来执行它。<br />对 G 来讲，P 相当于真正的 cpu 核心，只有绑定到 P 的 G 队列中，G 才会被调度执行。同时 P 还可以为 M 提供运行 G 所需要的上下文环境信息（跨 M 执行 G 的关键）。<br />而 M 作为真正执行代码逻辑的系统功能，只有在没有空闲的 M 来执行 G 的时候才会新建一个。<br />通过 schedule loop （事件循环）机制，这个调度可以一直持续下去直到 main 函数退出或者操作系统将程序终止。<br />当运行中的 goroutine 阻塞时，会被从 M 上解绑，让位给其他可与行的 goroutine，当其阻塞结束后，调度器会再将其放到 P 的 G 队列中等待调度执行。

那么调度是怎么开始的呢？
## go 的启动过程
> 我们总会被问到，当我们再浏览器地址栏输入了某个 url 之后会发生什么……那么，当我们执行一个 go build 出来的可执行文件的时候，会发生什么？

build 命令可以将 go 代码编译成可执行文件，编译过程如下：
```
graph LR
    0(写代码)--go程序--> 1(编译器)--汇编代码--> 2(汇编器)--.o目标程序-->3(链接器)--可执行文件-->4(结束)
```
直接在命令行访问打包的结果文件就可以运行了：
```shell
go build main.go
./main
```
这时候操作系统会加载这个可执行文件，创建进程和主线程，然后为主线程分配栈空间，并将命令行参数拷贝进去，最后将主线程放入操作系统的运行队列中等待调度执行。<br />当主线程被调度执行之后，go 程序就开始启动了。<br />启动时会获取命令行的参数、把内存物理页参数、cpu 核心数等信息。（调用 runtime 包中的 `runtime.args` 和 `runtime.osinit` 方法）<br />接下来就是 go runtime 的核心内容了：<br />初始化调度器，完成当前 G 系统线程初始化，通过上面获取到的 cpu 核心数和 `GOMAXPROCS` 配置创建多个 P，也会创建第一个系统线程 m0，并将 m0 和某一个 P 进行绑定。（调用 runtime 包中的 `runtime.schedinit` 方法）<br />M 和 P 绑定后，紧接着便是创建一个特殊的 goroutine，也就是 g0，每个 M 都会有一个 g0，它是 P 的 goroutine 队列的管理员，负责 goroutine 的调度和垃圾回收。<br />用来执行 `runtime.main`（这个方法是 go 运行时的入口，也是 go 程序真正的入口），并放到 P 的 G 队列中，等待调度。（调用 runtime 包中的 `runtime.newproc` 方法）
```go
// src/runtime/proc.go
func schedinit() {
    _g_ := getg()
    (...)

    // 栈、内存分配器、调度器相关初始化
    sched.maxmcount = 10000 // 限制最大系统线程数量
    stackinit()         // 初始化执行栈
    mallocinit()        // 初始化内存分配器
    mcommoninit(_g_.m)  // 初始化当前系统线程
    (...)

    gcinit()    // 垃圾回收器初始化
    (...)

    // 创建 P
    // 通过 CPU 核心数和 GOMAXPROCS 环境变量确定 P 的数量
    procs := ncpu
    if n, ok := atoi32(gogetenv("GOMAXPROCS")); ok && n > 0 {
        procs = n
    }
    procresize(procs)
    (...)
}
```
随后便是启动调度器的调度循环（runtime.mstart 方法），开始从 P 的 G 队列中取出 goroutine 来执行，也就是 main goroutine。<br />此时才会开始执行我们代码中的 main 函数。
```go
// The main goroutine.
func main() {
    g := getg()

    ...
    // 执行栈最大限制：1GB（64位系统）或者 250MB（32位系统）
    if sys.PtrSize == 8 {
        maxstacksize = 1000000000
    } else {
        maxstacksize = 250000000
    }
    ...

    // 启动系统后台监控（定期垃圾回收、抢占调度等等）
    systemstack(func() {
        newm(sysmon, nil)
    })

    ...
    // 让goroute独占当前线程， 
    // runtime.lockOSThread的用法详见http://xiaorui.cc/archives/5320
    lockOSThread()

    ...
    // runtime包内部的init函数执行
    runtime_init() // must be before defer

    // Defer unlock so that runtime.Goexit during init does the unlock too.
    needUnlock := true
    defer func() {
        if needUnlock {
                unlockOSThread()
        }
    }()
    // 启动GC
    gcenable()

    ...
    // 用户包的init执行
    main_init()
    ...

    needUnlock = false
    unlockOSThread()

    ...
    // 执行用户的main主函数
    main_main()

    ...
    // 退出
    exit(0)
    for {
        var x *int32
        *x = 0
    }
}
```
接下来便是 GPM 调度不断执行 goroutine 的过程了。<br />![image.png](https://cdn.nlark.com/yuque/0/2023/png/303138/1691577884345-0048178c-2006-4afd-a573-6a777de783ea.png#averageHue=%23fefefe&clientId=u22a4d355-7199-4&from=paste&height=446&id=A66pv&originHeight=892&originWidth=1602&originalType=binary&ratio=2&rotation=0&showTitle=false&size=179290&status=done&style=none&taskId=u6fa3a6c5-1ebc-4a1f-9c8e-2e621b6ecc0&title=&width=801)<br />以下行为会触发调度：

1. 使用关键字 go
2. 垃圾回收
3. 系统调用
4. 同步互斥操作，也就是 `Lock()`，`Unlock()` 等
# 并发安全
线程 A 可以操作这个地址的内存，线程 B、C、D 等也都可以，这也是所谓的竞态。<br />并发过程中，可能会出现同时有两个协程在操作相同的内存的情况，如 A 写入变量的过程中，B 来读，或者也来写入，就会导致程序 panic。
```go
func main() {
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 协程1，往 m 里加字段
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Printf("write %d\n", i)
			m[i] = i
		}
		wg.Done()
	}()
	// 协程2，往更后面写
	go func() {
		for i := 1000; i < 2000; i++ {
			fmt.Printf("write %d\n", i)
			m[i] = i
		}
		wg.Done()
	}()
	wg.Wait() // 等待上面的 wg Done，防止 main 直接结束
}
```
这段程序运行之后会导致 `fatal error: concurrent map writes` 的报错：
```
go run main.go 
...
write 1194
write 1195
write 221
fatal error: concurrent map writes

goroutine 18 [running]:
main.main.func1()
...
```
所谓安全并发，线程安全之类的概念，其实就是因为变量共享的问题。<br />我们定义的变量存在于内存之中，修改、读取变量时，其实就是在操作内存块，所以并发的安全问题其实是内存操作的安全问题。
### race 检测工具
当我们需要确认程序是否有竞态问题的时候可以使用 race detector 工具（加上 -race 参数即可）。
```shell
go run -race main.go > race.log
# 通过 > race.log 将 fmt.Println 输出的信息换个地方展示
# 方便查看控制台输出的报错信息
```
与直接 run 不同，这样运行程序之后，当出现竞态问题时，会在控制台输出 DATA RACE 的 WARNING：
```
==================
WARNING: DATA RACE
Write at 0x00c00009a180 by goroutine 7:
  runtime.mapaccess2_fast64()
      /usr/local/go/src/runtime/map_fast64.go:53 +0x1cc
  main.main.func2()
      /.../demo/main.go:24 +0xa4

Previous write at 0x00c00009a180 by goroutine 6:
  runtime.mapaccess2_fast64()
      /usr/local/go/src/runtime/map_fast64.go:53 +0x1cc
  main.main.func1()
      /.../demo/main.go:16 +0xa4
```
好了，现在我们知道同时出现读写操作可能会导致并发安全的问题了，那怎么解决呢？
## 锁
保证并发安全最常见的方案就是给操作加锁。<br />给一段代码加上锁之后，可以保证当前线程执行完该代码段之前，其他来执行该代码段的线程会被阻塞在这个代码段之前。这就可以避免同时有多个 go 协程在操作同一块内存。
```go
var count = mutex sync.Mutex

func () {
    mutex.Lock() // 在 unlock 之前，其他协程会停在这里
    // 这段被锁住的代码也被称为代码临界区
    // ...
    // 临界区结束
    mutex.Unlock()
}
```
加锁的方式也很简单，sync 包提供了不同类型的锁，下面来一一介绍。
### Mutex（互斥锁）
看看如何使用互斥锁解决上面的并发安全问题：
```go
func main() {
	m := make(map[int]int)
    
	var mu sync.Mutex
    
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			m[i] = i
            fmt.Printf("write %d\n", i)
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1000; i < 2000; i++ {
			mu.Lock()
			m[i] = i
            fmt.Printf("write %d\n", i)
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}
```
现在再运行就不会报错了。使用互斥锁时，当一个 goroutine 调用 `mu.Lock()` 成功拿到锁之后，其他 goroutine 在再次调用 `mu.Lock()` 时会被阻塞，直到获取到锁的 goroutine 调用 `mu.Unlock()` 释放锁。<br />看看输出的结果：
```
...
write 141
write 142
write 143
write 144
write 1000
write 1001
write 1002
...
write 1160
write 145
write 146
write 147
...
```
从输出的接口可以看到两个 goroutine 在交替执行临界区的代码逻辑。当一个 goroutine 成功调用 `mu.Lock()` 获取到锁之后，其他 go routine 就必须等待着，直到其调用 `mu.Unlock()`。
### RWMutex（读写锁）
一般除了往变量写入内容，我们还需要读取变量的内容。为了避免写的过程中有其他 goroutine 来读，需要也给读操作上锁，在写的 goroutine 调用 `mu.Unlock()` 之前，不能读。
```go
func main() {
	m := make(map[int]int)
    mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			m[i] = i
            fmt.Printf("write %d\n", i)
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			fmt.Printf("read %d value %d\n", i, m[i])
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}
```
但是这样做之后，当有大量的并发请求来读取 m 的数据的时候，不仅要等写操作释放锁，还要等其他读操作释放锁，而读之间的互斥是不必要的。因此 go 还提供了 RWMutex。<br />使用 RWMutex 可以稍微加快一点：
```go
func main() {
	m := make(map[int]int)
    mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			m[i] = i
            fmt.Printf("write %d\n", i)
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.RLock()
			fmt.Printf("read %d value %d\n", i, m[i])
			mu.RUnlock()
		}
		wg.Done()
	}()
	wg.Wait()
}
```
它除了一样有 Lock 只有，还提供了专门方便读操作用的 RLock 与 RUnlock 方法。<br />当多个 goroutine 都调用了 RLock，他们会都被写操作的 Lock 阻塞，但是他们之间不会互相阻塞。也就是读写是互斥的，但是读与读是可以同时执行的。
## atomic
因为 go 不能保证任何单独的操作是原子性的，即便是字符串的赋值语读取，都可能会因为 A 协程在写的过程中 B 协程读取了而导致意外的报错。<br />我们可以通过加锁来避免这个问题，但是锁是个低效的操作，会大大降低并发的效率。最好可以实现一种无锁的数据结构，自然解决并发读写的问题。<br />go 提供了 `sync/atomic` 包，里面提供了实现原子操作的一系列方法，它们也被称作原语。使用原语，可以帮我们实现这样的数据结构。<br />atomic 就是原子操作的意思。原子操作则是指在多个线程同时执行某操作时，不会被其它线程打断、干扰的操作（有点像事务？但是并不是）。可以想象到这其实是 CPU 的一个最小执行单位：当某个值的原子操作在被进行的过程中，CPU 绝不会再去进行其他的针对该值的操作。<br />刚开始接触 atomic 的时候，它提供的方法让我很困惑，为什么只能是类似 int32、uint32、int64、uint64 这种数据操作？
> 作为一名前端，我们的世界简单到只有 number，光是分 32  位、64位、有无符号整数就已经消耗掉了一半的心智 😂。

为什么要有 AddXXX，直接声明一个变量，然后赋值有什么不一样？
```go
var a int32
a = a + 1
// ？？？和 atomic 提供的方法有什么区别？
atomic.AddInt32(&a, 1)
```
答案当然是并发的问题，`a = a + 1` 的赋值操作实际上实际上被转成机器码之后，分成了好几句汇编指令：
```
mov eax, dword ptr [a]  
inc eax
mov dword ptr [a], eax
```
> 首先将变量 a 对应的内存值搬运到某个寄存器（如 eax）中，然后将该寄存器中的值自增 1，再将该寄存器中的值搬运回a的内存中。

当有多个 goroutine 并发执行的时候，这些语句并不能保证被三条三条地执行。<br />![image.png](https://cdn.nlark.com/yuque/0/2023/png/303138/1691654339966-e19e7910-5864-492d-a94b-ad8cd7f2909b.png#averageHue=%23fdfdfb&clientId=udb88b37f-9a89-4&from=paste&height=296&id=dD62v&originHeight=283&originWidth=451&originalType=binary&ratio=2&rotation=0&showTitle=false&size=45196&status=done&style=none&taskId=ua34814d3-a9b0-4d9b-b5cc-9055d2eaf0d&title=&width=471.5)<br />而 `atomic.AddInt32(&a, 1)` 则可能会被编译成如下的指令（不同 CPU 和操作系统架构下编译出来的机器码指令是不同的）：
```
TEXT runtime∕internal∕atomic·Xadd(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), BX // 注意第一个参数是一个指针类型，是64位，所以还是 MOVQ 指令
	MOVL	delta+8(FP), AX // 第二个参数32位的，所以是MOVL指令
	MOVL	AX, CX
	LOCK
	XADDL	AX, 0(BX)
	ADDL	CX, AX
	MOVL	AX, ret+16(FP)
	RET
```
看到有个 LOCK！
> LOCK 指令是一个指令前缀，其后是读-写性质的指令，在多处理器环境中， 可以确保在执行 LOCK 随后的指令时，处理器拥有对数据的独占使用，以此实现对共享内存独占访问。

看起来原子操作是最终还是通过在底层给内存加锁来实现的。<br />现在我们知道 atomic 的原子性是如何实现的了。但是它看起来只能加减数字，比较一下数值之类的简单操作，这够用干嘛的？
### WaitGroup
`sync` 包还提供了 `WaitGroup` 这种数据结构，它的 `Add`、`Wait` 和 `Done` 方法可以实现并发流程的控制。<br />Add 了 n 各等待标识之后，Wait 会被阻塞到 Done 被调用了 n 次之后。<br />看我们上面的示例代码，就是通过 wg 来保证 main 函数不马上退出，要等到 goroutine 都跑完。<br />但是这也仅仅是能做一定的流程控制。
### Once
同样也是 `sync` 包提供的方法，这个就和内存有点关系了，当我们有很多 goroutine 都可能会初始化某个数据结构（常用初始化配置来举例）的时候，他们可能会造成并发写操作的安全问题。
```go
var once sync.Once

func initConfig() {
    // 初始化配置信息
}

func main() {
    once.Do(initConfig)
    // 其他代码
}
```
这样做就可保证只会有一次的初始化。适合任何单例模式的初始化操作。<br />显然，这还不够，我们的业务代码又不是只有初始化的时候才会并发读写。
## channel
> 这又是一个挑战前端页面仔心智模型的东西。。。

channel 是消息通道， 是 goroutine 之间的通信方式。可以这样往 channel 中写入和读取数据：
```go
ch := make(chan int)
// 将一个数据 value 写入至 channel，这会导致阻塞，直到有其他 goroutine 从这个 channel 中读取数据
ch <- value
// 从 channel 中读取数据，如果 channel 之前没有写入数据，也会导致阻塞，直到 channel 中被写入数据为止
value := <-ch
// 箭头的指向就是数据的流向！
```
哦哦！会阻塞！当代码执行到上面这两种地方的时候，是有可能会停下来的。这样不就可以保证并发读写的安全了吗？<br />事实上 channel 本身就是一种内置的并发安全的数据结构，它可以在多个 goroutine 之间安全地传递数据。
```go
func main() {
	c := make(chan int)

	// 在一个 goroutine 中发送数据
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// 在另一个 goroutine 中接收数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(1000 * time.Millisecond)
}

```
就算有多个 goroutine 同时读写 channel 的数据，也不会导致并发安全问题，并且还能保证输出的顺序，因为在读出来之前，再也不会写进去了。<br />channel 其实还分了有缓冲（Buffered Channels）和上面这种无缓冲的。<br />区别在声明时就体现出来了：
```go
c1 := make(chan int) // 无缓冲
c2 := make(chan int, 1) // 有缓冲
```
有缓冲的 channel 有第二个参数：channel 的容量（capacity）。<br />有缓冲的 channel 像阻塞队列一样，读取时只会在没有数据时阻塞，写入时也只会在没有可用容量时阻塞。
### select-case
除了并发安全读写外，巧用 channel 的阻塞特性可以实现一些使用的技巧。<br />配合 `select-case` 语法，我们可以实现监听多个 channel 的数据写入，根据不同的 `case` 处理先获取到数据的 channel。比如下面的超时方法：
```go
func main() {
	c := make(chan struct{})

	go func() {
    	fmt.Println("do something...")
    	time.Sleep(2 * time.Second)
    	c <- struct{}{}
	}()

	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
```
`time.After(1 * time.Second)` 也会返回一个 channel，会在指定的延迟时间后写入数据。<br />select 语句将会阻塞到 `res := <-c` 或者`time.After(1 * time.Second)`中的任意一个能成功取到数据的时候。
# 并发控制
## 线程池（worker pool）
> worker，很形象，对前端页面仔毫无心智负担。

是的，就和我们写 js 时想要用 worker 来处理耗时任务一样，我们可以启动很多 goroutine 来处理耗时请求。<br />将启动 n 个 goroutine 并发处理耗时请求的逻辑封装一下，就是 Worker Pool 模式了。
```go

type job struct {
}
type result struct {
    err error
}
func main() {
	// step 1: specify the number of jobs
	var numJobs = 1

	// step 2: specify the job and result
	jobs := make(chan job, numJobs)
	results := make(chan result, numJobs)
    
	runtime.GOMAXPROCS(runtime.NumCPU())
    
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(jobs <-chan job, results chan<- result) {
			for j := range jobs {
				// step 3: specify the work for the worker
                // 次数应该运行耗时的操作
				var r result
				results <- someExpensiveCalc(j, r)
			}
		}(jobs, results)
	}

	// step 4: send out jobs
	for i := 0; i < numJobs; i++ {
		jobs <- job{}
	}
	close(jobs)

	// step 5: do something with results
	for i := 0; i < numJobs; i++ {
		r := <-results
		if r.err != nil {
			// do something with error
		}
	}
}

```

参考：<br />[详解 Go 程序的启动流程，你知道 g0，m0 是什么吗？](https://eddycjy.com/posts/go/go-bootstrap0/)<br />[Go Runtime Scheduler](https://speakerdeck.com/retervision/go-runtime-scheduler?slide=14)<br />[Go语言的原子操作atomic - ZhanLi - 博客园](https://www.cnblogs.com/ricklz/p/13648859.html)<br />[Go Channel 详解](https://colobu.com/2016/04/14/Golang-Channels/)<br />[Worker pools in Golang | schollz](https://schollz.com/tinker/worker-pool/)
