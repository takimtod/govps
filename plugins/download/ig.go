package download

import (
  "inc/lib"
  "net/http"
   "encoding/json"
  "net/url"
  "fmt"
  "strings"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(ig|instagram)",
    As:       []string{"instagram"},
    Tags:     "downloader",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      resp, err := http.Get("https://skizo.tech/api/igdl?url="+url.QueryEscape(m.Querry)+"&apikey=batu")

      if strings.Contains(m.Querry, "https://www.instagram.com/reel/") {
      type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, url := range media {
         
          bytes, err := client.GetBytes(url)
          if err != nil {
             fmt.Println("Error:", err)
            return
          }
          client.SendVideo(m.From, bytes, caption, m.ID)
             
        }

       } else if strings.Contains(m.Querry, "https://www.instagram.com/p/") {
        type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          m.Reply(err.Error())
          return
        }
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, ur := range media {
          bytes, err := client.GetBytes(ur)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          client.SendVideo(m.From, bytes, caption, m.ID)
           client.SendImage(m.From, bytes, caption, m.ID)
             
        }   

      } else if strings.Contains(m.Querry, "https://www.instagram.com/stories/") {
        type respon struct {
          Caption string   `json:"caption"`
          Media   []string `json:"media"`
        }
        if err != nil {
            fmt.Println("Error:", err)
            return
          }
          defer resp.Body.Close()
          var data respon
          err = json.NewDecoder(resp.Body).Decode(&data)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          // Mengambil media
            caption := data.Caption
          media := data.Media
          for _, ur := range media {
            bytes, err := client.GetBytes(ur)
            if err != nil {
              m.Reply(err.Error())
              return
            }
            client.SendVideo(m.From, bytes, caption, m.ID)
             client.SendImage(m.From, bytes, caption, m.ID)

          }   
          
      }
      
    },
  })
}