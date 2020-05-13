# Horcrux

Split your ~~soul~~ source code into multiple repositories.

This project was mostly a silly idea, 
but it seemed to have *just enough* use-case for me to 
seriously think about implementing.

The idea is, whenever you push changes to a remote, that remote sends
a webhook to `horcrux`, which then clones and pushes those changes
to any number of configured services.

For an example config, check out [horcrux.example.yml](horcrux.example.yml)

## Webhook Endpoints

The format for a webhook endpoint is `https://horcrux.domain.tld/<name>/<service>` where `name` is the name of your
configured repository and `service` is the type of service sending the webhook. (gitea, github, or gitlab)