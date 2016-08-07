**It's So Easy Slack** (i.s.e.s)

read and respond to slack messages from a terminal

```
go get github.com/andrewarrow/ises
cd that dir
go build
set these env vars
SLACK_TOKEN_0=xoxp-4422442222-3111111111-11111111118-11aeea211e
SLACK_TOKEN_1=xoxp-4422442222-3111111111-11111111118-11aeea211e
SLACK_TEAMS=2

you can get test tokens:

https://api.slack.com/docs/oauth-test-tokens

or make real ones with generate.rb
```

Your first time running ./ises it will:

1) download the name and id of each room your are in, in each team

2) download the names of every user in every team

3) download the last 10 messages in every room... in every team

and caches this all to disk.

You can get it to do this again someday with --init option.

The 2nd time you run ./ises it will boot quick and be in conversation mode.

`room_name>type your message here`

is the prompt you see. Press ] to advance a room.

type quit to exit

