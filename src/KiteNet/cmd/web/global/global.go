package global

import (
	"KiteNet/cmd/web/config"
	"KiteNet/tcp/client"
)

//WorldSS 全局的world tcp会话
var WorldSS *client.Client

var Conf *config.Configs