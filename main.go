package main

import (
	"fmt"
	"net"
)

const (
	SOCKET_NETWORK string = "tcp4"
	SOCKET_ADDR string = ":8101"
)



func main () {
	conn, err := net.Listen(SOCKET_NETWORK, SOCKET_ADDR)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	for {
		fmt.Println("等待客户端连接")
		c, err := conn.Accept()
		fmt.Println("客户连接上了")
		if err != nil {
			fmt.Println(err)
		}

		go handleData(c) //耗时10分钟
	}
}

// Message /** 前端发送来的消息类型
type Message struct {
	MsgType string `json:"msg_type" form:"msg_type"` //消息类型 决定显示样式
	/**
		text 文本消息类型
		emoji 表情消息类型
		location 位置消息类型
		image 图片消息类型
	 */
	MsgData string `json:"msg_data" form:"msg_data"` //消息正文 决定显示内容
}

type MsgResp struct {
	Code int8 `json:"code" form:"code"` //code = 0 success code != 0 fail
	Msg string `json:"msg" form:"msg"`
}

/**
	消息数据类型 JSON string
	消息结构 Message
 */
func handleData (conn net.Conn) *MsgResp {
	buf := make([]byte, 1024)
	//resp := &MsgResp{}
	for {
		conn.Read(buf)
		fmt.Println(string(buf))
		//if err != nil {
		//	resp.Code = -1
		//	resp.Msg = err.Error()
		//	fmt.Println(err)
		//	return resp
		//}
		//
		//msg := &Message{}
		//err = json.Unmarshal(buf, msg)
		//if err != nil {
		//	resp.Code = -2
		//	resp.Msg = err.Error()
		//	fmt.Println(err)
		//	return resp
		//}
		//
		//
		//
		//if msg.MsgType == "text" {
		//	return resp
		//}
	}

}
