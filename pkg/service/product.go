package service

import (
	"context"
	"fmt"
	"log"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/bolt"
	"github.com/nhamtybv/test_kit_bo/pkg/repository/oracle"
	"go.etcd.io/bbolt"
)

type ProductService interface {
	Syns(ctx context.Context) error
	FindByNumber(ctx context.Context, product_number string) (*entity.Product, error)
	FindAll(ctx context.Context) (*entity.ProductList, error)
	FindAllAgents(ctx context.Context) (*entity.Agent, error)
}

type productSrv struct {
	oraRepo  repository.ProductRepository
	boltRepo repository.ProductRepository
}

func NewProductService(db *bbolt.DB) ProductService {

	prdBolt := bolt.NewProductRepo(db)
	connStr, err := prdBolt.GetConnection(context.Background())
	if err != nil {
		log.Println("WARNING: Oracle connection wasnot setted up")
	}
	log.Printf("Oracle connection: %s", connStr)
	prdOra := oracle.NewProductRepoOrcl(connStr)

	return &productSrv{
		oraRepo:  prdOra,
		boltRepo: prdBolt,
	}
}

// FindAll implements ProductService
func (p *productSrv) FindAll(ctx context.Context) (*entity.ProductList, error) {
	data, err := p.boltRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get products, error: %w", err)
	}
	return data, nil
}

// FindByNumber implements ProductService
func (p *productSrv) FindByNumber(ctx context.Context, product_number string) (*entity.Product, error) {
	panic("unimplemented")
}

// Syns implements ProductService
func (p *productSrv) Syns(ctx context.Context) error {
	log.Println("SERVICE => DEBUG: call syns products")
	data, err := p.oraRepo.FindAll(ctx)
	if err != nil {
		return fmt.Errorf("getting product from database, error: %w", err)
	}

	prds := &entity.ProductList{
		Count:    0,
		Products: []entity.Product{},
	}

	for i := 0; i < data.Count; i++ {
		if data.Products[i].ParentID == -1 {
			p := data.Products[i]
			getChildrenProducts(&p, data)
			prds.Count += 1
			prds.Products = append(prds.Products, p)
		}
	}

	err = p.boltRepo.Save(ctx, prds)
	if err != nil {
		return fmt.Errorf("syns products error: %w", err)
	}

	log.Println("SERVICE => DEBUG: call syns agent")

	agent, err := p.oraRepo.FindAgent(ctx)
	if err != nil {
		return fmt.Errorf("getting product from database, error: %w", err)
	}

	err = p.boltRepo.SaveAgent(ctx, agent)
	if err != nil {
		return fmt.Errorf("syns agent error: %w", err)
	}
	log.Println("SERVICE => DEBUG: syns finished")
	return nil
}

// FindAllAgents implements ProductService
func (p *productSrv) FindAllAgents(ctx context.Context) (*entity.Agent, error) {
	data, err := p.boltRepo.FindAgent(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get products, error: %w", err)
	}
	return data, nil
}

func getChildrenProducts(p *entity.Product, ps *entity.ProductList) {
	for _, v := range ps.Products {
		if v.ParentID == p.ID {
			getChildrenProducts(&v, ps)
			p.Children = append(p.Children, v)
		}
	}
}
