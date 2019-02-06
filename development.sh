#!/bin/bash
function run {
    docker-compose up app
}

"$@"