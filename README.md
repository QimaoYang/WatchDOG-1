# WatchDOG
## API Design
### backend_inner
__1.generate QRcode__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/qrcode
```
request body(JSON):
{
    "seat_number":string,
    "time":{year-month-day} string
}
```
```
response:
status code: 200
response body{JSON}:
{
    "status": "success",
    "data": {
        "qr_code":{encrypted_string}    
    },
}
``` 
__2.encrypt code__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/encrpt
```
request body(JSON):
{
    "seat_number":string,
    "time":{year-month-day hh:mm:ss} string
}
```
```
response:
status code: 200
response body{JSON}:
{
    "status": "success",
    "data": {
        "encryp_text":[]byte    
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
    "status": "success",
    "data": {
        "region": "all"
        "seats":{
            [
                "A1" : "availble",
                "B2": "unavailble",
                ...
            ]  
        }
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
    "status": "success",
    "data": {
        "region":{region_code}
        "seats":{
            [
                "A1" : "availble",
                "B2": "unavailble",
                ...
            ]  
        }
    }
}
```
__3.user register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/user/register
```
request body(JSON):
{
    "user_email":string,
    "user_password":string
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
__3.user login__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/user/login
```
request body(JSON):
{
    "user_email":string,
    "user_password":string
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
__4.seat register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/seat/register
```
request body(JSON):
{
    "encrypted_qrcode":string,
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
__5.decrypt code__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/decrpt
```
request body(JSON):
{
    "encryp_text":[]byte
}
```
```
response:
status code: 200
response body{JSON}:
{
    "status": "success",
    "data": {
        "seat_number":{string},
        "time":{year-month-day hh:mm:ss}
    },
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
    "status": "success",
    "data": {
        "region": "all"
        "seats":{
            [
                "A1" : "availble",
                "B2": "unavailble",
                ...
            ]  
        }
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
    "status": "success",
    "data": {
        "region":{region_code}
        "seats":{
            [
                "A1" : "availble",
                "B2": "unavailble",
                ...
            ]  
        }
    }
}
```
__3.user register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/db/user/register
```
request body(JSON):
{
    "user_email":string,
    "user_password":string
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
__3.get user password__ 
    method: get
    URL: {host}:{port}/powercubicle/v1/db/user?user_email={user_email}
```
response:
status code: 200
response body{JSON}:
{
    "status": "success",
}
``` 
__4.seat register__ 
    method: post
    URL: {host}:{port}/powercubicle/v1/seat/register
```
request body(JSON):
{
    "seat_number":string,
    "time":{year-month-day} string
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