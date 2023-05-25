# Imports
import json
from web3 import Web3
from pprint import pprint
# Init full and archive provider
full_node_provider = Web3(
    Web3.HTTPProvider(
        "https://nd-479-987-415.p2pify.com/271156373b36700f7576cf46e68b1262"
    )
)
archive_node_provider = Web3(
    Web3.HTTPProvider(
        "https://nd-072-228-848.p2pify.com/3f7a80739e2f6739cae0256a2660725b"
    )
)

def to_checksum_address(address):
    return full_node_provider.toChecksumAddress(address)

def to_hex(string):
    return full_node_provider.toHex(string)

# Returns the current block number of a network
def get_block_number():
    return full_node_provider.eth.block_number

# Returns the ETH balance of a given address at a given block
def get_eth_balance(address, block):
    print(
        "[QUERYING] Fetching ETH balance from address {} at block {}".format(
            address, block
        )
    )
    try:
        print("[QUERYING] Attempting with full Node")
        return full_node_provider.eth.get_balance(address, block)
    except Exception as e:
        if "missing trie node" in str(e):
            print("[OLD-BLOCK-QUERY] Switching to archive query")
            return archive_node_provider.eth.get_balance(address, block)
        else:
            print("exception: ", e)
            return None

# Returns the storage of an address at a given position and block
def get_storage_at(address, position, block):
    try:
        print(
            "[QUERYING] Fetching storage at address {} at position {} at block {}".format(
                address, position, block
            )
        )
        return full_node_provider.eth.get_storage_at(address, position, block)
    except Exception as e:
        if "missing trie node" in str(e):
            print("[OLD-BLOCK-QUERY] Switching to archive query")
            return archive_node_provider.eth.get_storage_at(address, position, block)
        else:
            return None

# Returns the code at a given address and block
def get_code(address, block):
    try:
        print(
            "[QUERYING] Fetching code at address {} at block {}".format(address, block)
        )
        return full_node_provider.eth.get_code(address, block)
    except Exception as e:
        if "missing trie node" in str(e):
            print("[OLD-BLOCK-QUERY] Switching to archive query")
            return archive_node_provider.eth.get_code(address, block)
        else:
            return None

# Returns the mined transactions in a given block
def get_block_transactions(block, full_transactions=False):
    try:
        print("[QUERYING] Fetching transactions from block {}".format(block))
        return full_node_provider.eth.get_block(block, full_transactions)
    except Exception as e:
        if "missing trie node" in str(e):
            print("[OLD-BLOCK-QUERY] Switching to archive query")
            return archive_node_provider.eth.get_block(block, full_transactions)
        else:
            return None

def print_eth_balance_of(address, block):
    eth_balance = get_eth_balance(address, block)
    print(
        "[BALANCE-RESULTS] Eth balance of address {} at block {}: {} $ETH".format(
            address, block, eth_balance
        )
    ) if eth_balance is not None else print("Invalid Query")

def print_storage_at(address, position, block):
    pprint(address)
    pprint(position)
    pprint(block)
    storage_at = full_node_provider.toHex(get_storage_at(address, position, block))
    print(
        "[STORAGE-AT-RESULTS] Storage at {} at position {} at block {}: {}".format(
            address, position, block, storage_at
        )
    ) if storage_at is not None else print("Invalid query")

def print_code_at(address, block):
    code_at = full_node_provider.toHex(get_code(address, block))
    print(
        "[CODE-AT-RESULTS] Code at address {} at block {}: {}".format(
            address, block, code_at
        )
    ) if code_at is not None else print("Invalid query")

def print_block_transactions(block, full):
    block_transactions = get_block_transactions(block, full)
    print(
        "[TRANSACTIONS] Transactions at block {}: {}".format(block, block_transactions)
    ) if block_transactions is not None else print("Invalid Query")


if __name__ == '__main__':
    address = input('Please input the address: ')
    block = input('Please input the blocknum: ')
    position = input('input slot: ') 
    print_storage_at(address, position, block)
    # a = float(a)
    # b = float(b)
    # print('add result:', add(a, b))
    # print('multi result:', multi(a, b))



def get_pair_info_from_factory(self, src_addr, dst_addr):
        pancake_factory = contract_function_call(self.route_contract_instance, 'factory')
        factory_contract = self.w3.eth.contract(address=pancake_factory, abi=FACTORY_ABI)
        pair_addr = contract_function_call(factory_contract, 'getPair', src_addr, dst_addr)
        reverse0, reverse1, block_timestamp_last = contract_function_call(factory_contract, 'getReverses')
        token0 = contract_function_call(factory_contract, 'token0')
        token1 = contract_function_call(factory_contract, 'token1')