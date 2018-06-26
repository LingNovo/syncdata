package guide

// 树形结构
type Record struct {
	Id       string    // 唯一标识
	Name     string    // 名称
	FullName string    // 全路径名称
	IsValid  bool      // 是否有效
	Data     []*Record // 包含的数据
	Itmes    []*Record //子项
}
