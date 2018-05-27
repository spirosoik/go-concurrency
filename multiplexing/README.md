# Multiplexing 

The core idea is to not make a sequencing conversation but whosoever is ready to talk, should be able to do it.

So now goroutines can independently talk without Blocking and the FanIn channel will receive all the messages in order to process them. Check the following rough diagram:

John ---
       \
        ----- FanIn --- Independent Messages beeing sent
       /
Doe ---
