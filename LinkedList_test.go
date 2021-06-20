package linkedList

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const listLen = 5
var indexList = [...]int {listLen - 1, 2, 0}

func TestAdder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adder Suite")
}

var _ = Describe("Linked List", func() {

	Describe("TestPush", func() {

		It("test the push function directly", func() {
			var list = LinkedList{nil, nil, 0}
			pushToList(&list, listLen)
			cur := list.head
			for i := 0; cur != nil; cur, i = cur.next, i+1 {
				Expect(cur.val).To(Equal(i))
			}
		})
	})

	Describe("TestPop", func() {

		It("test the Pop function", func() {
			var list = LinkedList{nil, nil, 0}
			pushToList(&list, listLen)
			for i := listLen - 1; i <= 0; i-- {
				Expect(list.pop()).To(Equal(i))
				Expect(list.len).To(Equal(i + 1))
			}
		})
	})

	Describe("TestPeek", func() {

		It("test the Peek function", func() {
			var list = LinkedList{nil, nil, 0}
			pushToList(&list, listLen)
			for i := 0; i < listLen; i++ {
				Expect(list.peek(i)).To(Equal(i))
			}
		})
	})

	Describe("TestEraseList", func() {

		It("test the eraseList function", func() {
			var list = LinkedList{nil, nil, 0}
			pushToList(&list, listLen)
			err := list.eraseIndexList(indexList[0:3])
			Expect(err).To(BeNil())
			Expect(list.len).To(Equal(2))
			cur := list.head
			for i := 0; i < listLen; i++ {
				if !contains(indexList[0:3], i) {
					Expect(cur.val).To(Equal(i))
					cur = cur.next
				}
			}
		})
	})

	Describe("TestError", func() {
		It("asserts that the linked list returns error for illegal input", func() {
			var list = LinkedList{nil, nil, 0}
			_, err := list.pop()
			Expect(err).NotTo(BeNil())
			_, err = list.peek(-1)
			Expect(err).NotTo(BeNil())
			pushToList(&list, listLen)
			_, err = list.peek(listLen)
			Expect(err).NotTo(BeNil())
			err = list.eraseIndexList([]int{-1})
			Expect(err).NotTo(BeNil())
			err = list.eraseIndexList([]int{0, 0})
			Expect(err).NotTo(BeNil())
			err = list.eraseIndexList([]int{listLen})
			Expect(err).NotTo(BeNil())
		})
	})
})

func pushToList(list *LinkedList, amountToPush int) {
	for i := 0; i < amountToPush; i++ {
		list.push(i)
		Expect(list.len).To(Equal(i + 1))
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
