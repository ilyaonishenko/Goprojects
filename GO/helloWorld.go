package main
import "fmt"

func main(){
  //countTheString("hello")
  n:=[]int{5, -1, 0, 12, 3, 5}
  fmt.Printf("unsorted: %v\n",n)
  bubblesort(n)
  fmt.Printf("sorted: %v\n",n)

}
func printA(count int, a string){
  for i:=0;i<count;i++{
    print(a)
  }
  println("")
}
func countTheString(text string){
  println(len(text))
}
//q6
func average(xs []float64) (avg float64){
  sum:=0.0
  switch len(xs){
    case 0:
      avg=0
    default:
      for v:=range xs{
        sum+=float64(v)
      }
      avg = sum/float64(len(xs))
  }
  return
}
func rightOrder(a,b int)(int,int){
  if a>b{
    return b,a
  }
  return a,b
}
type stack struct{
  i int
  data [10]int
}
func (s *stack) push (k int){
  if s.i+1>9{
    return
  }
  s.data[s.i]=k
  s.i++
}
func (s *stack) pop ()int{
  s.i--
  return s.data[s.i]
}
func fibonacci(start int)[]int{
  x:=make([]int,start)
  x[0]=1;x[1]=1
  for i:=2;i<start;i++{
    x[i]=x[i-1]+x[i-2]
  }
  return x
}
func bubblesort(arr []int){
  for i:=0;i<len(arr)-1;i++{
    for j:=i+1;j<len(arr);j++{
      if arr[j]<arr[i]{
        arr[i],arr[j]=arr[j],arr[i]
      }
    }
  }
}
