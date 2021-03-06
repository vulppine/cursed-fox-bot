# Cursed Fox Bot
*a quickly done bot to create pictures of unsettling foxes, to post on Twitter*

---

## Building

Either:
- Use `go get github.com/vulppine/cursed-fox-bot/bot` in order to get the code, and then put `import "github.com/vulppine/cursed-fox-bot/bot"` in your code
- Use `go install github.com/vulppine/cursed-fox-bot` to install it as an application.

---

## Usage

**Desktop**

You will need to have the following:
- DeepAI key
- Twitter Application Consumer API key
- Twitter Application Consumer Secret key

Set these keys in your environment variables as:
```
DEEP_AIKEY = DeepAI key
TWITTER_APIKEY = Twitter Consumer API key
TWITTER_APISECRET = Twitter Consumer API secret
```

Afterwards, run the program. You will get a link to Twitter, with URL search parameters oauth_token and oauth_verifier. Copy these into these environmental variables:

```
TWUSER_TOKEN
TWUSER_SECRET
```

If successful, you should be able to run the program without error.
A picture of a red fox, generated by DeepAI's text2img API will be posted on your feed.

**Google Cloud Functions**

Get the above keys either through the provided method, or some other method.
Afterwards, upload a copy of the bot/ folder into your Google Cloud Function.
Set the entrypoint to be `GooglePubSubEntryPoint` and run.

---

# Thanks

Scott Ellison Reed and DeepAI for providing the Text to Image API

---

# License

Copyright (c) Flipp Syder under the MIT License
