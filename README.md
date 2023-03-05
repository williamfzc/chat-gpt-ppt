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

Or different render engine:

![](./doc/sample-remark.png)

## Usage

You can download prebuild binaries from [release page](https://github.com/williamfzc/chat-gpt-ppt/releases).

1. Download a executable binary
2. Add your topics to `topic.txt`
3. Add your token (official openai api key, with no extra white space, no empty line) to `token.txt`
4. run `./cgp`

Everything done. You can get some help about command line arguments with `cgp --help`. 

```bash
$ ./cgp_macos --help
Usage of ./cgp_macos:
  -client string
        gpt client type (default "GPT35")
  -output string
        out path (default "./output.html")
  -renderer string
        renderer type (default "REMARK")
  -rendererBin string
        binary file for renderer
  -token string
        token file path (default "./token.txt")
  -topic string
        topic file path (default "./topic.txt")
```

## Contribution

Thanks for your interest. This project is really simple to hack.

This project consists of two pluggable parts:

- Client: Send topics to GPT and get their responses
- Renderer: Build slides from these pairs

If you want to make some changes:

- git clone
- change code
- run `make` to build a binary file (Go installation required)
- check
- push to repo and send a PR

Feel free to send us PR/issues.

## Changelog

> [2023-03-02] Rewrite with golang. All in one file.
>
> [2023-01-13] https://github.com/williamfzc/chat-gpt-ppt/issues/2 OpenAI's services are not available in my country.
>
> [2022-12-06] Currently, ChatGPT has no official API. I am waiting for it to make this repo a real production.

## License

MIT
