# telegram

## Usage

```lua
local telegram = require("telegram")
bot , err := telegram.new("token")
if err then error(err) end

bot.send(chatId,chatMsg)

```

