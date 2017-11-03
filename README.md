<p align="center">
  <img src="media/Slick Icon@2x.png" width="128" />
  <h3 align="center">dnsrecordsbot</h3>
  <p align="center">A Telegram bot to fetch DNS records</p>
  <p align="center">
    <a href="https://t.me/dnsrecordsbot" target="_blank">
      <img src="media/Button@2x.png" width="128" />
    </a>
  </p>
</p>

## A bot for what?
dnsrecordsbot is a bot heavily inspired by Spatie's project [dnsrecords.io](https://dnsrecords.io), which allows you to retrieve all DNS records of a domain name.

## Tech stack
The bot is written in Go, relies on [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) to consume Telegram's Bot API and loves dig command that helps gathering DNS data.

## How to build
1. Clone the repo
   ```
   git clone this repo
   ```
2. Install packages
   ```
   go get -d
   ```
3. Build
   ```
   go build bot.go
   ```
4. Run
   ```
   BOT_TOKEN="YOUR-BOT-TOKEN" ./bot
   ```

## Environment vars
| Var | Value | Required |
| --------- | ---- | --- |
| BOT_TOKEN | none | Yes |
| DEBUG     | none | No  |

## License
The code in this repo and used modules are open-sourced software licensed under the [MIT license](LICENSE.md).
