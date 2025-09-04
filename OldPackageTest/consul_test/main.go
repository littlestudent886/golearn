package main

import "github.com/hashicorp/consul/api"

func Register(address string, port int, name string, tags []string, id string) (err error) {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8501"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Address = address
	registration.Port = port
	registration.Tags = tags

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://10.233.4.60:9021/health",
		Interval:                       "5s",
		Timeout:                        "3s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	_ = Register("10.233.4.60", 9021, "user-web", []string{"user-web"}, "user-web")
}
