#  Overview

This is an example of server running API service /json2xml and /xml2json
on HTTP port 80 serving GET request. It bounces back the sent structure
but in a different format.

The client is supposed to be general use client like curl. It is done in
the make file where example structures are used to test the server.

## Local tools used

The following tools shall be installed in order to run this example:
- go
- make
- jq
- curl

## Techniques demostrated

- Writing http server and clients
- Making request handlers
- Using many source files for a module
- Making a public structure with annotations
- Encoding/decoding JSON and XML formats
- Making unit tests
- Making integration tests
