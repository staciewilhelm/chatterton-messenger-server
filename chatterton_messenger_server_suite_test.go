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

var _ = Describe("Model", func() {
	Context("Utils", func() {
		Context("GetMessagesSQLWithQueryParams", func() {
			var (
				queryParams *domain.QueryParams
				sqlString   string

				response string

				itContainsDefaultLimit = func() {
					It("contains the default limit value", func() {
						Expect(response).To(ContainSubstring("LIMIT 100"))
					})
				}
			)

			Context("with limit in query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{Limit: "50"}
					sqlString = "select * from table"
					response = models.GetMessagesSQLWithQueryParams(sqlString, queryParams)
				})

				It("contains the limit value from the params", func() {
					Expect(response).To(ContainSubstring("LIMIT 50"))
				})
			})

			Context("with recipient and sender in query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{RecipientID: "userOne", SenderID: "userTwo"}
					sqlString = "select * from table"
					response = models.GetMessagesSQLWithQueryParams(sqlString, queryParams)
				})

				It("contains the limit value from the params", func() {
					Expect(response).To(ContainSubstring("recipient_id = userOne"))
					Expect(response).To(ContainSubstring("sender_id = userTwo"))
				})
			})

			Context("with limit, recipient, and sender in query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{Limit: "50", RecipientID: "userOne", SenderID: "userTwo"}
					sqlString = "select * from table"
					response = models.GetMessagesSQLWithQueryParams(sqlString, queryParams)
				})

				It("contains the limit value from the params", func() {
					Expect(response).To(ContainSubstring("recipient_id = userOne"))
					Expect(response).To(ContainSubstring("sender_id = userTwo"))
					Expect(response).To(ContainSubstring("LIMIT 50"))
				})
			})

			Context("with empty query params", func() {
				BeforeEach(func() {
					queryParams = &domain.QueryParams{}
					sqlString = "select * from table"
					response = models.GetMessagesSQLWithQueryParams(sqlString, queryParams)
				})

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
					request := httptest.NewRequest(http.MethodGet, "/messages?limit=50", nil)
					response = domain.GetQueryParams(request)
				})

				It("returns query params with limit", func() {
					Expect(response).To(Equal(&domain.QueryParams{Limit: "50"}))
				})
			})

			Context("with recipient and sender set", func() {
				BeforeEach(func() {
					targetURL := "/messages?recipient_id=userOne&sender_id=userTwo"
					request := httptest.NewRequest(http.MethodGet, targetURL, nil)
					response = domain.GetQueryParams(request)
				})

				It("returns query params with limit", func() {
					Expect(response).To(Equal(&domain.QueryParams{
						RecipientID: "userOne",
						SenderID:    "userTwo",
					}))
				})
			})

			Context("with limit, recipient, sender set", func() {
				BeforeEach(func() {
					targetURL := "/messages?recipient_id=userOne&sender_id=userTwo&limit=15"
					request := httptest.NewRequest(http.MethodGet, targetURL, nil)
					response = domain.GetQueryParams(request)
				})

				It("returns query params with limit, recipient, and sender", func() {
					Expect(response).To(Equal(&domain.QueryParams{
						Limit:       "15",
						RecipientID: "userOne",
						SenderID:    "userTwo",
					}))
				})
			})

			Context("with limit set incorrectly", func() {
				BeforeEach(func() {
					request := httptest.NewRequest(http.MethodGet, "/messages?limit=", nil)
					response = domain.GetQueryParams(request)
				})

				itReturnsEmptyQueryParams()
			})

			Context("with invalid limit set", func() {
				BeforeEach(func() {
					request := httptest.NewRequest(http.MethodGet, "/messages?limit=5000", nil)
					response = domain.GetQueryParams(request)
				})

				itReturnsEmptyQueryParams()
			})

			Context("with only recipient set", func() {
				BeforeEach(func() {
					request := httptest.NewRequest(http.MethodGet, "/messages?recipient_id=userOne", nil)
					response = domain.GetQueryParams(request)
				})

				itReturnsEmptyQueryParams()
			})

			Context("with only sender set", func() {
				BeforeEach(func() {
					request := httptest.NewRequest(http.MethodGet, "/messages?sender_id=userTwo", nil)
					response = domain.GetQueryParams(request)
				})

				itReturnsEmptyQueryParams()
			})
		})
	})

	XContext("Variables", func() {})
})
