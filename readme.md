# emarket
**emarket** is a blockchain built using Cosmos SDK and Tendermint and created with Ignite CLI. This blockchain tracks different products which can be bought and sold by users.

## Chain setup:

```
ignite scaffold chain emarket
```

### Assets setup

```
ignite scaffold list item name product_type amount:int price:int discounted:bool
```
This chain supports products which have the fields `name`, `product_type`, `amount`, `price` and `discounted`.

As list command was used, it setup the required Query and Transaction messages to interact with the chain.
Commands setup were : 
### Query command supported
emarketd q emarket list-item

emarketd q emarket show-item

### Transaction command supported
emarketd tx emarket create-item [name] [productType] [amount] [price] [discounted] --from=[account] --chain-id emarket

emarketd tx emarket update-item [id] [name] [productType] [amount] [price] [discounted] --from=[account] --chain-id emarket

emarketd tx emarket delete-item [id] --from=[account] --chain-id emarket


### Run the chain
To run this chain, clone the main branch:
and run in the project folder

```
ignite chain build
ignite chain serve
```

## Introducing consensus-breaking changes
### What is consensus breaking change?
A consensus-breaking change modifies the core rules that validators use to agree on the next block in the blockchain.
Hence, all validators are required to upgrade their software to continue participating in the network.
If some validators fail to upgrade, they may follow outdated rules, causing a network fork.
A fork splits the network into two paths, one following the old rules and the other following the new rules.
Its critical to handle consensus breaking changes carefully, as they can lead to network instability and loss of data.

### How to handle consensus breaking changes?
There are two broad approaches to handling consensus-breaking changes:
1. Migration: In this approach, the chain is upgraded to a new version that includes the consensus-breaking change. 
This requires all validators to upgrade their software to the new version. The chain can be upgraded using a simple `ignite chain upgrade` command.
Implementing a migration function to transform existing state/data into the new format, ensuring compatibility with the updated rules.
Example: Adding a mandatory field like tags to the product item structure. A migration function will populate this field with default values (i.e. an empty list) for existing items.
2. Handle both types of data (dual) approach: Update transaction logic to handle both old and new data formats. Modify functions to account for the absence of new fields in legacy data.
Example: Allow item objects with and without the tags field during the transition, ensuring backward compatibility.

### Implementation of consensus breaking changes in emarket
The changes for emarket chain are implemented in the 'consensus-break' branch.
In the emarket chain, I introduced a consensus-breaking change by adding a new field `tags` to the product item structure.
To test if the change is accepted by the chain, I ensure to update the following files to ensure compatibility for the new field: 
1. In `x/emarket/module/autocli.go`: Updated the item structures for functions to include the new field `tags` (for both create and update functions). 
2. In `x/emarket/types/message_item.go`: Updated the item structures for functions to include the new field `tags` (for both create and update functions). 
3. In `x/emarket/keeper/msg_server_item.go`: Updated the item structures for functions to include the new field `tags` (for both create and update functions). 
4. In `proto/emarket/emarket/item.proto`: Updated the item structures to include the new field `tags` (and then ran `ignite chain build` for change to reflect in `x/emarket/types/item.pub.go` file). 
5. In `proto/emarket/emarket/tx.proto`: Updated the item structures to include the new field `tags` (and then ran `ignite chain build` for change to reflect in `x/emarket/types/tx.pub.go` file).
This was done for both create and update sub-commands.

### Testing the consensus breaking changes
First we need to ensure that while 'main' branch is checked out and `ignite chain serve` is running, the chain is working as expected.
And we store a few items in the chain using the `create-item` command.
```
emarketd tx emarket create-item apple fruit 10 5 false --from alice --chain-id emarket
```
Post this we need to switch to the 'consensus-break' branch and run then add a few more products using the `create-item` command.
Note: the creete-item command now has an additional field `tags` which needs to be filled in.
```
emarketd tx emarket create-item carrot vegetable 6 8 false "orange,long" --from bob --chain-id emarket
```

Post that if we query the chain to list all items,
```
emarketd q emarket list-item
```
we should see both the old and new items listed.


If such a chain is deployed in cloud with multiple validators then all validators need to upgrade their software to the new version to continue participating in the network.
Otherwise some validators will continue to follow old rules where `tags` field is not present and some will follow new rules where `tags` field is present.
Validators running the older version will reject transactions containing the tags field, leading to a fork in the network.
This is where consensus breaking changes can lead to network instability and loss of data.