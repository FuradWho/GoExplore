package fabric_ca

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func InitCaClient()  {
	sdk, err := fabsdk.New(config.FromFile("./connect-config/org2-config.yaml"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("???")
	ctx := sdk.Context()

	client, err := msp.New(ctx)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.GetCAInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.CAName)

	//req := &msp.RegistrationRequest{
	//	Name: "User1",
	//	Type: "client",
	//	CAName: "",
	//	Secret: "123456",
	//}
	//
	//register, err := client.Register(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(register)
}