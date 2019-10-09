# Simple URL Shortener

## Run

1. ``docker-compose up``
2. Server runs on port 8080

## Create new shortened URL

***Endpoint:***

```bash
Method: POST
URL: http://localhost:8080/
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |



***Body:***

```js        
{
	"originalUrl": "https://google.com"
}
```

***Response:***
```
{
    "shortUrl": ""
}
```