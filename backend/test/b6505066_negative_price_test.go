package unit

import (
	"sut-final-lab/sut-final-lab/backend/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPrice(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Price is invalid`, func(t *testing.T) {
		product := entity.Products{
			Name:  "Siron",
			Price: 100002, //ผิดตรงนี้
			SKU:   "Test",
		}

		ok, err := govalidator.ValidateStruct(product)

		print("error", err)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(`Price must be between 1 and 1000`))

	})
}
