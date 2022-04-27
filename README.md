# Guestbook

An automation targeted simple guestbook web app implemented with Go, MongoDB and Bootstrap, inspired by [mthaler/guestbook](https://github.com/mthaler/guestbook)

# Design Goals

The focus of this project is to demonstrate how to implement a simple web app using Go, MongoDB and Bootstrap. Security is not a concern, thus it should not be used on production systems.

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

# Usage

Use the following command to run the guestbook web app from the guestbook directory:

```shell script
$ go run main.go
```

If the guestbook application can connect to the MongoDB database, it will print a message, otherwise it will print an error message.

Open a browser with the following URL:

```
http://localhost/guestbook
```

An empty guestbook should now be shown. New comments can be added by clicking on add comment.

# Manipulating the Database

MongoDB includes a shell that can be used to manipulate the database. See the MongoDB documentation for details.

To delete all comments, stop the guestbook web app and do the following

```
$ mongo
> use guestbook
switched to db guestbook
> db.comments.drop()
true
```

After starting the guestbook again, an empty guestbook is shown.

# Credits

The code is based on the guestbook example from the [Head First Go book](https://headfirstgo.com/).

# License

BSD 3-Clause License
