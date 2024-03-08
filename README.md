# Go

## Packages

## Command line

### Install

```bash
go install github.com/lwabish/go@latest
```

### Usage

[doc](docs/lwabish.md)

## Dev

### Extend

```bash
# 增加command
cobra-cli add image

# 增加sub command
cobra-cli add meta -p 'imageCmd'

# 为command增加flag
# 参考各个command的init函数
```

### Init

```bash
go install github.com/spf13/cobra-cli@latest

go mod init github.com/lwabish/go

cobra-cli init
```
