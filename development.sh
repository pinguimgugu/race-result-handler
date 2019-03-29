#!/bin/bash
function run {
    docker-compose up app
}

function dep {
    docker-compose up dep
}

function build {
    docker build -t race-base-image .
}

function run-unit {
    docker-compose up unit
}
"$@"
