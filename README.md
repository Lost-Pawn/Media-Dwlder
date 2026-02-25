# Media Downloader Telegram Bot

A Telegram bot written in Go that downloads media from multiple platforms using yt-dlp and gallery-dl.  
It supports direct link downloads as well as inline mode within Telegram chats.

The bot processes user-submitted URLs, determines the appropriate downloader backend, retrieves the media, and sends the file back using the Telegram Bot API.

---
## How to Operate
### 1. Build the Application
Make sure Go is installed.

```bash
go mod tidy
go build -o bot
```
---
### 2. Configure Environment Variable
Set your Telegram bot token:
Linux / Mac:
```bash
export BOT_TOKEN=your_telegram_bot_token
```
Windows (PowerShell):
```powershell
setx BOT_TOKEN "your_telegram_bot_token"
```
---
### 3. Run the Bot
```bash
./bot
```
On Windows:
```bash
bot.exe
```
---
### 4. Download via Direct Message
Send a supported media link to the bot in a private chat.  
The bot will download the media and send the file back to you.
---
### 5. Use Inline Mode
Enable inline mode via @BotFather.
In any Telegram chat, type:
```
@yourbotusername <media_link>
```
Select the result and send it directly in the chat.
---

Note: 1. Supported platforms depend on yt-dlp and gallery-dl compatibility.
      2. Deliverables size must be under 25-50mb as the source uses gv3
