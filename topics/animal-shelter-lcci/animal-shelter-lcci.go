// 面试题 03.06. 动物收容所
//动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，比如enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。
//
//enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
//
//dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。
//
//示例1:
//
// 输入：
//["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog", "dequeueAny"]
//[[], [[0, 0]], [[1, 0]], [], [], []]
// 输出：
//[null,null,null,[0,0],[-1,-1],[1,0]]
//示例2:
//
// 输入：
//["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny"]
//[[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
// 输出：
//[null,null,null,null,[2,1],[0,0],[1,0]]
//说明:
//
//收纳所的最大容量为20000
//通过次数11,498提交次数19,529
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/3/21 3:07 PM
package animal_shelter_lcci

type AnimalShelf struct {
	animals [2][]int
}

func Constructor() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	this.animals[animal[1]] = append(this.animals[animal[1]], animal[0])
}

func (this *AnimalShelf) DequeueAny() []int {
	cats, dogs := this.animals[0], this.animals[1]
	if len(cats) == 0 {
		if len(dogs) == 0 {
			return []int{-1, -1}
		} else {
			this.decreaseAnimal(1)
			return []int{dogs[0], 1}
		}
	} else {
		if len(dogs) == 0 || cats[0] < dogs[0] {
			this.decreaseAnimal(0)
			return []int{cats[0], 0}
		} else {
			this.decreaseAnimal(1)
			return []int{dogs[0], 1}
		}
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	return this.dequeue(1)
}

func (this *AnimalShelf) DequeueCat() []int {
	return this.dequeue(0)
}

func (this *AnimalShelf) dequeue(species int) []int {
	if len(this.animals[species]) == 0 {
		return []int{-1, -1}
	}
	animal := this.animals[species][0]
	this.decreaseAnimal(species)
	return []int{animal, species}
}

func (this *AnimalShelf) decreaseAnimal(species int) {
	this.animals[species] = this.animals[species][1:]
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
