# Linux User Administration

## build and run

`go build cmd/crudUsers.go`

`sudo ./crudUsers`

## Register Users

To create users in the system use the file called `users.txt` in `configfiles` directory. The struct of this file is in json, so the application can read de key `users` and impot them into the system.

ex:

```
{
  "users": [
    {
      "name": "bar",
      "directory": "/tmp/all",
      "group":  "foo",
      "shell": "/bin/bash"
    },
    {
      "name": "foo",
      "directory": "/tmp/tttt",
      "group": "bar",
      "shell": "/bin/bash"
    }
  ]
}
```

## Debug

The application generate `_application.log` log file to track information


## In development

* API support to create users
* Database layer
* application dockerizing