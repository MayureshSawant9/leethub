import (
    "container/heap"
    "sort"
)

type minHeap []int

func (h minHeap) Len() int {
    return len(h)
}

func (h minHeap) Less(x, y int) bool {
    return h[x] < h[y]
}

func (h minHeap) Swap(x, y int) {
    h[x], h[y] = h[y], h[x]
}

func (h *minHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
    old := *h
    x := old[len(old) - 1]
    *h = old[:len(old) - 1]
    return x
}


func maxEvents(events [][]int) int {
    sort.Slice(events, func (i, j int) bool {
        return events[i][0] < events[j][0]
    })

    mHeap := &minHeap{}
    heap.Init(mHeap)

    i, n, day, result, maxDay := 0, len(events), 1, 0, 0

    for _, event := range events {
        if maxDay < event[1] {
            maxDay = event[1]
        }
    }

    for day = 1; day <= maxDay; day++ {
        for i < n && events[i][0] == day {
            heap.Push(mHeap, events[i][1])
            i++
        }

        for mHeap.Len() > 0 && (*mHeap)[0] < day {
            heap.Pop(mHeap)
        }

        if mHeap.Len() > 0 {
            heap.Pop(mHeap)
            result++
        }
    }

    return result
}