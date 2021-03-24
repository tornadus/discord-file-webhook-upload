-- basically the same as batch but you really think I'm gonna write a whole Lua script just for this? LOL

local File = "FileName.FileExtension"
local Webhook = "webhook"

os.execute(string.format("curl -H \"Accept: application/json\" -H \"Content-Type:multipart/form-data\" -X POST -F \"file=@%s\" %s",File,Webhook))