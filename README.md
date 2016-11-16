# Truss Client Test

This repository hosts two [Truss](http://github.com/TuneLab/go-truss) generated
services: a generator and a printer. They don't do anything useful. But this was
an exercise in connecting two Truss-generated services using the generated HTTP
client.

In this case, the Printer service uses the Generator service to get a message
based on a name that is provided. The Printer service then decorates the returned
message with an exclamation point (just so that the Printer service isn't **only* a
passthrough).

## Notes

Truss generates a HTTP client that cannot reach the server that it attempts to
connect to. This happens because on the server side, all the HTTP routes are
specified as lowercase, and the client attempts to reach HTTP endpoints that are
title-cased. This results in a HTTP 404 being sent back to the client.

I have hand-modified the generated code for the Generator client to address this,
so that it attempts to connect to an endpoint with a lowercase path.

## To Run
```
make run
```

To use:
```
curl "locahost:11000/print?name=Test"
```

You should see the message in the JSON response.
