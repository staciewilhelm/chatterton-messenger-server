package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"chatterton-messenger-server/domain"
	"chatterton-messenger-server/models"
)

func TestChattertonMessengerServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChattertonMessengerServer Suite")
}

var _ = Describe("Models", func() {
	Context("Utils", func() {
		Context("GetSQLWithQueryParams", func() {
			var (
				queryParams *domain.QueryParams
				sqlString   string

				response string

				withSQLStringWithoutLimit = func() {
					BeforeEach(func() {
						sqlString = "select * from table"
					})
				}

				itContainsDefaultLimit = func() {
					It("contains the default limit value", func() {
						Expect(response).To(ContainSubstring("LIMIT 100"))
					})
				}
			)

			Context("with limit in query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{Limit: "50"}
					response = models.GetSQLWithQueryParams(sqlString, queryParams)
				})
				withSQLStringWithoutLimit()

				It("contains the limit value from the params", func() {
					Expect(response).To(ContainSubstring("LIMIT 50"))
				})
			})

			Context("with invalid limit in query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{Limit: "5000"}
					response = models.GetSQLWithQueryParams(sqlString, queryParams)
				})
				withSQLStringWithoutLimit()

				itContainsDefaultLimit()
			})

			Context("with empty query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{}
					response = models.GetSQLWithQueryParams(sqlString, queryParams)
				})
				withSQLStringWithoutLimit()

				itContainsDefaultLimit()
			})
		})
	})
})

var _ = Describe("Domain", func() {
	Context("QueryParams", func() {
		var (
			response *domain.QueryParams

			itReturnsEmptyQueryParams = func() {
				It("returns empty query params", func() {
					Expect(response).To(Equal(&domain.QueryParams{}))
				})
			}
		)

		Context("with no query params", func() {
			BeforeEach(func() {
				request := httptest.NewRequest(http.MethodGet, "/messages", nil)
				response = domain.GetQueryParams(request)
			})
			itReturnsEmptyQueryParams()
		})

		Context("with query params", func() {
			Context("with limit set correctly", func() {
				BeforeEach(func() {
					request := httptest.NewRequest(http.MethodGet, "/messages?limit=500", nil)
					response = domain.GetQueryParams(request)
				})

				It("returns query params with limit", func() {
					Expect(response).To(Equal(&domain.QueryParams{Limit: "500"}))
				})
			})

			Context("with limit set incorrectly", func() {
				BeforeEach(func() {
					request := httptest.NewRequest(http.MethodGet, "/messages?limit=", nil)
					response = domain.GetQueryParams(request)
				})

				itReturnsEmptyQueryParams()
			})
		})
	})
})
