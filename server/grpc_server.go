package main

import (
	"billsHotelService/domain/entity"
	"billsHotelService/infrastructure/database"
	"billsHotelService/proto"
	"billsHotelService/server/repository"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HotelServiceServer struct {
	proto.UnimplementedHotelServiceServer //grpc で構築していないメソッドの実装を許容する
	repo                                  *repository.MySQLHotelRepository
}

func NewHotelServiceServer(repo *repository.MySQLHotelRepository) *HotelServiceServer {
	return &HotelServiceServer{repo: repo}
}

func (s *HotelServiceServer) GetHotel(ctx context.Context, req *proto.HotelRequest) (*proto.HotelResponse, error) {
	hotel, err := s.repo.HotelGetById(int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &proto.HotelResponse{
		Hotel: &proto.Hotel{
			Id:             int32(hotel.ID),
			Name:           hotel.Name,
			PricePernight:  int32(hotel.PricePerNight),
			RoomsAvailable: int32(hotel.RoomsAvailable),
		},
	}, nil
}

func (s *HotelServiceServer) CreateHotel(ctx context.Context, req *proto.HotelRequest) (*proto.HotelResponse, error) {
	newHotel := entity.NewHotel(int(req.GetId()), req.GetName(), int(req.GetPricePernight()), int(req.GetRoomsAvailable()))
	err := s.repo.HotelSave(*newHotel)
	if err != nil {
		return nil, err
	}
	return &proto.HotelResponse{
		Hotel: &proto.Hotel{
			Id:             int32(newHotel.ID),
			Name:           newHotel.Name,
			PricePernight:  int32(newHotel.PricePerNight),
			RoomsAvailable: int32(newHotel.RoomsAvailable),
		},
	}, nil
}

func main() {
	db, err := database.NewMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err) //ログの出力
	}
	defer db.Close()

	repo := repository.NewMySQLHotelRepository(db)
	server := grpc.NewServer()
	proto.RegisterHotelServiceServer(server, NewHotelServiceServer(repo))

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("gRPC server is running on port 50051...")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
