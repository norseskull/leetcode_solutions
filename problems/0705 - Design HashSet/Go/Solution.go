type MyHashSet struct {
    List []bool
}


func Constructor() MyHashSet {
    return MyHashSet{make([]bool, 1000001)}
}


func (this *MyHashSet) Add(key int)  {
    this.List[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    this.List[key] = false
}


func (this *MyHashSet) Contains(key int) bool {
    return this.List[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */