package gee
type node struct {
	pattern string //待匹配路由，例如 /p/:lang
	part string     //路由中的一部分，例如:lang
	children []*node //子节点
	isWild bool //是否精确匹配 part含有：或*时为true
}
//返回第一个匹配成功的节点 用于插入
func (n *node)matchChild(part string)*node {
	for _,child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}
// 返回所有匹配成功的节点，用于查找
func (n *node)matchChildren(part string)[]*node{
	nodes := make([]*node,0)
	for _,child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes,child)
		}
	}
	return nodes
}
// 插入
/*
	插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，
则新建一个，有一点需要注意，/p/:lang/doc只有在第三层节点，
即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。
因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。
例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败
 */
