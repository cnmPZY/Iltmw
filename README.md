## 引言

因为川同学不太喜欢记单词，所以想让计算机记单词

本项目只供于学习使用，不得用来作弊（对，就是这样）

## 使用

* 访问 [我爱记单词](https://skl.hduhelp.com/#/english/list)
* 登录 (DD)
* 访问下面的这个页面：可以获得token
* ![image-20240523175731353](https://echin-h.oss-cn-hangzhou.aliyuncs.com/img/image-20240523175731353.png)

* token=48df6d44-13f5-4fa1-98ae-dec33343dab8

* 进入main.go，将token粘贴进去，同时把week(第几周)和mode(0为自测，1为考试)

* 本次AI使用的是百度智能云

* 进入api/work.go

    * Endpoint: https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ernie_speed

    * apiKey: 自己创建一个

    * secretKey: 自己创建一个

      不知道是什么就自己GPT

  ````go
  // 本人的设置，希望别滥用谢谢
  const (
  	apiKey      = "xcqrCYyJ76uSFqeHgb3i0IDw"
  	secretKey   = "renAUjjDoeTHjok6ViPS06ClP5yIIj4w"
  	apiEndpoint = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/ernie_speed"
  	)
  ````

* ```go
  go run main.go
  ```

## 注意： AI有风险，使用需谨慎
因为是AI的原因，通用性不是很高，可能会遇到一些关于序列化和反序列化或者越界啥的一堆问题

但是个人现在的配置是百分百的准确率的

所以如果你就考得太低，嗯，肯定不是我的问题....

## 总结

1. 这个东西其实只是川同学无聊写的
2. 学习了一些关于抓包的技巧
3. 学习了http包的使用
4. 测试，debug，接入ai，等等的学习
5. 接口json格式的接通
