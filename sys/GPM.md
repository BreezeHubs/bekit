M Machine: os线程(即操作系统内核提供的线程)
G Goroutine: 包含调度一个协程所需要的堆栈 以及 instruction pointer(IP指令指针)，以及其它一些重要调度信息
P Process: M与G的中介，实现m:n调度模型的关键，M必须拿到P才能对G进行调度，P其实限定了goroutine调度的最大并发度。一个逻辑处理器一个P绑定一个os线程

```
[M]              M线程
 |               G协程
[P] —— [G]       P调度中介
 |      |  
[G]    [G]       多个M 支持多个G队列
 :      |        M从线程池取出
 :     [G]
 :      :
 :      :
 :     代表正在等待的G
 :
代表当前执行的G
```  
工作流窃取
P队列上的goruntine全部调度完后，对应的M首先会尝试从`global runqueue`中获取goruntine进行调度
如果`global runqueue`中没有goruntine，当前M会从别的M对应P的`local runqueue`抢一半的goruntine放入自己的队列中进行调度