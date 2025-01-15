package unit

import (
	"sut-final-lab/sut-final-lab/backend/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestProducts(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Name correct`, func(t *testing.T) {
		product := entity.Products{
			Name:  "", //ผิดตรงนี้
			Price: 120.0,
			SKU:   "Test",
		}

		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(`Name is required`))

	})
}
