package stats

import (
	"github.com/umedjon-programm/bankavg/pkg/bank/types"
	"fmt"
)
 func ExampleAvg(){
	 payments:=[]types.Payment{
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
	 fmt.Println(Avg(payments))
 }
 //Output:
 //1000
