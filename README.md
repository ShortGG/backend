<img src="https://raw.githubusercontent.com/ShortGG/graphic-chart/master/logo.png" alt="logo-short-gg" width="25%" align="right" />

<br />

# backend of short.gg

Behind the scenes of short.gg

<hr />

# Routes

```
GET /shorten/:url
Return the shortened url key

GET /find/:key
Return the long url through the key.
Or Teapot if not found.

GET /:key
Return 301 with the long url in the location header.

```
