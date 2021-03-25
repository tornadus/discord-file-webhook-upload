using HTTP: HTTP,Form,Multipart;

println("Enter the webhook URL: ")

webhookURL = readline()

println("Enter the file path: ")

filePath = readline()

println("Enter the message content: ")

messageContent = readline()


part = Multipart(filePath,open(filePath, "r"))

res = HTTP.post(webhookURL, Form(["file" => part,"payload_json" => "{\"content\":\"$messageContent\"}"]))
