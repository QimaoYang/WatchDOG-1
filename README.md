# WatchDOG
## API Design
### backend_inner
__1.encrypt code__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/seat/encrpt
```
request body(JSON):
{
    "seat_number":string
}
```
```
response:
status code: 201
response body{JSON}:
{
    "data": {
        "encryp_text":string    
    },
}
``` 
### backend_outer
__1.check all seats status__
    method: get
    URL: {host}:{port}/powercubicle/v1/seat
```
response:
status code: 200
response body(JSON):
{
    "data": {
        "seats":
            [
                {"A" : 10},
                {"B": 19},
                ...
            ]  
    }
}
```
__2.check single region seats status__
    method: get
    URL: {host}:{port}/powercubicle/v1/seat?region={region_code}
```
response:
status code: 200
response body(JSON):
{
    "data": {
        "seats":[
            {"A1" : "availble"},
            {"A2": "Zhang3"},
            ...
        ]
    }
}
```
__3.user register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/user/register
```
request body(JSON):
{
    "user_name":string,
    "user_password":string
}
```
```
response:
status code: 201
response body{JSON}:
{
    "session_key":string
}
``` 
__4.user login__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/user/login
```
request body(JSON):
{
    "user_name":string,
    "user_password":string
}
```
```
response:
status code: 201
response body{JSON}:
{
    "session_key":string,
}
``` 
__5.seat register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/seat/register
```
request body(JSON):
header:{Auth:{$session_token}}
{
    "encrypted_qrcode":string,
}
```
```
response:
status code: 201
``` 
__5-1.seat release__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/seat/release
```
header:{Auth:{$session_token}}
```
```
response:
status code: 201
```  
__6.get current seat__ 
    method: get
    URL: {host}:{port}/powercubicle/v1/user/seat
```
header:{"session_key":string}
```
```
response:
status code: 200
response body{JSON}:
{
    "seat": "A1"/nil
}
``` 
### database middleware
__1.check all seats status__
    method: get
    URL: {host}:{port}/powercubicle/v1/db/seat
```
response:
status code: 200
response body(JSON):
{
    "data": {
        "seats":
            [
                {"A" : 10},
                {"B": 19},
                ...
            ]  
    }
}
```
__2.check single region seats status__
    method: get
    URL: {host}:{port}/powercubicle/v1/db/seat?region={region_code}
```
response:
status code: 200
response body(JSON):
{
    "data": {
        "seats":[
            {"A1" : "availble"},
            {"A2": "Zhang3"},
            ...
        ]
    }
}
```
__3.user register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/db/user/register
```
request body(JSON):
{
    "user_name":string,
    "user_password":string
}
```
```
response:
status code: 201
response body{JSON}:
{
    "session_key":string
}
``` 
__4.user login__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/db/user/login
```
request body(JSON):
{
    "user_name":string,
    "user_password":string
}
```
```
response:
status code: 201
response body{JSON}:
{
    "session_key":string,
}
``` 
__5.seat register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/db/seat/register
```
header:{Auth:{$session_token}}
request body(JSON):
{
    "encrypted_qrcode":string,
    "time_for_use":enum[-0.5, +0.5, 1]
}
```
```
response:
status code: 201
```
__5-1.seat release__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/seat/release
```
header:{Auth:{$session_token}}
```
```
response:
status code: 201
``` 
__6.get current seat__ 
    method: get
    URL: {host}:{port}/powercubicle/v1/db/user/seat
```
header:{"session_key":string}
```
```
response:
status code: 200
response body{JSON}:
{
    "seat": "A1"/nil
}
``` 
### Managemnt Account
__1.update key__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/management/key
```
request body(JSON):
{
    "key":string,
}
```
```
response:
status code: 200
response body{JSON}:
{
    "status": "success",
}
``` 
### error_handling
```
status code: 400
response body{JSON}:
{
    "status": "Bad request",
    "message": "{error_message}"
}
status code: 404
response body{JSON}:
{
    "status": "Request not found",
    "message": "{error_message}"
}
status code: 500
response body{JSON}:
{
    "status": "Internal server error",
    "message": "{error_message}"
}
```
