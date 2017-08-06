import logging
import requests
import json
import sys
BASE_URL_ERP = "http://xxxx:port/api/v1/"
BASE_URL_ELS = "http://xxxx:port/eslhttpservice/pushgoods/"
debug = 0
if len(sys.argv)>1:
    debug = sys.argv[1]
    print(debug)
def get_commodity():
        reqdata = {}
        reqdata['x-access-token'] = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjoiSFowODgwOCJ9.hpuBiupMSHaEu9LKOOms8PknZ1cXqxi1uBVoMMMiAIQ"
        reqdata['command'] = 'query_all_commodity'
        r = requests.post(BASE_URL_ERP + 'post_common/', json=reqdata)
        result = r.json()
        if r.status_code == 200:
            if debug == "1":
                print("get response == ", json.dumps(result, sort_keys=True, indent=4, ensure_ascii=False))
            return result
        else:
            logging.error('get_commodity fail! error msg is : %s' % json.dumps(result, sort_keys=True, ensure_ascii=False))
            if debug == "1":
                print("get_commodity fail! error msg is :", result)
            return None

def organize_commodity(commodity_resp):
        array = commodity_resp['output']
        ol = [[]for i in range(len(array))]
        comm_id =0 
        for t in array:
            ol[comm_id].append({'PropertyName': 'GoodsCode', 'Value': t['barcode']})
            ol[comm_id].append({'PropertyName': '商品名称', 'Value': t['name']})
            ol[comm_id].append({'PropertyName': "产地", 'Value': " "})
            ol[comm_id].append({'PropertyName': "规格", 'Value': " "})
            ol[comm_id].append({'PropertyName': "计价单位", 'Value': " "})
            ol[comm_id].append({'PropertyName': "等级", 'Value': " "})
            ol[comm_id].append({'PropertyName': "零售价", 'Value': t['price']})
            ol[comm_id].append({'PropertyName': "会员价", 'Value': " "})
            ol[comm_id].append({'PropertyName': "货号", 'Value': " "})
            ol[comm_id].append({'PropertyName': '条码', 'Value': t['barcode']})
            ol[comm_id].append({'PropertyName': "二维码", 'Value': " "})
            ol[comm_id].append({'PropertyName': "缺货标记", 'Value': " "})
            ol[comm_id].append({'PropertyName': "促销开始时间", 'Value': " "})
            ol[comm_id].append({'PropertyName': "促销结束时间", 'Value': " "})
            ol[comm_id].append({'PropertyName': "原价", 'Value': " "})
            ol[comm_id].append({'PropertyName': "现价", 'Value': " "})
            ol[comm_id].append({'PropertyName': "货架号", 'Value': " "})
            ol[comm_id].append({'PropertyName': "最低库存", 'Value': " "})
            ol[comm_id].append({'PropertyName': "最高库存", 'Value': " "})
            ol[comm_id].append({'PropertyName': "长宽高", 'Value': " "})
            ol[comm_id].append({'PropertyName': "当前库存", 'Value': " "})
            ol[comm_id].append({'PropertyName': "默认供应商代码", 'Value': " "})
            ol[comm_id].append({'PropertyName': "默认供应商名称", 'Value': " "})
            ol[comm_id].append({'PropertyName': "是否促销", 'Value': " "})
            ol[comm_id].append({'PropertyName': "活动时间", 'Value': " "})
            comm_id+=1
        if debug == "1":
            print("organize_commodity == ", json.dumps(ol, sort_keys=True, indent=4, ensure_ascii=False))
        return ol

def sync_commodity():
            ret = get_commodity()
            if ret is None:
                logging.info('Get commodity failed !')
                print("-3")
                return
            if int(ret['nums']) < 1:  # 一条记录也没有取到
                print("-2")
                logging.info('There is no commodity in erp system')
                return
            commodity_list = organize_commodity(ret)
            if debug == "1":
                print("send post == ", json.dumps(commodity_list, sort_keys=True, indent=4, ensure_ascii=False))
            for commodity in commodity_list:            
                try:
                    r = requests.post(BASE_URL_ELS, json=commodity)
                except Exception as e:
                    print("-4")
                    print(e)
                    serr = traceback.format_exc()
                    continue
                if r.status_code != 200:
                    print("-5")
                    if debug == "1":
                         print("sync error msg: ", r.text)
                    continue
                result = r.json()
                if result['result'] != "succeeded":
                   # 本次同步失败， 不保存任何数据
                      print("-6")
                      if debug == "1":
                           print("sync error msg: ", r.text,commodity[0])
                      continue
                if debug == "1":
                     print("got response == ", json.dumps(result, sort_keys=True, indent=4, ensure_ascii=False))

#import pdb
def main():
#      pdb.set_trace()
      sync_commodity()
main()
