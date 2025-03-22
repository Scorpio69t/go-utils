
# go-utils

**go-utils** 是一个用 Go 语言编写的实用工具库，封装了一系列常用功能，帮助开发者更高效地开发和维护 Go 项目。通过模块化的设计，`go-utils` 提供了常用的加解密、HTTP 请求封装、邮件发送、字符串处理、文件操作、数据结构转换等功能，帮助开发者更快地构建稳定的 Go 应用程序。

## 🌟 项目特色
- 🔒 支持常见的加密解密算法（AES、RSA、SHA 等）
- 🌐 简单封装的 HTTP 请求（GET、POST、PUT、DELETE）
- 📧 支持 SMTP 协议的邮件发送
- 🧰 实用的字符串和文件操作
- 🧪 覆盖完整的单元测试和基准测试
- 📖 详细的文档和示例

## 🗂️ 目录结构规划
```plaintext
go-utils/
├── .github/                   # GitHub 配置（PR模板、Issue模板等）
├── benchmark/                 # 基准测试
├── cmd/                       # 可执行程序示例
├── config/                    # 配置文件和环境变量解析
├── docs/                      # 项目文档
├── examples/                  # 具体的示例代码
├── internal/                  # 内部模块，非导出
├── pkg/                       # 工具包（公共导出模块）
│   ├── crypto/                # 加解密工具
│   ├── curl/                  # HTTP 请求封装
│   ├── email/                 # 邮件发送模块
│   ├── file/                  # 文件操作
│   ├── logger/                # 日志封装
│   ├── strutil/               # 字符串工具
│   ├── timeutil/              # 时间处理
├── scripts/                   # 构建和部署脚本
├── test/                      # 测试数据和工具
├── .gitignore                 # Git 忽略规则
├── go.mod                     # Go 依赖管理文件
├── go.sum                     # Go 依赖管理摘要文件
├── LICENSE                    # 许可证文件
├── README.md                  # 项目简介
└── Makefile                   # 构建和测试命令
```

## ✅ 单元测试和基准测试
> 单元测试覆盖率目标：**≥90%**

### 📋 单元测试示例（crypto）：
```go
func TestAesEncrypt(t *testing.T) {
    key := []byte("1234567890123456")
    data := "Hello, World!"
    encrypted, err := EncryptAES(data, key)
    if err != nil {
        t.Fatalf("EncryptAES failed: %v", err)
    }

    decrypted, err := DecryptAES(encrypted, key)
    if err != nil {
        t.Fatalf("DecryptAES failed: %v", err)
    }

    if decrypted != data {
        t.Errorf("Expected %s, got %s", data, decrypted)
    }
}
```

### 🚀 基准测试示例（crypto）：
```go
func BenchmarkAesEncrypt(b *testing.B) {
    key := []byte("1234567890123456")
    data := "Hello, World!"
    for i := 0; i < b.N; i++ {
        EncryptAES(data, key)
    }
}
```

## 🚀 目标
- 🔥 完整的工具集成封装
- 💡 可扩展、易于维护
- 🚀 适用于个人和企业项目  
