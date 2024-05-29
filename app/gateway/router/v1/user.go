package v1

func NewUserClient() protoUser.UserService {
	// 创建服务
	service := grpc.NewService()

	// 创建客户端
	userClient := protoUser.NewUserService("go.micro.srv.demo", service.Client())

	return userClient
}

//client := NewUserClient()
//
//// rpc 调用远程服务的方法
//resp, err := client.Login(context.TODO(), &protoUser.LoginRequest{Email: param.Email, Password: param.Password})
//if err != nil {
//  fmt.Println(err)
//}
