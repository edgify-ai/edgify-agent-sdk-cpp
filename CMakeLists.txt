set(CMAKE_CXX_STANDARD 17)
project(example)
cmake_minimum_required(VERSION 3.17)

# hint to gRpc installation
#set(GRPC_DIR "/home/linuxbrew/.linuxbrew/Cellar/grpc/1.38.1/include")
#include_directories(${GRPC_DIR})

set(_GRPC_GRPCPP grpc++)
set(_REFLECTION grpc++_reflection)
set(_PROTOBUF_LIBPROTOBUF libprotobuf)

add_executable(grpc_sync sync_example.cpp ../edgifySdk/prediction.grpc.pb.cc ../edgifySdk/prediction.pb.cc)
add_executable(grpc_async async_example.cpp ../edgifySdk/prediction.grpc.pb.cc ../edgifySdk/prediction.pb.cc)


target_link_libraries(grpc_sync
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
        )

target_link_libraries(grpc_async
        ${_REFLECTION}
        ${_GRPC_GRPCPP}
        ${_PROTOBUF_LIBPROTOBUF}
        )

