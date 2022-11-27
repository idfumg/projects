package server

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"myapp/models"
	"myapp/server/mocks"
	"myapp/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type TestLogger struct {
}

func (l *TestLogger) Printf(s string, params ...any) {

}

func TestGetAccount(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		buildMock     func(store *mocks.MockStore)
		chechResponse func(t testing.TB, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: account.ID,
			buildMock: func(store *mocks.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			chechResponse: func(t testing.TB, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "NOT FOUND",
			accountID: account.ID,
			buildMock: func(store *mocks.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(models.Account{}, sql.ErrNoRows)
			},
			chechResponse: func(t testing.TB, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.ID,
			buildMock: func(store *mocks.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(models.Account{}, sql.ErrConnDone)
			},
			chechResponse: func(t testing.TB, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			accountID: 0,
			buildMock: func(store *mocks.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			chechResponse: func(t testing.TB, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mocks.NewMockStore(ctrl)
			testCase.buildMock(store)

			logger := &TestLogger{}
			s, _ := NewServer(store, logger)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", testCase.accountID)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			s.ServeHTTP(recorder, request)

			testCase.chechResponse(t, recorder)
		})
	}
}

func randomAccount() models.Account {
	return models.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t testing.TB, body *bytes.Buffer, want models.Account) {
	t.Helper()

	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var got models.Account
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, got, want)
}
