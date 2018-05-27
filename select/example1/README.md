# Select

A handy way to handle multiple channels. It's like a switch but each step/case is a communication of a channel:
- All channels are evaluated
- Selection blocks until one communication can proceed, which then does
- Random selected if multiple channels can proceed simultaneously
- Default clause, if present, executes immediately if no channel is ready. If there is not then the select blocks forever till a channel can proceed.

So now we have a single go routine and we use select.