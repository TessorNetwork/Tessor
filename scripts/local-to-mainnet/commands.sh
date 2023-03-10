############################################
### WARNING: THIS FILE IS AUTOGENERATED. ###
###   ANY CHANGES WILL BE OVERWRITTEN.   ###
############################################
#### SETUP HOT WALLET (Only needs to be run once)
echo "$HOT_WALLET_1_MNEMONIC" | build/osmosisd keys add hot --recover --keyring-backend test 


#### START RELAYERS
# Create connections and channels
docker-compose -f scripts/local-to-mainnet/docker-compose.yml run --rm relayer rly transact link tessor-host 

# (OR) If the go relayer isn't working, use hermes (you'll have to add the connections to the relayer config though in `scripts/state/relaye/config/config.yaml`)
# docker-compose -f scripts/local-to-mainnet/docker-compose.yml run --rm hermes hermes create connection --a-chain osmosis-1 --b-chain local-test-1
# docker-compose -f scripts/local-to-mainnet/docker-compose.yml run --rm hermes hermes create channel --a-chain local-test-1 --a-connection connection-0 --a-port transfer --b-port transfer

# Ensure Relayer Config is updated (`scripts/state/relayer/config/config.yaml`)
#    paths:
#     tessor-host:
#       src:
#         chain-id: tessor-1
#         client-id: 07-tendermint-0
#         connection-id: connection-0
#       dst:
#         chain-id: cosmoshub-4
#         client-id: {CLIENT-ID}
#         connection-id: {CONNECTION-ID}

# Get channel ID created on the host
build/tessord --home s q ibc channel channels 
transfer_channel=$(build/tessord --home s q ibc channel channels | grep channel-0 -A 4 | grep counterparty -A 1 | grep channel | awk '{print $2}') && echo $transfer_channel

# Start Hermes Relayer
docker-compose -f scripts/local-to-mainnet/docker-compose.yml up -d hermes
docker-compose -f scripts/local-to-mainnet/docker-compose.yml logs -f hermes | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> scripts/logs/hermes.log 2>&1 &

# Configure the Go Relayer to only run ICQ
sed -i -E "s|rule: \"\"|rule: allowlist|g" scripts/state/relayer/config/config.yaml

# Start Go Relayer (for ICQ)
docker-compose -f scripts/local-to-mainnet/docker-compose.yml up -d relayer
docker-compose -f scripts/local-to-mainnet/docker-compose.yml logs -f relayer | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> scripts/logs/relayer.log 2>&1 &


#### REGISTER HOST
# IBC Transfer from HOST to tessor (from relayer account)
build/osmosisd tx ibc-transfer transfer transfer $transfer_channel tessor1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 4000000uosmo --from hot --chain-id osmosis-1 -y --keyring-backend test --node http://osmo-fleet-direct.main.tessornet.co:26657

# Confirm funds were recieved on tessor and get IBC denom
build/tessord --home s q bank balances tessor1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8

# Register host zone
IBC_DENOM=$(build/tessord --home s q bank balances tessor1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 | grep ibc | awk '{print $2}' | tr -d '"') && echo $IBC_DENOM
build/tessord --home s tx stakeibc register-host-zone \
    connection-0 uosmo osmo $IBC_DENOM channel-0 1 \
    --from admin --gas 1000000 -y

# Add validator
build/tessord --home s tx stakeibc add-validator osmosis-1 imperator osmovaloper1t8qckan2yrygq7kl9apwhzfalwzgc2429p8f0s 10 5 --chain-id local-test-1 --keyring-backend test --from admin -y

# Confirm ICA channels were registered
build/tessord --home s q stakeibc list-host-zone

#### FLOW
## Go Through Flow
# Liquid stake (then wait and LS again)
build/tessord --home s tx stakeibc liquid-stake 1000000 uosmo --keyring-backend test --from admin -y --chain-id local-test-1 -y

# Confirm stTokens, StakedBal, and Redemption Rate
build/tessord --home s q bank balances tessor1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
build/tessord --home s q stakeibc list-host-zone

# Redeem
build/tessord --home s tx stakeibc redeem-stake 1000 osmosis-1 osmo1c37n9aywapx2v0s6vk2yedydkkhq65zz38jfnc --from admin --keyring-backend test --chain-id local-test-1 -y

# Confirm stTokens and StakedBal
build/tessord --home s q bank balances tessor1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
build/tessord --home s q stakeibc list-host-zone

# Add another validator
build/tessord --home s tx stakeibc add-validator osmosis-1 notional osmovaloper1083svrca4t350mphfv9x45wq9asrs60c6rv0j5 10 5 --chain-id local-test-1 --keyring-backend test --from admin -y

# Liquid stake and confirm the stake was split 50/50 between the validators
build/tessord --home s tx stakeibc liquid-stake 1000000 uosmo --keyring-backend test --from admin -y --chain-id local-test-1 -y

# Change validator weights
build/tessord --home s tx stakeibc change-validator-weight osmosis-1 osmovaloper1t8qckan2yrygq7kl9apwhzfalwzgc2429p8f0s 1 --from admin -y
build/tessord --home s tx stakeibc change-validator-weight osmosis-1 osmovaloper1083svrca4t350mphfv9x45wq9asrs60c6rv0j5 49 --from admin -y

# LS and confirm delegation aligned with new weights
build/tessord --home s tx stakeibc liquid-stake 1000000 uosmo --keyring-backend test --from admin -y --chain-id local-test-1 -y

# Call rebalance to and confirm new delegations
build/tessord --home s tx stakeibc rebalance-validators osmosis-1 5 --from admin

# Clear balances
fee_address=$(build/tessord --home s q stakeibc show-host-zone osmosis-1 | grep feeAccount -A 1 | grep address | awk '{print $2}') && echo $fee_address
balance=$(build/osmosisd --home s q bank balances $fee_address | grep amount | awk '{print $3}' | tr -d '"') && echo $balance
build/tessord --home s tx stakeibc clear-balance osmosis-1 $balance $transfer_channel --from admin

# Update delegations (just submit this query and confirm the ICQ callback displays in the tessor logs)
# Must be submitted in ICQ window
build/tessord --home s tx stakeibc update-delegation osmosis-1 osmovaloper1t8qckan2yrygq7kl9apwhzfalwzgc2429p8f0s --from admin -y

#### MISC 
# If a channel closes, restore it with:
build/tessord --home s tx stakeibc restore-interchain-account osmosis-1 {DELEGATION | WITHDRAWAL | FEE | REDEMPTION} --from admin