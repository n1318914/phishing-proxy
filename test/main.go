package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func SetJSONVariable(body []byte, key string, value interface{}) ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	data[key] = value
	newBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return newBody, nil
}

func SetJSONVariableDeep(body []byte, key string, value interface{}) ([]byte, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	keys := strings.Split(key, ".")
	m := data
	for i := 0; i < len(keys)-1; i++ {
		if _, ok := m[keys[i]]; !ok {
			m[keys[i]] = map[string]interface{}{}
		}
		m = m[keys[i]].(map[string]interface{})
	}
	m[keys[len(keys)-1]] = value

	return json.Marshal(data)
}

func main() {
	body := []byte(`{"Detail":{"eventTimestamp":"2025-10-20T02:32:39.564Z","eventType":"payment_info_submitted","eventId":"sh-ff76e666-7687-4D8F-9473-6FDB53D55FDD","storeName":"chem-guys.myshopify.com","shopId":"chem-guys.myshopify.com","userId":"ig_205f6551293dd1c7f87b7fbaa657522d48d6","shopifyClientId":"3df08c07-cdb1-4bd9-b41c-c6642ea362b8","igVars":"{\"0b7426493aef\":\"ac39f188e23b\",\"57bebd2e3b31\":\"_UNASSIGNED\",\"d8a05dbf90fe\":\"_UNASSIGNED\",\"423d9d712620\":\"_UNASSIGNED\",\"a7e00be29341\":\"082bf88d0a05\",\"89017b46cbb8\":\"_UNASSIGNED\",\"cc8062e23adc\":\"_UNASSIGNED\",\"f2409b937cc0\":\"b180d8dfcb45\",\"a553baa1788d\":\"_UNASSIGNED\",\"83d450abd69b\":\"22e68aae27a7\",\"d65536a51ea7\":\"20a9e8bdc2f8\",\"73993a31c807\":\"_UNASSIGNED\",\"redirectedFrom\":\"\"}","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:144.0) Gecko/20100101 Firefox/144.0","httpReferrer":"https://www.chemicalguys.fake.com/","pageTitle":"Checkout - Chemical Guys","customer":null,"cart":{"id":"hWN4JfVD6oExShW9tr0WIxvN","igId":"ig_205f6551293dd1c7f87b7fbaa657522d48d6"},"location":{"origin":"https://www.chemicalguys.fake.com","host":"www.chemicalguys.fake.com","pathname":"/checkouts/cn/hWN4JfVD6oExShW9tr0WIxvN/en-us","search":"?auto_redirect=false&edge_redirect=true&skip_shop_pay=true","hash":""},"checkoutToken":"b6dbe1119b7a787690b3457e84aa5376","phone":"","email":"test@proton.me","transactions":[{"amount":{"amount":26.93,"currencyCode":"USD"},"gateway":"","paymentMethod":{"type":"creditCard","name":"MASTERCARD"}}],"billingLocation":{"city":"Los Angeles","country":"US","countryCode":"US","province":"CA","provinceCode":"CA","zip":"90001"},"shippingLocation":{"city":"Los Angeles","country":"US","countryCode":"US","province":"CA","provinceCode":"CA","zip":"90001"},"deliveryOptionsSelected":[{"cost":{"amount":4.99,"currencyCode":"USD"},"costAfterDiscounts":{"amount":4.99,"currencyCode":"USD"},"description":null,"handle":"b6dbe1119b7a787690b3457e84aa5376-98e2be58e947caa314ec75f5a2c0312d","title":"Standard Flat Rate","type":"shipping"}]}}`)

	newBody, err := SetJSONVariableDeep(body, "Detail.location.origin", 11111111)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(newBody))
}
