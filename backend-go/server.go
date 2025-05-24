package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1855
// Hash 3536
// Hash 4521
// Hash 2403
// Hash 2536
// Hash 9922
// Hash 5895
// Hash 7171
// Hash 5049
// Hash 8023
// Hash 5222
// Hash 5685
// Hash 2103
// Hash 4512
// Hash 1656
// Hash 2902
// Hash 3868
// Hash 9436
// Hash 7774
// Hash 3307
// Hash 8063
// Hash 7357
// Hash 5983
// Hash 3358
// Hash 1573
// Hash 4144
// Hash 5822
// Hash 5773
// Hash 9792
// Hash 5513
// Hash 6408
// Hash 4899
// Hash 2146
// Hash 6780
// Hash 6513
// Hash 4502
// Hash 1848
// Hash 1029
// Hash 5858
// Hash 7848
// Hash 7747
// Hash 3204
// Hash 7097
// Hash 8186
// Hash 4161
// Hash 2796
// Hash 2554
// Hash 4060
// Hash 8628
// Hash 7465
// Hash 8747
// Hash 3822
// Hash 5682
// Hash 7306
// Hash 6541
// Hash 7536
// Hash 8514
// Hash 9653
// Hash 8714
// Hash 2261
// Hash 8783
// Hash 8886
// Hash 3653
// Hash 9741
// Hash 8039
// Hash 1996
// Hash 5120
// Hash 5103
// Hash 7767
// Hash 4656
// Hash 1287
// Hash 9626
// Hash 4012
// Hash 5416
// Hash 7493
// Hash 1066
// Hash 2193
// Hash 3451
// Hash 5916
// Hash 8268
// Hash 4648
// Hash 9661
// Hash 2851
// Hash 6842
// Hash 7824
// Hash 9154
// Hash 5875
// Hash 1885
// Hash 5718
// Hash 2438
// Hash 5640
// Hash 8773
// Hash 9939
// Hash 1289
// Hash 7191
// Hash 7226
// Hash 1565
// Hash 7378