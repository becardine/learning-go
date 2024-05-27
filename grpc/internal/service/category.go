package service

import (
	"context"

	"github.com/becardine/learning-go/grpc/internal/database"
	"github.com/becardine/learning-go/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (s *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := s.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}
	return categoryResponse, nil
}
