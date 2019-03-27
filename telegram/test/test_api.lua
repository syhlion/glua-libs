local tg = require("telegram")

local token = ""
local chatId = 0
local message = "helloworld"
bot,err = tg:send(token)
if err then error(err) end
print(err)

msg,err = bot:send(chatId,message)

print(msg,err)



