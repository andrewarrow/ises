**It's So Easy Slack** (i.s.e.s)

read and respond to slack messages from a command line

```
go get github.com/andrewarrow/ises
cd that dir
go build
set these env vars
SLACK_TEAMS=2
SLACK_TOKEN_1=xoxp-4422442222-3111111111-11111111118-11aeea211e
SLACK_TOKEN_0=xoxp-4422442222-3111111111-11111111118-11aeea211e

you can get test tokens:

https://api.slack.com/docs/oauth-test-tokens

or make real ones with generate.rb

./ises -s # to sync last 10 messages from every room to disk
./ises -r # to read the most recent messages from every room in every team

touch cache/0_roomname/mute     # this will mute this room from now on
vi cache/1_otherim/say          # puts lines in this file u want to say in room, then sync again
                                # say file is deleted after it's sent
```

Where does ises name come from?

https://github.com/andrewarrow/paradise_ftp

https://github.com/andrewarrow/jungle_smtp

see the pattern?

