package notifier_test

import (
	"github.com/netice9/notifier-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoNotifier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Notifier Suite")
}

var _ = Describe("Notifier", func() {

	var n *notifier.Notifier

	BeforeEach(func() {
		n = notifier.NewNotifier("first")
	})

	It("should notify listeners of changes", func(done Done) {
		c := n.AddListener(0)
		notification := <-c
		Expect(notification).To(Equal("first"))
		close(done)
	})

	Context("when listener is added", func() {

		var l1 chan interface{}

		BeforeEach(func() {
			l1 = n.AddListener(0)
		})

		Describe("RemoveListener()", func() {
			It("should remove existing listener", func() {
				n.RemoveListener(l1)
				Expect(n.NumberOfListeners()).To(Equal(0))
			})
		})
	})

	Context("when there is already one value being sent to listeners", func() {
		BeforeEach(func() {
			n.Notify("test1")
		})

		Context("When a new listener is added", func() {
			var l1 chan interface{}

			BeforeEach(func() {
				l1 = n.AddListener(0)
			})

			It("should send the last notification to the new listener", func(done Done) {
				notification := <-l1
				Expect(notification).To(Equal("test1"))
				close(done)
			})
		})
	})

})
