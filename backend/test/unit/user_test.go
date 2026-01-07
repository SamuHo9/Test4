package unit

import (
	"testing"
	
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
	
	"github.com/SamuHo9/Test4/backend/entity"
)
func TestUser(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run("sssfsfsf",func (t *testing.T){ 

		user := entity.User {
			Name:  "sdsdsdsdsd",
			Email: "dddgg@gmail.com",
			Age:   30,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		
	})
}
