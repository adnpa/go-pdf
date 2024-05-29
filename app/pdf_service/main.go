package pdf_service

type UserInfo struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

//func main() {
//	r := gin.Default()
//
//	// 连接 RabbitMQ
//	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	if err != nil {
//		panic(err)
//	}
//	defer ch.Close()
//
//	// 声明交换机和队列
//	q, err := ch.QueueDeclare(
//		"login_success", // name
//		false,           // durable
//		false,           // delete when unused
//		false,           // exclusive
//		false,           // no-wait
//		nil,             // arguments
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	// 缓存已登录用户信息
//	userInfoCache := make(map[string]UserInfo)
//
//	// 消费登录成功消息
//	msgs, err := ch.Consume(
//		q.Name, // queue
//		"",     // consumer
//		true,   // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	go func() {
//		for msg := range msgs {
//			var userInfo UserInfo
//			err := json.Unmarshal(msg.Body, &userInfo)
//			if err != nil {
//				fmt.Printf("failed to unmarshal user info: %v\n", err)
//				continue
//			}
//			userInfoCache[userInfo.UserId] = userInfo
//			fmt.Printf("Received and cached user info: %+v\n", userInfo)
//		}
//	}()
//
//	r.GET("/protected", func(c *gin.Context) {
//		// 检查用户是否已登录
//		userId := c.Query("userId")
//		if _, ok := userInfoCache[userId]; !ok {
//			c.JSON(401, gin.H{"error": "unauthorized"})
//			return
//		}
//
//		// 处理业务逻辑
//		c.JSON(200, gin.H{"message": "Access granted"})
//	})
//
//	r.Run()
//}
