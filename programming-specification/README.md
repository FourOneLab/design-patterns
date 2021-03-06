# Programming Specification

## Naming and Comments

命名的好坏，对于代码的可读性来说非常重要，甚至可以说是起决定性作用的。命名能力也体现了一个程序员的基本编程素养。

> 对于影响范围比较大的命名，比如包名、接口、类名，一定要反复斟酌、推敲。实在想不到好名字的时候，可以去 GitHub 上用相关的关键词联想搜索一下，看看类似的代码是怎么命名的。

### 命名

命名的一个原则就是以能准确达意为目标。命名的时候，一定要学会换位思考，假设自己不熟悉这块代码，从代码阅读者的角度去考量命名是否足够直观。

#### 命名长度

- 用很长的命名方式，觉得命名一定要准确达意，哪怕长一点也没关系：尽管长的命名可以包含更多的信息，但是导致语句过长会影响代码的可读性
- 用短的命名方式，能用缩写就尽量用缩写：在足够表达其含义的情况下，命名当然是越短越好。
  - 对于一些默认的、大家都比较熟知的词，推荐用缩写
  - 对于作用域比较小的变量，使用相对短的命名
  - 对于类名这种作用域比较大的，推荐用长的命名方式

#### 利用上下文简化命名

如下面的例子：

```go
package main

// UserV1 简化前
type UserV1 struct {
 UserName      string
 UserPassword  string
 UserAvatarURL string
}

// UserV2 简化后
type UserV2 struct {
 Name      string
 Password  string
 AvatarURL string
}


// 在使用这些属性时候，我们能借助对象这样一个上下文，表意也足够明确。
func demo(){
  user := new(UserV2)
  _ = user.Name
}
```

除了类之外，函数参数也可以借助函数这个上下文来简化命名。

```go
package main

// UploadUserAvatarImageV1 简化前
func UploadUserAvatarImageV1(userAvatarImageUrl string){}

// UploadUserAvatarImageV2 简化后
func UploadUserAvatarImageV2(imageUrl string){}
```

#### 命名要可读，可搜索

- 可读：不要用一些特别生僻、难发音的英文单词来命名。
- 可搜索：在 IDE 中编写代码，经常会用“关键词联想”的方法来自动补全和搜索。所以，在命名的时候，最好能符合整个项目的命名习惯：
  - 大家都用“selectXXX”表示查询，就不要用“queryXXX”
  - 大家都用“insertXXX”表示插入一条数据，就要不用“addXXX”

统一规约是很重要的，能减少很多不必要的麻烦。

#### 命名接口和抽象类的方法

- 接口：一般有两种比较常见的方式。
  - 一种是加前缀“I”，表示一个 Interface。比如 IUserService，对应的实现类命名为 UserService。
  - 另一种是不加前缀，比如 UserService，对应的实现类加后缀“Impl”，比如 UserServiceImpl。
- 抽象类：有两种方式，
  - 一种是带上前缀“Abstract”，比如 AbstractConfiguration；
  - 另一种是不带前缀“Abstract”。

对于接口和抽象类，选择哪种命名方式都是可以的，只要项目里能够统一就行。

### 注释

命名很重要，注释跟命名同等重要。命名再好，毕竟有长度限制，不可能足够详尽，而这个时候，注释就是一个很好的补充。

#### 如何写注释

注释的目的就是让代码更容易看懂。只要符合这个要求的内容，你就可以将它写到注释里。总结一下，注释的内容主要包含这样三个方面：做什么、为什么、怎么做。

如下面的示例：

```go
package main

// BeansFactory 
// (what) Bean factory to create beans.
//
// (why) The class likes Spring IOC framework, but is more lightweight.
//
// (how) Create objects from different sources sequentially:
// user specified object > SPI > configuration > default object.
//
type BeansFactory struct {
 // ...
}
```

