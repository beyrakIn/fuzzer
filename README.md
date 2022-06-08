# Fuzzer

``Tool help menu``
```shell
[!] Usage: ./fuzzer -u https://google.com/ -w wordlist.txt -o output.txt
[!] Usage: ./fuzzer -u google.com -w wordlist.txt -o output.txt
[!] Usage: -u https://google.com
[!] Usage: -w wordlist.txt
[-] Usage: -o output.txt
[+] Help: -h
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

