package rocket

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("get a rocket by an id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockController)
		id := "uuid-1"

		rocketStoreMock.
			EXPECT().
			GetRocket(id).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService, err := NewService(rocketStoreMock)

		assert.NoError(t, err)

		rocket, err := rocketService.GetRocket(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		assert.Equal(t, rocket, Rocket{
			ID: id,
		})
	})

	t.Run("insert a rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockController)
		id := "uudi-1"

		rocketStoreMock.
			EXPECT().
			AddRocket(Rocket{
				ID: id,
			}).
			Return(Rocket{
				ID: id,
			}, nil)

		rocketService, err := NewService(rocketStoreMock)

		assert.NoError(t, err)

		rocket, err := rocketService.AddRocket(
			context.Background(),
			Rocket{
				ID: id,
			},
		)

		assert.NoError(t, err)
		assert.Equal(t, rocket, Rocket{
			ID: id,
		})
	})

	t.Run("delete mocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockController)
		id := "uuid-1"

		rocketStoreMock.
			EXPECT().
			DelRocket(id).
			Return(nil)

		rocketService, err := NewService(rocketStoreMock)

		assert.NoError(t, err)

		err = rocketService.DelRocket(context.Background(), id)

		assert.NoError(t, err)
	})
}
