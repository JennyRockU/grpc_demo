package storytel

import (
	"log"

	"golang.org/x/net/context"
)

// an abstract represntation of Server type, for DI
type Server struct {
}

// FetchCourses creates a response to a Courses' request
func (s *Server) Fetch(ctx context.Context, in *CoursesRequest) (*CoursesResponse, error) {
	log.Printf("Received message, with userId: %s\n", in.Userid)

	courses := coursesForUser(in.Userid)
	return &CoursesResponse{Courses: courses}, nil
}

func coursesForUser(userid string) []string {

	switch userid {
	case "897146":
		return []string{"Global Finance", "Marketing Management", "Product Innovation"}

	case "7622327":
		return []string{"Entrepreneurial Management", "Financial Analysis", "Innovation Equity", "Big Data Basics"}

	default:
		return []string{}
	}

}
