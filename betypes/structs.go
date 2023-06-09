package betypes

type BotMessage struct {
	Message struct {
		Message_Id int
		From       struct {
			Username string
			Id       int
		}
		Chat struct {
			Id int
		}
		Text string
	}
}

type BotSendMessageID struct {
	Result struct {
		Message_id int
	}
}

type Photos struct {
	Entries []struct {
		FullPath string
		Mime     string
	}
}
