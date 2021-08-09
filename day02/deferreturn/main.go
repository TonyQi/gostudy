package main

//go语言的return带defer操作是分成两步执行的，不是原子操作，在底层中分步操作，
//第一步  返回值赋值
//第二步  执行defer
//第三步  执行RET操作
