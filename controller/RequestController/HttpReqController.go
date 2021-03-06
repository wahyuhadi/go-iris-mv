package RequestController

import (
	"encoding/json"
	"fmt"
	"log"

	"go-iris-mv/service"

	"github.com/kataras/iris"
)

//---------------------------------------------------
// Berikut adalah contoh hasil dari api yang akan diakses
//---------------------------------------------------
/**

	example api get from url
	{
    "message": "success",
    "people": [
        {
            "craft": "ISS",
            "name": "Oleg Artemyev"
        },
        {
            "craft": "ISS",
            "name": "Andrew Feustel"
        },
        {
            "craft": "ISS",
            "name": "Richard Arnold"
        },
        {
            "craft": "ISS",
            "name": "Sergey Prokopyev"
        },
        {
            "craft": "ISS",
            "name": "Alexander Gerst"
        },
        {
            "craft": "ISS",
            "name": "Serena Aunon-Chancellor"
        }
    ],
    "number": 6
}

**/

//---------------------------------------------------
// membuat file struct dari hasil
//---------------------------------------------------
// thanks to https://play.golang.org/p/rAJfkD1i7n or  https://stackoverflow.com/questions/34489887/go-unmarshal-json-nested-array-of-objects
type DataPoeple struct {
	Number int `json:"number"` // get field number from url
	People []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
}

//---------------------------------------------------
// Fungsi untuk Get HTTP
//---------------------------------------------------
func GetHttpReq(ctx iris.Context) {
	var (
		result iris.Map
	)
	data := service.HttpReqGet("http://api.open-notify.org/astros.json")
	ctx.ReadJSON(&data)

	// parsing response json
	Data := DataPoeple{}
	jsonErr := json.Unmarshal(data, &Data)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(Data.Number)
	fmt.Println(Data.People[0].Name) // yuou can check name in array ....

	result = iris.Map{
		"code":   iris.StatusOK,
		"result": Data,
	}
	ctx.JSON(result)
	return
}
