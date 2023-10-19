# base64uuid

`base64uuid` encode and decode base64 UUIDs 

`go install github.com/fpuc/base64uuid@latest`

Usage:

```shell
base64uuid 
# S0kaxR7hRuiCcsJjXr7sWw==

base64uuid 0acf3688-09b8-4c98-bdfc-a5418eb2db03
# Cs82iAm4TJi9/KVBjrLbAw==

base64uuid -decode Cs82iAm4TJi9/KVBjrLbAw==
# 0acf3688-09b8-4c98-bdfc-a5418eb2db03

base64uuid -d Cs82iAm4TJi9/KVBjrLbAw==
# 0acf3688-09b8-4c98-bdfc-a5418eb2db03
```

