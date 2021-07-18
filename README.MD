# Edgify Agent SDK for C++
The C++ SDK requires the auto-generated protobuf files 
These files are provided with this SDK, but may also need to be re-generated in case of a version mismatch.
Instructions on how to generate the files is provided below.

## Generate Protobuf files

In order to generate the protobuf files, you must first open the Edgify-Agent project
###prerequisites
A gRpc installation is required. For the simplest installation, use Brew:
https://docs.brew.sh/Homebrew-on-Linux

After installing Brew, install gRPC using

    brew tap grpc/grpc
    brew install grpc

cd to 'edgify-agent/sdk/prediction/c++' and run the following commands:

    protoc -I ../../../src/edgify/edgifypb/ --grpc_out=. --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin ../../../src/edgify/edgifypb/prediction.proto
    protoc -I ../../../src/edgify/edgifypb/ --cpp_out=. ../../../src/edgify/edgifypb/prediction.proto
    
    protoc -I ../../../src/edgify/edgifypb/ --grpc_out=. --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin ../../../src/edgify/edgifypb/samples.proto
    protoc -I ../../../src/edgify/edgifypb/ --cpp_out=. ../../../src/edgify/edgifypb/samples.proto
    
    protoc -I ../../../src/edgify/edgifypb/ --grpc_out=. --plugin=protoc-gen-grpc=/usr/local/bin/grpc_cpp_plugin ../../../src/edgify/edgifypb/analytics.proto
    protoc -I ../../../src/edgify/edgifypb/ --cpp_out=. ../../../src/edgify/edgifypb/analytics.proto

* Note that you may need to change the value of the `protoc-gen-grpc` property.  To learn where it's installed on your machine, run the command: `which grpc_cpp_plugin`

After the required files were generated in the Agent directory - copy them over to the **edgifySdk** folder in the **edgify-agent-sdk-cpp** project.

## Running the example

Once you have the required protobuf files, cd to the **edgify-agent-sdk-cpp* directory and run the following commands:
To build:
    
    cmake --build ./cmake-build-debug  --target grpc_async -- -j 4

To run:

    ./cmake-build-debug/grpc_test 


