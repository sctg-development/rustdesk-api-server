# rustdesk-api-server

RustDesk API Server Go version, supports sqlite3, mysql databases

After logging in, you can return all active hosts under the same account starting with this number

[![Build Release](https://github.com/sctg-development/rustdesk-api-server/actions/workflows/build.yml/badge.svg)](https://github.com/sctg-development/rustdesk-api-server/actions/workflows/build.yml)

## Compilation

Install Golang

Install GCC and configure PATH

## Usage

### Modify Database Connection

Modify the configuration items in conf/config.yml

```yaml
dbtype: mysql # supports mysql or sqlite3
mysql:
  host: '127.0.0.1'
  port: 3306 
  database: 'rustdesk' # database name
  username: 'root' # database username
  password: '' # database password
app:
  authkey: 123456 # authorization password for adding accounts or modifying passwords
  cryptkey: 123123123123  # password encryption salt value, recommended not to change after the first modification
```

### S3 Client Download
by adding the following configuration to the configuration file, you can download obtain dynamically from the S3 server a download link for the RustDesk client:

```yaml
s3:
  Endpoint: https://random.objectstorage.eu-par-1.oraclecloud.com
  Region: eu-par-1
  AccessKey: 2b3e1f4a5c6d7e8f9a0b1c2d3e4f5g6h7i8j9k0l
  SecretKey: R5sT8uVwX1yZ2aB3cD4eF5gH6iJ7kL8mN9oP0qR1s
  Bucket: randombucket123
  Windows64Key: master/sctgdesk-releases/sctgdesk-2.0.1-x86_64.exe
  Windows32Key: master/sctgdesk-releases/sctgdesk-2.0.1-i686.exe
  OSXKey: master/sctgdesk-releases/sctgdesk-2.0.1.dmg
  OSXArm64Key: master/sctgdesk-releases/sctgdesk-2.0.1.dmg
```

The api endpoint for the download link is `/api/software/client-download-link/os` where `os` can be `windows64`, `windows32`, `osx` or `osxarm64`.
the api will return a json object with the download link for the client. Links are valid for 15 minutes.

```json
{
    "url": "https://random.objectstorage.eu-par-1.oraclecloud.com/master/sctgdesk-2.0.1.dmg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=949a57ffdc586d72961ccab7618b9d58a7372d40%2F20240324%2Feu-marseille-1%2Fs3%2Faws4_request&X-Amz-Date=20240324T175308Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=43cebc953df3f9cc1ed8ba51191956f0e3ade9db27684107cbad8c9c9605394b"
}
```

### Set Up and Run

1. Run the program
   1. Running the program will automatically create tables

2. Port mapping (recommended):
   > Directly map port 21114 out, seems like rustdesk default port is 21114, at least for mobile it is ~~

   Nginx server:
   > Use reverse proxy, reverse proxy can use port 80 or whatever

   ```nginx
   #PROXY-START/

   location ^~ /
   {
       proxy_pass http://127.0.0.1:21114;
       proxy_set_header Host $host;
       proxy_set_header X-Real-IP $remote_addr;
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
       proxy_set_header REMOTE-HOST $remote_addr;
       proxy_set_header Upgrade $http_upgrade;
       proxy_set_header Connection $connection_upgrade;
       proxy_http_version 1.1;
       # proxy_hide_header Upgrade;

       add_header X-Cache $upstream_cache_status;

       #Set Nginx Cache

       set $static_file10BAHqk7 0;
       if ( $uri ~* "\.(gif|png|jpg|css|js|woff|woff2)$" )
       {
           set $static_file10BAHqk7 1;
           expires 1m;
       }
       if ( $static_file10BAHqk7 = 0 )
       {
           add_header Cache-Control no-cache;
       }
   }

   #PROXY-END/
   ```

### RustDesk Configuration

ID server and relay server can be found and installed using Docker

Here only the API server configuration is described
![img.png](img.png)

On Android, do not need to fill in the `http://` prefix and must be `21114` port

### Create Account

Request
<http://127.0.0.1:21114/api/reg?username=test&password=test&auth_key=123456>

UI
http://127.0.0.1:21114/

### Modify Password

Request
<http://127.0.0.1:21114/api/set-pwd?username=test&password=test&auth_key=123456>

## Note

When saving the address book, if `username` equals `----`, it will not be saved

## About

This project is for learning and is used for API server interaction with rustdesk remote assistance software

Using framework:
[beego](https://github.com/beego/beego)

Because RustDesk interface returns fixed content = =, the interface return structure is not very uniform

## Sponsorship
