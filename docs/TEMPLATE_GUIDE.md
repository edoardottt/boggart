# Template guide

This is a basic example (Same as examples/basic-raw/basic-raw.yaml):
```yaml
type: "raw"
requests:
  - request:
    id: "1"
    methods: ["GET", "POST"]
    endpoint: "/"
    response-type: "raw"
    content-type: "text/html"
    content:
      |
      <html>
      <head>
        <title>Hello</title>
      </head>
      <body>
        <h1>Hello!</h1>
      </body>
      </html>
    headers:
      - "server: Akamai Resource Optimizer"
      - "server-timing: cdn-cache; desc=HIT"
  - request:
    id: "default"
    response-type: "file"
    content-type: "text/html"
    content: "404.html"
    headers:
      - "server: Akamai Resource Optimizer"
      - "server-timing: cdn-cache; desc=HIT"

ignore: ["/favicon.ico"]
```

As you can see boggart templates are YAML files.  
The first line tells the template engine that this is a 'raw' configuration (The other choice is 'shodan', so like copying exposed hosts with shodan, but for now it's not implemented).  
Then at the second line we have the definition of the requests the honeypot can accept.  
A request must have an id (identifier) that must be unique. Then you can specify for which HTTP methods it should replies to the client (it is an array of methods!).  
Then we have the definition of the endpoint and the response type. We have two types of response:
  - "raw" means you have to add the content-type (as shown, whatever it is: json, html...) and the 'content' with its raw content (as shown).
  - "file" means you have to add the content-type (as shown, whatever it is: json, html...) and the 'content' will be the path to the file that boggart will read and use as reply to the client.

There is always a default request definition with the id 'default' used for 404 responses.  
Then we have the ignore array, that is the paths boggart must ignore.  
The last thing is the headers configuration, here you can add custom headers. 

If you need more examples see [/examples](https://github.com/edoardottt/boggart/tree/main/examples).  
Be aware that if you need files to be read with 'file' as response-type (as shown here, 404.html) you need to put them in the folder `public/honepot`.  
Use
```console
./scripts/set-apache-httpd.sh
```
to create an Apache httpd honeypot (`./scripts/unset-apache-httpd.sh` back to basic configuration).
