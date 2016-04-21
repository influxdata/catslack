# CatSlack

Inspired by [SlackCat](https://github.com/rlister/slackcat) but rewritten in Go for easy portability.

This is a command line utility to post `stdin` to a configured [Incoming Webhook](https://api.slack.com/incoming-webhooks) on your slack channel. To use it have a script that outputs something useful to `stdin`. Maybe something like `myCoolScript.sh`...

```sh
$ cat myCoolScript.sh
#!/bin/bash
echo "I'm going to be posted in Slack!"

# Set your Slack Incoming Webhook URL:
$ export URL="https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX"

# Experience pleasure by using the catslack!
$ ./myCoolTestScript.sh | catslack -url $URL
```
