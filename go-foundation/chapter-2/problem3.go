package main
import "fmt"
import "golang.org/x/crypto/bcrypt"
func main() {
	pass:="mypasswd"
	res,err:=bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
 if(err != nil){
	fmt.Println("failed to bcrypt")
	return
 }
	fmt.Printf("%x",res)
	isVerified := bcrypt.CompareHashAndPassword(res,[]byte(pass))
	if isVerified !=nil{
		fmt.Println("Not Verified")
		return
	}
	fmt.Println("\nVerified Status: Success ")

}