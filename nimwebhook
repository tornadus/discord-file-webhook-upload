import httpclient

var client = newHttpClient()
var data = newMultipartData()

write(stdout, "Insert File Location > ")

var file = readLine(stdin)

write(stdout, "Insert Webhook URL > ")

var url = readLine(stdin)

data.addFiles({file: file})

echo client.postContent(url, multipart=data)
