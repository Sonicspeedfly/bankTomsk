package main

import (
	"github.com/Sonicspeedfly/bankTomsk/v1/pkg/credit"
	"github.com/Sonicspeedfly/bankTomsk/v1/pkg/types"
	"fmt"
)

func main() {
	Names := []types.Credit{
		{Name: "Умаров Алишер Алёшавич",
		City: "Томск",
		Summa: 10_000,
		Year: 2020,
		},
		{Name: "Иванов Джалилдин",
		City: "Душанбе",
		Summa: 9000,
		Year: 2020,
		},
		{Name: "Умаров Алексей Тошмахмадович",
		City: "Душанбе",
		Summa: 100_000,
		Year: 2020,
		},
		{Name: "Умарова Амина",
		City: "Душанбе",
		Summa: 10_000,
		Year: 2019,
		},
		{Name: "Тоджниссо Нурслатулоевна",
		City: "Томск",
		Summa: 10_000,
		Year: 2019,
		},		
		{Name: "Шукурова Муновара",
		City: "Томск",
		Summa: 50_000,
		Year: 2019,
		},
		{Name: "Суворов Юрий",
		City: "Томск",
		Summa: 90_000,
		Year: 2019,
		},
	}
	Search := credit.SearchCreditTomsk(Names,"Томск",10_000,90_000,2019)
	fmt.Println(Search) 
}