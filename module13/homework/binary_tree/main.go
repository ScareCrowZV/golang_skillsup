package main

import (
	"fmt"
	"strings"
)

type BTreeN struct {
	value *int
	left  *BTreeN
	right *BTreeN
}

func (t *BTreeN) add(newValue int) {

	if t.look(newValue) != nil {
		fmt.Println("Элемент уже существует в дереве")
		return
	}

	if t.left == nil {
		t.left = &BTreeN{value: &newValue}
	} else if t.right == nil {
		t.right = &BTreeN{value: &newValue}

	} else if t.left.tallLevel() > t.right.tallLevel() {
		t.right.add(newValue)
	} else {
		t.left.add(newValue)
	}
}

func (t *BTreeN) tallLevel() int {
	if t == nil {
		return -1
	}

	leftH := t.left.tallLevel()
	rightH := t.right.tallLevel()

	if leftH > rightH {
		return leftH + 1
	} else {
		return rightH + 1
	}
}

func (t *BTreeN) changeParent(elementTree *BTreeN) {

	if t.left == nil {
		t.left = elementTree
		return
	}

	if t.right == nil {
		t.right = elementTree
		return
	}

	if t.left != nil {
		t.left.changeParent(elementTree)
	}
}

func (t *BTreeN) remove(root *BTreeN) {

	if t.value == nil && t.left == nil && t.right == nil {
		return
	}

	parent := root.lookParentByElement(t)

	if t.left.right != nil {
		t.left.changeParent(t.left.right)

	}

	if t.right != nil {
		t.left.right = t.right
	}

	if parent.left == t {
		fmt.Println("Удаляем элемент слева")

		parent.left = t.left

	}

	if parent.right == t {
		fmt.Println("Удаляем элемент справа")
		parent.right = t.left
	}

	t.left = nil
	t.right = nil
	t.value = nil

}

func (t *BTreeN) lookParentByElement(searchElement *BTreeN) (parent *BTreeN) {

	if t.left == nil && t.right == nil {
		return nil
	}

	if t.left == searchElement || t.right == searchElement {
		parent = t
	}

	if parent == nil && t.left != nil {
		parent = t.left.lookParentByElement(searchElement)
		if parent == nil && t.right != nil {
			parent = t.right.lookParentByElement(searchElement)
		}
	}

	return parent
}

func (t *BTreeN) look(searchValue int) *BTreeN {
	if t == nil {
		return nil
	}

	if *t.value == searchValue {
		return t
	}

	if res := t.left.look(searchValue); res != nil {
		return res
	}

	return t.right.look(searchValue)
}

func (t *BTreeN) print(lvl int) {

	if t.value != nil {
		fmt.Printf("%s '%d'(lvl=%d)->\n", strings.Repeat(" ", lvl), *t.value, lvl)
	}

	if t.left != nil {
		t.left.print(lvl + 1)
	}
	if t.right != nil {
		t.right.print(lvl + 1)
	}
}

func main() {

	// Создаём корень дерева со значением 0
	var vl int = 0
	var root BTreeN = BTreeN{value: &vl}

	// Наполняем дерево элементами
	for i := 1; i < 15; i++ {
		root.add(i)
	}

	// fmt.Println("Ищем значение 5")
	// findingTree := root.look(5)
	// fmt.Println(findingTree)

	fmt.Println("Вызвали функцию печати")
	root.print(0)

	fmt.Println("Будем удалять. Ищём значение")
	deleteElement := root.look(4)
	fmt.Println("Удаляемое значение", *deleteElement.value)
	fmt.Println("Вызвали функцию удаления")
	deleteElement.remove(&root)

	fmt.Println("Вызвали функцию печати снова")
	root.print(0)

	// fmt.Println(*root.lookParentByElement(0, &bTree).value)

	// Проверка что потомки точно удалились
	// fmt.Println(vlLeft)
	// fmt.Println(vlRight)

}
