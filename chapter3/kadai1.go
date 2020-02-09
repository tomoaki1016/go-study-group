package chapter3

// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

func (o Kadai1) ID() int {
	return o.id
}

func (o *Kadai1) SetID(v int) {
	o.id = v
} 

func (o Kadai1) Name() string {
	return o.name
}

func (o *Kadai1) SetName(v string) {
	o.name = v
}
