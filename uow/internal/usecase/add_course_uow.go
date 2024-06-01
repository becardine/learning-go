package usecase

import (
	"context"

	"github.com/becardine/learning-go/uow/internal/entity"
	"github.com/becardine/learning-go/uow/internal/repository"
	"github.com/becardine/learning-go/uow/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) ExecuteUow(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		// tudo que for executado aqui dentro será uma transação com begin e commit/rollback
		category := entity.Category{
			Name: input.CategoryName,
		}

		categoryRepository := a.getCategoryRepository(ctx)
		err := categoryRepository.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: category.ID,
		}

		courseRepository := a.getCourseRepository(ctx)
		err = courseRepository.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	categoryRepository, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	return categoryRepository.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	courseRepository, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}

	return courseRepository.(repository.CourseRepositoryInterface)
}
