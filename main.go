package main

import	(
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"strings"
)
var phone,captcha string
func main(){
	fmt.Println("【流量领取助手】By：吾爱破解-永彬后花园")
	fmt.Printf("联通用户每次可领取30MB全国流量,每月可领取10次。提交领取申请后,流量会在24小时内充值到账，领取成功后可登录联通手机营业厅APP查询,流量自到帐后即时生效。有效期为三个月,每月第一天和最后一天不能领取流量,其他时间均可正常领取。\n\n")
	runmain()
}
func runmain(){

	for	{
	fmt.Printf("请输入手机号码：")
		_, err := fmt.Scanf("%s\n", &phone)
		if err != nil {
			fmt.Println(err)
		}else if len(phone)!=11{
			fmt.Println("手机号码格式错误，请重新输入。")
		}else{
			getCode()
			break
		}

	}
}

func getCode(){
	url_getCode := "https://m.10010.com/god/AirCheckMessage/sendCaptcha"
	form := url.Values{}
	form.Add("phoneVal", phone)
	form.Add("type", "21")
	resp, err := http.PostForm(url_getCode, form)
	if err != nil {
			fmt.Println(err)
	} else {
		getRes, gerErr := ioutil.ReadAll(resp.Body)
		if gerErr != nil {
			fmt.Println(gerErr)
		} else {
			//fmt.Println(string(getRes))
			var GetCode_state bool = strings.Contains(string(getRes), "0000")
			//fmt.Println("(验证码发送状态：",strings.Contains(string(getRes), "0000")) //true
			if GetCode_state == true {
				fmt.Println("验证码发送成功")
				getCOde_scanf()
			}else{
				if string(getRes) == "{}"{
				fmt.Println("验证码发送失败，60秒内不可重复发送验证码")
				runmain()
				}else{
				fmt.Printf("号码:%s，一天最多只能兑换三次流量，请明天再来\n",phone)
				runmain()
				}
			}


		}

	}
}
func getCOde_scanf(){
	for	{
		fmt.Printf("请输入验证码：")
		_, err := fmt.Scanf("%s\n", &captcha)
		if err != nil {
			fmt.Println(err)
		}else if len(captcha)!=4{
			fmt.Println("验证码格式错误")
			getCOde_scanf()
		}else{
			flowExchange()
			break
		}

	}
}

	

func flowExchange(){
	url_flowExchange := fmt.Sprintf("https://m.10010.com/god/qingPiCard/flowExchange?number=%s&type=%s&captcha=%s", phone, "21", captcha)
	//fmt.Printf("%s\n", url)
	get, err := http.Get(url_flowExchange)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		getRes, err1 := ioutil.ReadAll(get.Body)
		if err != nil {
			fmt.Println(err1)
			return
		} else if strings.Contains(string(getRes), "1001") {
			fmt.Println("验证码输入错误")
			getCOde_scanf()
		}else if strings.Contains(string(getRes), "0000") {
			fmt.Printf("已成功向号码%s充值30M流量，3个月内有效。\n",phone)
			runmain()
		}else {
			fmt.Printf("未知错误，请稍后再试。\n", string(getRes))
			runmain()
		}
	}
}
