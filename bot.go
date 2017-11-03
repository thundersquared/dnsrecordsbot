package main

import (
  "log"
  "os"
  "regexp"
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

  command := regexp.MustCompile("/?(start|help)")
  domain := regexp.MustCompile("(https?:?/?/?)?[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9].[a-zA-Z]{2,}(.*)")
  domainType := regexp.MustCompile("(https?:?/?/?)?[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9].[a-zA-Z]{2,}(.*) (.*)")

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }

    if (bot.Debug) {
      log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
    }

    if (command.MatchString(update.Message.Text)) {
      var helper = []string {
        "Hey! I'm the DNS Records Bot.\n",
        "Use /start or /help to see this message.\n",
        "To check a domain's records, send the domain name.",
        "Example: thundersquared.com\n",
        "To check specific records send the domain name followed by the record type.",
        "Example: thundersquared.com A",
      }

      msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Join(helper, "\n"))
      msg.ReplyToMessageID = update.Message.MessageID

      bot.Send(msg)
    }

    if (domain.MatchString(update.Message.Text)) {
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

    if (domainType.MatchString(update.Message.Text)) {
      msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Work in progress!")
      msg.ReplyToMessageID = update.Message.MessageID

      bot.Send(msg)
    }
  }
}
