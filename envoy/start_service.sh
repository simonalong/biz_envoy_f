#!/usr/bin/env bash

envoy -c /etc/envoy/envoy.yaml --service-cluster biz-f-service&
sleep 10s
/app/server
