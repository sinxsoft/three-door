package app

type Door struct {
	isSelect bool
	target  int /*1表示车，2，表示羊*/
	IsInit bool
}

func (this * Door) GetType() int{
	return this.target
}

func (this *Door) Select() {
	this.isSelect = true
}

func (this *Door) SetType( i int){
	this.target = i
}

