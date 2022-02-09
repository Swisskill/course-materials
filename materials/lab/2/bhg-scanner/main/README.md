In here? try > 
- go build
- time ./main

1) The code is a port scanner. You can provide an address (followed by a colon)
    as well as a range of what ports you want to look through. 
2) Modification to the code: 
    Todo 1: Added my name and some detail on how to run the code
    Todo 2: Replaced the original timing mech with dialTimeout
    Todo 3: Added an int that tracked closed ports. Chose an int instead of a slice to save memory
    Todo 4: No changes. 100 is a fine number of channels.
    Todo 5: Added an TUI that can take an address, along with a range of what ports to scan
    Todo 6: Returned info on ports scanned for testing and output

3) To run, you can either do 
        cd C:\Users\wrbra\Desktop\COSC\Cyber\course-materials\materials\lab\2\bhg-scanner\main\
        go build
        ./main
    And then follow the prompts or
        cd C:\Users\wrbra\Desktop\COSC\Cyber\course-materials\materials\lab\2\bhg-scanner\scanner
        go test