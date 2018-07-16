#!/bin/bash

./rspamdInfluxDB -interval="$INTERVAL" -rspamdUrl="$RSPAMD_URL" -rspamdPassword="$RSPAMD_PASSWORD" -influxHost="$INFLUXDB_URL" -influxDB="$INFLUXDB_DB" -influxUser="$INFLUXDB_USER" -influxPwd="$INFLUXDB_PWD"