- 注释比代码承载的信息更多：命名的主要目的是解释“做什么”，函数和变量命名得好，可以不用在注释中解释它是做什么的。但是，对于类来说，包含的信息比较多，一个简单的命名就不够全面详尽了。这个时候，在注释中写明“做什么”就合情合理了。
- 注释起到总结性作用、文档的作用：代码之下无秘密。阅读代码可以知道如何实现，但是在注释中，关于具体的代码实现思路，可以写一些总结性的说明、特殊情况的说明。这样能够让阅读代码的人通过注释就能大概了解代码的实现思路，阅读起来就会更加容易。**对于有些比较复杂的类或者接口，我们可能还需要在注释中写清楚“如何用”**。
- 一些总结性注释能让代码结构更清晰：对于逻辑比较复杂的代码或者比较长的函数，如果不好提炼、不好拆分成小的函数调用，那我们可以借助总结性的注释来让代码结构更清晰、更有条理。

#### 注释的量

注释太多和太少都有问题。

- 太多，有可能意味着代码写得不够可读，需要写很多注释来补充。除此之外，注释太多也会对代码本身的阅读起到干扰。而且，后期的维护成本也比较高，代码改了，注释忘了同步修改，就会让代码阅读者更加迷惑。
- 太少，如果代码中一行注释都没有，那只能说明这个程序员很懒，要适当督促添加一些必要的注释。

一般来说，类和函数/方法一定要写注释，而且要写得尽可能全面、详细，而函数内部的注释要相对少一些，一般都是靠好的命名、提炼函数、解释性变量、总结性注释来提高代码的可读性。

## Code Style

说起代码风格，很难说哪种风格更好。最重要的，也是最需要做到的，是在团队、项目中保持风格统一，让代码像同一个人写出来的，整齐划一。这样能减少阅读干扰，提高代码的可读性。这才是在实际工作中应该实现的目标。

### 类和函数的体量

- 类或函数的代码行数太多，一个类上千行，一个函数几百行，逻辑过于繁杂，阅读代码的时候，很容易就会看了后面忘了前面。
- 类或函数的代码行数太少，在代码总量相同的情况下，被分割成的类和函数就会相应增多，调用关系就会变得更复杂，阅读某个代码逻辑的时候，需要频繁地在 n 多类或者 n 多函数之间跳来跳去，阅读体验也不好。

> 对于函数代码行数的最大限制，网上有一种说法，不要超过一个显示屏的垂直高度，如果要让一个函数的代码完整地显示在 IDE 中，那最大代码行数不能超过 50。这个说法挺有道理的，超过一屏之后，在阅读代码时，为了串联前后的代码逻辑，就需要频繁地上下滚动屏幕，阅读体验不好还容易出错。

通过反向例子来说明类的大小：

- 当一个类的代码读起来让你感觉头大了，
- 实现某个功能时不知道该用哪个函数了，
- 想用哪个函数翻半天都找不到了，
- 只用到一个小功能要引入整个类（类中包含很多无关此功能实现的函数）的时候，

这就说明类的行数过多了。

### 一行代码的长度

Google 的 Java 规范中说，一行代码最长限制为 100 个字符。

不同的编程语言、不同的规范、不同的项目团队，对此的限制可能都不相同。

总体原则是：一行代码最长不能超过 IDE 显示的宽度。需要滚动鼠标才能查看一行的全部代码，显然不利于代码的阅读。当然，这个限制也不能太小，太小会导致很多稍长点的语句被折成两行，也会影响到代码的整洁，不利于阅读。

### 善用空行

对于比较长的函数，如果逻辑上可以分为几个独立的代码块，在不方便将这些独立的代码块抽取成小函数的情况下，为了让逻辑更加清晰：

- 用总结性注释
- 用空行来分割各个代码块

在类的成员变量与函数之间、各成员变量之间、各函数之间，可以通过添加空行的方式，让这些不同模块的代码之间，界限更加明确。

**写代码就类似写文章，善于应用空行，可以让代码的整体结构看起来更加有清晰、有条理**。

### 关于缩进的争论

不同编程语言，对于缩进的规范也不尽相同，只要项目内部能够统一就行了。

与业内推荐的风格统一、跟著名开源项目统一。如 Golang 自带的 `gofmt` 工具一样，保证了全球的 Go 代码都是格式统一的。

