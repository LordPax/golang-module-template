package media

import (
	"context"
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"
	"io"
	"os"
	"strings"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/config"
	"github.com/gophercloud/gophercloud/v2/openstack/config/clouds"
	"github.com/gophercloud/gophercloud/v2/openstack/objectstorage/v1/containers"
	"github.com/gophercloud/gophercloud/v2/openstack/objectstorage/v1/objects"
)

type IOpenstackService interface {
	core.IProvider
	Authenticate() error
	CreateContainerIfNotExist(containerName string) error
	UploadFile(file io.Reader, objectName string, containerName string) (string, error)
}

type OpenstackService struct {
	*core.Provider
	dotenvService dotenv.IDotenvService
	client        *gophercloud.ServiceClient
	ctx           context.Context
}

func NewOpenstackService(module core.IModule) *OpenstackService {
	return &OpenstackService{
		Provider:      core.NewProvider("OpenstackService"),
		dotenvService: module.Get("DotenvService").(dotenv.IDotenvService),
		ctx:           context.Background(),
	}
}

func (o *OpenstackService) OnInit() error {
	return o.Authenticate()
}

func (o *OpenstackService) Authenticate() error {
	fmt.Println("Authenticating with OpenStack")
	os.Setenv("OS_CLOUD", o.dotenvService.Get("OS_CLOUD"))
	authOpts, endpointOpts, tlsConfig, err := clouds.Parse()
	if err != nil {
		return fmt.Errorf("error parsing cloud configuration: %w", err)
	}

	provider, err := config.NewProviderClient(o.ctx, authOpts, config.WithTLSConfig(tlsConfig))
	if err != nil {
		return fmt.Errorf("error creating OpenStack provider client: %w", err)
	}

	client, err := openstack.NewObjectStorageV1(provider, endpointOpts)
	if err != nil {
		return fmt.Errorf("error creating Swift service client: %w", err)
	}

	o.client = client
	return nil
}

func (o *OpenstackService) CreateContainerIfNotExist(containerName string) error {
	result := containers.Get(o.ctx, o.client, containerName, nil)
	if result.Err == nil {
		return nil // Container exists
	}

	if !strings.Contains(result.Err.Error(), "404") {
		return fmt.Errorf("error checking container %s: %w", containerName, result.Err)
	}

	_, err := containers.Create(o.ctx, o.client, containerName, containers.CreateOpts{}).Extract()
	if err != nil {
		return fmt.Errorf("error creating container %s: %w", containerName, err)
	}

	return nil
}

func (o *OpenstackService) UploadFile(
	file io.Reader,
	objectName string,
	containerName string,
) (string, error) {
	createOpts := objects.CreateOpts{
		Content:     file,
		ContentType: "application/octet-stream",
	}

	// TODO : Fix file upload issue
	result := objects.Create(o.ctx, o.client, containerName, objectName, createOpts)
	if err := result.Err; err != nil {
		return "", fmt.Errorf("error uploading file: %w", err)
	}

	url := fmt.Sprintf("%s/%s/%s", o.client.Endpoint, containerName, objectName)
	return url, nil
}
