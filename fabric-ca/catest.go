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

	//resp, err := client.GetCAInfo()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp.CAName)
	//
	//affiliations, err := client.GetAllIdentities()
	//if err != nil {
	//	return
	//}

	//req := &msp.RegistrationRequest{
	//	Name: "User8",
	//	Type: "client",
	//	CAName: "",
	//	Secret: "123456",
	//}
	////
	//register, err := client.Register(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(register)
	//
	//err = client.Enroll(register)
	//if err != nil {
	//	fmt.Println(err)
	//}

	info, err := client.GetCAInfo()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(info.CAName)
	fmt.Println(info.Version)
	fmt.Println(info.CAChain)
	fmt.Println(info.IssuerPublicKey)
	fmt.Println(info.IssuerRevocationPublicKey)

	//identity, err := client.CreateIdentity(&msp.IdentityRequest{ID: "123", Affiliation: "org2",
	//	Attributes: []msp.Attribute{{Name: "attName1", Value: "attValue1"}}})
	//if err != nil {
	//	fmt.Printf("Create identity return error %s\n", err)
	//	return
	//}
	//fmt.Printf("identity '%s' created\n", identity.ID)


	//identity, err := client.GetIdentity("User4")
	//if err != nil {
	//	fmt.Printf("Get Identity : %s \n",err)
	//}
	//fmt.Println(identity.ID)
	//

	req := &msp.RegistrationRequest{
		Name: "User10",
		Type: "client",
		CAName: "",
		Secret: "123456",
	}
	//
	register, err := client.Register(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(register)

	err = client.Enroll("User10",msp.WithSecret("123456"))
	if err != nil {
		fmt.Println(err)
	}

	signingIdentity, err := client.GetSigningIdentity("User10")
	if err != nil {
		fmt.Printf("GetSigningIdentity : %s \n",err)
	}
	fmt.Println(signingIdentity.PrivateKey())




	//identity, err := client.GetSigningIdentity("User5")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//key := identity.Identifier().ID
	//fmt.Println(key)

	
	//for _ , info := range affiliations{
	//	fmt.Println(info.ID)
	//	fmt.Println(info.Type)
	//	fmt.Println(info.Attributes)
	//	fmt.Println(info.CAName)
	//
	//	fmt.Println("----------------------")
	//}

	
}