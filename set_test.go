package app_test

import (
	"app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	Describe("Vazio", func() {
		Context("Quando o set não contém itens", func() {
			It("Deve estar vazio", func() {
				s := app.NewSet()
				Expect(s.IsEmpty()).To(BeTrue())
			})
		})
	})
})
