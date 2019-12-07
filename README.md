# go-url-shortener
A simple Golang URL Shortener using native net/http for GCI 2019

## How to use
Simply type `make run` to run this app.
The app will be served on `127.0.0.1:9000`

## Adding redirect path
All redirection rules are stored in the `urls.json` file.
Simply add a new object under `urls` array with parameter of `base_url` and `redirect_url`