package v1

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"sideEcho/stats"
)

func Test_wrapContextMiddleware(t *testing.T) {
	t.Parallel()

	// 1. mock setup
	ctrl := gomock.NewController(t)
	mockStats := stats.NewMockStats(ctrl)

	// 2. echo context setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// 3. echo context setup
	middleware := wrapContextMiddleware(mockStats)
	h := middleware(func(c echo.Context) error {
		ctx, ok := c.(*customContext)
		// 4. custom context setting 확인
		assert.True(t, ok)
		assert.Equal(t, mockStats, ctx.stats)
		return c.NoContent(http.StatusOK)
	})
	err := h(c)

	// 5. error 여부 확인
	assert.NoError(t, err, "wrapContextMiddleware는 성공해야 합니다.")
}

func Test_customWrapper(t *testing.T) {
	t.Parallel()

	type fields struct {
		mockStats *stats.MockStats
	}
	type expected struct {
		err error
	}
	testcases := []struct {
		name       string
		prepareCtx func(f *fields) echo.Context
		expected   expected
	}{
		{
			name: "Success Case - custom context를 넘기면 custom context를 받는 handler를 실행시킵니다.",
			prepareCtx: func(f *fields) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				return &customContext{
					Context: c,
					stats:   f.mockStats,
				}
			},
			expected: expected{
				err: nil,
			},
		},
		{
			name: "Failure Case - echo context를 넘기면 에러를 반환합니다.",
			prepareCtx: func(f *fields) echo.Context {
				e := echo.New()
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				return e.NewContext(req, rec)
			},
			expected: expected{
				err: echo.NewHTTPError(http.StatusInternalServerError, "failed to wrap customContext"),
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
				mockStats: stats.NewMockStats(ctrl),
			}

			// 2. echo context setup
			c := tc.prepareCtx(f)

			// 3. custom wrapper에서는 context 변환을 해줘야 함
			h := customWrapper(func(c *customContext) error {
				// 4. custom context setting 확인
				//    custom context가 넘어오지 않으면 실행되지 않음
				assert.Equal(t, f.mockStats, c.stats)
				return c.NoContent(http.StatusOK)
			})
			err := h(c)

			// 5. error 여부 확인
			if tc.expected.err != nil {
				assert.EqualError(t, err, tc.expected.err.Error())
			} else {
				assert.NoError(t, err, "customWrapper는 성공해야 합니다.")
			}
		})
	}
}

func Test_requestStatMiddleware(t *testing.T) {
	t.Parallel()

	type fields struct {
		mockStats *stats.MockStats
	}
	type expected struct {
		err error
	}
	testcases := []struct {
		name     string
		prepare  func(f *fields)
		handler  func(ctx *customContext) error
		expected expected
	}{
		{
			name: "핸들러에서 에러를 반환하지 않으면 success request count가 증가합니다.",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.mockStats.EXPECT().
						IncreaseRequestCount(),
					f.mockStats.EXPECT().
						IncreaseSuccessRequestCount(),
				)
			},
			handler: func(c *customContext) error {
				return c.NoContent(http.StatusOK)
			},
			expected: expected{
				err: nil,
			},
		},
		{
			name: "핸들러에서 에러를 반환하면 failure request count가 증가합니다.",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.mockStats.EXPECT().
						IncreaseRequestCount(),
					f.mockStats.EXPECT().
						IncreaseFailureRequestCount(),
				)
			},
			handler: func(c *customContext) error {
				return echo.NewHTTPError(http.StatusInternalServerError, "test error")
			},
			expected: expected{
				err: echo.NewHTTPError(http.StatusInternalServerError, "test error"),
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
				mockStats: stats.NewMockStats(ctrl),
			}
			tc.prepare(f)

			// 2. echo context setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			cc := &customContext{
				Context: c,
				stats:   f.mockStats,
			}

			// 3. request stat middleware 생성
			middleware := requestStatMiddleware()
			h := middleware(tc.handler)
			err := h(cc)

			// 5. error 여부 확인
			if tc.expected.err != nil {
				assert.EqualError(t, err, tc.expected.err.Error())
			} else {
				assert.NoError(t, err, "requestStatMiddleware는 성공해야 합니다.")
			}
		})
	}
}
