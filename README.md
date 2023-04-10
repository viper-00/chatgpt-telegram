# ChatGPT-bot

> Interact with ChatGPT

English - [中文](./README_zh.md)

Go CLI to fuels a Telegram bot that lets you interact with [ChatGPT](https://openai.com/blog/chatgpt/), a large language model trained by OpenAI.

## Installation

Clone the repository and navigate into the folder.

```
git clone https://github.com/viper-00/chatgpt-telegram.git

cd chatgpt-telegram
```

Then, open the `env.example` file with a text editor and fill in your credentials, including:

- `TELEGRAM_TOKEN`: Your Telegram Bot token.
  - Follow [this guide](https://core.telegram.org/bots/tutorial#obtain-your-bot-token) to create a bot and obtain the token.
- `TELEGRAM_ID` (Optional): Your Telegram User ID
  - If you set this, only you will be able to interact with the bot.
  - To get your ID, message `@userinfobot` on Telegram.
  - Multiple IDs can be provided, separated by commas.
- `EDIT_WAIT_SECONDS` (Optional): Amount of seconds to wait between edits
  - This is set to `1` by default, but you can increase if you start getting a lot of `Too Many Requests` errors.
- Save the file, and rename it to `.env`.
> **Note** Make sure you rename the file to _exactly_ `.env`! The program won't work otherwise.

Finally, open the terminal in your computer (if you're on Windows, look for `PowerShell`), navigate to the directory where you saved the above file using the "cd" command, and run:

```
go run main.go -c chatgpt.json
```

or build and run:

```
go build -o chatgpt-telegram
./chatgpt-telegram -c chatgpt.json
```

<!-- ### Running with Docker

If you're trying to run this on a server with an existing Docker setup, you might want to use our Docker image instead.

```sh
docker pull ghcr.io/m1guelpf/chatgpt-telegram
```

Here's how you'd set things up with `docker-compose`:

```yaml
services:
  chatgpt-telegram:
    image: ghcr.io/m1guelpf/chatgpt-telegram
    container_name: chatgpt-telegram
    volumes:
      # your ".config" local folder must include a "chatgpt.json" file
      - .config/:/root/.config
    environment:
      - TELEGRAM_ID=
      - TELEGRAM_TOKEN=
```

> **Note** The docker setup is optimized for the Browserless authentication mechanism, described below. Make sure you update the `.config/chatgpt.json` file in this repo with your session token before running. -->

## Authentication

The program use an `API key` to authenticate your ChatGPT account information, so you need to generate a unique key on the offcial ChatGPT site in order to bind your account.

> **Note** If you are not familiar with ChatGPT API keys, you can check the [official site](https://platform.openai.com/account/api-keys) for more information.

After you enter the folder, open the `chatgpt.json.example` file with a text editor and replace the 'your_chatgpt_api_key' text with your generated API key:

```
{ "openaiauthorization": "your_chatgpt_api_key" }
```

Save the file and rename it to `chatgpt.json`.

## License

This repository is licensed under the [MIT License](LICENSE).
