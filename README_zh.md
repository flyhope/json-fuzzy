# json-fuzzy 中文文档

[English Documentation](README.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/flyhope/json-fuzzy.svg)](https://pkg.go.dev/github.com/flyhope/json-fuzzy)

一个用于模糊JSON反序列化的Go语言库，能够优雅地处理类型不匹配的情况。该包扩展了标准`encoding/json/v2`的功能，提供宽松的解析选项和自定义反序列化器，自动在不同JSON类型之间进行转换。

## 特性

- **模糊类型转换**：在反序列化时自动在不同JSON类型间转换：
  - 数字、布尔值和null值转换为字符串
  - 字符串、布尔值和null值转换为整数和其他数值类型
  - 各种字符串表示转换为布尔值
  - 字符串和数字转换为浮点值
- **宽松的JSON解析**：包含默认的宽松解析选项：
  - `AllowDuplicateNames`：允许JSON对象中存在重复的字段名
  - `AllowInvalidUTF8`：处理字符串中的无效UTF-8序列
- **无缝集成**：提供与标准库相似的API，可以作为直接替代品使用
- **可扩展设计**：可以轻松添加或修改自定义的反序列化器

## 模糊反序列化模式

该库提供两种可配置的反序列化模式：

- **FuzzyUnmarshalerOrigin**：处理基本类型转换（字符串、数字、布尔值），开销最小。适用于性能关键且不使用`any`类型的场景。
  
- **FuzzyUnmarshalerFull**（默认）：在基础模式上通过`FuzzyAny`反序列化器扩展对`any`类型的支持。提供处理异构数据的最大灵活性，但有轻微的性能开销（基准测试显示约10-15%的额外开销）。

### 修改全局默认行为

可以在程序初始化时通过修改`fuzzy.FuzzyUnmarshaler`变量来改变全局默认行为：

```go
package main

import (
    "github.com/flyhope/json-fuzzy/fuzzy"
)

func init() {
    // 全局切换到基础模式以获得更好的性能
    fuzzy.FuzzyUnmarshaler = fuzzy.FuzzyUnmarshalerOrigin
}

// 现在所有的Unmarshal调用都将默认使用基础模式
func main() {
    var result struct{ Value int }
    err := jsonfuzzy.Unmarshal([]byte(`{"Value":"42"}`), &result)
}
```

> **注意**：此操作仅应在程序初始化阶段进行，因为并发修改是不安全的。

### 代码示例

```go
package main

import (
    "github.com/flyhope/json-fuzzy"
    "github.com/flyhope/json-fuzzy/fuzzy"
)

func main() {
    // 使用基础模式(FuzzyUnmarshalerOrigin)以获得更好的性能
    // 当不需要'any'类型支持时
    basicOpts := jsonfuzzy.WithOptions(fuzzy.FuzzyUnmarshalerOrigin())
    var basicResult struct{ Value int }
    err := jsonfuzzy.Unmarshal([]byte(`{"Value":"42"}`), &basicResult, basicOpts)
    
    // 使用完整模式(默认)支持'any'类型
    var fullResult struct{ Value any }
    err = jsonfuzzy.Unmarshal([]byte(`{"Value":true}`), &fullResult)
    // 等效于:
    // err = jsonfuzzy.Unmarshal([]byte(`{"Value":true}`), &fullResult, fuzzy.FuzzyUnmarshalerFull())
}
```

## 安装

```bash
go get github.com/flyhope/json-fuzzy
```

## 使用

该库提供了与标准库JSON API相似的函数，但具有模糊解析能力：

```go
package main

import (
    "github.com/flyhope/json-fuzzy"
)

func main() {
    type testStruct struct {
        StringField string  `json:"string_field"`
        IntField    int     `json:"int_field"`
        FloatField  float64 `json:"float_field"`
        BoolField   bool    `json:"bool_field"`
    }

    // 包含类型不匹配的JSON输入
    jsonData := []byte(`{
        "string_field": 123,
        "int_field":    "456",
        "float_field":  "78.9",
        "bool_field":   "true"
    }`)

    var result testStruct
    err := jsonfuzzy.Unmarshal(jsonData, &result)
    // 即使类型与结构体字段不匹配，也能成功解析
}
```

## API

- `Unmarshal(data []byte, v any, opts ...json.Options) error`：使用模糊类型转换反序列化JSON数据
- `Marshal(v any, opts ...json.Options) ([]byte, error)`：将Go值序列化为JSON（标准行为）
- `MarshalString(v any, opts ...json.Options) (string, error)`：将Go值序列化为JSON字符串
- `WithOptions(opts ...json.Options) json.Options`：将提供的选项与默认的模糊选项组合
- `DefaultOptions() json.Options`：返回包含模糊反序列化器和宽松解析设置的默认选项

## 许可证

MIT License