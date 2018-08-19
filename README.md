# godbms

Simple key-value storage with records expiring.

## Commands

**kv \<storage_name\> create** - create key-value storage

**kv \<storage_name\> set \<key\> \<expire\> \<value\>** - set key in storage with expire time (-1 - not expires)

**kv \<storage_name\> get <key>** - get value from storage by key