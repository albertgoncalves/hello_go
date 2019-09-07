package main

import (
    "errors"
    "fmt"
    "log"
    "strings"
)

type Data interface {}

type Value interface {
    toString() string
    equalTo(Value) bool
    lessThan(Value) bool
    greaterThan(Value) bool
}

type StringValue struct {
    str string
}

func (s StringValue) toString() string {
    return s.str
}

func (s StringValue) equalTo(t Value) bool {
    return s.str == t.toString()
}

func (s StringValue) lessThan(t Value) bool {
    return s.str < t.toString()
}

func (s StringValue) greaterThan(t Value) bool {
    return s.str > t.toString()
}

type Node struct {
    Value Value
    Data  Data
    Left  *Node
    Right *Node
}

func (n *Node) Insert(value Value, data Data) error {
    if n == nil {
        return errors.New("Cannot insert a value into a nil tree")
    }
    switch {
    case value.equalTo(n.Value):
        return nil
    case value.lessThan(n.Value):
        if n.Left == nil {
            n.Left = &Node{Value: value, Data: data}
            return nil
        }
        return n.Left.Insert(value, data)
    case value.greaterThan(n.Value):
        if n.Right == nil {
            n.Right = &Node{Value: value, Data: data}
            return nil
        }
        return n.Right.Insert(value, data)
    }
    return nil
}

func (n *Node) Find(s Value) (Data, bool) {
    if n == nil {
        return nil, false
    }
    switch {
    case s.equalTo(n.Value):
        return n.Data, true
    case s.lessThan(n.Value):
        return n.Left.Find(s)
    default:
        return n.Right.Find(s)
    }
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
    if n == nil {
        return nil, parent
    }
    if n.Right == nil {
        return n, parent
    }
    return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
    if n == nil {
        return errors.New("replaceNode() not allowed on a nil node")
    }
    if n == parent.Left {
        parent.Left = replacement
        return nil
    }
    parent.Right = replacement
    return nil
}

func (n *Node) Delete(s Value, parent *Node) error {
    if n == nil {
        return errors.New("Value to be deleted does not exist in the tree")
    }
    switch {
    case s.lessThan(n.Value):
        return n.Left.Delete(s, n)
    case s.greaterThan(n.Value):
        return n.Right.Delete(s, n)
    default:
        if n.Left == nil && n.Right == nil {
            n.replaceNode(parent, nil)
            return nil
        }
        if n.Left == nil {
            n.replaceNode(parent, n.Right)
            return nil
        }
        if n.Right == nil {
            n.replaceNode(parent, n.Left)
            return nil
        }
        replacement, replParent := n.Left.findMax(n)
        n.Value = replacement.Value
        n.Data = replacement.Data
        return replacement.Delete(replacement.Value, replParent)
    }
}

type Tree struct {
    Root *Node
}

func (t *Tree) Insert(value Value, data Data) error {
    if t.Root == nil {
        t.Root = &Node{Value: value, Data: data}
        return nil
    }
    return t.Root.Insert(value, data)
}

func (t *Tree) Find(s Value) (Data, bool) {
    if t.Root == nil {
        return nil, false
    }
    return t.Root.Find(s)
}

func (t *Tree) Delete(s Value) error {
    if t.Root == nil {
        return errors.New("Cannot delete from an empty tree")
    }
    fakeParent := &Node{Right: t.Root}
    err := t.Root.Delete(s, fakeParent)
    if err != nil {
        return err
    }
    if fakeParent.Right == nil {
        t.Root = nil
    }
    return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
    if n == nil {
        return
    }
    t.Traverse(n.Left, f)
    f(n)
    t.Traverse(n.Right, f)
}

func printTree(tree *Tree) {
    tree.Traverse(
        tree.Root,
        func(n *Node) { fmt.Print(n.Value, ": ", n.Data, ", ") },
    )
    fmt.Println("\n")
}

func bold(s string) string {
    return strings.Join([]string{"\033[1m", s, "\033[0m"}, "")
}

func main() {
    values := []StringValue{
        {"d"},
        {"b"},
        {"c"},
        {"e"},
        {"a"},
    }
    data := []string{
        "delta",
        "bravo",
        "charlie",
        "echo",
        "alpha",
    }
    tree := &Tree{}
    for i := 0; i < len(values); i++ {
        err := tree.Insert(values[i], data[i])
        if err != nil {
            log.Fatalf("Error inserting value '%v'\nError: %s", values[i], err)
        }
    }
    {
        fmt.Println(bold("MULTI-NODE TREE\n"))
        fmt.Println(bold("Sorted values"))
        printTree(tree)
    }
    x := StringValue{"e"}
    {
        fmt.Println(bold(fmt.Sprintf("Find node %v", x)))
        d, found := tree.Find(x)
        if !found {
            log.Fatalf("Cannot find %v", x)
        }
        fmt.Printf("Found: %v %v\n\n", x, d)
    }
    {
        err := tree.Delete(x)
        if err != nil {
            log.Fatalf("Error deleting %v\nError: %s", x, err)
        }
        fmt.Println(bold(fmt.Sprintf("After deleting %v", x)))
        printTree(tree)
    }
    y := StringValue{"a"}
    {
        fmt.Println(bold("\nSINGLE-NODE TREE\n"))
        tree = &Tree{}
        tree.Insert(y, "alpha")
        fmt.Println(bold("After insert:"))
        printTree(tree)
    }
    {
        tree.Delete(y)
        fmt.Println(bold(fmt.Sprintf("After deleting %v", y)))
        printTree(tree)
    }
}
