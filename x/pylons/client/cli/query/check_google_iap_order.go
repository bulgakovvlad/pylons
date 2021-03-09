package query

import (
	"fmt"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// CheckGoogleIAPOrder check if google iap order is already used
func CheckGoogleIAPOrder() *cobra.Command {
	ccb := &cobra.Command{
		Use:   "check_google_iap_order <purchase_token>",
		Short: "check if google iap order is given to user with purchase token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			tokenReq := &types.CheckGoogleIAPOrderRequest{
				PurchaseToken: args[0],
			}

			res, err := queryClient.CheckGoogleIAPOrder(cmd.Context(), tokenReq)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(
				fmt.Sprintf(
					"purchaseToken: %s \nexist: %t",
					res.GetPurchaseToken(), res.Exist),
			)
		},
	}
	return ccb
}
