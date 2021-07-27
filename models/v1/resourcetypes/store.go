package resourcetypes

import "fabric-go-rest-api/models"

func Store(name string) (resourcetype *models.ResourceType, err error) {
	resourcetype, err = models.NewResourceType(name)
	if err != nil {
		return
	}

	mockResourceTypes = append(mockResourceTypes, *resourcetype)

	return
}
