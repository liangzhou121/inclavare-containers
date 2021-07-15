# 可重复构建测试平台
为了方便，做RBI测试，镜像构建等的脚本集合
## 文件说明
*   `rbi.sh` 脚本主入口了，具体的功能可以执行`./rbi.sh help`查看
*   `kata-agent/` 和kata-agent相关的镜像构建、执行测试、拉代码仓库等脚本  

## 操作说明
### kata-agent可重复构建
先构建RBCI
```bash
./rbi.sh agent-image
```
测试目录为`/path/to/kata-containers`源码的RB
```
./rbi.sh agent-local /path/to/kata-containers
```
测试git源码
```bash
./rbi.sh agent-git
```
上述两个命令，均将会以二进制文件+报告的形式将结果输出到`report`下
删除RBCI
```bash
./rbi.sh agent-image
```
清除所有临时文件
```bash
./rbi clean
```


# Reproducible Build Infrastructure
For convenience, some scripts are collected for automatically building.
## Files
*   `rbi.sh` main script. Use `./rbi.sh help` to see details.
*   `kata-agent/` scripts related to RB of kata-agent.

## Instructions
### RB for kata-agent
Firstly, build RBCI(Reproducible Build Container Image) for kata-agent
```bash
./rbi.sh agent-image
```
Check the reproducibility of source code in `/path/to/kata-containers`.
```
./rbi.sh agent-local /path/to/kata-containers
```
Or, check the reproducibility of source code from github.com.
```bash
./rbi.sh agent-git
```
Above 2 operations can both produce a report and an artifest in `report/`.

Delete RBCI for kata-agent
```bash
./rbi.sh agent-image
```

Clean all tempfiles
```bash
./rbi clean
```