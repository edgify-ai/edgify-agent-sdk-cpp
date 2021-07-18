//#include <grpcpp/grpcpp.h>

#include <grpc/grpc.h>
#include <grpcpp/channel.h>
#include <grpcpp/client_context.h>
#include <grpcpp/create_channel.h>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/server.h>
#include <grpcpp/server_builder.h>
#include <grpcpp/server_context.h>

#include "edgifySdk/prediction.pb.h"
#include "edgifySdk/prediction.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using grpc::CreateChannel;


using edgify::EdgifyService;
using edgify::PredictionRequest;
using edgify::PredictionResponse;
using edgify::GroundTruthRequest;
using edgify::GroundTruthResponse;
using edgify::GroundTruth;
using edgify::Prediction;
using edgify::PredictionItem;
using edgify::Image;


class ExampleClient {
public:
    explicit ExampleClient(std::shared_ptr<Channel> channel)
            : stub_(EdgifyService::NewStub(channel)) {}


    Status GetPrediction(PredictionResponse* response) {
        ClientContext context;
        PredictionRequest request;

        Status status = stub_->GetPrediction(&context, request, response);

        if (status.ok()) {
            const Prediction prediction = response->prediction();

            int predictions_size = prediction.predictions_size();
            for (int i=0; i<predictions_size; i++){
                const PredictionItem item = prediction.predictions(i);
                int data_size = item.data_size();

                if (data_size>0) {
                    const std::string label = item.data(0);
                    const std::string accuracy = item.data(1);
                    std::cout << "prediction " << label << ": " << accuracy << ": " << std::endl;
                }
            }

            const Image image = prediction.image();
            const std::string image_str = image.image();
            // std::cout << "image: " << image_str << std::endl;

            const std::string uuid = prediction.uuid();
            std::cout << "prediction uuid: " << uuid << std::endl;
        } else {
            std::cout << status.error_code() << ": " << status.error_message() << std::endl;
        }

        return status;
    }

    Status CreateGroundTruth(Prediction* prediction, const std::string& label) {
        ClientContext context;
        GroundTruthRequest request;
        GroundTruthResponse response;

        GroundTruth ground_truth;// = request.mutable_ground_truth();
        ground_truth.set_allocated_prediction(prediction);
        ground_truth.set_label(label);
        std::cout << "prediction uuid: " << prediction->uuid() << std::endl;

        request.set_allocated_ground_truth(&ground_truth);
        Status status = stub_->CreateGroundTruth(&context, request, &response);

        ground_truth.release_prediction();
        request.release_ground_truth();

        return status;
    }

private:
    std::unique_ptr<EdgifyService::Stub> stub_;
};


int main() {
    std::string target_str = "localhost:50051";
    ExampleClient client = ExampleClient(grpc::CreateChannel(
            target_str, grpc::InsecureChannelCredentials()));

    PredictionResponse response;
    Status status = client.GetPrediction(&response);

    Status gt_status = client.CreateGroundTruth(response.mutable_prediction(), "banana");

    std::cout << "prediction status: " << gt_status.ok() << std::endl;
}

