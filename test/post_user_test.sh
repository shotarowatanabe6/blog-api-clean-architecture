#!/bin/bash
curl -s -X POST http://localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"Name":"test-user","Email":"test@example.com"}'
