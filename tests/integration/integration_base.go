package integrationtests

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	api "shorty/api/app"
	"shorty/configs"
	repository "shorty/repositories"
	"time"

	"github.com/stretchr/testify/suite"
)

var settings = configs.GetSettings()

type BaseTest struct {
	suite.Suite
	Server          *api.App
	CacheRepository repository.CacheRepository
	BaseTestURL     string
}

func NewBaseTest() *BaseTest {
	baseTest := new(BaseTest)
	baseTest.CacheRepository = *repository.NewCacheRepository()
	baseTest.Server = baseTest.StartServer()
	baseTest.BaseTestURL = baseTest.GetBaseTestURL()

	return baseTest
}

func (b *BaseTest) TearDownTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, ok := b.CacheRepository.CleanCache(ctx)
	if !ok {
		fmt.Println("Error to clean cache in TearDown.")
	}

	err := b.Server.Stop(ctx)
	if err != nil {
		fmt.Println("Error closing server in TearDown.")
	}
}

func (b *BaseTest) StartServer() *api.App {
	server := api.MakeApp()
	go func() {
		err := server.Start()
		if err != nil {
			log.Print(err)
		}
	}()
	return server
}

func (b *BaseTest) GetBaseTestURL() string {
	return fmt.Sprintf("%s:%s%s", settings["BASE_URL"], settings["PORT"], settings["API_V1"])
}

func (b *BaseTest) Request(verb, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending GET:", err)
		return nil, err
	}

	return resp, nil
}
