前提条件：
启动etcd
$etcd

测试命令及正确结果输出：
$ go run server.go
2019-12-08 16:28:04.769923 I | Transport [http] Listening on [::]:50106
2019-12-08 16:28:04.769923 I | Broker [http] Connected to [::]:50107
2019-12-08 16:28:04.816925 I | Registry [etcd] Registering node: lp.srv.eg1-7f2d04df-8eb2-4e24-a991-fc1112d49cef
received hello server

allen@allen-PC MINGW64 /f/GoWorkSpace/micro-etcd (master)
$ go run client.go
msg:"hello world" values:"a" values:"b" header:<key:"name" value:<key:1 values:"abc" > > type:DESCEND
