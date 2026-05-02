#!/bin/bash
curl -s http://localhost:8080/api/v1/users/${1:?"Usage: $0 <user-id>"}
