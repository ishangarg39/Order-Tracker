package batch

import (
	dbconnection "OrderTracker/dbConnection"
	"OrderTracker/models"
	"fmt"

	"github.com/Deeptiman/go-batch"
	"github.com/gin-gonic/gin"
)

type SalesResource struct {
	// This struct will represent a sales batch in the database.
	ID         int64 `json:"id"`
	SupplierId int64 `json:"supplierId"`
	TotalSales int   `json:"totalSales"`
}

func SalesBatching(context *gin.Context) {
	// This function will batch sales in the database.
	// This function will be called from the controllers.go file.

	// Batch the sales here.
	orderSlice, err := models.GetOrderQuery()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	bt := batch.NewBatch(batch.WithMaxItems(5))
	go bt.StartBatchProcessing()

	salesMap := make(map[int64]int)

	//go routine for sales
	go func() {
		for item := range bt.Item {
			order := item.(models.Order)
			salesMap[order.SupplierId]++
		}
	}()

	for _, order := range orderSlice {
		//passing to channel to do batch processing
		bt.Item <- order
	}

	bt.Close()
	batchTransaction(salesMap)
}

func batchTransaction(salesMap map[int64]int) error {
	// This function will batch the transactions in the database.
	// This function will be called from the SalesBatching function.

	// Batch the transactions here.

	tx, err := dbconnection.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO sales (supplier_id, total_sales) VALUES"

	//interface slice to hold the all values
	values := []interface{}{}

	//placeholder to create place for transaction
	placeholder := ""

	for supplierId, totalSales := range salesMap {

		placeholder = fmt.Sprintf("($%d, $%d),", len(values)+1, len(values)+2)

		values = append(values, supplierId, totalSales)
	}

	//removing the last comma
	query += placeholder[:len(placeholder)-1]

	_, err = tx.Exec(query, values...)

	if err != nil {
		return err
	}
	return tx.Commit()
}
