# 🌱 Grass

[![GitHub license](https://img.shields.io/github/license/levinion/grass)](https://github.com/levinion/grass/blob/main/LICENSE)
![GitHub stars](https://img.shields.io/github/stars/levinion/grass?style=social)

**Grass** 是一个用于进程管理的命令行工具。🚀

## 特性

- 💡 基于 Unix Socket 的 IPC 通信，实现高效的进程间通信。
- ⚙️ 支持进程的启动、停止和状态监控，方便管理和控制。
- ➕ 具备动态添加和移除进程管理的能力，灵活适应不同需求。
- 🖥️ 提供简洁易用的命令行界面，方便操作和管理进程。

## 安装

```bash
go install github.com/levinion/grass@latest
```

## 使用方法

### 启动 Grass 服务

通过以下命令启动 Grass 服务：
```
grass
```

### 添加进程进行管理

使用以下命令将进程添加到管理中：
```
grass add <进程名称>
```

### 停止指定进程 

通过以下命令停止指定的进程：
```
grass stop <进程名称>
```

### 启动指定进程

通过以下命令启动指定的进程：
```
grass start <进程名称> 
```

### 重新加载进程

通过以下命令重新加载指定的进程：
```
grass reload <进程名称> 
```

### 移除进程管理

通过以下命令移除进程管理：
```
grass remove <进程名称>
```

### 查看所有进程状态

使用以下命令查看所有进程的状态：
```
grass show
```

### 在 Hyprland 中运行

在 Hyprland 中，可以通过以下命令运行 Grass：
```
exec-once = grass
exec-once = grass add "fcitx5 -d" "mako" "hyprpaper" "waybar"
```

## 授权许可

该项目使用 MIT 授权协议，详细信息请查看 LICENSE 文件。

## 贡献

欢迎提出问题（Issue）或贡献代码（Pull Request）参与贡献。🤝