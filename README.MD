# Edgify Agent SDK for C++
The C++ SDK requires the auto-generated protobuf files. 
These files are provided with this SDK, but may also need to be re-generated in case of a version mismatch.

If you get any errors during the build stage, it's advised to re-generate the protobuf files and try again. 
It will likely solve your issue.

Instructions on how to generate the files is provided below.

##prerequisites

#### gRpc

Install gRpc as described in the gRpc installation instructions:
https://grpc.io/docs/languages/cpp/quickstart/
  
#### protobuf-compiler
protobuf-compiler is also required and can be installed using
    
    apt install protobuf-compiler

## Generate Protobuf files

In order to generate the protobuf files, run the following commands:

    protoc -I ./protoFiles/ --grpc_out=./edgifySdk --plugin=protoc-gen-grpc=/home/.local/bin/grpc_cpp_plugin protoFiles/prediction.proto
    protoc -I ./protoFiles/ --cpp_out=./edgifySdk protoFiles/prediction.proto

The files will be generated in the edgifySdk directory.

* Note that you may need to change the value of the `protoc-gen-grpc` property.  To learn where it's installed on your machine, run `which grpc_cpp_plugin`


## Building and running the examples

We provide 2 examples for the use of the SDK - an async prediction example and a sync prediction example.

To build the examples:

    cmake . -DCMAKE_BUILD_TYPE=Debug -B cmake-build-debug

    cmake --build ./cmake-build-debug  --target grpc_async -- -j4
    cmake --build ./cmake-build-debug  --target grpc_sync -- -j4


To run:

    ./cmake-build-debug/grpc_async 
    ./cmake-build-debug/grpc_sync 



