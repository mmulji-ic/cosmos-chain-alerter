package models

import "time"

// https://api.mintscan.io/v1/cosmos/status
type CosmosStatsu struct {
	ChainID              string    `json:"chain_id"`
	BlockHeight          int       `json:"block_height"`
	BlockTime            float64   `json:"block_time"`
	TotalTxsNum          int       `json:"total_txs_num"`
	JailedValidatorNum   int       `json:"jailed_validator_num"`
	UnjailedValidatorNum int       `json:"unjailed_validator_num"`
	TotalValidatorNum    int       `json:"total_validator_num"`
	Timestamp            time.Time `json:"timestamp"`
	CommunityPool        []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"community_pool"`
	BondedTokens      int64 `json:"bonded_tokens"`
	NotBondedTokens   int64 `json:"not_bonded_tokens"`
	TotalSupplyTokens struct {
		Supply []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"supply"`
		Pagination struct {
			NextKey interface{} `json:"next_key"`
			Total   string      `json:"total"`
		} `json:"pagination"`
	} `json:"total_supply_tokens"`
	TotalCirculatingTokens struct {
		Supply []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"supply"`
		Pagination struct {
			NextKey interface{} `json:"next_key"`
			Total   string      `json:"total"`
		} `json:"pagination"`
	} `json:"total_circulating_tokens"`
	Inflation string `json:"inflation"`
}

// https://api.mintscan.io/v1/cosmos/validators
type ValidatorInfo []struct {
	Rank              int       `json:"rank"`
	OperatorAddress   string    `json:"operator_address"`
	ConsensusPubkey   string    `json:"consensus_pubkey"`
	Jailed            bool      `json:"jailed"`
	Status            int       `json:"status"`
	Tokens            string    `json:"tokens"`
	DelegatorShares   string    `json:"delegator_shares"`
	Moniker           string    `json:"moniker"`
	Identity          string    `json:"identity"`
	Website           string    `json:"website"`
	Details           string    `json:"details"`
	UnbondingHeight   string    `json:"unbonding_height"`
	UnbondingTime     time.Time `json:"unbonding_time"`
	UpdateTime        time.Time `json:"update_time"`
	MinSelfDelegation string    `json:"min_self_delegation"`
	KeybaseURL        string    `json:"keybase_url"`
	AccountAddress    string    `json:"account_address"`
	Rate              string    `json:"rate"`
	MaxRate           string    `json:"max_rate"`
	MaxChangeRate     string    `json:"max_change_rate"`
	Uptime            struct {
		Address      string `json:"address"`
		MissedBlocks int    `json:"missed_blocks"`
		OverBlocks   int    `json:"over_blocks"`
	} `json:"uptime"`
}

// https://api.mintscan.io/v1/cosmos/validators/votes
type VotesInfo struct {
	TargetCount string `json:"target_count"`
	Validators  []struct {
		Address         string `json:"address"`
		OperatorAddress string `json:"operator_address"`
		VoteCount       string `json:"vote_count"`
	} `json:"validators"`
}

// https://api.mintscan.io/v2/utils/market/price/uatom
type UAtomInfo struct {
	Denom                     string    `json:"denom"`
	CoinGeckoID               string    `json:"coinGeckoId"`
	CurrentPrice              float64   `json:"current_price"`
	MarketCap                 int64     `json:"market_cap"`
	DailyVolume               int       `json:"daily_volume"`
	DailyPriceChangeInPercent float64   `json:"daily_price_change_in_percent"`
	LastUpdated               time.Time `json:"last_updated"`
}

// https://api.mintscan.io/v1/cosmos/node_info
type NodeInfo struct {
	NodeInfo struct {
		ProtocolVersion struct {
			P2P   string `json:"p2p"`
			Block string `json:"block"`
			App   string `json:"app"`
		} `json:"protocol_version"`
		ID         string `json:"id"`
		ListenAddr string `json:"listen_addr"`
		Network    string `json:"network"`
		Version    string `json:"version"`
		Channels   string `json:"channels"`
		Moniker    string `json:"moniker"`
		Other      struct {
			TxIndex    string `json:"tx_index"`
			RPCAddress string `json:"rpc_address"`
		} `json:"other"`
	} `json:"node_info"`
	ApplicationVersion struct {
		Name             string   `json:"name"`
		ServerName       string   `json:"server_name"`
		Version          string   `json:"version"`
		Commit           string   `json:"commit"`
		BuildTags        string   `json:"build_tags"`
		Go               string   `json:"go"`
		BuildDeps        []string `json:"build_deps"`
		CosmosSdkVersion string   `json:"cosmos_sdk_version"`
	} `json:"application_version"`
}

// https://api.mintscan.io/v1/cosmos/block/blocktime
type BlockTime struct {
	BlockTime float64 `json:"block_time"`
}

// https://api.mintscan.io/v1/relayer/cosmoshub-4/paths
type RelayerPaths struct {
	Sendable []struct {
		ChainID string `json:"chain_id"`
		Paths   []struct {
			ChannelID    string `json:"channel_id"`
			PortID       string `json:"port_id"`
			ChannelState string `json:"channel_state"`
			CounterParty struct {
				ChannelID    interface{} `json:"channel_id"`
				PortID       string      `json:"port_id"`
				ChannelState string      `json:"channel_state"`
			} `json:"counter_party"`
			Auth  bool `json:"auth"`
			Stats struct {
				Current struct {
					TxNum struct {
						Transfer int `json:"transfer"`
						Receive  int `json:"receive"`
					} `json:"tx_num"`
					Vol struct {
						Transfer interface{} `json:"transfer"`
						Receive  interface{} `json:"receive"`
					} `json:"vol"`
				} `json:"current"`
				Past struct {
					TxNum struct {
						Transfer int `json:"transfer"`
						Receive  int `json:"receive"`
					} `json:"tx_num"`
					Vol struct {
						Transfer interface{} `json:"transfer"`
						Receive  interface{} `json:"receive"`
					} `json:"vol"`
				} `json:"past"`
			} `json:"stats"`
			CreatedAt interface{} `json:"created_at"`
		} `json:"paths"`
	} `json:"sendable"`
}

// https://lcd-cosmos.cosmostation.io/cosmos/liquidity/v1beta1/pools
type Pools struct {
	Pools []struct {
		ID                    string   `json:"id"`
		TypeID                int      `json:"type_id"`
		ReserveCoinDenoms     []string `json:"reserve_coin_denoms"`
		ReserveAccountAddress string   `json:"reserve_account_address"`
		PoolCoinDenom         string   `json:"pool_coin_denom"`
	} `json:"pools"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

// https://api.mintscan.io/v2/utils/market/prices?currency=usd
type DenomPrice []struct {
	Denom                     string    `json:"denom"`
	CoinGeckoID               string    `json:"coinGeckoId"`
	CurrentPrice              float64   `json:"current_price"`
	MarketCap                 int       `json:"market_cap"`
	DailyVolume               int       `json:"daily_volume"`
	DailyPriceChangeInPercent float64   `json:"daily_price_change_in_percent"`
	LastUpdated               time.Time `json:"last_updated"`
}
