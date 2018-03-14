# Ethereum  API Document

* <a href="#node" class="scrollto">nodeInfo</a>
* <a href="#block" class="scrollto">blockInfo</a>
* <a href="#transation" class="scrollto">transationInfo</a>
* <a href="#startminer" class="scrollto">startminer</a>
* <a href="#stopminer" class="scrollto">stopminer</a>
* <a href="#sendtrans" class="scrollto">sendtrans</a>



<a name="node"></a>
# nodeInfo
## <span class="label label-default">GET /node
<pre>
Input:
    EX: /node

Output:
{
    "enode": "enode://fa3d2bf4cf297fd851b961182042ad0e65b10e1026ecca9c2e3fda249647e38e78847e456cb779b695e4e5f571dd3e65391a498a09887c5fce36f4085f889db4@xxx.xxx.xx.xx:30303",
    "name": "Geth/v1.8.2-stable/darwin-amd64/go1.10"
}
</pre>


<a name="block"></a>
# blockInfo
## <span class="label label-default">GET /block/:block_number
<pre>
{

Input:
    EX: /block/1000

Output:
{
    "hash": "0x5b4590a9905fa1c9cc273f32e6dc63b4c512f0ee14edc6fa41c26b416a7b5d58",
    "parentHash": "0xc31b362e591aa07faa977dbc492ae43cd47eef291920435153bbbf3acaf2fc2f",
    "miner": "0xbb7b8287f3f0a933474a79eae42cbca977791171",
    "gasLimit": 5000,
    "gasUsed": 0,
    "difficulty": 27800789535,
    "totalDifficulty": 22019797038325
}
</pre>


<a name="transation"></a>
# transationInfo
## <span class="label label-default">GET /transation/:transation_hash
<pre>
Input:
   EX: /transation/0xe9e91f1ee4b56c0df2e9f06c2b8c27c6076195a88a7b8537ba8313d80e6f124e

Output:
{
    "hash": "0xe9e91f1ee4b56c0df2e9f06c2b8c27c6076195a88a7b8537ba8313d80e6f124e",
    "nonce": 17387,
    "blockHash": "0x8e38b4dbf6b11fcc3b9dee84fb7986e29ca0a02cecd8977c161ff7333329681e",
    "blockNumber": 1000000,
    "from": "0x32be343b94f860124dc4fee278fdcbd38c102d88",
    "to": "0xdf190dc7190dfba737d7777a163445b7fff16133",
    "gas": 50000,
    "gasPrice": 60000000000,
    "value": 437194980000000000
}
</pre>


<a name="startminer"></a>
# startminer
## <span class="label label-default">PUT /startminer
<pre>
Input:
   EX: /startminer

</pre>


<a name="stopminer"></a>
# stopminer
## <span class="label label-default">DELETE /stopminer
<pre>
Input:
   EX: /stopminer

</pre>


<a name="sendtrans"></a>
# sendtrans
## <span class="label label-default">POST /sendtrans
<pre>
Input:
   EX: POST Body:
       {
         "transation": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
       }

</pre>









