package module

type person struct {
	name   string
	age    int
	status string
}

func NewSturct() *person {
	return &person{
		name:   "小红",
		age:    18,
		status: "原始信息",
	}
}
