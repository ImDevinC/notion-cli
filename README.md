# notion-cli
This is a simple application that allows you to add content to a page in Notion.
Note that this is very much an alpha project, use at your own risk!

## Why?
I wanted a very simple way from my command line to send a message to Notion quickly. This gives me that solution

## Usage
1. To start, you need to setup an integration and get your integration key: https://www.notion.so/my-integrations
2. You will also need to make sure to add your integration as a connection on the page you want to edit: https://developers.notion.com/docs/create-a-notion-integration
3. Lastly, get the page ID you want to use by copying the ID out of the url: https://www.notion.so/imdevinc/MyTestPage-7c0ccbd307d34a2fbadd66a717aa5c95
   + In this example, `7c0ccbd307d34a2fbadd66a717aa5c95` is the page ID
4. Then run `notion-cli cfg -t <token> -p <page>`
5. Now simply run `notion-cli <message>` to have your message sent to Notion and stored on your page


