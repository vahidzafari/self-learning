# WebSocket

The WebSocket protocol is a computer communications protocol that provides full-duplex (transmission of data in two directions simultaneously) communication channels over a single TCP connection. The WebSocket protocol is defined in RFC 6455 (<https://tools.ietf.org/html/rfc6455>) and uses ws:// and wss:// instead of <http://> and <https://>, respectively. Therefore, the client should begin a WebSocket connection by using a URL that starts with ws://.

The golang.org/x/net/websocket package offers another way of developing WebSocket clients and servers. However, according to its documentation, golang. org/x/net/websocket lacks some features and it is advised that you use <https://godoc.org/github.com/gorilla/websocket,> the one used here, or <https://godoc.org/nhooyr.io/websocket> instead.

The advantages of the WebSocket protocol include the following:

- A WebSocket connection is a full-duplex, bidirectional communications channel. This means that a server does not need to wait to read from a client to send data to the client and vice versa.
- WebSocket connections are raw TCP sockets, which means that they do not have the overhead required to establish an HTTP connection.
- WebSocket connections can also be used for sending HTTP data. However, plain HTTP connections cannot work as WebSocket connections.
- WebSocket connections live until they are killed, so there is no need to reopen them all the time.
- WebSocket connections can be used for real-time web applications.
- Data can be sent from the server to the client at any time, without the cient even requesting it.
- WebSocket is part of the HTML5 specification, which means that it is supported by all modern web browsers.
