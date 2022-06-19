package oracle

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"

	"github.com/nhamtybv/test_kit_bo/pkg/database"
	"github.com/nhamtybv/test_kit_bo/pkg/entity"
	"github.com/nhamtybv/test_kit_bo/pkg/repository"
)

var productQuery = `
		select json_object(
			  'product_id' value p.id
			, 'product_type' value p.product_type
			, 'contract_type' value p.contract_type
			, 'parent_id' value nvl(p.parent_id, -1)
			, 'inst_id' value p.inst_id
			, 'status' value p.status
			, 'product_number' value p.product_number
			, 'split_hash' value p.split_hash
			, 'product_name' value get_text(
				i_table_name  => 'prd_product'
				, i_column_name => 'label'
				, i_object_id   => p.id
				, i_lang        => 'LANGENG'
			)
			, 'product_services' value (select json_arrayagg(json_object( 'service_id' value service_id, 'entity_type' value pst.entity_type, 'service_name' value get_text(i_table_name  => 'prd_service'
																		, i_column_name => 'label'
																		, i_object_id   => service_id
																		, i_lang        => 'LANGENG'
																))) from prd_product_service pps 
																    join prd_service psr on pps.service_id = psr.id 
																	join prd_service_type pst on psr.service_type_id = pst.id 
																   where pps.product_id = p.id)
			, 'product_account_types' value (select json_arrayagg(json_object('account_type' value account_type, 'currency' value currency)) from acc_product_account_type aat where aat.product_id = p.id)
			, 'card_types' value (select json_arrayagg(json_object('card_type_id' value card_type_id, 'card_type' value get_text (
																							i_table_name  => 'net_card_type'
																							, i_column_name => 'name'
																							, i_object_id   => pct.card_type_id
																							, i_lang        => 'LANGENG'
																							))) from iss_product_card_type pct where pct.product_id = p.id) absent on null
			)
			as json_data
			
			from prd_product p
			where p.product_type = 'PRDT0100'
			order by p.parent_id nulls first, p.id
	`

type productRepo struct {
	strConn string
}

func NewProductRepoOrcl(strConn string) repository.ProductRepository {
	return &productRepo{strConn: "oracle://" + strConn}
}

// FindAll implements repository.ProductRepository
func (p *productRepo) FindAll(ctx context.Context) (*entity.ProductList, error) {
	log.Println("REPO => DEBUG: call syns function")
	tdb, err := database.NewOracleConnection(p.strConn)
	if err != nil {
		return nil, fmt.Errorf("prepare oracle connection error: %w", err)
	}
	defer tdb.Close()

	log.Println("REPO => DEBUG: prepare query")

	stmt, err := tdb.Prepare(productQuery)
	if err != nil {
		return nil, fmt.Errorf("prepare query error: %w", err)
	}
	defer stmt.Close()

	log.Println("REPO => DEBUG: running query")
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	lst := &entity.ProductList{
		Count:    0,
		Products: []entity.Product{},
	}

	for rows.Next() {
		tmp := ""
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning product error: %w", err)
		}

		tmpPrd := entity.Product{}
		err = json.Unmarshal([]byte(tmp), &tmpPrd)
		if err != nil {
			return nil, fmt.Errorf("unmarshal product error: %w", err)
		}

		lst.Count += 1

		lst.Products = append(lst.Products, tmpPrd)
	}
	log.Println("REPO => DEBUG: query finished")
	return lst, nil
}

// FindByNumber implements repository.ProductRepository
func (p *productRepo) FindByNumber(ctx context.Context, product_number string) (*entity.Product, error) {
	return nil, fmt.Errorf("unimplemented")
}

// Save implements repository.ProductRepository
func (p *productRepo) Save(ctx context.Context, c *entity.ProductList) error {
	return fmt.Errorf("unimplemented")
}

// GetConnection implements repository.ProductRepository
func (*productRepo) GetConnection(ctx context.Context) (string, error) {
	panic("unimplemented")
}
