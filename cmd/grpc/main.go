package main 

import(
	"weather/internal/grpcserver"
	"weather/weatherpb"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	weatherpb.RegisterWeatherServiceServer(grpcServer, &grpcserver.Server{})

	log.Println("gRPC Weather server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}