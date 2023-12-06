TODO
====

protbuf support
---------------

packet structure
----------------

```protobuf
message packet {
    required int64  packetno = 1; // 包编号
    optional string msg = 2;
    optional int64  sendtime = 3; // 消息发送时间戳
}

```
