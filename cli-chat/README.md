### Open two CLI tools

**In Alice's terminal**

```bash
cd ~/go/src/github.com/stefanhans/cli-chat/cli-chat/

go build
./cli-chat alice

```


**In Bob's terminal**

```bash
cd ~/go/src/github.com/stefanhans/cli-chat/cli-chat/

./cli-chat alice

```


### Start the chat interactively

**Create your memberlist node**

- `memberlistconfigure` creates a memberlist configuration
- `memberlistcreate` creates the memberlist specified by the configuration


**Get, and possibly join, the bootstrap peers**

- `bootstrapjoin` joins calling peer to bootstrap peers


**Join the memberlist**

- `memberliststart` starts broadcasting between the members
- `memberlistjoin` joins bootstrap peers to memberlist


**Join the chat and say hi**

- `chatjoin` start chat listener and join the chat
- `msg hi` send a message to all chat members


**See all available commands**

- `<tab><tab>` lists all commands using code completion
- `help` or other not existing command shows commands and descriptions


**Leave the cli-tool**

- `quit` 

