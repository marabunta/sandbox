# Mutual authentication (mTLS  example)

clone/download and run `make`

If error try to get the dependencies (gRPC):

    $ cd src/client
    $ go get

Test by running the server:

    $ ./server

Then the client:

    $ ./client

Or

    $ ./client foo

Output should be something like:

    2018/11/07 23:22:15 Greting: Hello foo

On the server:

    2018/11/07 23:22:10 Listening on port :1415
    [127.0.0.1:63583 - client.example.com
