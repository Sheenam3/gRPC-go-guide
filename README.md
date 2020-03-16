# gRPC-go-guide
To get a basic understanding of gRPC Client &amp; Server Services


<h3> Server Output: </h3>
<pre>
root@sheenam:~/go-workspace/src/github.com/quotes/quote_server# go run server.go
2020/03/16 13:32:02 -------------Simple RPC-----------
2020/03/16 13:32:02 Client sends: quote1 is here
2020/03/16 13:32:02 Recieving quote from client for formatting: quote1 is here
2020/03/16 13:32:02 -------------Client Streaming-----------
2020/03/16 13:32:02 Printing the value of r quote 1 streaming client
2020/03/16 13:32:02 Printing the value of r quote 1 streaming client
2020/03/16 13:32:02 -------------Bi-Directional-----------
2020/03/16 13:32:02  From Client: quoteone:"Hello, I am Client"
2020/03/16 13:32:02  From Client: quoteone:"Hello, I am Client"
2020/03/16 13:32:02  From Client: quoteone:"Hello, I am Client"
</pre>


<h3> Client Output: </h3>
<pre>
root@sheenam:~/go-workspace/src/github.com/quotes/quote_client# go run client.go
2020/03/16 13:32:02 -------------Simple RPC-----------
2020/03/16 13:32:02  Quote of the Day: quote1 is here
2020/03/16 13:32:02 -------------Server Streaming-----------
2020/03/16 13:32:02 quote1_response:"quote 1" quote2_response:"quote 2" quote3_response:"quote 3"
2020/03/16 13:32:02 quote1_response:"quote 1" quote2_response:"quote 2" quote3_response:"quote 3"
2020/03/16 13:32:02 -------------Client Streaming -----------
2020/03/16 13:32:02 resp: quote1: quote1_response:"response 1" quote2_response:"response 2" quote3_response:"response 3"
2020/03/16 13:32:02 -------------Bi-Directional-----------
2020/03/16 13:32:02 resp: Client send : quoteone:"Hello, I am Client"  /n Server Responded: quote:"Hi, from the Server"
2020/03/16 13:32:02 resp: Client send : quoteone:"Hello, I am Client"  /n Server Responded: quote:"Hi, from the Server"
2020/03/16 13:32:02 resp: Client send : quoteone:"Hello, I am Client"  /n Server Responded: quote:"Hi, from the Server"
</pre>
