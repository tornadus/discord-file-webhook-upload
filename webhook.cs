using System;
using System.IO;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading.Tasks;

namespace AAA
{
    class Program
    {
        static async Task<int> Main(string[] args)
        {
            Console.WriteLine("Enter the webhook URL: ");
            string webhookURL;
            webhookURL = Console.ReadLine();

            Console.WriteLine("Enter the file path: ");
            string filePath;
            filePath = Console.ReadLine();


            HttpClient httpClient = new HttpClient();

            var form = new MultipartFormDataContent();
            var fileContent = new ByteArrayContent(await File.ReadAllBytesAsync(filePath));

            fileContent.Headers.ContentType = MediaTypeHeaderValue.Parse("multipart/form-data");
            form.Add(fileContent, "file", Path.GetFileName(filePath));

            await httpClient.PostAsync(webhookURL, form);

            return 0;
        }
    }
}
