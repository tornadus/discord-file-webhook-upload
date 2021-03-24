# Written by tornadus
import requests

webhook = input("Insert Webhook Url: ")
file = input("Insert Path To File: ")

payload = {
    "username": "Test",
    "content": "lol"
}

with open(file, "rb") as openfile:
    multipart = {
        "file": (openfile.name, openfile, "application/octet-stream")
    }
    requests.post(url=webhook, files=multipart, data=payload)
