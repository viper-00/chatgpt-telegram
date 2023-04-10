# ChatGPT-bot

> 与 ChatGPT 交互

[English](./README.md) - 中文

使用Go CLI来驱动一个Telegram机器人，让你可以与[ChatGPT](https://openai.com/blog/chatgpt/)进行互动，ChatGPT是由OpenAI训练的一个大型语言模型。

## 安装

克隆该仓库并进入文件夹。

```
git clone https://github.com/viper-00/chatgpt-telegram.git

cd chatgpt-telegram
```

然后，使用文本编辑器打开 `env.example` 文件并填写你的凭据，包括：

- `TELEGRAM_TOKEN`: 您的Telegram Bot令牌。
  - 按照[这篇指南](https://core.telegram.org/bots/tutorial#obtain-your-bot-token)创建一个Bot并获得令牌。
- `TELEGRAM_ID`(可选项): 您的Telegram用户ID。
  - 如果设置了此项，则只有您才能与该Bot进行交互。
  - 要获取您的ID，请在Telegram上发送信息 `@userinfobot`。
  - 多个ID可以用逗号分隔提供。
- `EDIT_WAIT_SECONDS` (可选项): 在编辑之间等待的秒数。
  - 默认设置为`1`，但如果您开始遇到很多 `Too Many Requests` 错误，可以增加此值。
- 保存文件，并将其重命名为 `.env`.
> **注意** 确保将文件重命名为 _exactly_ `.env`! 否则程序将无法工作。

最后，打开你的电脑终端（如果你在Windows上，可以找 `PowerShell`), 使用 "cd" 命令导航到保存上述文件的目录，然后运行：

```
go run main.go -c chatgpt.json
```

或者编译后运行：

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

## 认证

程序使用 `API key` 来验证你的 ChatGPT 账户信息，因此你需要在官方 ChatGPT 网站上生成唯一的密钥以绑定你的账户。

> **注意** 果你不知道 ChatGPT API key，你可以查看 [官网](https://platform.openai.com/account/api-keys) 获取更多信息。

设置后进入文件夹，用文本编辑器打开 `chatgpt.json.example` 文件，并将 `your_chatgpt_api_key` 文本替换为生成的API密钥：

```
{ "openaiauthorization": "your_chatgpt_api_key" }
```

保存文件并将其重命名为 `chatgpt.json`。

## 许可证

本代码库使用 [MIT](LICENSE) 许可证。
