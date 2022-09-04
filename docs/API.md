# API

### Routes

| Route | Description  | Implemented |
|---|---|---|
| `/api/health` | This is the health status endpoint. Performing a request to this endpoint the client should receive a lightweight response 'OK' with 200 as status code if everything is behaving correctly. | ✅ |
| `/api/info/{ip}` | This endpoint is intended to serve a summary of the information available for the IP taken as input. Parameter: top. Which info? Number of requests, timestamp of last activity, top X (default 10) methods, path, headers. | ✅ |
| `api/logs` | This endpoint is intended to serve information about logs, so the requests were made to the Honeypot. These are the parameters the endpoint accepts: id, ip, method, header, path, date (YYYY-MM-DD), lt (less than YYYY-MM-DD-HH-MM-SS), gt (greater than YYYY-MM-DD-HH-MM-SS). | ❌ |
| `api/detect` | This endpoint is intended to perform a heavy and accurate scan on the logs. It takes as input these parameters: regex (Go format), attack (use a list of well known regex), target (where to apply the regex), ip, method, header, path, date, lt (less than YYYY-MM-DD-HH-MM-SS), gt (greater than YYYY-MM-DD-HH-MM-SS). | ❌ |
| `api/stats` | This endpoint gives a general overview of the system. | ❌ |
| `api/stats/db` | This endpoint gives a detailed overview of the data stored in the DB. | ❌ |
