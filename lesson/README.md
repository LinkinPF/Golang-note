# lesson : Golang template包学习

## 学习内容

1. 使用模板传递一个变量、结构体、map，以及同时传递结构体和map
2. 模板使用自定义函数和模板嵌套
3. 修改模板的默认标识符，html/template和text/template的区别，html/template可以在模板渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击。方式就是通过在模板中自定义函数，之后在模板中{{. | function }}就可以了。