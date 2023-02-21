A UNIX Domain Socket or Inter-Process Communication (IPC) socket is a data

# UNIX Domain Socket

communications endpoint that allows you to exchange data between processes that run on the same machine. You might ask, why use UNIX domain sockets instead of TCP/IP connections for processes that exchange data on the same machine? First, because UNIX domain sockets are faster than TCP/IP connections and second, because UNIX domain sockets require fewer resources than TCP/IP connections. So, you can use UNIX domain sockets when both the clients and the server are on the same machine.
