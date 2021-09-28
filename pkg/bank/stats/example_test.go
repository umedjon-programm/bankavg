package stats

import (
	"github.com/umedjon-programm/bankavg/pkg/bank/types"
	"fmt"
)
 func ExampleTotalInCategore(){
	 payments:=[]types.Payment{
		 {
			 ID: 1,
			 Amount: 1000,
			 Category: "oli",
		 },
		 {
			ID: 1,
			Amount: 1000,
			Category: "miyona",
		},
		{
			ID: 1,
			Amount: 1000,
			Category: "oli",
		},
		{
			ID: 1,
			Amount: 1000,
			Category: "oli",
		},

	 }
	 fmt.Println(TotalInCategory(payments,"oli"))
 }
 //Output:
 //3000
