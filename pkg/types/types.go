package types

//Name ИФО клиента, который берёт кредит
type Name string

//City город, в котором был взят кредит
type City string

//Summa размер кредита
type Summa int64

//Year год, когда был выдан кредит
type Year int64

//Credit данные кредита
type Credit struct {
	Name Name
	City City
	Summa Summa
	Year Year
}