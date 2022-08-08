# github.com/lwabish/go-snippets

## 命令行入口创建
项目初始化
```bash
go install github.com/spf13/cobra-cli@latest
go mod init github.com/lwabish/snippets
cobra-cli init
```

cobra
```bash
# 增加command
cobra-cli add image
# 增加sub command
cobra-cli add meta -p 'imageCmd'
# 为command增加flag
# 参考meta command
```
