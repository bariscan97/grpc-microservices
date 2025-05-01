package server

import (
	"time"

	"github.com/bariscan97/ecomm/ecomm-grpc/pb"
	"github.com/bariscan97/ecomm/ecomm-grpc/repository"
	"github.com/bariscan97/ecomm/pkg/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func torepositoryProduct(p *pb.ProductReq) *repository.Product {
	return &repository.Product{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}

func toPBProductRes(p *repository.Product) *pb.ProductRes {
	res := &pb.ProductRes{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
		CreatedAt:    timestamppb.New(p.CreatedAt),
	}
	if p.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*p.UpdatedAt)
	}

	return res
}

func patchProductReq(product *repository.Product, p *pb.ProductReq) {
	if p.Name != "" {
		product.Name = p.Name
	}
	if p.Image != "" {
		product.Image = p.Image
	}
	if p.Category != "" {
		product.Category = p.Category
	}
	if p.Description != "" {
		product.Description = p.Description
	}
	if p.Rating != 0 {
		product.Rating = p.Rating
	}
	if p.NumReviews != 0 {
		product.NumReviews = p.NumReviews
	}
	if p.Price != 0 {
		product.Price = p.Price
	}
	if p.CountInStock != 0 {
		product.CountInStock = p.CountInStock
	}
	product.UpdatedAt = toTimePtr(time.Now())
}

func toTimePtr(t time.Time) *time.Time {
	return &t
}

func torepositoryOrder(o *pb.OrderReq) *repository.Order {
	return &repository.Order{
		PaymentMethod: o.PaymentMethod,
		TaxPrice:      o.TaxPrice,
		ShippingPrice: o.ShippingPrice,
		TotalPrice:    o.TotalPrice,
		UserID:        o.UserId,
		Items:         torepositoryOrderItems(o.Items),
	}
}

func torepositoryOrderItems(items []*pb.OrderItem) []repository.OrderItem {
	var res []repository.OrderItem
	for _, i := range items {
		res = append(res, repository.OrderItem{
			Name:      i.Name,
			Quantity:  i.Quantity,
			Image:     i.Image,
			Price:     i.Price,
			ProductID: i.ProductId,
		})
	}
	return res
}

func toPBOrderStatus(os repository.OrderStatus) pb.OrderStatus {
	switch os {
	case repository.Pending:
		return pb.OrderStatus_PENDING
	case repository.Shipped:
		return pb.OrderStatus_SHIPPED
	case repository.Delivered:
		return pb.OrderStatus_DELIVERED
	default:
		return 0
	}
}

func toPBOrderRes(o *repository.Order) *pb.OrderRes {
	res := &pb.OrderRes{
		Id:            o.ID,
		Items:         toPBOrderItems(o.Items),
		PaymentMethod: o.PaymentMethod,
		TaxPrice:      o.TaxPrice,
		ShippingPrice: o.ShippingPrice,
		TotalPrice:    o.TotalPrice,
		Status:        toPBOrderStatus(o.Status),
		CreatedAt:     timestamppb.New(o.CreatedAt),
	}
	if o.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*o.UpdatedAt)
	}

	return res
}

func toPBOrderItems(items []repository.OrderItem) []*pb.OrderItem {
	var res []*pb.OrderItem
	for _, i := range items {
		res = append(res, &pb.OrderItem{
			Name:      i.Name,
			Quantity:  i.Quantity,
			Image:     i.Image,
			Price:     i.Price,
			ProductId: i.ProductID,
		})
	}
	return res
}

func torepositoryUser(u *pb.UserReq) *repository.User {
	return &repository.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		IsAdmin:  u.IsAdmin,
	}
}

func toPBUserRes(u *repository.User) *pb.UserRes {
	return &pb.UserRes{
		Id:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		IsAdmin:  u.IsAdmin,
	}
}

func patchUserReq(user *repository.User, u *pb.UserReq) {
	if u.Name != "" {
		user.Name = u.Name
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Password != "" {
		hashed, err := util.HashPassword(u.Password)
		if err != nil {
			panic(err)
		}
		user.Password = hashed
	}
	if u.IsAdmin {
		user.IsAdmin = u.IsAdmin
	}
	user.UpdatedAt = toTimePtr(time.Now())
}
