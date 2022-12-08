package simple_factory

import "fmt"

// 简单工厂模式
// 1. 当一个类不知道它所必须创建的对象的类的时候
// 2. 当一个类希望由它的子类来指定它所创建的对象的时候
// 3. 当类将创建对象的职责委托给多个帮忙子类的中的某一个，并且你希望将哪一个帮助子类是代理者者一信息局部化的时候

// ParseConfig
type ParseConfig interface {
	Parse(file string) string
}

// NewParseConfig
func NewParseConfig(t string) ParseConfig {
	if t == "ini" {
		return &parseIni{}
	} else if t == "yaml" {
		return &parseYaml{}
	} else if t == "json" {
		return &parseJson{}
	}
	return nil
}

type parseIni struct{}

func (*parseIni) Parse(file string) string {
	return fmt.Sprintf("解析ini配置")
}

type parseYaml struct{}

func (*parseYaml) Parse(file string) string {
	return fmt.Sprintf("解析yaml配置")
}

type parseJson struct{}

func (*parseJson) Parse(file string) string {
	return fmt.Sprintf("解析json配置")
}
