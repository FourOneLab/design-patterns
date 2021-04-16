package object_oriented

// Golang 模拟继承总觉得不太像，可真尴尬。

// AbstractBird 鸟类的抽象类，细分的鸟都继承这个类。
type abstractBird struct{}

// Fly 给鸟类定一个飞的方法，但是并不是所有的鸟都会飞啊，
//     如，下面的鸵鸟。
func (b abstractBird) Fly() {}

type Ostrich struct {
	abstractBird
}

// Fly 重写Fly方法，表示鸵鸟不会飞，索然能解决问题，但是不够优雅。
//     因为不会飞的鸟很多，如，企鹅，每一种都重写Fly方法么？
func (o Ostrich) Fly() string {
	return "Sorry, I can't Fly."
}

// ---------------------------------------------

// 使用接口解决继承的问题，因为接口只声明来方法，每一个类都要自己去实现方法具体的逻辑，
// 比如下蛋的逻辑可能都是一样的，就会造成代码的重复。

type Flyable interface {
	Fly()
}

type Tweetable interface {
	Tweet()
}

type EggLayable interface {
	LayEgg()
}

// NewOstrich 避免和上面的鸵鸟冲突
type NewOstrich struct{}

func (o NewOstrich) Tweet() {
	panic("implement me")
}

func (o NewOstrich) LayEgg() {
	panic("implement me")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	panic("implement me")
}

func (s Sparrow) Tweet() {
	panic("implement me")
}

func (s Sparrow) LayEgg() {
	panic("implement me")
}

// 为了解决接口实现中，类之间重复代码的问题，
// 为每一个接口分别写一个实现类，然后通过组合和委托的方式来解决。

type FlyAbility struct{}

func (f FlyAbility) Fly() {
	panic("implement me")
}

type TweetAbility struct{}

func (t TweetAbility) Tweet() {
	panic("implement me")
}

type EggLayAbility struct{}

func (e EggLayAbility) LayEgg() {
	panic("implement me")
}

type AnotherOstrich struct {
	// 组合两种能力，这样的话，
	// 这个类已经实现类上面的两个接口
	// 可以覆盖，也可以再这两种能力的基础上增加或修改
	TweetAbility
	EggLayAbility
}

// ---------------------------------------------

// 没有继承关系，使用组合更好的业务场景

type Url struct{}

// Join URL 拼接
func (u Url) Join() {}

// Split URL 分割
func (u Url) Split() {}

type Crawler struct {
	Url
}

type PageAnalyzer struct {
	Url
}
