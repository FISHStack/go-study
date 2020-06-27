
# 运行
go run select.go

# 笔记

```
for i := 0; i < 2; i++ {
		//阻塞接收 一接收到信息就会退出
		select {
		case msg1 := <-c1:
			fmt.Println("recived:", msg1)
		case msg2 := <-c2:
			fmt.Println("recived:", msg2)
		}
	}
```

select会进行一次阻塞直到从chan中接收到消息,然后进入下一次循环,接收第二条消息

而c1会睡眠1s发送消息，c2会睡眠2s发送消息，预期新效果就是各接收到一条消息，然后退出。