# design-patterns
The beauty of design patterns.

## 可维护性

- 代码易维护：在不破坏原有代码设计、不引入新的Bug的情况下，能够快速地修改或添加代码
- 代码不易维护：修改或者添加代码需要冒着极大地引入新Bug的风险，并且需要话费很长的时间才能完成

>  易维护的代码：分层清晰、模块化好、高内聚低耦合、遵从基于接口而非实现编程的设计原则。

## 可读性

代码是否易读、易理解，代码的可读性很大程度上影响代码的可维护性。

> 可读性好的代码：符合编码规范、命名达意、注释详尽、函数长短合适、模块划分清晰、符合高内聚低耦合。

## 可扩展性

代码应对未来需求变化的能力（在不修改或少量修改原有代码的情况下，通过扩展的方式添加新的功能代码），很大程度上决定了代码的可维护性。

> 扩展性好的代码：预留一些功能扩展点，可以把新功能代码直接插到扩展点上，而不需要因为要添加一个新功能而大动干戈，改动大量的原有代码。

## 简洁性

KISS原则："Keep It Simple，Stupid"。

> 简洁的代码：代码简单、逻辑清晰，也就意味着易读、易维护。

## 可复用性

DRY原则："Don't Repeat Yourself"。

减少重复代码的编写，复用已有的代码。

> 可复用的代码：
> - 继承、多态存在的目的是提高代码的复用性
> - 单一职责原则与代码复用性相关
> - 解耦、高内聚、模块化能提高代码复用性

## 可测试性

单元测试难不难写，能很好的衡量可测试性。

## 灵活性

符合上面6个特性的代码就是灵活的代码。

----------

## 面向对象

面向对象编程因为具有丰富的特性（封装、抽象、继承、多态），可以实现复杂的设计思路，是很多设计原则和设计模式编码实现的基础。

> 面向对象编程是一种编程范式/风格，以类或对象作为组织代码的基本单元，并将封装、抽象、继承、多态作为代码设计和实现的基石。
> 
> 只要某种编程语言支持类或对象的语法概念，并且以此作为组织代码的基本单元，就可以粗略地认为它是面向对象编程语言。

## 设计原则

设计原则是指导代码设计的经验总结，比较抽象，了解设计的初衷，能解决的编程问题和应用场景。对于某些场景下，是否应该应用某种设计模式具有指导意义。

- SOLID 原则：SRP 单一职责原则
- SOLID 原则：OCP 开闭原则
- SOLID 原则：LSP 里式替换原则
- SOLID 原则：ISP 接口隔离原则
- SOLID 原则：DIP 依赖倒置原则
- DRY 原则、KISS 原则、YAGNI 原则、LOD 法则

## 设计模式

设计模式是针对软件开发中经常遇到的一些设计问题，总结出的一套解决方案或者设计思路。大部分设计模式要解决的都是代码的**可扩展性**问题。了解能解决哪些问题，掌握典型应用场景，不过度应用。

## 编程规范

编程规范主要解决的是代码的**可读性**问题，相对于设计原则，设计模式，更具体、更偏重代码细节，更能够落地。偏重于记忆，在编码过程中照做即可，持续的小重构依赖的理论基础就是编程规范。

## 代码重构

> 软件在不停迭代，就没有一劳永逸的设计，需求变更，代码堆砌，原有的设计必然会存在问题。针对这些问题需要进行代码重构，保证代码质量。

代码重构的工具就是面向对象设计思想、设计原则、设计模式、编码规范。开发初期一定不要过度设计，在出现问题后，更具原则和模式进行重构。保证重构不出错的技术：单元测试和代码的可测试性。