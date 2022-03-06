package db

// 定义数据库的访问模型
type Faucet struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	ID      int        `sql:"id" json:"id" form:"id"`                // 唯一标识的
	Address   string     `sql:"address" json:"address" form:"address"`       // 请求的地址
	Amount   int       `sql:"amount" json:"amount" form:"amount"`       // 已请求的数量
}
type FaucetNum struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	EtherNum   int       `sql:"etherNum" json:"etherNum" form:"etherNum"`       // 已请求的数量
}
