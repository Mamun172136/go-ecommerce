package repo

import (
	"github.com/jmoiron/sqlx"
	"database/sql"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(prodctID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	// productList []Product
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	// generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	// 	p.ID = len(r.productList) + 1
	// r.productList =  append(r.productList, p)
	// return &p, nil

	query := `INSERT INTO products (title, description, price, img_url) 
              VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err:= row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	var p Product

	// Use sqlx's Get method to fetch a single product by ID
	query := `SELECT id, title, description, price, img_url FROM products WHERE id = $1`
	err := r.db.Get(&p, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No product found
			return nil, nil
		}
		// Return any other error
		return nil, err
}
return &p,nil

}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product

	// Use sqlx's Select method to fetch all products
	query := `SELECT id, title, description, price, img_url FROM products`
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

	func (r *productRepo) Delete(productID int) error {
	query := `DELETE FROM products WHERE id = $1`

	// Execute the DELETE statement
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}
func (r productRepo) Update(product Product) (*Product, error) {
	// for idx, p := range r.productList {
	// 	if p.ID == product.ID {
	// 		r.productList[idx] = product
	// 	}
	// }
	// return &product, nil

	query := `
		UPDATE products
		SET title = $1,
		    description = $2,
		    price = $3,
		    image_url = $4
		WHERE id = $5
		RETURNING id, title, description, price, img_url
	`
	row := r.db.QueryRow(query,
		product.Title,
		product.Description,
		product.Price,
		product.ImgUrl,
		product.ID,
	)
	err:= row.Err()
	if err != nil{
		return nil,err
	}
	return &product,nil
}

// func generateInitialProducts(r *productRepo) {
// 	prd1 := Product{
// 		ID:          1,
// 		Title:       "Orange",
// 		Description: "Orange is red. I love orange.",
// 		Price:       100,
// 		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
// 	}

// 	prd2 := Product{
// 		ID:          2,
// 		Title:       "Apple",
// 		Description: "Apple is green. I hate apple.",
// 		Price:       40,
// 		ImgUrl:      "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg",
// 	}

// 	prd3 := Product{
// 		ID:          3,
// 		Title:       "Banana",
// 		Description: "Banana is boring. I feel bored eating banana.",
// 		Price:       5,
// 		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
// 	}

// 	// prd4 := Product{
// 	// 	ID: 4,
// 	// 	Title: "Angur Fol",
// 	// 	Description: "Angur Fol Tastes good.",
// 	// 	Price: 140,
// 	// 	ImgUrl: "https://cdn.dhakapost.com/media/imgAll/BG/2022January/angur-2-20220215152127.jpg",
// 	// }

// 	// prd5 := Product{
// 	// 	ID: 5,
// 	// 	Title: "Mango",
// 	// 	Description: "Mango is my favorite. I love it very much.",
// 	// 	Price: 1000000,
// 	// 	ImgUrl: "https://www.dole.com/sites/default/files/styles/512w384h-80/public/media/dole-blog-03-maerz-mango-05.jpg?itok=qXHJMEAz-PEthlz_-",
// 	// }

// 	// prd6 := Product{
// 	// 	ID: 6,
// 	// 	Title: "Strawberry",
// 	// 	Description: "Strawberries are sweet, juicy, and bursting with flavor.",
// 	// 	Price: 500,
// 	// 	ImgUrl: "https://snaped.fns.usda.gov/sites/default/files/styles/crop_ratio_7_5/public/seasonal-produce/2018-05/strawberries.jpg.webp?itok=B4LFd4vV",
// 	// }

// 	r.productList = append(r.productList, prd1)
// 	r.productList = append(r.productList, prd2)
// 	r.productList = append(r.productList, prd3)
// 	// productList = append(productList, prd4)
// 	// productList = append(productList, prd5)
// 	// productList = append(productList, prd6)
// }
