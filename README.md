# Command Line Tool Collection: lwabish

## Usage
```bash
# image
lwabish image -h

# k8s
lwabish k8s -h

# all commands
lwabish -h
```

## install
```bash
go install github.com/lwabish/go-snippets@latest
```

## dev
### cobra命令创建
```bash
# 增加command
cobra-cli add image

# 增加sub command
cobra-cli add meta -p 'imageCmd'

# 为command增加flag
# 参考各个command的init函数
```

### 项目初始化
```bash
go install github.com/spf13/cobra-cli@latest

go mod init github.com/lwabish/snippets

cobra-cli init
```
