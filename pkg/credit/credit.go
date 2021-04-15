package credit

import (
	"github.com/Sonicspeedfly/bankTomsk/v1/pkg/types"
)

//SearchCreditTomsk ищет клинтов, чьи кредит были зарегестрированы - в том или ином городе, тот или иной год, от одной суммы до данной суммы
func SearchCreditTomsk(Names []types.Credit, city types.City, summa1 types.Summa, summa2 types.Summa, year types.Year) []types.Credit {
	var TomskCredit []types.Credit
	for _, Name := range Names {
		if((Name.City==city)&&(Name.Year==year)&&(Name.Summa>=summa1)&&(Name.Summa<=summa2)){
			TomskCredit = append(TomskCredit, types.Credit{Name: Name.Name, City: Name.City, Summa: Name.Summa, Year: Name.Year})
		}
	}
	return TomskCredit
}