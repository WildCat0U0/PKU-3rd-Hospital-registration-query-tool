## 北医三院 挂号查询接口
### 你需要做什么
- Golang 开发环境
- Charles
- 手机

### 步骤
- 手机连接Charles用于抓包
- 打开北医三院挂号查询页面，查询某个科室，如：神经科
- 点击查询按钮，查看Charles抓包，找到请求地址 https://mp.mhealth100.com/gateway/registration/appointment/ 下面的内容
- 查看cookie信息
- 修改代码中的cookie信息，运行程序 如果有号的话，会用不断播放提示音
- 轮询时间和响铃时间可以自行修改# PKU-3rd-Hospital-registration-query-tool
