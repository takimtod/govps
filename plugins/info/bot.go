package info

import (
    //"fmt"
  "inc/lib"
  
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "bot",
    As:       []string{"bot"},
    Tags:     "info",
   // IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {
      m.Reply("Saya Aktif kak")
    },
  })
}
