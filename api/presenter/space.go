package presenter

import (
	"net/url"
	"time"

	"code.cloudfoundry.org/korifi/api/repositories"
)

const (
	// TODO: repetition with handler endpoint?
	spacesBase = "/v3/spaces"
)

type SpaceResponse struct {
	Name          string        `json:"name"`
	GUID          string        `json:"guid"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	Links         SpaceLinks    `json:"links"`
	Metadata      Metadata      `json:"metadata"`
	Relationships Relationships `json:"relationships"`
}

type SpaceLinks struct {
	Self         *Link `json:"self"`
	Organization *Link `json:"organization"`
}

func ForSpaceList(spaces []repositories.SpaceRecord, apiBaseURL, requestURL url.URL) ListResponse {
	spaceResponses := make([]interface{}, 0, len(spaces))
	for _, space := range spaces {
		spaceResponses = append(spaceResponses, ForSpace(space, apiBaseURL))
	}

	return ForList(spaceResponses, apiBaseURL, requestURL)
}

func ForSpace(space repositories.SpaceRecord, apiBaseURL url.URL) SpaceResponse {
	return SpaceResponse{
		Name:      space.Name,
		GUID:      space.GUID,
		CreatedAt: space.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: space.CreatedAt.UTC().Format(time.RFC3339),
		Metadata: Metadata{
			Labels:      emptyMapIfNil(space.Labels),
			Annotations: emptyMapIfNil(space.Annotations),
		},
		Relationships: Relationships{
			"organization": Relationship{
				Data: &RelationshipData{
					GUID: space.OrganizationGUID,
				},
			},
		},
		Links: SpaceLinks{
			Self: &Link{
				HRef: buildURL(apiBaseURL).appendPath(spacesBase, space.GUID).build(),
			},
			Organization: &Link{
				HRef: buildURL(apiBaseURL).appendPath(orgsBase, space.OrganizationGUID).build(),
			},
		},
	}
}