> 值得强调的是，不管是用两格缩进还是四格缩进，**一定不要用 tab 键缩进**。
>
> 在不同的 IDE 下，tab 键的显示宽度不同，有的显示为四格缩进，有的显示为两格缩进。如果在同一个项目中，不同的同事使用不同的缩进方式（空格缩进或 tab 键缩进），有可能会导致有的代码显示为两格缩进、有的代码显示为四格缩进。

### 大括号的位置

不同编程语言，对于打括号起始位置的规范也不尽相同。在 Golang 中，对大括号的起始位置有这严格的要求，将括号放在和语句同一行的位置。

### 类中成员的排列顺序

Golang 中提供 `goimport` 等工具，对导入的类库进行排序，通常是按照字母顺序从小到大排列的。

在类中，相关性强的成员变量放在一起，先可导出变量（首字母大写），再内部变量（首字母小写）的顺序。

## Coding Tips

### 把代码分割成更小的单元块

通常阅读代码的习惯是：先看整体再看细节。

写代码的时候要有模块化和抽象思维，善于将大块的复杂逻辑提炼成类或者函数，屏蔽掉细节，这样阅读代码不至于迷失在细节中，能极大地提高代码的可读性。

通常代码逻辑比较复杂的时候，才建议提炼类或者函数。如果提炼出的函数只包含两三行代码，在阅读代码时，还要跳转反倒增加了阅读成本。

如[示例](contrast.go)中 Invest 函数。

### 避免函数参数过多

函数的参数一般在5个以内，太多就会影响代码的可读性了，使用起来也不方便。针对函数参数过多的情况，一般有2种处理方法：

- 考虑函数是否**职责单一**，是否能通过拆分成多个函数的方式来减少参数
- 将函数的参数封装成对象

> 如果函数是对外暴露的远程接口，将参数封装成对象，还可以提高接口的兼容性。在往接口中添加新的参数时，老的远程接口调用者有可能就不需要修改代码来兼容新的接口了。

### 勿用函数参数来控制逻辑

不要在函数中使用布尔类型的标识参数来控制内部逻辑：

- true：走这块逻辑
- false：走另一块逻辑

这明显违背了**单一职责原则**和**接口隔离**原则，建议将其拆成两个函数，可读性上也要更好。

如果是内部函数，影响范围有限，或者拆分之后的两个函数经常同时被调用，可以酌情考虑保留标识参数。

除了布尔类型作为标识参数来控制逻辑的情况外，还有一种“根据参数是否为 nil”来控制逻辑的情况。针对这种情况，也应该将其拆分成多个函数。拆分之后的函数职责更明确，不容易用错。

### 函数设计要职责单一

- 针对类、模块这样的应用对象可以使用单一职责原则。
- 对于函数的设计，更要满足单一职责原则。

相对于类和模块，函数的粒度比较小，代码行数少，所以在应用单一职责原则的时候，没有像应用到类或者模块那样模棱两可，能多单一就多单一。

### 移除过深的嵌套层次

代码嵌套层次过深往往是因为 if-else、switch-case、for 循环过度嵌套导致的。

通常，嵌套最好不超过两层，超过两层之后就要思考一下是否可以减少嵌套。过深的嵌套本身理解起来就比较费劲，除此之外，嵌套过深很容易因为代码多次缩进，导致嵌套内部的语句超过一行的长度而折成两行，影响代码的整洁。

解决嵌套过深的方法也比较成熟，有下面 4 种常见的思路。

1. 去掉多余的 if 或 else 语句。
2. 使用编程语言提供的 continue、break、return 关键字，提前退出嵌套。
3. 调整执行顺序来减少嵌套。
4. 将部分嵌套逻辑封装成函数调用，以此来减少嵌套。

如[示例](contrast.go)所示。

除此之外，常用的还有通过使用多态来替代 if-else、switch-case 条件判断的方法。这个思路涉及代码结构的改动。

### 学会使用解释性变量

常用的用解释性变量来提高代码的可读性的情况有下面 2 种。

- 常量取代魔法数字
- 使用解释性变量来解释复杂表达式

如[示例](contrast.go)所示。

**最后，非常重要的就是，项目、团队，甚至公司，一定要制定统一的编码规范，并且通过 Code Review 督促执行，这对提高代码质量有立竿见影的效果。**
