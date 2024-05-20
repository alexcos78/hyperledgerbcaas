package chaincode

import (
        "encoding/json"
        "fmt"
        "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an RawDataset
type SmartContract struct {
        contractapi.Contract
}

type Dataset struct {
        ID             string `json:"ID"`
        DatasetName    string `json:"DatasetName"`
        Hash           string `json:"Hash"`
        Owner          string `json:"Owner"`
        FromHash       string `json:"FromHash"`
        Version        string `json:"Version"`
        Algorithm      string `json:"Algorithm"`
        Timestamp      string `json:"Timestamp"`
}

// CreateRawDataset issues a new dataset to the world state with given details.
func (s *SmartContract) CreateDataset(ctx contractapi.TransactionContextInterface, id string, datasetName string, hash string, owner string, fromHash string, version string, algorithm string, timestamp string) error {
        exists, err := s.DatasetExists(ctx, id)
        if err != nil {
                return err
        }
        if exists {
                return fmt.Errorf("the dataset %s already exists", id)
        }

        dataset := Dataset{
                ID:             id,
                DatasetName:    datasetName,
                Hash:           hash,
                Owner:          owner,
                FromHash:       fromHash,
                Version:        version,
                Algorithm:      algorithm,
                Timestamp:      timestamp,
        }
        datasetJSON, err := json.Marshal(dataset)
        if err != nil {
                return err
        }

        return ctx.GetStub().PutState(id, datasetJSON)
}

// ReadRawDataset returns the dataset stored in the world state with given id.
func (s *SmartContract) ReadDataset(ctx contractapi.TransactionContextInterface, id string) (*Dataset, error) {
        datasetJSON, err := ctx.GetStub().GetState(id)
        if err != nil {
                return nil, fmt.Errorf("failed to read from world state: %v", err)
        }
        if datasetJSON == nil {
                return nil, fmt.Errorf("the dataset %s does not exist", id)
        }

        var dataset Dataset
        err = json.Unmarshal(datasetJSON, &dataset)
        if err != nil {
                return nil, err
        }

        return &dataset, nil
}


func (s *SmartContract) ReadDatasetsFromHash(ctx contractapi.TransactionContextInterface, fromHash string) ([]*Dataset, error) {
        // range query with empty string for startKey and endKey does an
        // open-ended query of all datasets in the chaincode namespace.
        resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
        if err != nil {
                return nil, err
        }
        defer resultsIterator.Close()

        var datasets []*Dataset
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return nil, err
                }

                var dataset Dataset
                err = json.Unmarshal(queryResponse.Value, &dataset)
                if err != nil {
                        return nil, err
                }
                if dataset.FromHash == fromHash {
                        datasets = append(datasets, &dataset)
                    }


        }

        return datasets, nil
}

func (s *SmartContract) ReadDatasetsOwner(ctx contractapi.TransactionContextInterface, owner string) ([]*Dataset, error) {
        // range query with empty string for startKey and endKey does an
        // open-ended query of all datasets in the chaincode namespace.
        resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
        if err != nil {
                return nil, err
        }
        defer resultsIterator.Close()

        var datasets []*Dataset
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return nil, err
                }

                var dataset Dataset
                err = json.Unmarshal(queryResponse.Value, &dataset)
                if err != nil {
                        return nil, err
                }
                if dataset.Owner == owner {
                        datasets = append(datasets, &dataset)
                    }


        }

        return datasets, nil
}


// RawDatasetExists returns true when dataset with given ID exists in world state
func (s *SmartContract) DatasetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
        datasetJSON, err := ctx.GetStub().GetState(id)
        if err != nil {
                return false, fmt.Errorf("failed to read from world state: %v", err)
        }

        return datasetJSON != nil, nil
}

// ProcessedDatasetExists returns true when dataset with given ID exists in world state
func (s *SmartContract) ProcessedDatasetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
        datasetJSON, err := ctx.GetStub().GetState(id)
        if err != nil {
                return false, fmt.Errorf("failed to read from world state: %v", err)
        }

        return datasetJSON != nil, nil
}


// GetAllRawDatasets returns all datasets found in world state
func (s *SmartContract) GetAllDatasets(ctx contractapi.TransactionContextInterface) ([]*Dataset, error) {
        // range query with empty string for startKey and endKey does an
        // open-ended query of all datasets in the chaincode namespace.
        resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
        if err != nil {
                return nil, err
        }
        defer resultsIterator.Close()

        var datasets []*Dataset
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return nil, err
                }

                var dataset Dataset
                err = json.Unmarshal(queryResponse.Value, &dataset)
                if err != nil {
                        return nil, err
                }
                datasets = append(datasets, &dataset)
        }

        return datasets, nil
}


type Operation struct {
        ID             string `json:"ID"`
        Operation_type string `json:"Operation_type"`
        Executed_by    string `json:"Executed_by"`
        Timestamp      string `json:"Timestamp"`
}


// CreateRawDataset issues a new dataset to the world state with given details.
func (s *SmartContract) CreateOperation(ctx contractapi.TransactionContextInterface, id string, operationType string, executed_by string, timestamp string) error {
        operation := Operation{
                ID:             id,
                Operation_type: operationType,
                Executed_by:    executed_by,
                Timestamp:      timestamp,
        }
        operationJSON, err := json.Marshal(operation)
        if err != nil {
                return err
        }

        return ctx.GetStub().PutState(id, operationJSON)
}

// ReadRawDataset returns the dataset stored in the world state with given id.
func (s *SmartContract) ReadOperation(ctx contractapi.TransactionContextInterface, id string) (*Operation, error) {
        operationJSON, err := ctx.GetStub().GetState(id)
        if err != nil {
                return nil, fmt.Errorf("failed to read from world state: %v", err)
        }
        if operationJSON == nil {
                return nil, fmt.Errorf("the operation %s does not exist", id)
        }

        var operation Operation
        err = json.Unmarshal(operationJSON, &operation)
        if err != nil {
                return nil, err
        }

        return &operation, nil
}