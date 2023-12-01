# Footies

Hello, welcome on footies repo ✨

Footies is a live betting football application, allowing users to bet on live matches and get up to date informations.

## Structure

Footies has 3 main services :

-   A front-end written in Vue
-   An API written in Node
-   A worker written in Golang

### Front-end

TODO

### API

The Restful API is the bridge between the front-end and the worker. It communicates with the front-end with websockets, and with workers using redis pub/sub or REST API.

### Worker

The Worker communicates with a third party API to get live informations about currently playing football games.
It exposes informations about matches and leagues with a GET endpoint, and also listen for some messages to trigger job in a redis pub/sub channel.

#### Job specs

TODO

## Specs

### Versions

-   Node 20.10.0
-   Vue 2.7
-   Go 1.20.5
