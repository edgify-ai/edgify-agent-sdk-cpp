#include <grpcpp/grpcpp.h>

#include "edgifySdk/prediction.pb.h"
#include "edgifySdk/prediction.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using grpc::CreateChannel;
using grpc::ClientAsyncResponseReader;
using grpc::CompletionQueue;

using edgify::EdgifyService;
using edgify::PredictionRequest;
using edgify::PredictionResponse;
using edgify::GroundTruthRequest;
using edgify::GroundTruthResponse;
using edgify::GroundTruth;
using edgify::Prediction;
using edgify::PredictionItem;
using edgify::Image;


class ExampleAsyncClient {
  public:
    explicit ExampleAsyncClient(std::shared_ptr<Channel> channel)
            : stub_(EdgifyService::NewStub(channel)) {}


    Status GetPrediction(PredictionResponse* response) {
        ClientContext context;
        PredictionRequest request;
        CompletionQueue cq;
        Status status;

        // stub_->PrepareAsyncGetPrediction() creates an RPC object, returning
        // an instance to store in "call" but does not actually start the RPC
        // Because we are using the asynchronous API, we need to hold on to
        // the "call" instance in order to get updates on the ongoing RPC.
        std::unique_ptr<ClientAsyncResponseReader<PredictionResponse> > rpc(
                stub_->PrepareAsyncGetPrediction(&context, request, &cq));

        // StartCall initiates the RPC call
        rpc->StartCall();

        // Request that, upon completion of the RPC, "reply" be updated with the
        // server's response; "status" with the indication of whether the operation
        // was successful. Tag the request with the integer 1.
        rpc->Finish(response, &status, (void*)1);
        void* got_tag;
        bool ok = false;
        // Block until the next result is available in the completion queue "cq".
        // The return value of Next should always be checked. This return value
        // tells us whether there is any kind of event or the cq_ is shutting down.
        GPR_ASSERT(cq.Next(&got_tag, &ok));

        // Verify that the result from "cq" corresponds, by its tag, our previous
        // request.
        GPR_ASSERT(got_tag == (void*)1);
        // ... and that the request was completed successfully. Note that "ok"
        // corresponds solely to the request for updates introduced by Finish().
        GPR_ASSERT(ok);

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
        CompletionQueue cq;
        Status status;

        GroundTruth ground_truth;// = request.mutable_ground_truth();
        ground_truth.set_allocated_prediction(prediction);
        ground_truth.set_label(label);
        std::cout << "prediction uuid: " << prediction->uuid() << std::endl;

        request.set_allocated_ground_truth(&ground_truth);

        std::unique_ptr<ClientAsyncResponseReader<GroundTruthResponse> > rpc(
                stub_->PrepareAsyncCreateGroundTruth(&context, request, &cq));

        rpc->StartCall();
        rpc->Finish(&response, &status, (void*)1);

        void* got_tag;
        bool ok = false;

        GPR_ASSERT(cq.Next(&got_tag, &ok));
        GPR_ASSERT(got_tag == (void*)1);
        GPR_ASSERT(ok);

        ground_truth.release_prediction();
        request.release_ground_truth();

        return status;
    }

private:
    std::unique_ptr<EdgifyService::Stub> stub_;
};

int main() {
    std::string target_str = "localhost:50051";
    ExampleAsyncClient client = ExampleAsyncClient(grpc::CreateChannel(
            target_str, grpc::InsecureChannelCredentials()));

    PredictionResponse response;
    Status status = client.GetPrediction(&response);
    Status gt_status = client.CreateGroundTruth(response.mutable_prediction(), "banana");

    std::cout << "prediction status: " << gt_status.ok() << std::endl;
}