<?php
    // make sure your uploaded file content name is "file"
    $webhookURL = "webhookURL"; // set your webhook URL here
    $MessageContent = "github.com/Not-Cyrus"; // put content here or whatever

    function createFile($TmpFile) {
        return [
            "file" => $TmpFile["tmp_name"],
            "name" => $TmpFile["name"],
        ];
    }

    if ($_SERVER["REQUEST_METHOD"] == "POST") {

        if (isset($_FILES["file"])) {

            $File = createFile($_FILES["file"]);

            $curl = curl_init($webhookURL);
            curl_setopt_array($curl,[
                CURLOPT_POST => 1,
                CURLOPT_HTTPHEADER => ["Content-Type: multipart/form-data"],
                CURLOPT_POSTFIELDS => [
                    "payload_json" => json_encode([
                        'content' => $MessageContent,
                    ]),
                    curl_file_create($File["file"],null,$File["name"]),
                ],
                CURLOPT_SAFE_UPLOAD => 1,
                CURLOPT_RETURNTRANSFER => 1,
            ]);

            $Response = curl_exec($curl);
            curl_close($curl);
            
           die($Response); // you're welcome to read the response
        }

        die("You need to upload a file"); // when they don't upload a file
    }

    die("Invalid HTTP method (Required: POST)")
?>
