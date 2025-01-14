package test
import(
	"exam/backend/entity"
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameRequire(t *testing.T){
	g:=NewGomegaWithT(t)
	branch:=entity.Branch{
		Name: "",
	}

	ok,err := govalidator.ValidateStruct(branch)
	
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("name is require"))
}
