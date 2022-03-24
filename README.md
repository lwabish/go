# github.com/lwabish/snippets

## 命令行入口创建
```bash
go install github.com/spf13/cobra-cli@latest
go mod init github.com/lwabish/snippets
cobra-cli init

# 增加command
cobra-cli add image
# 增加sub command
cobra-cli add meta -p 'imageCmd'
```
