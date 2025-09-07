# Consistent Hashing

My implementation of the consistent hashing algorithm for learning
- hash_ring is the main data structure, allowing the addition of servers onto the ring and routing requests to their appropriate server

## Hash

I use sha256 to hash servers and requests into a big.Int object
- big.Int is not the most optimal, an implementation of a fixed-size 256-bit uint is preferred 

## Replicas

To achieve a more uniform distribution of the servers, there is an option to pass in #replicas
- #replicas will be placed on the ring when a server is added

## Routing Requests

I use an external BST library to keep the servers in sorted order
- When a request comes in, it gets hashed and will get routed to the first replica where *request_hash <= replica_hash*
- The time complexity of this operation can be improved by using binary search instead of iterating through every replica, but inefficient for a small # of replicas