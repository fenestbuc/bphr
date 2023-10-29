package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	sc "github.com/hyperledger/fabric-protos-go/peer"
)

type BPHRChaincode struct {
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Transaction struct {
	UserID    string  `json:"userId"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

type Reward struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type Outlet struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Purchase struct {
	UserID         string `json:"userID"`
	OutletID       string `json:"outletID"`
	DateOfPurchase string `json:"dateOfPurchase"`
	IsApproved     bool   `json:"isApproved"`
}

// Helper struct to sort purchases by date
type PurchasesByDate []Purchase

func (a PurchasesByDate) Len() int           { return len(a) }
func (a PurchasesByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PurchasesByDate) Less(i, j int) bool { return a[i].DateOfPurchase < a[j].DateOfPurchase }

func (s *BPHRChaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *BPHRChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case "addUser":
		return s.addUser(stub, args)
	case "updateUser":
		return s.updateUser(stub, args)
	case "addTransaction":
		return s.addTransaction(stub, args)
	case "addReward":
		return s.addReward(stub, args)
	case "registerOutlet":
		return s.registerOutlet(stub, args)
	case "registerPurchase":
		return s.registerPurchase(stub, args)
	case "approvePurchase":
		return s.approvePurchase(stub, args)
	case "redeemReward":
		return s.redeemReward(stub, args)
	default:
		return shim.Error("Invalid function name.")
	}
}

func (s *BPHRChaincode) addUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("incorrect number of arguments. expecting 4")
	}

	var user = User{ID: args[0], Name: args[1], Age: atoi(args[2]), Address: args[3]}
	userAsBytes, _ := json.Marshal(user)
	stub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *BPHRChaincode) updateUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("incorrect number of arguments. expecting 4")
	}

	userAsBytes, _ := stub.GetState(args[0])
	if userAsBytes == nil {
		return shim.Error("could not locate user")
	}
	var user User
	json.Unmarshal(userAsBytes, &user)

	user.Name = args[1]
	user.Age = atoi(args[2])
	user.Address = args[3]

	userAsBytes, _ = json.Marshal(user)
	stub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *BPHRChaincode) addTransaction(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("incorrect number of arguments. expecting 4")
	}

	var txn = Transaction{UserID: args[0], Type: args[1], Amount: atof(args[2]), Timestamp: args[3]}
	txnID := "txn" + args[3] // Using timestamp for unique transaction ID
	txnAsBytes, _ := json.Marshal(txn)
	stub.PutState(txnID, txnAsBytes)

	return shim.Success(nil)
}

func atoi(arg string) int {
	val, err := strconv.Atoi(arg)
	if err != nil {
		return 0
	}
	return val
}

func atof(arg string) float64 {
	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0
	}
	return val
}

func (s *BPHRChaincode) addReward(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 (RewardID, RewardName)")
	}
	reward := Reward{ID: args[0], Name: args[1], Owner: "Ashoka University"}
	rewardAsBytes, _ := json.Marshal(reward)
	stub.PutState(args[0], rewardAsBytes)
	return shim.Success(nil)
}

func (s *BPHRChaincode) registerOutlet(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 (OutletID, OutletName)")
	}
	outlet := Outlet{ID: args[0], Name: args[1]}
	outletAsBytes, _ := json.Marshal(outlet)
	stub.PutState(args[0], outletAsBytes)
	return shim.Success(nil)
}

func (s *BPHRChaincode) registerPurchase(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3 (UserID, OutletID, DateOfPurchase)")
	}
	purchase := Purchase{
		UserID:         args[0],
		OutletID:       args[1],
		DateOfPurchase: args[2],
		IsApproved:     false,
	}
	purchaseAsBytes, _ := json.Marshal(purchase)
	purchaseKey := "purchase" + args[0] + args[1] + args[2]
	stub.PutState(purchaseKey, purchaseAsBytes)
	return shim.Success(nil)
}

func (s *BPHRChaincode) approvePurchase(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3 (UserID, OutletID, DateOfPurchase)")
	}
	purchaseKey := "purchase" + args[0] + args[1] + args[2]
	purchaseAsBytes, err := stub.GetState(purchaseKey)
	if err != nil {
		return shim.Error("Purchase not found.")
	}
	purchase := Purchase{}
	json.Unmarshal(purchaseAsBytes, &purchase)
	purchase.IsApproved = true
	updatedPurchaseAsBytes, _ := json.Marshal(purchase)
	stub.PutState(purchaseKey, updatedPurchaseAsBytes)
	return shim.Success(nil)
}

func (s *BPHRChaincode) redeemReward(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 (UserID, RewardID)")
	}
	userID := args[0]
	rewardID := args[1]
	if !hasSevenConsecutivePurchases(stub, userID) {
		return shim.Error("User does not have 7 consecutive validated purchases.")
	}
	rewardAsBytes, err := stub.GetState(rewardID)
	if err != nil {
		return shim.Error("Reward not found.")
	}
	reward := Reward{}
	json.Unmarshal(rewardAsBytes, &reward)
	reward.Owner = userID
	updatedRewardAsBytes, _ := json.Marshal(reward)
	stub.PutState(rewardID, updatedRewardAsBytes)
	return shim.Success(nil)
}

func hasSevenConsecutivePurchases(stub shim.ChaincodeStubInterface, userID string) bool {
	resultsIterator, err := stub.GetStateByPartialCompositeKey("purchase", []string{userID})
	if err != nil {
		return false
	}
	defer resultsIterator.Close()
	var purchases []Purchase
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return false
		}
		var purchase Purchase
		json.Unmarshal(queryResponse.Value, &purchase)
		if purchase.IsApproved {
			purchases = append(purchases, purchase)
		}
	}
	sort.Sort(PurchasesByDate(purchases))
	consecutiveCount := 0
	for i := 0; i < len(purchases)-1; i++ {
		t1, _ := time.Parse("2006-01-02", purchases[i].DateOfPurchase)
		t2, _ := time.Parse("2006-01-02", purchases[i+1].DateOfPurchase)
		if t2.Sub(t1).Hours() == 24 {
			consecutiveCount++
			if consecutiveCount == 6 {
				return true
			}
		} else {
			consecutiveCount = 0
		}
	}
	return false
}

func main() {
	err := shim.Start(new(BPHRChaincode))
	if err != nil {
		fmt.Printf("Error starting BPHRChaincode: %s", err)
	}
}
