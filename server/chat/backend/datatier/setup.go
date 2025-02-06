package datatier

type DataTier struct {
	chatDB *chatDB
}

func NewCreator() *DataTier {
	c := &DataTier{
		chatDB: newChatDB(),
	}

	return c
}

func (x DataTier) NewChatClient() ChatClient {
	return ChatClient{
		chatDB: x.chatDB,
	}
}
