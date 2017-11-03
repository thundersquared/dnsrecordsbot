package main

import (
  "log"
  "os"
  "strings"

  "./dns"
  
  "gopkg.in/telegram-bot-api.v4"
)

func main() {
  bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
  if err != nil {
    log.Panic(err)
  }

  bot.Debug = os.Getenv("DEBUG") == "true"

  log.Printf("Authorized on account %s", bot.Self.UserName)

  _, err = bot.RemoveWebhook()
  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }

    if (bot.Debug) {
      log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
    }

    // Fetch dig data
    d := dns.Dns(update.Message.Text)
    records := d.GetRecords()

    /*for _, element := range records {
      fmt.Printf("%s\n", element)
    }*/

    // Join and reply
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(records, "\n"))
    msg.ReplyToMessageID = update.Message.MessageID

    bot.Send(msg)
  }
}
