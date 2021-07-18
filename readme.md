install grpc from source to ./grpc (or change the path in cmake file)
instructions : https://github.com/grpc/grpc/blob/v1.34.0/BUILDING.md#pre-requisites

put in this folder the generated grpc files (the sdk files)

build with:
cmake --build ./cmake-build-debug --target grpc_test -- -j 4

run with:
./cmake-build-debug/grpc_test 


install grpc using:
`sudo apt install libgrpc++-dev`


sudo apt-get install -y libprotobuf-dev
