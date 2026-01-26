# vidpovid-bot-go

Telegram-bot for communication with AI-person. It can transcribe voice messages.

## Prepare

```bash
cp .env.example .env
```

And then edit `.env` file.

## Run Telegram-bot

```bash
make run
```

## Building

Build binary:

```bash
make build
```

Clean:

```bash
make clean
```

# Run with docker

Build image:

```bash
docker build -t vidpovid-bot-go .
```

Run:

```bash
docker run --name vidpovid-bot-go --rm \
-e "TELEGRAM_BOT_TOKEN=<TELEGRAM_BOT_TOKEN>" \
-e "OPENAI_API_KEY=<OPENAI_API_KEY>" \
-e "OPENAI_MODEL=gpt-4.1" \
-e "OPENAI_ASSISTANT_MESSAGE=\"Відповідай львівський батяр\"" \
vidpovid-bot-go
```
