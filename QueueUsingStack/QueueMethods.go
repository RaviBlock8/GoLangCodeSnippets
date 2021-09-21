package main

import "errors"

type queueStruct struct{
	queue []int
	front int
	back int
	cap int
}

func (q *queueStruct) enqueue(val int)(bool,error){
	if(q.back==(q.cap-1)){
		return false,errors.New("Overflow!!")
	}
	q.back+=1
	q.queue[q.back]=val
	return true,nil
}

func (q *queueStruct) dequeue()(int,error){
	if(q.front==-1){
		return 0,errors.New("Underflow!!")
	}
	popped:=q.queue[q.front]
	if(q.front==q.back){
		q.front=0
		q.back=0
	}else{
		q.front+=1
	}
	return popped,nil
}