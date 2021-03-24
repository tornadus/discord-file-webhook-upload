use reqwest;
use std::fs;
use std::io;

#[tokio::main]
async fn main() {
    let mut webhook = String::new();
    let mut file_path = String::new();

    println!("Enter your webhook: ");
    io::stdin()
    .read_line(&mut webhook)
    .expect("Failed to read the webhook");

    println!("Enter the file you want to upload (include the path): ");
    io::stdin()
    .read_line(&mut file_path)
    .expect("Failed to read the filepath");


    let url = format!("{}",webhook);

    let data = fs::read(file_path.clone().trim()).unwrap();

    let part = reqwest::multipart::Part::bytes(data).file_name(file_path);

    let form = reqwest::multipart::Form::new()
    .text("content","test")
    .part("file",part);

    reqwest::Client::new()
    .post(&url)
    .multipart(form)
    .send()
    .await
    .expect("Failed to send");
}
