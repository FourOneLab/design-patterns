package design_principles

import (
	"encoding/json"
	"strings"
)

// Serialization 类实现了一个简单协议的序列化和反序列功能。

const IdentifierString = "UEUEUE;"

type Serialization struct {
	identifierString string
}

func NewSerialization() *Serialization {
	return &Serialization{identifierString: IdentifierString}
}

func (s *Serialization) Serialize(object map[string]string) (string, error) {
	var builder strings.Builder
	_, _ = builder.WriteString(s.identifierString)
	marshal, err := json.Marshal(&object)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func (s *Serialization) Deserialize(text string) map[string]string {
	if !s.startWith(text) {
		return nil
	}

	res := make(map[string]string)
	err := json.Unmarshal([]byte(text[len(s.identifierString):]), &res)
	if err != nil {
		return nil
	}

	return res
}

func (s *Serialization) startWith(text string) bool {
	header := text[:len(s.identifierString)]
	return header == s.identifierString
}
