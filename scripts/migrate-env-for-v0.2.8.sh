#!/bin/bash

set -e 

echo "Start Migrating Process"
echo "-----------------------"
echo "Backup env and docker-compose files"

date=$(date '+%Y-%m-%d')

cp .env .env_backup_$date
cp docker-compose.yml docker-compose-backup-$date.yml

echo "Finished backup for 2 files"

# To new docker-compose config in v0.2.8 
# Following docs here https://docs.roninchain.com/validators/setup/mainnet/run-combined

# - LISTENERS__RONIN__RPCURL=${RPC_ENDPOINT}
# - LISTENERS__RONIN__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY=${LISTENERS__RONIN__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY}
# - RONIN_LEGACY_BRIDGE_OPERATOR_KEY=${LEGACY_BRIDGE_OPERATOR_PRIVATE_KEY}
# - RONIN_BRIDGE_VOTER_KEY=${BRIDGE_VOTER_PRIVATE_KEY}
# - RONIN_RELAYER_KEY=${RONIN_RELAYER_KEY}
# - LISTENERS__ETHEREUM__RPCURL=${LISTENERS__ETHEREUM__RPCURL}
# - ETHEREUM_VALIDATOR_KEY=${ETHEREUM_VALIDATOR_KEY}
# - ETHEREUM_RELAYER_KEY=${ETHEREUM_RELAYER_KEY}
# - DATABASE__HOST=db
# - DATABASE__DBNAME=${DATABASE__DBNAME}
# - DATABASE__PORT=5432
# - DATABASE__USER=${DATABASE__USER}
# - DATABASE__PASSWORD=${DATABASE__PASSWORD}
# - VERBOSITY=${VERBOSITY}
# - CONFIG_PATH=${CONFIG_PATH}
# - LISTENERS__RONIN__TASKINTERVAL=${LISTENERS__RONIN__TASKINTERVAL}
# - LISTENERS__RONIN__TRANSACTIONCHECKPERIOD=${LISTENERS__RONIN__TRANSACTIONCHECKPERIOD}
# - LISTENERS__RONIN__MAXPROCESSINGTASKS=${LISTENERS__RONIN__MAXPROCESSINGTASKS}
# - LISTENERS__ETHEREUM__GETLOGSBATCHSIZE=${LISTENERS__ETHEREUM__GETLOGSBATCHSIZE}
# - LISTENERS__RONIN__GASLIMITBUMPRATIO=${LISTENERS__RONIN__GASLIMITBUMPRATIO}
# - LISTENERS__RONIN__STATS__NODE=${LISTENERS__RONIN__STATS__NODE}
# - LISTENERS__RONIN__STATS__HOST=${LISTENERS__RONIN__STATS__HOST}
# - LISTENERS__RONIN__STATS__SECRET=${LISTENERS__RONIN__STATS__SECRET}

echo "-------------------"
echo "Start replacing Ronin and Etherum Endpoint configs"
sed -i -e 's/RPC_ENDPOINT/LISTENERS__RONIN__RPCURL/g' .env
sed -i -e 's/BRIDGE_OPERATOR_PRIVATE_KEY/LISTENERS__RONIN__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY/g' .env
sed -i -e 's/ETHEREUM_ENDPOINT/LISTENERS__ETHEREUM__RPCURL/g' .env

sed -i -e 's/RONIN_RPC/LISTENERS__RONIN__RPCURL/g' docker-compose.yml
sed -i -e 's/RPC_ENDPOINT/LISTENERS__RONIN__RPCURL/g' docker-compose.yml
sed -i -e 's/RONIN_BRIDGE_OPERATOR_KEY/LISTENERS__RONIN__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY/g' docker-compose.yml
sed -i -e 's/BRIDGE_OPERATOR_PRIVATE_KEY/LISTENERS__RONIN__SECRET__BRIDGEOPERATOR__PLAINPRIVATEKEY/g' docker-compose.yml
sed -i -e 's/ETHEREUM_RPC/LISTENERS__ETHEREUM__RPCURL/g' docker-compose.yml
sed -i -e 's/ETHEREUM_ENDPOINT/LISTENERS__ETHEREUM__RPCURL/g' docker-compose.yml

echo "-------------------"

echo "Start replacing DATABASE config"

sed -i -e 's/DB_HOST/DATABASE__HOST/g' .env
sed -i -e 's/DB_NAME/DATABASE__DBNAME/g' .env
sed -i -e 's/DB_PORT/DATABASE__PORT/g' .env
sed -i -e 's/DB_USERNAME/DATABASE__USER/g' .env
sed -i -e 's/DB_PASSWORD/DATABASE__PASSWORD/g' .env

sed -i -e 's/DB_HOST/DATABASE__HOST/g' docker-compose.yml
sed -i -e 's/DB_NAME/DATABASE__DBNAME/g' docker-compose.yml
sed -i -e 's/DB_PORT/DATABASE__PORT/g' docker-compose.yml
sed -i -e 's/DB_USERNAME/DATABASE__USER/g' docker-compose.yml
sed -i -e 's/DB_PASSWORD/DATABASE__PASSWORD/g' docker-compose.yml

echo "-------------------"


echo "Start replacing Ronin Task Process config"

sed -i -e 's/RONIN_TASK_INTERVAL/LISTENERS__RONIN__TASKINTERVAL/g' .env
sed -i -e 's/RONIN_TRANSACTION_CHECK_PERIOD/LISTENERS__RONIN__TRANSACTIONCHECKPERIOD/g' .env
sed -i -e 's/RONIN_MAX_PROCESSING_TASKS/LISTENERS__RONIN__MAXPROCESSINGTASKS/g' .env

sed -i -e 's/RONIN_TASK_INTERVAL/LISTENERS__RONIN__TASKINTERVAL/g' docker-compose.yml
sed -i -e 's/RONIN_TRANSACTION_CHECK_PERIOD/LISTENERS__RONIN__TRANSACTIONCHECKPERIOD/g' docker-compose.yml
sed -i -e 's/RONIN_MAX_PROCESSING_TASKS/LISTENERS__RONIN__MAXPROCESSINGTASKS/g' docker-compose.yml

echo "-------------------"

echo "Start replacing Etherum Task Process config"

sed -i -e 's/ETHEREUM_GET_LOGS_BATCH_SIZE/LISTENERS__ETHEREUM__GETLOGSBATCHSIZE/g' .env
sed -i -e 's/ETHEREUM_GET_LOGS_BATCH_SIZE/LISTENERS__ETHEREUM__GETLOGSBATCHSIZE/g' docker-compose.yml

echo "-------------------"

echo "Start replacing Bridge Stats config"

sed -i -e 's/BRIDGE_STATS_URL/LISTENERS__RONIN__STATS__HOST/g' .env
sed -i -e 's/BRIDGE_STATS_NODE_NAME/LISTENERS__RONIN__STATS__NODE/g' .env
sed -i -e 's/BRIDGE_STATS_SECRET/LISTENERS__RONIN__STATS__SECRET/g' .env


sed -i -e 's/BRIDGE_STATS_URL/LISTENERS__RONIN__STATS__HOST/g' docker-compose.yml
sed -i -e 's/BRIDGE_STATS_NODE_NAME/LISTENERS__RONIN__STATS__NODE/g' docker-compose.yml
sed -i -e 's/BRIDGE_STATS_SECRET/LISTENERS__RONIN__STATS__SECRET/g' docker-compose.yml

echo "-------------------"
echo "Happy for new version v0.2.8, please double check with the final doc, and start bridge."
