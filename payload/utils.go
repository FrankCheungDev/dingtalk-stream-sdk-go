package payload

import (
	"fmt"
	"github.com/google/uuid"
)

/**
 * @Author linya.jj
 * @Date 2023/4/7 15:13
 */

const (
	DataFrameHeaderKTopic       = "topic"
	DataFrameHeaderKContentType = "contentType"
	DataFrameHeaderKMessageId   = "messageId"
	DataFrameHeaderKTime        = "time"

	DataFrameContentTypeKJson   = "application/json"
	DataFrameContentTypeKBase64 = "base64String"

	DataFrameResponseStatusCodeKOK              = 200
	DataFrameResponseStatusCodeKInternalError   = 500
	DataFrameResponseStatusCodeKHandlerNotFound = 404

	BotMessageCallbackTopic = "/v1.0/im/bot/messages/get" //机器人消息统一回调topic
)

func GenerateMessageId(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.New().String())
}
