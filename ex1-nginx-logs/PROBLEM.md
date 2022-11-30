# nginx logs

There are two files in this repository, which are both Nginx access logs.

- `access.log`
- `access.log.1`

How to read these log lines:

2600:1700:1430:4818:c812:c26f:9ebe:3bf5 - - [30/Nov/2022:00:00:05 +0000] mlskickofftimes.com "GET /favicon.ico HTTP/1.1" 404 87 0.003 "-" "Mozilla/5.0 (iPhone; CPU iPhone OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Mobile/15E148 Safari/604.1"

2600:1700:1430:4818:c812:c26f:9ebe:3bf5: Remote IP address

30/Nov/2022:00:00:05 +0000: Request date

mlskickofftimes.com: Request host

GET /favicon.ico HTTP/1.1: First line of HTTP request sent by client

404: Status code

87: How many bytes are in the response body.

0.003: How much time it took the server to respond.

"Mozilla/5.0 (iPhone; CPU iPhone OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Mobile/15E148 Safari/604.1": User Agent string.

### Questions

See if you can write a command to answer each of these questions.

- How many lines are there in both files, combined?

- What is the last line in `access.log` ?

- What is the second to last line in `access.log` ?

- How many lines are there with status code 500?

- How many different web domains are there? (mlskickofftimes.com, kevin.burke.dev,
  burke.dev are all different hosts). What is the most popular domain?

- What are the ten most popular request IP addresses?
