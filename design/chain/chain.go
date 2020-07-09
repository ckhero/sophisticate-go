package chain
import "fmt"
/**
学生请假责任链  2天的老师可以批，超过两天的校长批
 */
type Handle interface {
	HaveRight(days int) bool
	Exec(days int)
}

type VocateChain struct {
	Handle
	Next *VocateChain
}

func (v *VocateChain) SetNext(next *VocateChain)  {
	v.Next = next
}

func (v *VocateChain) Exec(days int)  {
	if v.HaveRight(days) {
		v.Handle.Exec(days)
	} else {
		v.Next.Exec(days)
	}
}

/**
老师
 */
type Teacher struct {}

func NewTeacher() *VocateChain {
	return &VocateChain{
		Handle: &Teacher{},
	}
}

func (t *Teacher) HaveRight(days int) bool  {
	return 2 >= days
}

func (t *Teacher) Exec(days int)  {
	if t.HaveRight(days) {
		fmt.Println("teach agree!")
	}
}

/**
校长
 */
type Headermaster struct {

}

func NewHeadermaster() *VocateChain {
	return &VocateChain{
		Handle: &Headermaster{},
	}
}
func (*Headermaster) HaveRight(days int) bool  {
	return 2 < days
}

func (*Headermaster) Exec(days int)  {
	fmt.Println("Headermaster agree!")
}