/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-04-29
 */
package main

import (
	"log"
	"time"

	tb "github.com/riversgo007/telebot"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "1706227378:AAE1i8R1CgpG_WJd7wr6y_bned_ggt0NJ7o",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/url", func(m *tb.Message) {
		url, err:=b.CreateInviteLink(m.Chat)
		if err!=nil{
			log.Print("get link err:", err.Error())
		}
		log.Print("create invite link:",url)
		b.Send(m.Sender, url)
	})
	b.Start()
}