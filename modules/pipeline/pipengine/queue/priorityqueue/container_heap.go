// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package priorityqueue

// implement pkg container heap.Interface
func (pq priorityQueue) Len() int { return len(pq.items) }

func (pq priorityQueue) Less(i, j int) bool {
	if pq.items[i].Priority() == pq.items[j].Priority() {
		return pq.items[i].CreationTime().Before(pq.items[j].CreationTime())
	}
	return pq.items[i].Priority() > pq.items[j].Priority()
}

func (pq priorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].SetIndex(i)
	pq.items[j].SetIndex(j)
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(Item)
	item.SetIndex(n)
	pq.items = append(pq.items, item)
	pq.itemByKey[item.Key()] = item
}

func (pq *priorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	item.SetIndex(-1)
	pq.items = old[0 : n-1]
	delete(pq.itemByKey, item.Key())
	return item
}
