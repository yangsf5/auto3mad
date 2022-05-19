docker run --rm -it -e SQLCONN="root:@tcp(host.docker.internal:3306)/auto3mad?charset=utf8" -e HTTPADDR="0.0.0.0" -p 1128:1127 auto3mad:test
