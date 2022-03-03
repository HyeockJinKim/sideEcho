package v1

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"sideEcho/exchange"
)

func TestHandler_buy(t *testing.T) {
	t.Parallel()

	type fields struct {
		mockExchangeCtrl *exchange.MockController
	}
	type expected struct {
		buyValue uint64
		err      error
	}
	testcases := []struct {
		name      string
		prepare   func(f *fields, expected expected)
		inputJSON string
		expected  expected
	}{
		{
			name: "Success Case - inputJSON이 정해진 대로 들어온다면 성공합니다.",
			prepare: func(f *fields, expected expected) {
				f.mockExchangeCtrl.EXPECT().
					Buy(gomock.Eq(expected.buyValue)).
					Return(nil)
			},
			inputJSON: `{"value":2}`,
			expected: expected{
				buyValue: 2,
				err:      nil,
			},
		},
		{
			name: "Success Case - 빈 input이 들어오더라도 요청은 성공합니다.",
			prepare: func(f *fields, expected expected) {
				f.mockExchangeCtrl.EXPECT().
					Buy(gomock.Eq(expected.buyValue)).
					Return(nil)
			},
			inputJSON: ``,
			expected: expected{
				buyValue: 0,
				err:      nil,
			},
		},
		{
			name:      "Failure Case - 잘못된 inputJSON이 들어온다면 실패합니다.",
			inputJSON: `{`,
			expected: expected{
				buyValue: 0,
				err:      echo.NewHTTPError(http.StatusBadRequest, "invalid request"),
			},
		},
		{
			name:      "Failure Case - integer inputJSON이 들어온다면 실패합니다.",
			inputJSON: `23`,
			expected: expected{
				buyValue: 0,
				err:      echo.NewHTTPError(http.StatusBadRequest, "invalid request"),
			},
		},
		{
			name:      "Failure Case - string inputJSON이 들어온다면 실패합니다.",
			inputJSON: `"abc"`,
			expected: expected{
				buyValue: 0,
				err:      echo.NewHTTPError(http.StatusBadRequest, "invalid request"),
			},
		},
		{
			name: "Failure Case - controller의 buy가 실패하면 에러를 반환합니다.",
			prepare: func(f *fields, expected expected) {
				f.mockExchangeCtrl.EXPECT().
					Buy(gomock.Eq(expected.buyValue)).
					Return(errors.New("test error"))
			},
			inputJSON: ``,
			expected: expected{
				buyValue: 0,
				err:      echo.NewHTTPError(http.StatusInternalServerError, "test error"),
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// 1. mock setup
			ctrl := gomock.NewController(t)
			f := &fields{
				mockExchangeCtrl: exchange.NewMockController(ctrl),
			}
			if tc.prepare != nil {
				tc.prepare(f, tc.expected)
			}
			handler := NewHandler(f.mockExchangeCtrl)

			// 2. context 생성
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/buy", strings.NewReader(tc.inputJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			cc := &customContext{
				Context: c,
			}

			// 3. buy 테스트
			err := handler.buy(cc)
			if tc.expected.err != nil {
				assert.EqualError(t, err, tc.expected.err.Error())
			} else {
				assert.NoError(t, err, "buy는 성공해야 합니다.")
			}
		})
	}
}

func TestHandler_sell(t *testing.T) {
	t.Parallel()

	type fields struct {
		mockExchangeCtrl *exchange.MockController
	}
	type expected struct {
		sellValue uint64
		err       error
	}
	testcases := []struct {
		name      string
		prepare   func(f *fields, expected expected)
		inputJSON string
		expected  expected
	}{
		{
			name: "Success Case - inputJSON이 정해진 대로 들어온다면 성공합니다.",
			prepare: func(f *fields, expected expected) {
				f.mockExchangeCtrl.EXPECT().
					Sell(gomock.Eq(expected.sellValue)).
					Return(nil)
			},
			inputJSON: `{"value":2}`,
			expected: expected{
				sellValue: 2,
				err:       nil,
			},
		},
		{
			name: "Success Case - 빈 input이 들어오더라도 요청은 성공합니다.",
			prepare: func(f *fields, expected expected) {
				f.mockExchangeCtrl.EXPECT().
					Sell(gomock.Eq(expected.sellValue)).
					Return(nil)
			},
			inputJSON: ``,
			expected: expected{
				sellValue: 0,
				err:       nil,
			},
		},
		{
			name:      "Failure Case - 잘못된 inputJSON이 들어온다면 실패합니다.",
			inputJSON: `{`,
			expected: expected{
				sellValue: 0,
				err:       echo.NewHTTPError(http.StatusBadRequest, "invalid request"),
			},
		},
		{
			name:      "Failure Case - integer inputJSON이 들어온다면 실패합니다.",
			inputJSON: `23`,
			expected: expected{
				sellValue: 0,
				err:       echo.NewHTTPError(http.StatusBadRequest, "invalid request"),
			},
		},
		{
			name:      "Failure Case - string inputJSON이 들어온다면 실패합니다.",
			inputJSON: `"abc"`,
			expected: expected{
				sellValue: 0,
				err:       echo.NewHTTPError(http.StatusBadRequest, "invalid request"),
			},
		},
		{
			name: "Failure Case - controller의 sell이 실패하면 에러를 반환합니다.",
			prepare: func(f *fields, expected expected) {
				f.mockExchangeCtrl.EXPECT().
					Sell(gomock.Eq(expected.sellValue)).
					Return(errors.New("test error"))
			},
			inputJSON: ``,
			expected: expected{
				sellValue: 0,
				err:       echo.NewHTTPError(http.StatusInternalServerError, "test error"),
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// 1. mock setup
			ctrl := gomock.NewController(t)
			f := &fields{
				mockExchangeCtrl: exchange.NewMockController(ctrl),
			}
			if tc.prepare != nil {
				tc.prepare(f, tc.expected)
			}
			handler := NewHandler(f.mockExchangeCtrl)

			// 2. context 생성
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/sell", strings.NewReader(tc.inputJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			cc := &customContext{
				Context: c,
			}

			// 3. sell 테스트
			err := handler.sell(cc)
			if tc.expected.err != nil {
				assert.EqualError(t, err, tc.expected.err.Error())
			} else {
				assert.NoError(t, err, "sell은 성공해야 합니다.")
			}
		})
	}
}
