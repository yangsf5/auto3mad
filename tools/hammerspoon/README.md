Hammerspoon 配置管理

```sh
cd ~
ln -s ~/code/go/src/github.com/yangsf5/auto3mad/tools/hammerspoon .hammerspoon
```

- Hammerspoon 的 init.lua 等配置脚本在本文件夹。
- 官方指定的 ~/.hammerspoon 是本文件夹的软链。
- 这样安排，是方便在这里直接进行 Git 版本管理，不再需要进行拷贝同步。
