import json

def HEXArraytoStringArray(input):
    ret = []
    for i in range(len(input)):
        ret.append("0x%X" % input[i])
    return json.dumps(ret)

arr = [0x31ECC21A745E3968A04E9570E4425BC18FA8019C68028196B546D1669C200C68,
0x61035B26E3E9EEE00E0D72FD1EE8DDCA6894550DCA6916EA2AC6BAA90D11E510,
0x82A75BDEEAE8604D839476AE9EFD8B0E15AA447E21BFD7F41283BB54E22C9A82,
0x9B22D3D61959B4D3528B1D8BA932C96FBE302B36A1AAD1D95CAB54F9E0A135EA,
0x71beda120aafdd3bb922b360a066d10b7ce81d7ac2ad9874daac46e2282f6b45,
0x46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e94354,
0x7901cb5addcae2d210a531c604a76a660d77039093bac314de0816a16392aff1,
0x7ef464cf5a521d70c933977510816a0355b91a50eca2778837fb82da8448ecf6,
0x72a152ddfb8e864297c917af52ea6c1c68aead0fee1a62673fcc7e0c94979d00,
0x550d3de95be0bd28a79c3eb4ea7f05692c60b0602e48b49461e703379b08a71a,
0x28afdd85196b637a3c64ff1f53af1ad8de145cf652297ede1b38f2cbd6a4b4bf,
0x47197230e1e4b29fc0bd84d7d78966c0925452aff72a2a121538b102457e9ebe]

part1 = 0
part2 = 0

for i in range(int(len(arr)/2)):
    part1 = part1 ^ arr[2*i]
    part2 = part2 ^ arr[2*i+1]

part1_str =("%X" % part1)
part2_str =("%X" % part2)

print("0x"+part1_str)
print("0x"+part2_str)
concat_str = part1_str + part2_str
print("CONCAT: " + concat_str)

bytes4 = []
for i in range(0, len(concat_str), 8):
    bytes4.append("0x"+concat_str[i:i+8])

print(json.dumps(bytes4))

temp = [0x2fcaf464,0x8add2bfc,0xcd6bdb65,0x7e1a928a,0xc5698359,0x457c45f1,0xc7bf00b0,0xb93bcb1a,0xa0685161,0xf0d911d3,0x649f17ff,0x727e4267]
print(HEXArraytoStringArray(temp))