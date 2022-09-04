# Template guide

This is a basic example (Same as examples/basic-raw/basic-raw.yaml):
```
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
  - request:
    id: "default"
    response-type: "file"
    content-type: "text/html"
    content: "404.html"

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
