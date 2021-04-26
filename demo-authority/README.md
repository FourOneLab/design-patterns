# 接口鉴权

面向对象分析（OOA）、面向对象设计（OOD）、面向对象编程（OOP），是面向对象开发的三个主要环节。

> 整个过程主要包括，如何做需求分析，如何做职责划分？需要定义哪些类？每个类应该具有哪些属性、方法？类与类之间该如何交互？如何组装类成一个可执行的程序？
> 
> 整个软件开发本来就是一个迭代、修修补补、遇到问题解决问题的过程，是一个不断重构的过程。

## 背景

通过HTTP协议暴露接口的微服务系统，为了保证接口调用的安全性，设计实现一个接口调用鉴权功能，经过认证之后的系统才能调用接口。

存在的问题：

- 需求不明确：需求过于模糊、笼统，不够具体、细化，离落地到设计、编码还有一定的距离；需要通过沟通、挖掘、分析、假设、梳理，搞清楚具体的需求有哪些，哪些是现在要做的，哪些是未来可能要做的，哪些是不用考虑做的。
- 相对于CURD，鉴权难度更高：作为与具体业务无关的功能，可以开发成独立的框架，继承到多个业务系统中；通过框架肯定比业务代码难，代码质量也有更高要求，对需求分析能力、设计能力、编码能力，甚至逻辑思维能力的要求，都是比较高的。

## 需求分析

产出：详细的需求描述。

> 针对框架、组件、类库等非业务系统的开发，一定要有组件化意识、框架意识、抽象意识，开发出来的东西要足够通用，不能局限于单一的某个业务需求，但并不代表可以脱离具体的应用场景，闷头拍脑袋做需求分析。

需求分析的时候，先从简单的方案开始，然后再优化。

### 基础方案

通过用户名加密码来做认证：

1. 允许访问的服务调用方，派发一个应用名（AppID）和一个密钥。
2. 调用方进行接口请求时带上 AppID 和密钥。
3. 服务端接收到请求后解析 AppID 和密钥，并与数据库中保存的数据比对。
4. 一致则认证成功，否则拒绝请求。

### 优化

