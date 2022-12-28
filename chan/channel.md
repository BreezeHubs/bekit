[缓冲] make(chan int, 1024)，等待-非阻塞
[非缓冲] make(chan int)，等待-阻塞

写入：ch <-1
读取：val := <-1

