package LinkedList

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const listLen = 5

var indexList = [...]int{listLen - 1, 2, 0}

func TestAdder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adder Suite")
}

var _ = Describe("Linked List", func() {

	Describe("TestPushBack", func() {

		It("test the pushBack function directly", func() {
			var list = LinkedList{nil, nil, 0}
			pushBackToList(&list, listLen)
			cur := list.head
			for i := 0; cur != nil; cur, i = cur.next, i+1 {
				Expect(cur.val).To(Equal(i))
			}
		})
	})

	Describe("TestPushFront", func() {

		It("test the pushFront function directly", func() {
			var list = LinkedList{nil, nil, 0}
			pushFrontToList(&list, listLen)
			cur := list.head
			for i := listLen - 1; cur != nil; cur, i = cur.next, i-1 {
				Expect(cur.val).To(Equal(i))
			}
		})
	})

	Describe("TestPopBack", func() {

		It("test the PopBack function", func() {
			var list = LinkedList{nil, nil, 0}
			pushBackToList(&list, listLen)
			for i := listLen - 1; i <= 0; i-- {
				Expect(list.PopBack()).To(Equal(i))
				Expect(list.len).To(Equal(i + 1))
			}
		})
	})

	Describe("TestPopFront", func() {

		It("test the PopFront function", func() {
			var list = LinkedList{nil, nil, 0}
			pushBackToList(&list, listLen)
			for i := 0; i < listLen; i++ {
				Expect(list.PopFront()).To(Equal(i))
				Expect(list.len).To(Equal(listLen - i - 1))
			}
		})
	})

	Describe("TestPeek", func() {

		It("test the Peek function", func() {
			var list = LinkedList{nil, nil, 0}
			pushBackToList(&list, listLen)
			for i := 0; i < listLen; i++ {
				Expect(list.Peek(i)).To(Equal(i))
			}
		})
	})

	Describe("TestEraseList", func() {

		It("test the eraseList function", func() {
			var list = LinkedList{nil, nil, 0}
			pushBackToList(&list, listLen)
			err := list.EraseIndexList(indexList[0:3])
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
			_, err := list.PopBack()
			Expect(err).NotTo(BeNil())
			_, err = list.PopFront()
			Expect(err).NotTo(BeNil())
			_, err = list.Peek(-1)
			Expect(err).NotTo(BeNil())
			pushBackToList(&list, listLen)
			_, err = list.Peek(listLen)
			Expect(err).NotTo(BeNil())
			err = list.EraseIndexList([]int{-1})
			Expect(err).NotTo(BeNil())
			err = list.EraseIndexList([]int{0, 0})
			Expect(err).NotTo(BeNil())
			err = list.EraseIndexList([]int{listLen})
			Expect(err).NotTo(BeNil())
		})
	})
})

func pushBackToList(list *LinkedList, amountToPush int) {
	for i := 0; i < amountToPush; i++ {
		list.PushBack(i)
		Expect(list.len).To(Equal(i + 1))
		Expect(list.tail.val).To(Equal(i))
	}
}

func pushFrontToList(list *LinkedList, amountToPush int) {
	for i := 0; i < amountToPush; i++ {
		list.PushFront(i)
		Expect(list.len).To(Equal(i + 1))
		Expect(list.head.val).To(Equal(i))
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
