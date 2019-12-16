package util

import (
	"github.com/gin-gonic/gin"
)

//var order = 0

func JsonNested(messageSlice []Message) []gin.H {
	//order++
	var messageJsons []gin.H
	//fmt.Printf("第%d层开始", order)
	//fmt.Println()
	var messageJson gin.H
	for _, messages := range messageSlice {
		//fmt.Println("分解过程", messages)
		message := *messages.ChildMessage
		//fmt.Println("分解过程的的子留言", message)
		if messages.ChildMessage != nil {
			messageJson = gin.H{
				"user_id":         messages.Username,
				"message":         messages.Message,
				"ChildrenMessage": JsonNested(message),
			}
		} else {
			messageJson = gin.H{
				"user_id": messages.Username,
				"message": messages.Message,
				"ChildrenMessage":"null",
			}
		}
		messageJsons = append(messageJsons, messageJson)
	}
	//fmt.Printf("第%d层结束。", order)
	//fmt.Println()
	//order--
	return messageJsons
}
