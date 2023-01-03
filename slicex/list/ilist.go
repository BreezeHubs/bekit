package list

type IArrayList[T any] interface {
	// Get 获取对应下标的数据
	Get(index int) (T, error)

	// Set 对已存在的下标更新数据
	Set(index int, value T) error

	// Append 新增数据到最后
	Append(value T) int

	// QueueJump 数据插队
	QueueJump(index int, value T) error

	// Delete 删除数据
	Delete(index int) (T, error)

	// Len 获取数据长度
	Len() int

	// Cap 获取数据容量
	Cap() int

	// Range 遍历数据
	Range(f func(index int, value T) error) error

	// GetSlice 获取所有数据
	GetSlice() []T
}
