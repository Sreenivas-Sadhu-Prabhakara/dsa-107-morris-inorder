# Morris Inorder Traversal

> Striver SDE Sheet 107 · Day 17 Binary Tree — O(1) space: thread the predecessor's nil right pointer back to the root, walk it, unthread on return. Go, stdlib only, tested.

## Problem

Inorder traversal without recursion *and without a stack* — O(1) auxiliary space. The stack existed to remember the way back up; Morris stores that information in the tree itself.

## Approach

Before descending left, find the current node's inorder predecessor (the rightmost node of its left subtree) and point its unused nil `Right` at the current node — a *thread*. Descend. When the traversal later runs off that predecessor, the thread teleports it back to the node, and finding `pred.Right == cur` on the second visit is the signal: the left side is done, so visit, remove the thread, and go right. Every edge is walked at most a constant number of times (the predecessor hunt re-walks each right spine once per thread) — amortized linear. The test asserts the threads are gone afterwards: a Morris traversal that forgets to unthread leaves a corrupted, cyclic tree.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
