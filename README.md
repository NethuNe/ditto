# Ditto!

Ditto is a dependency mocking service. It uses a simple JSON configuration schema to identify what request
is being sent and how to reply. 

This currently requires a configuration change in the client service to redirect it to Ditto,
but that can also be accomplished through IPTables manipulation.

## todo
- define the JSON config w/ a simple example
- define the wildcard config-handler

## Features

### Completed

none.

### Planned

- HTTP config driven request handling
- byte vomit to enable testing large and/or unexpected responses / mangled packets
- on-the-fly config upload / changing & local config storage
- kubernetes support
- library options such as parameterization, latency injection, request dropping, etc..
- HTTPS (will only work if the certificate wherever it's deployed is the same as the original requester... or you spoof it or allow insecure connections)
- UDP
- Redirection component, meant to be run separately, to redirect your service for you to wherever this one is deployed

## How to Run

`./.build/buildAndRun.sh`

## How to Maintain