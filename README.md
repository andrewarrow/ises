=It's So Easy Slack (i.s.e.s)=

read and respond to slack messages from arch linux

in a terminal

with no window manager

but a nice console font

```
from real history:
  444  ./ises                                    # lists all conversations in all teams
  445  ./ises -t 0 -id C0K5LNG2H                 # list message history from that id
  458  ./ises -p 1                               # list all people in all teams
  487  ./ises -t 0 -id U0FRP3RCN -s test         # say test in that place
```

to run:

```
go get github.com/andrewarrow/ises
cd that dir
go build
./ises
set these env vars
SLACK_TEAMS=2
SLACK_TOKEN_1=xoxp-4422442222-3111111111-11111111118-11aeea211e
SLACK_TOKEN_0=xoxp-4422442222-3111111111-11111111118-11aeea211e

you can get tokens:

https://api.slack.com/docs/oauth-test-tokens
```


```
future ideas:
/ises u - unread
/ises t - tail last file messages of x
/ises s - speak
/ises e - edit last message
```

Where does ises name come from?

https://github.com/andrewarrow/paradise_ftp
https://github.com/andrewarrow/jungle_smtp
see the pattern?

