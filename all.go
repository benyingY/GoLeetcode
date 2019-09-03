package GoLeetcode

//70 爬楼梯
func climbStairs(n int) int {
	a := 1
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b

}

//120 三角形最小路径和
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= len(triangle[row])-1; col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

//152 乘积最大子序列
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	maximum, big, small := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			big, small = small, big
		}

		big = max(nums[i], nums[i]*big)
		small = min(nums[i], nums[i]*small)
		maximum = max(maximum, big)
	}
	return maximum
}

//300 最长上升子序列
import "sort"

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}

//146 LRU Cache
type LRUCache struct {
    head *node
    tail *node
    store map[int]*node
    cap int
}

type node struct {
    prev *node
    next *node
    key int
    val int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{cap:capacity, store:make(map[int]*node)}
}

func (c *LRUCache) removeFromChain(n *node) {
    if n == c.head {
        c.head = n.next
    }
    
    if n == c.tail {
        c.tail = n.prev
    }
    
    if n.next != nil {
        n.next.prev = n.prev
    }
    
    if n.prev != nil {
        n.prev.next = n.next
    }
}

func (c *LRUCache) addToChain(n *node) {
    n.prev = nil
    n.next = c.head
    if n.next != nil {
        n.next.prev = n
    }
    c.head = n
    if c.tail == nil {
        c.tail = n
    }
}

func (c *LRUCache) Get(key int) int {
    n, ok := c.store[key]
    if !ok {
        return -1
    }
    
    c.removeFromChain(n)
    c.addToChain(n)
    return n.val
}


func (c *LRUCache) Put(key int, value int)  {
    n, ok := c.store[key]
    if !ok {
        n = &node{val:value, key:key}
        c.store[key] = n
    } else {
        n.val = value
        c.removeFromChain(n)
    }
    
    c.addToChain(n)
    if len(c.store) > c.cap {
        n = c.tail
        if n != nil {
            c.removeFromChain(n)
            delete(c.store, n.key)
        }
    }
}
