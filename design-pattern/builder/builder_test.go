package builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	//首先新建一个ConcreteBuilder
	builder := NewConcreteBuilder()
	//之后把创建好的ConcreteBuilder给Director，因为ConcreteBuilder实现了接口
	director := NewDirector(&builder)
	//开始构造
	director.Construct()
	result := builder.GetResult()
	fmt.Println(result.built)
}


