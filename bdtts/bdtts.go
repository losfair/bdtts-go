package bdtts

import "net/http"
import "net/url"
import "math/rand"
import "strconv"
import "errors"
import "io/ioutil"

var cached_output map[string]interface{}

func Request(access_token,text string) ([]byte,error) {
	cc,ok := cached_output[text]
	if ok {
		return cc.([]byte),nil
	}

	u,_ := url.Parse("http://tsn.baidu.com/text2audio")
	q := u.Query()
	q.Set("tok",access_token)
	q.Set("tex",text)
	q.Set("lan","zh")
	q.Set("ctp","1")
	q.Set("cuid",strconv.Itoa(rand.Intn(100000)))
	u.RawQuery = q.Encode()
	res,err := http.Get(u.String())
	if err!=nil {
		return nil,err
	}
	result,err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err!=nil {
		return nil,err
	}
	if string(result[0:1])=="[" || string(result[0:1])=="{" {
		return nil,errors.New("TTS backend failed")
	}
//	log.Println("New cache for:",text)
	cached_output[text]=result
	return result,nil
}

