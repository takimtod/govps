package convert

import (
  "inc/lib"
"fmt"
  "inc/lib/api"
  "inc/lib/typings"
  "math/rand"
  "time"
    "net/url"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "qc",
    As:       []string{"qc"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  false,
     IsWaitt:  true,
     IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {
  
      warna := []string{"000000", "FFFFFF", "c27ba0", "bcbcbc", "1f2c34"}
      rand.Seed(time.Now().UnixNano())
      index := rand.Intn(len(warna))
      hasil := warna[index]
      res := "https://skizo.tech/api/qc?text="+url.QueryEscape(m.Querry)+"&username="+m.PushName+"&avatar=https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png?q=60&apikey=batu&hex="+ hasil
      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      
      s := api.StickerApi(&typings.Sticker{
        File: bytes,
        Tipe: func() typings.MediaType {
          if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
            return typings.IMAGE
          } else if m.IsVideo || m.IsQuotedVideo {
            return typings.VIDEO
          } else {
            return typings.TEKS
          }
          
        }(),
      }, &typings.MetadataSticker{
        Author:    m.PushName,
        Pack:      "https://s.id/ryuubot",
        KeepScale: true,
        Removebg:  "true",
        Circle: func() bool {
          if m.Querry == "-c" {
            return true
          } else {
            return false
          }
        }(),
      })

      client.SendSticker(m.From, s.Build(), m.ID)
    },
  })
}
