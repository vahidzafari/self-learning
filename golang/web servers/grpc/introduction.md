# Introduction

gRPC, which stands for gRPC Remote Procedure Calls, is an alternative to RESTful services that was developed by Google. The main advantage of gRPC is that it is faster than working with REST and JSON messages.

The process for creating a gRPC server and client has three main steps. The first step is creating the interface definition language (IDL) file, the second step is the development of the gRPC server, and the third step is the development of the gRPC client that can interact with the gRPC server.

The advantages of gRPC include the following:

- The use of binary format for data exchange makes gRPC faster than services that work with data in plain text format
- The command-line tools provided make your work simpler and faster
- Once you have defined the functions and the messages of a gRPC service, creating servers and clients for it is simpler than RESTful services
- gRPC can be used for streaming services
- You do not have to deal with the details of data exchange because gRPC takes care of the details

## Protocol buffers

A protocol buffer (protobuf) is basically a method for serializing structured data. A part of protobuf is the IDL. As protobuf uses binary format for data exchange, it takes up less space than plain text serialization formats.

Use the following command to generate

```bash
protoc --go_out=plugins=grpc:. protoapi.proto
```
