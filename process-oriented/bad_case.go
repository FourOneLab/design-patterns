package process_oriented

// 滥用 getter setter 方法
// 通常在写 JAVA 的时候，IDE的插件，顺手就给每个属性都生成了， 这样违反了面向对象的封装原则。
// 这个问题在 Golang 中体现为，结构体内的字段都直接大写了，尴尬。

type ShoppingCartItem struct {
	Name  string
	Price float64
}

// ShoppingCart 虽然都定义为包内属性，但是都有包外可访问都方法来进行修改，
// 这完全违背了封装特性：通过访问权限控制，隐藏内部数据。
type ShoppingCart struct {
	itemsCount int
	totalPrice float64
	items      []ShoppingCartItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		itemsCount: 0,
		totalPrice: 0,
		items:      make([]ShoppingCartItem, 0),
	}
}
func (c *ShoppingCart) GetItemsCount() int {
	return c.itemsCount
}

func (c *ShoppingCart) SetItemsCount(itemsCount int) {
	c.itemsCount = itemsCount
}

func (c *ShoppingCart) GetTotalPrice() float64 {
	return c.totalPrice
}

func (c *ShoppingCart) SetTotalPrice(totalPrice float64) {
	c.totalPrice = totalPrice
}

// GetItem 返回的 items 切片内的数据可以通过索引的方式直接修改，
// 这样会导致 item 中保存的数据和 count、totalPrice 不一致。
// 正确的做法是将原有数据克隆一份，对结果的任何修改不会影响现有对象
func (c *ShoppingCart) GetItem() []ShoppingCartItem {
	// bad case
	//return c.items

	// good case
	res := make([]ShoppingCartItem, len(c.items))
	// builtin copy(dst.src) copies min(len(dst),len(src)) elements
	// https://golang.org/ref/spec#Appending_and_copying_slices
	copy(res, c.items)
	return res
}

func (c *ShoppingCart) AddItem(cartItem ShoppingCartItem) {
	c.items = append(c.items, cartItem)
	c.itemsCount++
	c.totalPrice += cartItem.Price
}
