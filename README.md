# Insights
> Generates insights for a whatsapp chat.
 
## Usage
```bash
   # if you have go set up, install with the go get
   λ go get -u github.com/umayr/insights/cmd/insights
   # make sure you have $GOBIN defined in your environment
   λ insights -pretty ./path/to/chat/file.txt
   # you can also clone the repository and make the binary yourself
   λ git clone git@github.com:umayr/insights
   λ cd insights
   λ make
   λ ./bin/insights ./path/to/chat/file.txt
   # for different timezones you can use -timezone flag
   λ insights -pretty -timezone=Asia/Dubai ./path/to/chat/file.txt
```
 
## Metrics

- First Message
- Last Message
- Duration
- Frequency (how many messages in every hour of day) 
- Total Messages
- Total Words
- Total Letters
- Average Words Per Message
- Average Letters Per Message
- Average Messages Per Day
- Average Words Per Day
- Average Letters Per Day
- Participants
- Contribution Per Participant
- Contribution Count Per Participant
- Contribution Words Per Participant
- Contribution Letters Per Participant
- Contribution Frequency Per Participant
- Timeline (how many messages per day/week/month/year)
- Timeline Count
- Timeline Words
- Timeline Letters
- Most Active Day
- Most Active Count
- Least Active Day
- Least Active Count
- Emoji Used

## Graph

It also comes with a server that would represent the data extracted from the chat in graphical format it would require `-server` argument during execution, it would look like this:

![](Example.png)

## Contribution

This is a very crude implementation that I cooked up within a day so there are tons of improvements that could be done, please feel free to send a PR or raise an issue if you find anything. 