> 上面方案存在的问题：明文/加密后传输密钥，请求如果被拦截，那么黑客可以发起[重放攻击](https://zh.wikipedia.org/wiki/%E9%87%8D%E6%94%BE%E6%94%BB%E5%87%BB)。

可以借助 OAuth 的思路来解决：

1. 调用方将请求接口的 URL 和 AppID 及密钥拼接在一起，然后使用加密算法生成 Token。
2. 调用方进行请求的时候，将 Token 和 AppID，随 URL 一起发送给服务端。
3. 服务端从请求中获取 AppID，从数据库中获取对应的密钥，使用与步骤1相同的加密算法生成 Token，并与请求中的 Toke 比对。
4. 一致则认证成功，否则拒绝请求。

### 再优化

> 上面方案存在的问题：请求被拦截后，黑客依然会发起重放攻击。

可以通过优化加密算法的方式来解决：

1. 原先的加密算法针对（URL、AppID、密钥）三者加密，现在引入一个随机变量（如，时间戳）再进行加密，将生成的 Token 和时间戳随请求一起发送给服务端。
2. 服务端收到请求后，根据请求中的时间戳先进行判断，超过一定时间窗口（一分钟）的请求，认为 Token 过期了，直接就拒接。
3. Token 没有过期的请求，则使用步骤1中的加密算法生成 Token 进行比对。
4. 一致则认证成功，否则拒绝请求。

### 再再优化

> 上面方案存在的问题：在 Token 过期的时间段内，依然会出现重放攻击的问题。

权衡安全性、开发成本、对系统性能的影响，这个方案算是比较折中、比较合理的了。

> 另一个可以优化的点：调用方的 AppID 及对应密钥的存储位置，保存在数据库（如，MySQL）可能不太好。因为，开发像鉴权这样的非业务功能，最好不要与具体的第三方系统有过度的耦合。

因此，可以提供灵活的配置，来支持不同的存储方式：ZooKeeper、本地配置文件、配置中心、MySQL、Redis等。

在设计和开发的时候留下可扩展的点，保证系统有足够的灵活性和扩展性，在切换存储的时候，也能尽可能少的改动代码。

### 确定需求

- 调用方进行接口请求的时候，将 URL、AppID、密钥和时间戳拼接在一起，通过加密算法生成 Token，并将 Token、AppID 和时间戳拼接在 URL 中，发送给服务端。
- 服务端接收到调用方的请求后，从请求中解析出 Token、AppID和时间戳。
- 服务端先根据时间戳判断 Token 是否过期，如果过期则直接拒接请求。
- 服务端再根据 AppID获 取到存储中密钥，然后使用相同到算法生成一个新 Token，比对两个 Token 是都相同，一致则成功，否则拒接。

## 面向对象设计

产出：将需求描述转化为具体的类。

### 划分职责进而识别出有哪些类

> 类是现实世界中事物的一个建模。并不是每个需求都能映射到现实世界，也并不是每个类都与现实世界中的事物一一对应。对于一些抽象的概念，无法通过映射现实世界中的事物的方式来定义类的。

方法：

1. 把需求描述中的名词罗列出来，作为可能的候选类，然后再进行筛选。
2. 根据需求描述，把其中涉及的功能点，罗列出来，然后再去看哪些功能点职责相近，操作同样的属性，是否应该归为同一个类。

#### 功能点

> 拆解出来的每个功能点要尽可能小，每个功能点只负责做一件很小的事情（**单一职责**）。

1. 把 URL、AppID、密码、时间戳拼接为一个字符串；
2. 对字符串通过加密算法加密生成 token；
3. 将 token、AppID、时间戳拼接到 URL 中，形成新的 URL；
4. 解析 URL，得到 token、AppID、时间戳等信息；
5. 从存储中取出 AppID 和对应的密码；
6. 根据时间戳判断 token 是否过期失效；
7. 验证两个 token 是否匹配；

可以分为三类：

- Token 相关类（`AuthToken`）：1、2、6、7，负责 Token 的生成和验证
- URL 相关类（`Url`）：3、4，负责URL拼接和解析; 为了更通用写， 将 Url 类升级为 ApiRequest 类，因为还会存在 rpc 类型地请求
- 密钥存储相关类（`CredentialStorage`）：5，负责密钥的存储和根据 AppID 获取

> 注意：针对复杂的需求开发，首先进行模块划分，将需求划分成几个小的、独立的功能模块，然后在模块内部，进行面向对象设计。而模块的划分和识别（或者领域驱动设计），跟类的划分和识别，是类似的套路。

### 定义类及其属性和方法

- 属性的识别：将功能点中涉及的名词，作为候选的属性，再进一步过滤筛选。
- 方法的识别：将需求描述中的动词，作为候选的方法，再进一步过滤筛选。

> 从业务模型上来说，不应该属于这个类的属性和方法，不应该被放到这个类里。
>
> 在设计类具有哪些属性和方法的时候，不能单纯地依赖当下的需求，还要分析这个类从业务模型上来讲，理应具有哪些属性和方法。这样可以一方面保证类定义的完整性，另一方面不仅为当下的需求还为未来的需求做些准备。

### 定义类与类之间的交互关系

UML 统一建模语言中定义了六种类之间的关系。它们分别是：泛化、实现、聚合、组合、关联、依赖。

- **泛化**（Generalization）：可以理解为继承关系
- **实现**（Realization）：一般指接口和实现类之间的关系
- 聚合（Aggregation）：包含关系，如 A 类对象包含 B 类对象，B 类对象的生命周期可以不依赖 A 类对象的生命周期，也就是说可以单独销毁 A 类对象而不影响 B 对象，【New A 的时候入参有 B】
- **组合**（Composition）：包含关系，如 A 类对象包含 B 类对象，B 类对象的生命周期依赖 A 类对象的生命周期，B 类对象不可单独存在，【New A 的时候，新 New 一个 B】
- 关联（Association）：一种非常弱的关系，包含聚合、组合两种关系
- **依赖**（Dependency）：一种比关联关系更加弱的关系，包含关联关系，只要 B 类对象和 A 类对象有任何使用关系，都称它们有依赖关系

### 将类组装起来并提供执行入口

可以是一个main()函数，也可以是一组给外部使用的API接口。

这里设计的接口鉴权不是一个独立运行的系统，是一个集成在系统上运行的组建，所以封装所有细节，设计一个最顶层的 ApiAuthenticator 接口类，暴露一组给外部调用者使用的API接口，作为触发执行鉴权逻辑的入口。