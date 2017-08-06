package main
import (
	"fmt"
	"os/exec"
	"encoding/json"
)
const BASE_URL_ELS = "http://xxxx:port/eslhttpservice/pushgoods/"
const BASE_URL_ERP = "http://xxxx:port/api/v1/"
type ErpData struct{
	Command string `json:"command"`
	Output []CommInfo `json:"output"`
	Nums string `json:"nums"`
	Result string `json:"result"`
	Message string `json:"message"`
}
type CommInfo struct{
	Barcode string `json:"barcode"`
	Name string    `json:"name"`
	Price string   `json:"price"`
}
type EslData struct{
	PropertyName string `json:"PropertyName"`
	Value        string `json:"Value"`
}
func main(){
	fmt.Println("/////Get ERP data/////")
	Cmdstring := fmt.Sprintf("curl -X POST %spost_common -H 'cache-control: no-cache' -H 'content-type: application/json' -H 'postman-token: 431adb81-557c-5664-9470-f95d883dbb51' -d '{\n\"x-access-token\":\"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjoiSFowODgwOCJ9.hpuBiupMSHaEu9LKOOms8PknZ1cXqxi1uBVoMMMiAIQ\",\n\"command\":\"query_all_commodity\"}'", BASE_URL_ERP)
	outstr,err:=exec.Command("/bin/sh", "-c", Cmdstring).Output()
	if err == nil{
		//fmt.Println(outstr)
	}
	erpdata := &ErpData{}
	err = json.Unmarshal([]byte(outstr), &erpdata)
	fmt.Println(erpdata)
	fmt.Println("///////reform ERP data to ELS data//////")
	var esllist [][]EslData
	tmpdata := EslData{}
	output := erpdata.Output
	for _, val := range output{
		 var elsdata []EslData
                 tmpdata.PropertyName = "GoodsCode"
                 tmpdata.Value        = val.Barcode
		 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "商品条码"
                 tmpdata.Value        = val.Name
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "产地"
                 tmpdata.Value        = " "
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "规格"
                 tmpdata.Value        = " "
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "计价单位"
                 tmpdata.Value        = " "
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "等级"
                 tmpdata.Value        = " "
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "会员价"
                 tmpdata.Value        = " "
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "零售价" 
                 tmpdata.Value        = val.Price 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "货号" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "条码" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "二维码" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "缺货标记" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "促销开始时间" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "促销结束时间" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "原价" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "现价" 
                 tmpdata.Value        = " "  
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "货架号" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "最低库存" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "最高库存" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "长宽高" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "当前库存" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "默认供应商代码" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "默认供应商名称" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "是否促销" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
                 tmpdata.PropertyName = "活动时间" 
                 tmpdata.Value        = " " 
                 elsdata = append(elsdata, tmpdata)
		 esllist = append(esllist, elsdata)
	}
	fmt.Println("////////sent data to esl///////////")
	for i, val :=range esllist {
		jsondata, err:=json.Marshal(val)	
		if err == nil {
			Cmd := fmt.Sprintf("curl -X POST %s -H 'cache-control: no-cache' -H 'content-type: application/json' -H 'postman-token: 99260967-1b0c-0b4c-e969-53f28244199e' -d '%s'", BASE_URL_ELS, jsondata)
			outstr,err:=exec.Command("/bin/sh", "-c", Cmd).Output()
        		if err == nil{
                		fmt.Println(outstr)
        		}
			fmt.Println(i,jsondata)
		}
	}
}
