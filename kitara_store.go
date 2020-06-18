package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type product struct {
	Name          string `json:"name" binding:"required"`
	StockQuantity int    `json:"stock_quantity"`
}

type orderProduct struct {
	Product       product `json:"product" binding:"required"`
	OrderQuantity int     `json:"order_quantity" binding:"required"`
}

type order struct {
	CustomerName string         `json:"customer_name" binding:"required"`
	OrderProduct []orderProduct `json:"order_product" binding:"required"`
}

func (o *order) validate() {
	var op []orderProduct
	for _, x := range o.OrderProduct {
		if len(op) != 0 {
			f := true
			for i, y := range op {
				if x.Product.Name == y.Product.Name {
					y.OrderQuantity += x.OrderQuantity
					op[i] = y
					f = false
				}
			}
			if f {
				op = append(op, x)
			}
		} else {
			op = append(op, x)
		}
	}
	o.OrderProduct = op
}

func (o order) isOrderProductExist(p orderProduct) (b bool) {
	for _, x := range o.OrderProduct {
		if x.Product.Name == p.Product.Name {
			b = true
			break
		}
	}
	return b
}

type store struct {
	Name             string    `json:"store_name" binding:"required"`
	AvailableProduct []product `json:"available_product"binding:"required"`
	Order            []order   `json:"order" binding:"required"`
}

func (s *store) addProduct(p product) (b bool) {
	exist, _ := s.isProductExist(p)
	if !exist {
		s.AvailableProduct = append(s.AvailableProduct, p)
		b = true
	}
	return b
}

func (s *store) editProduct(p product) (b bool) {
	exist, index := s.isProductExist(p)
	if exist {
		s.AvailableProduct[index] = p
		b = true
	}
	return b
}

func (s *store) deleteProduct(p product) (b bool) {
	exist, i := s.isProductExist(p)
	if exist {
		s.AvailableProduct = append(s.AvailableProduct[:i], s.AvailableProduct[i+1:]...)
		b = true
	}
	return b
}

func (s *store) isProductExist(p product) (b bool, i int) {
	for d, x := range s.AvailableProduct {
		if x.Name == p.Name {
			b = true
			i = d
			break
		}
	}
	return b, i
}

func (s *store) addOrder(o order) (err error) {
	if len(o.OrderProduct) != 0 {
		noe := true
		for _, x := range o.OrderProduct {
			if x.OrderQuantity == 0 {
				noe = false
				break
			}
		}
		if noe {
			o.validate()
			for _, x := range o.OrderProduct {
				for i, y := range s.AvailableProduct {
					if y.Name == x.Product.Name {
						y.StockQuantity -= x.OrderQuantity
						s.AvailableProduct[i] = y
					}
				}
			}
			s.Order = append(s.Order, o)
		} else {
			err = errors.New("please input quantity of the product you ordered")
		}
	} else {
		err = errors.New("order product should not empty")
	}
	return err
}

func main() {
	// List Product of Kitara Store
	hs := product{
		Name:          "Hijab Syar'i",
		StockQuantity: 500,
	}
	hp := product{
		Name:          "Hijab Pendek",
		StockQuantity: 500,
	}
	bg := product{
		Name:          "Baju Gamis",
		StockQuantity: 100,
	}
	bgs := product{
		Name:          "Baju Gamis Syar'i",
		StockQuantity: 200,
	}
	bs := product{
		Name:          "Baju Muslim",
		StockQuantity: 300,
	}

	// Available product
	st := store{
		Name:             "Kitara Store",
		AvailableProduct: []product{},
		Order:            []order{},
	}

	// Add Product to Kitara Store
	st.addProduct(hs)
	st.addProduct(hp)
	st.addProduct(bg)
	st.addProduct(bgs)
	st.addProduct(bs)

	//sarah := order{
	//	CustomerName: "Sarah",
	//	OrderProduct: []orderProduct{
	//		{
	//			Product:       hs,
	//			OrderQuantity: 1,
	//		},
	//		{
	//			Product:       hp,
	//			OrderQuantity: 2,
	//		},
	//		{
	//			Product:       hs,
	//			OrderQuantity: 5,
	//		},
	//	},
	//}
	//
	//st.addOrder(sarah)

	fmt.Println(st.AvailableProduct)
	st.deleteProduct(bgs)
	fmt.Println(st.AvailableProduct)

	r := gin.Default()
	// The JSON of Product for CREATE, EDIT and DELETE should be like this

	/*
		{
			"name": "Baju Anak",
			"stock_quantity": 1
		}
	*/

	/*
	 Yes of course, DELETE also used that JSON because no ID provided, but you don't need to provide
	 stock_quantity while deleting product
	*/
	r.POST("/product", func(c *gin.Context) {
		var p product
		if err := c.Bind(&p); err != nil {
			c.JSON(400, gin.H{"message": "Incorrect JSON body"})
			return
		} else {
			if st.addProduct(p) {
				c.JSON(200, gin.H{"message": "Success"})
			} else {
				c.JSON(500, gin.H{"message": "Product with name " + p.Name + " is exist"})
			}
		}
	})
	r.PUT("/product", func(c *gin.Context) {
		var p product
		if err := c.Bind(&p); err != nil {
			c.JSON(400, gin.H{"message": "Incorrect JSON body"})
			return
		} else {
			if st.editProduct(p) {
				c.JSON(200, gin.H{"message": "Success edit product"})
			} else {
				c.JSON(500, gin.H{"message": "Product with name " + p.Name + " is not exist"})
			}
		}
	})
	r.DELETE("/product", func(c *gin.Context) {
		var p product
		if err := c.Bind(&p); err != nil {
			c.JSON(400, gin.H{"message": "Incorrect JSON body"})
			return
		} else {
			if st.deleteProduct(p) {
				c.JSON(200, gin.H{"message": "Success delete product"})
			} else {
				c.JSON(500, gin.H{"message": "Product with name " + p.Name + " is not exist"})
			}
		}
	})
	r.GET("/product", func(c *gin.Context) {
		c.JSON(200, st.AvailableProduct)
	})

	// The JSON of Order should be like this

	/*
		{
			"customer_name": "Sarah",
			"order_product": [
				{
					"product": {
						"name": "Hijab Pendek"
					},
					"order_quantity": 2
				}
			]
		}
	*/

	r.POST("/order", func(c *gin.Context) {
		var o order
		if err := c.Bind(&o); err != nil {
			c.JSON(400, gin.H{"message": "Incorrect JSON body"})
			return
		} else {
			err := st.addOrder(o)
			if err == nil {
				c.JSON(200, gin.H{"message": "Success create an order"})
			} else {
				c.JSON(500, gin.H{"message": err.Error()})
			}
		}
	})
	r.GET("/order", func(c *gin.Context) {
		c.JSON(200, st.Order)
	})
	_ = r.Run()
}
