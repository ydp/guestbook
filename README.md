# Guestbook

A simple guestbook web app implemented with Go, MongoDB and Bootstrap.

# Design Goals

The focus of this project is to demonstrate how to implement a simple web app using Go, MongoDB and Bootstrap.

# Requirements

## Installing MongoDB

To build and run the guestbook web app, MongoDB needs to be installed, see https://docs.mongodb.com/manual/installation/

Note: MongoDB is not included in recent versions of Debian, follow the installation instructions above.

Make sure that mongod is started before running the guestbook web app.

## Installing MongoDB Go Driver

The driver can easily be installed with go get:

```shell script
$ go get go.mongodb.org/mongo-driver
```

go get will print a warning that there are no go files. This can be ignored.