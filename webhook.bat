@echo off

curl -H "Accept: application/json" -H "Content-Type:multipart/form-data" -X POST -F "file=@FileName.FileExtension" -F "payload_json={\"content\":\"Message Content\"}" webhook here

pause
