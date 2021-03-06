# 面向过程

面向过程是一种编程范式/风格，以过程（方法、函数、操作）作为组织代码的基本单元，以数据（成员变量、属性）与方法相分离为最主要的特点。面向过程是一种流程化的风格，通过拼接一组顺序执行的方法来操作数据完成一项功能。

面向过程编程语言最大的特点是不支持类和对象两个语法概念，不支持面向对象的特性，仅支持面向过程编程。

> 面向过程和面向对象最基本的区别就是，代码的组织方式不同。
>
> - 面向过程风格的代码被组织成了一组方法集合及其数据结构，方法和数据结构的定义是分开的。
> - 面向对象风格的代码被组织成一组类，方法和数据结构被绑定一起，定义在类中。

## 面向对象比面向过程的优势

### OOP更能应对大规模复杂的程序开发

在需求简单的时候，整个程序的处理流程只有一条主线，面向过程就是这种流程化、线性的思维方式。

但是复杂程序的处理流程是一个网状结构，因此面向对象的编程以类为思考对象，将复杂流程拆分为一个个方法，先思考业务建模，将需求翻译为类，再建立类之间的关系，这种开发模式和思考方式，让我们能更好的应对复杂程序的开发，思路更清晰。

> 类是一种非常好地组织函数和数据结构的方式，是一种将代码模块化的有效手段。

面向过程的编程语言也可以写出面向对象风格的代码，只是付出的代价更高。

### OOP更易复用、扩展和维护

面向过程是一种非常简单的风格，没有面向对象提供的丰富特性，而面向对象的这些特性能极大地满足复杂的编程需求。

- 封装（易维护性）：将数据和方法绑定，通过访问权限控制，只允许外部调用者通过类暴露的有限方法访问数据，而不是面向过程那种任意方法随意修改。
- 抽象（可扩展性）：函数就是一种抽象，隐藏具体细节，基于接口的抽象，在不修改现有实现的情况下，轻松替换新地实现逻辑。
- 继承（复用性）：将相同的属性和方法抽取出来。
- 多态（可扩展性）：在实际代码运行中调用新的逻辑而不用修改原有代码，遵从对修改关闭，对扩展开放的设计原则。

### OOP更人性化，更高级，更智能

```shell
二进制指令 ---> 汇编语言 ---> 面向过程编程 ---> 面向对象编程
```

与机器交互的越多，需要越计算机思维的编程语言，与人交互的越多，需要越人类思维的编程语言。

交互方式从计算机思维转变为人类思维，编程语言的发展只会越来越智能化。

## 面向对象中写面向过程的代码

### 滥用getter和setter

见：[`bad_case.go`](bad_case.go)

### 全局变量

在面向对象编程中，常见的全局变量有单例类对象、静态成员变量、常量等。常见的全局方法有静态方法。

- 单例类对象：在全局代码中只有一份，相当于全局变量
- 静态成员变量：归属于类上的数据，被所有的实例化对象所共享，相当于一定程度上的全局变量
- 常量：最常见的全局变量

见：[`bad_global_var.go`](bad_global_var.go)

### 全局方法

静态方法：在不用创建对象的情况下，拿来就用，静态方法将方法与数据分离，破坏类封装特性，是典型的面向过程风格。

在 Golang 中，全局方法表现为 util 包。

> util 包出现的背景：业务上A和B类之间没有继承关系，但是又有相同地处理逻辑，把这些共同地处理逻辑放在 util 包中。

这就是一种面向过程的编程风格，可以有效地解决代码复用的问题。 util 这样的包，不是不用而是不能滥用。 util 类设计的时候，也可以根据不同的功能细分为不同的 utils，如 FileUtils、IOUtils 等。

### 定义数据和方法分离

基于 MVC 进行 Web 开发是会出现这样的问题。传统的 MVC 结构分为 Model 层、Controller 层、View 层这三层。

前后端分离后，三层结构在后端开发中，调整为 Controller 层、Service 层、Repository 层。

- Controller 层负责暴露接口给前端调用
- Service 层负责核心业务逻辑
- Repository 层负责数据读写

在每一层中，会定义相应的 VO（View Object）、BO（Business Object）、Entity。一般情况下，VO、BO、Entity 中只会定义数据，不会定义方法，所有操作这些数据的业务逻辑都定义在对应的 Controller 类、Service 类、Repository 类中。

这就是典型的面向过程的编程风格。这种开发模式叫作基于贫血模型的开发模式，也是现在非常常用的一种 Web 项目的开发模式。
