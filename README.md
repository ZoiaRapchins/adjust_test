# README

## Description

This program executes GET requests to specified urls.
Requests are fired concurrently. Number of max. concurrent requests can be set via command line.
Default value is 10.

## Build

Program can be built by executing:

`go build -o adjust_test`

Note: You would need a working Go setup on your machine. [Info](https://golang.org/doc/install)

## Usage

Program can be run by executing:

`./adjust_test -parallel=3 http://google.com facebook.com`

- parallel - is optional parameter. It sets max. concurrent requests.

