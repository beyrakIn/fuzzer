# Fuzzer

``Tool help menu``
```markdown
.__           .__          
|  |__   ____ |  | ______  
|  |  \_/ __ \|  | \____ \ 
|   Y  \  ___/|  |_|  |_> >
|___|  /\___  >____/   __/ 
     \/     \/     |__|    
[!] Usage: ./fuzzer -u https://google.com -w wordlist.txt -o output.txt
[!] Usage: ./fuzzer -u google.com -w wordlist.txt -o output.txt
[!] -u url is required
[!] -w wordlist is required
[!] -o output is optional
[!] -h help
[+] Thanks for using this tool
```


``Example output``
```shell

➜ ./fuzzer -u http://google.com/ -w small.txt
[+] http://google.com/2001
[+] http://google.com/about
[+] http://google.com/account
[+] http://google.com/appliance
[+] http://google.com/blog
...
```

``Options``
```shell
-u: url
-w: wordlist
-o: output
-h: help
```

