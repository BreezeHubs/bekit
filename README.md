# bekit

个人开发 支持泛型的工具库

### 目前可以使用的工具方法：
- file
  - 判断文件或文件夹是否存在：PathOrFileExists(path string) (bool, error)

- ginutil(在gin中使用的便捷方法)
  - http
    - 文件流下载：FileDataStream(ctx *gin.Context, fileName string)
    - 文件流下载：FileStream(ctx *gin.Context, fileName string)
  - pkg
    - 跨域：Cors(url ...string) gin.HandlerFunc
    - 处理并自定义gin的参数绑定验证 FilterBindErr(errs error, r any) error
    - 用于gin.Use，方便捕捉异常并统一返回错误信息：Recover(isPrintErr, isStack bool, f func(c *gin.Context, err error)) gin.HandlerFunc

- mode(设计模式相关工具方法)
  - optional
    - 为结构体绑定optional方法：BindOpt[T any](t *T, opts ...Opt[T])、BindOptWithExcept[T any](t *T, opts ...OptWithExcept[T]) error
  
- request
  - get请求：Get(url string) *curl (c *curl) Do(headers map[string]string, t time.Duration) ([]byte, error)
  - post请求：Post(url string, data map[string]string) *curl (c *curl) Do(headers map[string]string, t time.Duration) ([]byte, error)

- slice
  - list(数组、并发安全数组、链表) 
    - 数组：NewArrayList() Get Set Append QueueJump Delete Len Cap Range GetSlice
    - 并发安全数组：NewConSafeArrayList() Get Set Append QueueJump Delete Len Cap Range GetSlice
    - 双向循环链表存储的数组：NewLinkedList() Get Set Append QueueJump Delete Len Cap Range GetSlice
  - 生成新数组：ForEach[T arrayT](arr []T, f func(T, int) any) ([]T, error)
  - 判断map是否包含：MapHas[T comparable, K comparable](array map[K]T, has T) bool
  - 判断map是否符合对比：MapHasFunc[T any, K comparable](array map[K]T, has T, f func(src, has T) bool) bool
  - 判断map是否包含多个：MapHasAny[T comparable, K comparable](m map[K]T, has []T) bool
  - 判断map是否多个符合对比：MapHasAnyFunc[T any, K comparable](m map[K]T, has []T, f func(src, has T) bool) bool
  - 缩容成新数组：ArrayShrink[T any](src []T) []T
  - 判断数组是否包含：ArrayHas[T comparable](array []T, has T) bool
  - 判断数组是否符合对比：ArrayHasFunc[T any](array []T, has T, f func(src, has T) bool) bool
  - 判断数组是否包含多个：ArrayHasAny[T comparable](array []T, has []T) bool
  - 判断数组是否多个符合对比：ArrayHasAnyFunc[T any](array []T, has []T, f func(src, has T) bool) bool
  - 获取对应下标的数据：FindByIndex[T any](array []T, idx int) (val T, err error)
  - 去除对应下标的数组，返回新数组：ExcludeByIndex[T any](array []T, idx int) (res []T, val T, err error)
  - 数组转map：ArrayToMap[T any](array []T) map[int]T
  - map转数组：MapToArray[T any, K comparable](m map[K]T) []T

- sqlx
  - 根据date format语句转换sql：DateFormat(field string, format string) string

- sys(系统相关工具函数)
  - 监听服务退出信号：NewListenExitSignal() *eSignal IsExit
  - 获取当前线程的ID：GetCurrentThreadIDByWin() int
  - 设置当前go env的线程数：SetThreadNum(n int) int
  - 获取当前go env的线程数：GetThreadNum() int
  - 获取当前系统的CPU核数量：GetCPUNum() int