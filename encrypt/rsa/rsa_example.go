package rsa
//rsa 原理
//主要就是利用 大数因式分解的困难度来保证数学操作的单向性，即可以从私钥推导公钥，不能从公钥推导私钥
//http://www.ruanyifeng.com/blog/2013/06/rsa_algorithm_part_one.html
//http://www.ruanyifeng.com/blog/2013/07/rsa_algorithm_part_two.html
//https://blog.csdn.net/li396864285/article/details/79865806

//如何生成rsa密钥对
//how to find two primes for generation of rsa keys:
//https://crypto.stackexchange.com/questions/1970/how-are-primes-generated-for-rsa
