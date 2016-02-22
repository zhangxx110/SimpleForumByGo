package logger
import(
	"fmt"
	"utils"
)
func Println(a ...interface{}){
	if utils.ISDEBUG{
		fmt.Println(a)
	}
}