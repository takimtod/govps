package lib

import (
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type Event struct {
	WA *whatsmeow.Client
}

type ICommand struct {
	Name        string
	As          []string
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	IsMedia     bool
	IsQuerry    bool
	IsGroup     bool
	IsAdmin     bool
	IsBotAdmin  bool
	IsWaitt     bool
	IsPrivate   bool
	After       func(client *Event, m *IMessage)
	Exec        func(client *Event, m *IMessage)
}

type IMessage struct {
  Info			      types.MessageInfo
	From            types.JID
	IsBot           bool
	Sender          types.JID
	OwnerNumber     []string
	PushName        string
	IsOwner         bool
	IsGroup         bool
	Querry          string
	Body            string
	Command         string
	IsImage         bool
	IsVideo         bool
	IsQuotedImage   bool
	IsQuotedVideo   bool
	IsQuotedSticker bool
	IsAdmin         bool
	IsBotAdmin      bool
	Media           whatsmeow.DownloadableMessage
	ID              *waProto.ContextInfo
	QuotedMsg       *waProto.ContextInfo
	Reply           func(text string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error)
}
