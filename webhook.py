# Written by tornadus
import requests

webhook = input("Insert Webhook Url: ")
file = input("Insert Path To File: ")

payload = {
    "username": "github.com/tornadus",
    "content": "github.com/Not-Cyrus"
}

with open(file, "rb") as openfile:
    multipart = {
        "file": (openfile.name, openfile, "application/octet-stream")
    }
    requests.post(url=webhook, files=multipart, data=payload)
