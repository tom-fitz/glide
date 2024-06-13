package schemas

import (
	"sync"
)

var chatRequestPool = sync.Pool{
	New: func() interface{} {
		return &ChatRequest{}
	},
}

var chatResponsePool = sync.Pool{
	New: func() interface{} {
		return &ChatResponse{}
	},
}

// GetChatRequest get objects from the pool
func GetChatRequest() *ChatRequest {
	return chatRequestPool.Get().(*ChatRequest)
}

// ReleaseChatRequest release objects from the pool
func ReleaseChatRequest(req *ChatRequest) {
	*req = ChatRequest{}
	chatRequestPool.Put(req)
}

// GetChatResponse get objects from the pool
func GetChatResponse() *ChatResponse {
	return chatResponsePool.Get().(*ChatResponse)
}

// ReleaseChatResponse release objects from the pool
func ReleaseChatResponse(res *ChatResponse) {
	*res = ChatResponse{}
	chatResponsePool.Put(res)
}

var chatStreamMessagePool = sync.Pool{
	New: func() interface{} {
		return &ChatStreamMessage{}
	},
}

var chatStreamRequestPool = sync.Pool{
	New: func() interface{} {
		return &ChatStreamRequest{}
	},
}

// GetChatStreamMessage gets objects from the pool
func GetChatStreamMessage() *ChatStreamMessage {
	return chatStreamMessagePool.Get().(*ChatStreamMessage)
}

// ReleaseChatStreamMessage releases objects back to the pool
func ReleaseChatStreamMessage(msg *ChatStreamMessage) {
	*msg = ChatStreamMessage{}
	chatStreamMessagePool.Put(msg)
}

// GetChatStreamRequest gets objects from the pool
func GetChatStreamRequest() *ChatStreamRequest {
	return chatStreamRequestPool.Get().(*ChatStreamRequest)
}

// ReleaseChatStreamRequest releases objects back to the pool
func ReleaseChatStreamRequest(req *ChatStreamRequest) {
	*req = ChatStreamRequest{}
	chatStreamRequestPool.Put(req)
}
