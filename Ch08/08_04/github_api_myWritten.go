package main
import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)
const url="https://api.github.com/users/"

type User struct{
	Name string `json:"name"`
	PublicRepos int `json:"public_repos"`
}

func userInfo(login string) (*User,error){
	if login==""{
		return nil,fmt.Errorf("Login can't be empty")
	}
	reqUrl:=url+login
	resp,err:=http.Get(reqUrl)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	user:=&User{}
	enc:=json.NewDecoder(resp.Body)
	if err:=enc.Decode(user);err != nil {
		return nil,err
	}
	return user,nil
}

func main(){
	user,err:=userInfo("zohaibsoomro")
	if err != nil {
		log.Fatalf("Error: %s",err)
	}
	fmt.Printf("%+v",user)
}