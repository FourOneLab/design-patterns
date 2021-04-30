package design_principles

type Serializations struct {
}

func (s *Serializations) Serialize(obj interface{}) string {
	return ""
}

func (s *Serializations) Deserialize(str string) interface{} {
	return nil
}

// ------------------------------
// 拆分为更小的类

type Serializer struct{}

func (s *Serializer) Serialize(obj interface{}) string {
	return ""
}

type Deserializer struct{}

func (d *Deserializer) Deserialize(str string) interface{} {
	return nil
}

// ------------------------------
// 引入接口，实现高内聚和LOD原则

type Serializers interface {
	Serialize(obj interface{}) string
}

type Deserializers interface {
	Deserialize(str string) interface{}
}

type DefaultSerialization struct{}

func (s *DefaultSerialization) Serialize(obj interface{}) string {
	return ""
}

func (d *DefaultSerialization) Deserialize(str string) interface{} {
	return nil
}
