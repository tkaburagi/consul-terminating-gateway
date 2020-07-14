#!/bin/sh
consul agent -server -bind=127.0.0.1 \
-client=0.0.0.0 \
-data-dir=data-dir \
-bootstrap-expect=1 -ui
