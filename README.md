# godbms

Simple key-value storage with records expiring.

## Commands

**_kv \<storage_name\> create_** - create key-value storage

**_kv \<storage_name\> set \<key\> \<expire\> \<value\>_** - set key in storage with expire time (-1 - not expires)

**_kv \<storage_name\> get \<key\>_** - get value from storage by key