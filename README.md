# chat-gpt-ppt

Use ChatGPT to generate PPT automatically, all in one single file.

## Showcase

1. Some topics for presentation named `topic.txt`:

```
what's OpenAI?
how OpenAI works?
what is the future of OpenAI?
```

2. Save your openai token to `token.txt`.
3. Generate a ppt in seconds:

```
./cgp
```

And you get one:

![](./doc/sample.png)

With multi languages support:

![](./doc/sample-chi.png)

## Usage

You can download prebuild binaries from [release page](https://github.com/williamfzc/chat-gpt-ppt/releases).

1. Download a executable binary
2. Add your topics to `topic.txt`
3. Add your token to `token.txt`
4. run `./cgp`

Everything done.

## Contribution

Feel free to send us PR/issues.

## Changelog

> [2023-03-02] Rewrite with golang. All in one file.
>
> [2023-01-13] https://github.com/williamfzc/chat-gpt-ppt/issues/2 OpenAI's services are not available in my country.
>
> [2022-12-06] Currently, ChatGPT has no official API. I am waiting for it to make this repo a real production.

## License

MIT
