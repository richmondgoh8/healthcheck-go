package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/rlc4u/health-check/internal/model"
	"github.com/rlc4u/health-check/pkg/client"
	"github.com/rlc4u/health-check/pkg/logger"
)

type HealthSvc interface {
	HealthCheck() (model.Response, error)
}

func NewHealthSvc(customClient client.CustomHTTP) HealthSvc {
	return HealthSvcImpl{
		customClient: customClient,
	}
}

type HealthSvcImpl struct {
	customClient client.CustomHTTP
}

func (h HealthSvcImpl) serviceCaller(nameList []string) chan model.PredictionChannel {
	results := make(chan model.PredictionChannel, len(nameList))
	defer close(results)

	var wg sync.WaitGroup
	wg.Add(len(nameList))

	for _, name := range nameList {
		go func(name string) {
			defer wg.Done()
			logger.Info(name)
			url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
			resp, err := h.customClient.Get(url)
			if err != nil {
				fmt.Println("error", err)
			} else {
				var responseObject model.RandomResponse
				err = json.Unmarshal(resp, &responseObject)
				if err == nil {
					results <- model.PredictionChannel{
						Name: responseObject.Name,
						Age:  responseObject.Age,
					}
				}
			}
		}(name)
	}

	wg.Wait()
	return results
}

func (h HealthSvcImpl) HealthCheck() (model.Response, error) {

	nameList := []string{"michael", "john"}
	dataSet := h.serviceCaller(nameList)
	for data := range dataSet {
		logger.Info(fmt.Sprintf("%s is destined to live to %d", data.Name, data.Age))
	}
	respData := model.GenerateResponse(http.StatusOK, model.SvcSuccessMsg)
	return respData, nil
}
