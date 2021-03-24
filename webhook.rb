require 'net/http'
require 'uri'

puts "Enter the webhook URL"

webhookURL = gets.chomp

puts "Enter the file path"

filePath = gets.chomp

uri = URI(webhookURL)

request = Net::HTTP::Post.new(uri,"Content-Type" => "multipart/form-data")

form_data = [["file",File.open(filePath)]]


request.set_form form_data, "multipart/form-data"

response = Net::HTTP.start(uri.host, uri.port, use_ssl: true) do |http|
    http.request(request)
end
