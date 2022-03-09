# smuggler
A tool for detecting HTTP request smuggling vulnerabilities

Example test on Portswigger labs:
```
$ ./smuggler -urls targets.test 
https://ac291f541e8bb695c0338e2700cd00fb.web-security-academy.net/ is vulnerable
https://ac761f071e92b6b8c0448fee001c0096.web-security-academy.net/ is vulnerable
https://acb71f6e1f079d25c08104be00290033.web-security-academy.net/ is vulnerable
```

Usage:
```
$ ./smuggler -h
Usage of ./smuggler:
  -proxy string
    	Set the Golang proxy, for example: http://example.com:8080
  -u string
    	Target URL
  -urls string
    	List of URLs

```
