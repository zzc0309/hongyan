package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Character struct {
	name string
	hp int
	mp int
	attack_num int
}

func random_number()int{
	rand.Seed(time.Now().Unix())
	//fmt.Println(rand.Intn(2))
	return rand.Intn(2)
}

func A_attack_B(A *Character,B *Character)(){
	time.Sleep(1000000 * time.Microsecond)
B.hp-=A.attack_num
B.hp-=10
if random_number()==1{
	time.Sleep(1000000 * time.Microsecond)
	A_attack_B(A,B)}
//fmt.Println("B的生命值:",B.hp)
	return
}

func B_attack_A(B *Character,A *Character)(){
	if B.mp==50{
		A.attack_num-=A.attack_num/10
	}
	A.hp-=B.attack_num
	//fmt.Println("A的生命值:",A.hp)
	B.mp+=10
}

func battle(A *Character,B *Character)bool{
	for i:=0;i<=1000;i++{
		A_attack_B(A,B)
		if B.hp==0{return true}
		B_attack_A(B,A)
		if A.hp==0{return false}
	}
	return true
}
func battle2(A *Character,B *Character)bool{
	for i:=0;i<=1000;i++{
		B_attack_A(B,A)
		if A.hp<=0{return false}
		A_attack_B(A,B)
		if B.hp<=0{return true}
	}
	return true
}
//func A_skill
func main() {
	A := Character{name: "A角色", hp: 100, mp: 0, attack_num: 10}
	B := Character{name: "B角色", hp: 300, mp: 0, attack_num: 20}
	A_count:=0
	B_count:=0
	for i:=0;i<=4;i++ {
		if battle(&A, &B) {
			A_count++
		} else {
			B_count++
		}
		A.hp=100
		A.attack_num=10
		B.hp=300
		B.mp=0
	}
	for i:=0;i<=4;i++ {
		if battle2(&A, &B) {
			A_count++
		} else {
			B_count++
		}
		A.hp=100
		A.attack_num=10
		B.hp=300
		B.hp=0
	}
	fmt.Println("A的胜率:",A_count*10,"%","B的胜率:",B_count*10,"%")//差不多强
}