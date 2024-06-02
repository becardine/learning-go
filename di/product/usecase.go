package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// This product was not supposed to be exported. We should return a DTO instead.
// This is just an example.
func (uc *ProductUseCase) GetProduct(id int) (Product, error) {
	return uc.repository.GetProduct(id)
}
