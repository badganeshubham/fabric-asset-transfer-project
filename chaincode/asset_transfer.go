// chaincode.go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Asset struct {
    DEALERID    string `json:"DEALERID"`
    MSISDN      string `json:"MSISDN"`
    MPIN        string `json:"MPIN"`
    BALANCE     int    `json:"BALANCE"`
    STATUS      string `json:"STATUS"`
    TRANSAMOUNT int    `json:"TRANSAMOUNT"`
    TRANSTYPE   string `json:"TRANSTYPE"`
    REMARKS     string `json:"REMARKS"`
}

type SmartContract struct {
    contractapi.Contract
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    assets := []Asset{
        {DEALERID: "D001", MSISDN: "9990001111", MPIN: "1234", BALANCE: 1000, STATUS: "Active", TRANSAMOUNT: 0, TRANSTYPE: "", REMARKS: "Initial Dealer"},
        {DEALERID: "D002", MSISDN: "9990002222", MPIN: "5678", BALANCE: 1500, STATUS: "Active", TRANSAMOUNT: 0, TRANSTYPE: "", REMARKS: "Initial Dealer"},
    }

    for _, asset := range assets {
        assetJSON, err := json.Marshal(asset)
        if err != nil {
            return err
        }

        err = ctx.GetStub().PutState(asset.MSISDN, assetJSON)
        if err != nil {
            return fmt.Errorf("failed to put asset to world state: %v", err)
        }
    }

    return nil
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID, msisdn, mpin string, balance int, status, transtype, remarks string, transAmount int) error {
    asset := Asset{
        DEALERID:    dealerID,
        MSISDN:      msisdn,
        MPIN:        mpin,
        BALANCE:     balance,
        STATUS:      status,
        TRANSAMOUNT: transAmount,
        TRANSTYPE:   transtype,
        REMARKS:     remarks,
    }

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(msisdn, assetJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, msisdn string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(msisdn)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("asset with MSISDN %s does not exist", msisdn)
    }

    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, msisdn, dealerID, mpin string, balance int, status, transtype, remarks string, transAmount int) error {
    asset, err := s.ReadAsset(ctx, msisdn)
    if err != nil {
        return err
    }

    asset.DEALERID = dealerID
    asset.MPIN = mpin
    asset.BALANCE = balance
    asset.STATUS = status
    asset.TRANSAMOUNT = transAmount
    asset.TRANSTYPE = transtype
    asset.REMARKS = remarks

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(msisdn, assetJSON)
}

func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, msisdn string) error {
    exists, err := s.AssetExists(ctx, msisdn)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("asset %s does not exist", msisdn)
    }

    return ctx.GetStub().DelState(msisdn)
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, msisdn string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(msisdn)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }

    return assetJSON != nil, nil
}

func (s *SmartContract) GetAssetHistory(ctx contractapi.TransactionContextInterface, msisdn string) ([]*Asset, error) {
    resultsIterator, err := ctx.GetStub().GetHistoryForKey(msisdn)
    if err != nil {
        return nil, fmt.Errorf("failed to get history for asset: %v", err)
    }
    defer resultsIterator.Close()

    var history []*Asset
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset Asset
        if len(response.Value) > 0 {
            err = json.Unmarshal(response.Value, &asset)
            if err != nil {
                return nil, err
            }
            history = append(history, &asset)
        }
    }
    return history, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        fmt.Printf("Error creating chaincode: %s", err)
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting chaincode: %s", err)
    }
}
