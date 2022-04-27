package relay

import (
	logs "github.com/danbai225/go-logs"
	"go-rustdesk-server/common"
	"go-rustdesk-server/model/model_proto"
	"google.golang.org/protobuf/proto"
)

func Start() {
	go registered()
	common.NewMonitor("tcp", ":21117", handlerMsg).Start()
}
func registered() {
	//dial, err := net.Dial("tcp", ":21116")
	//if err != nil {
	//	logs.Err(err)
	//	return
	//}
	//dial.Write(model_proto.RegisterPk{})
}
func handlerMsg(messageData []byte, write func(data []byte) error) {
	message := model_proto.Message{}
	err := proto.Unmarshal(messageData, &message)
	if err != nil {
		logs.Err(err)
	}
	logs.Info(message.Union)
}