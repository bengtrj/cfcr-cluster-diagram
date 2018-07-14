#!/usr/bin/env bash

instance_groups="$(yq read cfcr-manifest.yml "instance_groups" -j)"
master="$(echo ${instance_groups} | jq '.[] | select(.name | contains("master"))')"
master_instances="$(echo ${master} | jq '.instances')"

