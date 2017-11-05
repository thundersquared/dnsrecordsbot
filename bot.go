package main

import (
  "log"
  "os"
  "regexp"
  "strings"

  "./dns"

  "gopkg.in/telegram-bot-api.v4"
)

func stringInSlice(a string, list []string) bool {
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
}

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

  QueryCommand := regexp.MustCompile("(?i)((\\w+\\.?)+)\\s?((@(\\w+\\.?)+)?(A{1,4}|NS|SOA|MX|TXT|DNSKEY)?)")

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }

    if (bot.Debug) {
      log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
    }

    // Help commands
    if (update.Message.IsCommand()) {
      if (update.Message.Command() == "start" || update.Message.Command() == "help") {
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
    } else if (QueryCommand.MatchString(update.Message.Text)) {
      query := strings.ToLower(update.Message.Text)
      match := QueryCommand.FindAllStringSubmatch(query, -1)

      // Fetch dig data
      dig, err := dns.Dns(match[0][1])

      if err == nil {
        var message string

        isRecType := stringInSlice(strings.ToUpper(match[0][3]), dig.RecordTypes)
        isFrom := strings.HasPrefix(match[0][3], "@")

        if (isRecType && isFrom) {
          records := dig.GetRecordsFrom(match[0][3])
          message = strings.Join(records, "\n")
        } else if (isRecType) {
          message = dig.GetRecordsOfType(strings.ToUpper(match[0][3]))
        } else if (isFrom) {
          records := dig.GetRecordsFrom(match[0][3])
          message = strings.Join(records, "\n")
        } else {
          records := dig.GetRecords()
          message = strings.Join(records, "\n")
        }

        // Join and reply
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
        msg.ReplyToMessageID = update.Message.MessageID

        bot.Send(msg)
      }
    } else {
      text := "Looks like an unsupported markup. Did you check /help?"
      msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
      msg.ReplyToMessageID = update.Message.MessageID
      bot.Send(msg)
    }
  }
}
