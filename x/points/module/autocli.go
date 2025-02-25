package points

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "sixswingchain/api/sixswingchain/points"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetOwnerRequest",
					Use:            "get-owner-request [name]",
					Short:          "Query get-owner-request",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Mint",
					Use:            "mint [wallet] [amount]",
					Short:          "Send a mint tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "wallet"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "SetOwner",
					Use:            "set-owner [new-owner]",
					Short:          "Send a set-owner tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "newOwner"}},
				},
				{
					RpcMethod:      "Burn",
					Use:            "burn [owner] [amount]",
					Short:          "Send a burn tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner"}, {ProtoField: "amount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
