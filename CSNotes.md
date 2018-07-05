CSNotes.md

## Polymorphism

```
Polymorphism is the provision of a single interface to entities of different types. A polymorphic type is one whose operations can also be applied to values of some other type, or types.
```

## Classless Interdomain Routing (CIDR)
https://www.youtube.com/watch?v=Q1U9wVXRuHA

CLASSFUL IP Adressing

- Class A: 16,777,214 Hosts
- Class B: 655,234 Hosts
- Class C: 254 Hosts

ex. If you have 2000 computers(2000 IP addresses)?

- Class B: wastes 63,000+ addresses
- Class C: using 8 of this class produces cluttered routing.


### Subnet Masks
- Decimal Subnet Mask  => ex. 255.255.255.0
- Binary Subnet Mask   => 11111111.11111111.11111111.00000000


Variable Length Subnet Masks (VLSM)
```
11111111 => 255
11111110 => 254
11111100 => 252
11111000 => 248
11110000 => 240
11100000 => 224
11000000 => 192
10000000 => 128
00000000 => 0
```
Subnet Mask
 NetworkID        HostID
|                |                 |
11111111.11111111.00000000.00000000

With Subnetting
 NetworkID        SubnetID  HostID
|                |        |        |
11111111.11111111.11111000.00000000

Formulas
Number of Subnets:
```
2n            n = Number of 1's in the Subnet ID
```
Number of available host addresses:
```
2n - 2        n = Number of 0's in the Host ID
```
ex. Class B
```
2**16 - 2 = 655,234 Host
```
* Host ID cannnot be all 0's or all 1's.

Example:

### Solution:

```
11111111.11111111.11111000.00000000
=> 2**5 = 32 Subnets
   2**11 - 2 = 2046 Available hosts in each subnet
```






