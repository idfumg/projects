package test

import (
	"context"
	v1 "myapp/internal/proto/rocket/v1"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()

		_, _ = client.DelRocket(
			context.Background(),
			&v1.DelRocketRequest{
				Id: "1",
			},
		)

		resp, err := client.AddRocket(
			context.Background(),
			&v1.AddRocketRequest{
				Rocket: &v1.Rocket{
					Id:   "1",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotNil(t, resp.Rocket)
		assert.Equal(t, "1", resp.Rocket.Id)
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
