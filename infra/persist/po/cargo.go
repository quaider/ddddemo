// Package po 根据实际情况决定是放入gorm内还是在外层
// 理论上来说，不同数据库其对应的po可能有差别，因此一般 po 和具体仓储实现绑定
// PO 需要对外公开，原因是 查询的时候会用到(CQRS)
package po

// CargoPo 数据库持久化对象
type CargoPo struct {
}

// 其他持久化对象
