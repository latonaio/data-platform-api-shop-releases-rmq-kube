package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-shop-releases-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-shop-releases-rmq-kube/DPFM_API_Output_Formatter"

	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.Shop = %d ", input.Header.Shop)
	rows, err := c.db.Query(
		`SELECT 
			header.Shop
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_shop_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
