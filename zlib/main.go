package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"

	"github.com/weilinfox/youmu-hyouibana-spectacle/lib"
)

func main() {

	rd := []byte{0x78, 0x9c, 0x9d, 0x52, 0x6d, 0x0e, 0xc2, 0x20, 0x0c, 0x2d, 0x46, 0x8d, 0x3a, 0x4d, 0xf6, 0xc7, 0x43, 0x78, 0x0b, 0x6f, 0xb2, 0xe0, 0xc6, 0x16, 0x22, 0x0c, 0x03, 0xc3, 0x64, 0xf7, 0xf2, 0x80, 0xb6, 0x73, 0x1f, 0xcd, 0x12, 0x8d, 0xfa, 0xab, 0xed, 0xeb, 0x7b, 0x8f, 0x52, 0x48, 0x01, 0x36, 0x5b, 0x00, 0xa8, 0xa4, 0x55, 0x99, 0x75, 0x85, 0x5a, 0x00, 0xac, 0x04, 0x02, 0x29, 0x36, 0x96, 0x18, 0x09, 0x3b, 0x63, 0x8e, 0x38, 0x50, 0x0f, 0x66, 0x51, 0xb0, 0x1a, 0x73, 0x41, 0xba, 0x84, 0x74, 0x32, 0x34, 0xca, 0x67, 0x35, 0xfa, 0xce, 0xe5, 0x44, 0xa1, 0xc4, 0x2b, 0x6d, 0x23, 0x3f, 0x6f, 0xdd, 0xe9, 0xbc, 0x0e, 0x72, 0x70, 0xda, 0x21, 0x12, 0x8c, 0xbc, 0xab, 0xb7, 0x46, 0x93, 0x86, 0x3b, 0x8d, 0xf6, 0x83, 0x11, 0xd1, 0x2e, 0x95, 0xcd, 0x74, 0x41, 0xb4, 0xd8, 0xd3, 0x68, 0xd0, 0x9b, 0x91, 0xed, 0x87, 0x41, 0xf9, 0x3d, 0x87, 0x9a, 0x4f, 0x57, 0xe8, 0xb2, 0xd4, 0x79, 0x34, 0x4d, 0xcb, 0x55, 0xc9, 0x38, 0x77, 0xee, 0x8c, 0xf3, 0xdf, 0x2c, 0x50, 0x30, 0xe3, 0xfd, 0xb4, 0xc0, 0xbf, 0xf4, 0xd4, 0x0c, 0x37, 0x65, 0xcc, 0xaf, 0x2f, 0x77, 0xa0, 0x35, 0xc9, 0xfc, 0x5a, 0x79, 0x17, 0xeb, 0xa2, 0xdf, 0xd6, 0x89, 0x7d, 0x87, 0xa0, 0x54, 0x87, 0x3d, 0x8e, 0x2f, 0xd1, 0x13, 0x87, 0xcb, 0x3b, 0x6f}
	b := bytes.NewBuffer(rd)
	r, err := zlib.NewReader(b)
	if err != nil {
		fmt.Println(err)
	} else {
		ans := make([]byte, 2048)
		n, err := r.Read(ans)
		r.Close()
		fmt.Println(n, err == io.EOF, err, ans[:n])

		fmt.Println("Decode:")
		fmt.Println(lib.ZlibDataDecode(n, rd))
	}

}

/*
华33233-10800对战
华42237-10800对战
华60513-10800拒绝 beta-
华44555-10800拒绝 -beta
华59169-10800对战 beta-beta 似乎他们在协议上没啥区别

33233->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 76 5c 00 00
42237->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 f5 61 02 00
60513->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 bf 83 00 00
44555->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 85 d6 01 00
59169->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 5d 64 00 00

10800->33233 04
10800->42237 04
10800->60513 04
10800->44555 04
10800->59169 04
33233->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 76 5c 00 00
            00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00
            78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd b8 f3 78 25 96 82 42 3f 50 eb e7 4c 5e 52 93 e2 6a 86 17 c2 4c
            80 ee 00 b8 bb 64 52 80 9a 18 46 06 f2 3e 5d e6 a2 62 57 cd 5c bd 65 8f 14 9c 74 ae 90 1c 38 17 2a c5 3b 13
            8f d0 00 9d fa 2f 3d e7 e9 f2 be a5 28 fc 3e 73 65 59 b3 c9 a7 2c 28 df b2 fa 01 83 4c 14 6b
42237->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 f5 61 02 00
            00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00
            78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd b8 f3 78 25 96 82 42 3f 50 eb e7 4c 5e 52 93 e2 6a 86 17 c2 4c
            80 ee 00 b8 bb 64 52 80 9a 18 46 06 f2 3e 5d e6 a2 62 57 cd 5c bd 65 8f 14 9c 74 ae 90 1c 38 17 2a c5 3b 13
            8f d0 00 9d fa 2f 3d e7 e9 f2 be a5 28 fc 3e 73 65 59 b3 c9 a7 2c 28 df b2 fa 01 83 4c 14 6b
60513->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 bf 83 00 00
            00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00
            78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd b8 f3 78 25 96 82 42 3f 50 eb e7 4c 5e 52 93 e2 6a 86 17 c2 4c
            80 ee 00 b8 bb 64 52 80 9a 18 46 06 f2 3e 5d e6 a2 62 57 cd 5c bd 65 8f 14 9c 74 ae 90 1c 38 17 2a c5 3b 13
            8f d0 00 9d fa 2f 3d e7 e9 f2 be a5 28 fc 3e 73 65 59 b3 c9 a7 2c 28 df b2 fa 01 83 4c 14 6b
44555->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 85 d6 01 00
            00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00
            78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd 74 e7 d9 5c 49 94 82 42 3f 50 ab f5 4e 5e d2 a4 b8 9a e1 85 30
            06 d0 1d 00 fb e4 44 0a 50 86 61 64 20 e7 62 59 0a e5 6d d7 cc d5 5b f6 40 de 4a e7 0a c9 81 73 a5 9c 9d 5d
            c2 e5 1b a0 53 ff a5 e7 bc 6d 3a 8f 18 84 e7 77 aa 2c 6b 5b 74 31 09 ca b7 ac 7e 70 9a 13 c6
59169->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 5d 64 00 00
            00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00
            78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd 74 e7 d9 5c 49 94 82 42 3f 50 ab f5 4e 5e d2 a4 b8 9a e1 85 30
            06 d0 1d 00 fb e4 44 0a 50 86 61 64 20 e7 62 59 0a e5 6d d7 cc d5 5b f6 40 de 4a e7 0a c9 81 73 a5 9c 9d 5d
            c2 e5 1b a0 53 ff a5 e7 bc 6d 3a 8f 18 84 e7 77 aa 2c 6b 5b 74 31 09 ca b7 ac 7e 70 9a 13 c6

10800->33233 0b 00 00 00 76 5c 00 00
            02 00 81 d1 c0 a8 01 0b 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 84 00 84 00 00 01 b4 00 00 00
            78 9c 55 8e d1 0a c2 30 0c 45 ef c4 e1 a6 3e f4 0f 43 d6 05 15 ba 56 9a e9 f0 d1 4f de 07 08 a6 65 3e f8 94 e4 e4 e4 12 07 74 3d 80
            87 0a 85 34 0c af 0e 68 e0 8c ee 8d de 7c 8a 8d 01 b7 59 99 e3 48 2a 32 ee 80 f6 b3 be ab 78 2e
            a2 d2 9d b3 c4 99 9e 5a 12 ea 45 6b 0b 9f 42 ca c5 3e da f0 8b 8d 3c 49 e9 b1 b1 83 d5 49 54 f9
            f2 87 4f 56 39 84 b4 d0 c2 b3 bf d6 dc f2 cd 17 c5 df 1d 2f
10800->42237 0b 00 00 00 f5 61 02 00
            02 00 a4 fd c0 a8 01 0b 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 85 00 85 00 00 01 b4 00 00 00
            78 9c 55 8e 4b 0e c2 30 0c 05 5f 11 15 e5 b3 c8 35 b8 0a 97 88 dc d4 02 a4 34 41 71 a1 e2 ea 5d d5 8e ca 82 95 ed f1 f8 c9 0e e8 8e
            00 de c2 3e e6 be ff 76 40 03 a7 74 af f4 19 72 6a 14 b8 cd 2a 94 06 2f cc c3 0e 68 6f d7 a5 8a
            17 13 c5 bf a8 70 9a fc 47 2c a1 5e b4 ba 08 39 e6 62 f6 49 87 5f 6c a2 91 ad c7 c6 0e 5a 47 16
            a1 fb 1f 3e 6b a5 18 f3 ec 67 9a c2 a3 e6 da 37 2b 4b c3 1c 2d
10800->60513 0c 00 00 00 00 00 00 00 00 00 28 00 28 00 00 01 22 00 00 00
            78 9c 13 60 60 e0 60 67 60 60 c8 4d 2d 2e 4e 4c 4f 15 80 72 cb 52 8b 8a 33 f3 f3 18 19 18 18 01 61 20 06 2c
10800->44555 0c 00 00 00 00 00 00 00 00 00 28 00 28 00 00 01 22 00 00 00
            78 9c 13 60 60 e0 60 67 60 60 c8 4d 2d 2e 4e 4c 4f 15 80 72 cb 52 8b 8a 33 f3 f3 18 19 18 18 01 61 20 06 2c
16 0 0 8 7 0 0 0 109 101 115 115 97 103 101 16 0 0 8 7 0 0 0 118 101 114 115 105 111 110 1 0 0 1
10800->59169 0b 00 00 00 5d 64 00 00
            02 00 e7 21 c0 a8 01 0b 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 83 00 83 00 00 01 b4 00 00 00
            78 9c 55 8e 51 0e c2 30 0c 43 3d b4 89 31 f8 e8 35 38 55 94 75 11 20 75 2d 6a c6 26 ce cb 45 48 ab f1 c1 97
            93 67 cb b2 03 fa 13 80 97 0a 85 34 8e ef 1e 68 e0 8c b6 46 1f 3e c5 c6 80 db 53 99 e3 44 2a 32 1d 80 ee 43
            d7 6a 5c 4a 50 e9 c9 59 e2 42 ab 96 86 6a 74 66 f8 14 52 2e e9 c1 9e 5f 6d e4 59 ca 8d 9d 1d 4d 67 51 e5 db
            1f 3e 9b 72 08 69 a3 8d 17 7f af bd 65 cd 17 4c b3 1c 2c


33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 5f 60 00 00
10800->33233 00 00 00 00 76 5c 00 00 5f 60 00 00
42237->10800 06 00 71 00 f5 61 02 00 03 00 00 00 03 00 00 00 de 65 02 00
10800->42237 00 00 00 00 f5 61 02 00 de 65 02 00

10800->33233 05 00 00 00 76 5c 00 00 2d 85 04 00
33233->10800 01 00 00 00 76 5c 00 00 2d 85 04 00
10800->42237 05 00 00 00 f5 61 02 00 be c8 00 00
42237->10800 01 00 00 00 f5 61 02 00 be c8 00 00

10800->33233 12 04 02 01 00 01 00 00 00   len==9 一直重复
10800->42237 12 04 02 01 00 01 00 00 00

33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 48 64 00 00
10800->33233 00 00 00 00 76 5c 00 00 48 64 00 00
42237->10800 06 00 71 00 f5 61 02 00 03 00 00 00 03 00 00 00 c8 69 02 00
10800->42237 00 00 00 00 f5 61 02 00 c8 69 02 00

10800->33233 12 04 02 01 00 01 00 00 00   len==9 一直重复

33233->10800 13 00 06 01 00 00 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00 01 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 01 00 00 00 01 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 02 00 00 00 01 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 03 00 00 00 01 00 00 00 重复一次
33233->10800 13 00 06 01 00 00 00 00 00 00 00 00 00 00 01 00 00 00 03 00 00 00 02 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 03 00 00 00 02 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 04 00 00 00 02 00 00 00 重复一次
33233->10800 13 00 06 01 00 00 00 00 00 00 00 00 00 00 01 00 00 00 04 00 00 00 03 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 04 00 00 00 03 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 05 00 00 00 03 00 00 00
33233->10800 13 00 06 01 00 00 00 00 00 00 00 00 00 00 01 00 00 00 05 00 00 00 04 00 00 00
......
33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 32 68 00 00
10800->33233 00 00 00 00 76 5c 00 00 32 68 00 00
......
10800->33233 05 00 00 00 76 5c 00 00 01 8d 04 00
33233->10800 01 00 00 00 76 5c 00 00 01 8d 04 00
10800->42237 05 00 00 00 f5 61 02 00 91 d0 00 00
42237->10800 01 00 00 00 f5 61 02 00 91 d0 00 00
......
33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 1b 6c 00 00
10800->33233 00 00 00 00 76 5c 00 00 1b 6c 00 00
42237->10800 06 00 71 00 f5 61 02 00 03 00 00 00 03 00 00 00 b2 6d 02 00
10800->42237 00 00 00 00 f5 61 02 00 b2 6d 02 00
......
10800->33233 05 00 00 00 76 5c 00 00 ec 90 04 00
33233->10800 01 00 00 00 76 5c 00 00 ec 90 04 00
......
33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 05 70 00 00
10800->33233 00 00 00 00 76 5c 00 00 05 70 00 00
......
10800->33233 05 00 00 00 76 5c 00 00 d6 94 04 00
33233->10800 01 00 00 00 76 5c 00 00 d6 94 04 00
......
33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 ef 73 00 00
10800->33233 00 00 00 00 76 5c 00 00 ef 73 00 00
......
10800->33233 05 00 00 00 76 5c 00 00 c2 98 04 00
33233->10800 01 00 00 00 76 5c 00 00 c2 98 04 00
......
33233->10800 06 00 71 00 76 5c 00 00 03 00 00 00 03 00 00 00 d8 77 00 00
10800->33233 00 00 00 00 76 5c 00 00 d8 77 00 00
......
10800->33233 05 00 00 00 76 5c 00 00 ad 9c 04 00
33233->10800 01 00 00 00 76 5c 00 00 ad 9c 04 00
......
10800->33233 12 04 02 01 00 02 00 00 00
10800->33233 12 04 02 01 00 02 00 00 00   len==9 一直重复
10800->42237 12 04 02 01 00 02 00 00 00


10800->33233 0f 00 00 00 76 5c 00 00 00 00 00 00 结束
10800->42237 0f 00 00 00 f5 61 02 00 00 00 00 00
59169->10800 10 f9 fc 0e 5d 64 00 00 00 00 00 00


不一样的数据
             13 00 06 01 10 00 00 00 00 00 00 00 10 00 01 00 00 00 91 00 00 00 90 00 00 00

10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 01 00 00 00 01 00 00 00
10800->33233 12 06 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 02 00 00 00 01 00 00 00
10800->33233 12 06 00 02 00 02 00 42 00 42 00 42 00 02 00 00 00 74 03 00 00 73 03 00 00
10800->33233 12 06 00 02 00 02 00 02 00 02 00 02 00 02 00 00 00 7a 03 00 00 78 03 00 00


*/

/*

华54074-10800对战

54074->10800 [固定 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00] [random id c6 73 01 00]
10800->54074 04
54074->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 c6 73 01 00 00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00 78 9c 2d cd 4b 0a 80 30
             0c 04 d0 a9 28 16 dd b8 f3 78 25 96 82 42 3f 50 eb e7 4c 5e 52 93 e2 6a 86 17 c2 4c 80 ee 00 b8 bb 64 52 80 9a 18 46 06 f2 3e 5d e6 a2 62 57 cd
             5c bd 65 8f 14 9c 74 ae 90 1c 38 17 2a c5 3b 13 8f d0 00 9d fa 2f 3d e7 e9 f2 be a5 28 fc 3e 73 65 59 b3 c9 a7 2c 28 df b2 fa 01 83 4c 14 6b
10800->54074 0b 00 00 00 [c6 73 01 00] 02 00 [d3 3a 54074d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 85 00 85 00 00 01
             b4 00 00 00 78 9c 55 8e 41 0e c2 30 0c 04 b7 a8 15 a5 70 c8 33 f8 55 e4 a6 16 20 a5 31 8a 0b 15 07 be c8 9b 70 a2 72 e0 64 7b 3c 5e d9 01 fd 01
             c0 43 d9 47 19 c7 57 0f 34 70 46 5b a3 b7 20 a9 31 e0 36 2b 53 9a bc 32 4f 3b a0 3b 7f de 55 3c 15 51 fd 9d 32 a7 c5 3f b5 24 d4 8b ce 16 41 a2
             e4 62 0f 36 fc 62 13 cd 5c 7a 6c 6c 6f 75 66 55 ba fc e1 a3 55 8a 51 56 bf d2 12 ae 35 b7 7c f3 05 4e ee 1c 33

49872也尝试加入对战

49872->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 [random id 4f b1 01 00]
10800->49872 07 00 4c 12 02 00 [d3 3a 54074d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
10800->54074 07 00 4c 12 02 00 [c2 d0 49872d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
49872->54074 07 01 00 00 02 00 [2a 30 10800d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->49872 07 01 01 00 02 00 [2a 30 10800d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
49872->10800 07 02 01 00 02 00 [d3 3a 54074d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->10800 07 02 01 00 02 00 [c2 d0 49872d] [0a 00 02 02  10.0.2.2    ] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00

49872->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 4f b1 01 00
10800->49872 07 00 4c 12 02 00 [d3 3a 54074d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
10800->54074 07 00 4c 12 02 00 [c2 d0 49872d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
49872->54074 07 01 01 00 02 00 [2a 30 10800d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
10800->49872 04
54074->49872 07 01 01 00 02 00 [2a 30 10800d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
49872->10800 07 02 01 00 02 00 [d3 3a 54074d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->10800 07 02 01 00 02 00 [c2 d0 49872d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
49872->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 4f b1 01 00 00 00 00 00 00 00 6b 00 6b 00 00 01 87 00 00 00 78 9c 2d cd 4b 0a 80 30
             0c 04 d0 a9 28 16 dd b8 f3 78 25 96 82 42 3f 50 eb e7 4c 5e 52 93 e2 6a 86 17 c2 4c 80 ee 00 b8 bb 64 52 80 9a 18 46 06 f2 3e 5d e6 a2 62 57 cd
             5c bd 65 8f 14 9c 74 ae 90 1c 38 17 2a c5 3b 13 8f d0 00 9d fa 2f 3d e7 e9 f2 be a5 28 fc 3e 73 65 59 b3 c9 a7 2c 28 df b2 fa 01 83 4c 14 6b
10800->49872 0c 00 00 00 00 00 00 00 00 00 28 00 28 00 00 01 1f 00 00 00 78 9c 13 60 60 e0 60 67 60 60 c8 4d 2d 2e 4e 4c 4f 15 00 72 59 80 dc a4 d2 e2 4a 46 06 06 46 00 4a a5 04 e6
line 4749

重新尝试加入观战 port 37129 (random id 改变了)

line 5614

37129->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 [random id 32 ce 01 00]
10800->37129 07 [00] [4c 12 02 00] [d3 3a 54074d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
10800->54074 07 [00] [4c 12 02 00] [91 09 37129d] [c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->54074 07 [01] [00 00 02 00] 2a 30 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->37129 07 [01] [01 00 02 00] 2a 30 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 07 [02] [01 00 02 00] d3 3a c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->10800 07 [02] [01 00 02 00] 91 09 0a 00 02 02 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00

37129->10800 08 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 32 ce 01 00
10800->37129 07 [00] [4c 12 02 00] d3 3a c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
10800->54074 07 [00] [4c 12 02 00] 91 09 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
10800->37129 04
37129->54074 07 [01] [01 00 02 00] 2a 30 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->37129 07 [01] [01 00 02 00] 2a 30 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 07 [02] [01 00 02 00] d3 3a c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
54074->10800 07 [02] [01 00 02 00] 91 09 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 09 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00 32 ce 01 00 00 00 00 00 00 00 71 00 71 00 00 01 9c 00 00 00 78 9c 45 cd 4b 0a 80 30
             0c 04 d0 54 2c 8a 6e ba f3 78 12 a5 a0 d0 0f d4 fa 39 93 97 d4 89 88 ae 66 78 49 88 21 aa 6b 22 9a 97 7e e7 3c 4e e8 8a 0c 50 03 ed 91 13 2b 88
             40 0b 60 e7 e2 fe 2f 3e 5e c2 03 7b 2b 1d f5 39 6e 90 03 e7 ec 6c 1f 56 5f 10 69 f5 4e 2a e4 66 d3 32 c7 20 7c 9d dd f7 6d 8c 2e 26 41 b9 96 af
             37 7b 90 17 e6
10800->37129 0b 01 00 00 32 ce 01 00 02 00 91 09 c0 a8 00 68 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 44 00 44 00 00 01
            49 00 00 00 78 9c 13 60 60 e0 e0 61 60 60 c8 2c 8e 2f 48 2c 4a cd 2b 89 2f 2b e6 60 60 60 64 14 00 4a b0 03 25 72 53 8b 8b 13 d3 53 41 5c 20 8f
            01 44 73 40 d4 97 27 96 24 67 80 d5 32 02 09 00 dc 8e 0b d6
07 打洞两次， 08->04->09->0b 才是整个请求过程

line 5670
10800->37129 05 00 00 00 32 ce 01 00 18 0c 02 00
37129->10800 01 01 00 00 32 ce 01 00 18 0c 02 00

line 5732
37129->10800 13 01 09 00 00 00 00 00 00 00 00 00 00 00
10800->37129 12 0a 02 00 00 02 00 00 00 00 01 00 00 ca 00 ca 00 40 02 00 00 78 9c 9d 52 6d 0e c2 20 0c 2d 46 8d 3a 4d f6 c7 43 78 0b 6f b2 e0 c6 16 22 0c 03
             c3 64 f7 f2 80 b6 73 1f cd 12 8d fa ab ed eb 7b 8f 52 48 01 36 5b 00 a8 a4 55 99 75 85 5a 00 ac 04 02 29 36 96 18 09 3b 63 8e 38 50 0f 66 51 b0
             1a 73 41 ba 84 74 32 34 ca 67 35 fa ce e5 44 a1 c4 2b 6d 23 3f 6f dd e9 bc 0e 72 70 da 21 12 8c bc ab b7 46 93 86 3b 8d f6 83 11 d1 2e 95 cd 74
             41 b4 d8 d3 68 d0 9b 91 ed 87 41 f9 3d 87 9a 4f 57 e8 b2 d4 79 34 4d cb 55 c9 38 77 ee 8c f3 df 2c 50 30 e3 fd b4 c0 bf f4 d4 0c 37 65 cc af 2f
             77 a0 35 c9 fc 5a 79 17 eb a2 df d6 89 7d 87 a0 54 87 3d 8e 2f d1 13 87 cb 3b 6f
line 5740
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [00 00 00 00 00 00 00 00 00d]
10800->37129 12 0b 02 00 00 02 00 00 00
             [00 00 00 00 00d] [18 00 00 00 24d] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             [00 00 00 00 00d] [18 00 00 00 24d] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
line 5751
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [18 00 00 00 18 00 00 00 24d]
10800->37129 12 0b 02 00 00 02 00 00 00
             [18 00 00 00 24d] [30 00 00 00 48d] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             [18 00 00 00 24d] [30 00 00 00 48d] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 06 [01 71 00] 32 ce 01 00 [03 00 00 00 03 00 00 00] 1c d2 01 00
10800->37129 00 00 00 00 32 ce 01 00 1c d2 01 00
line 5761
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 30 00 00 00 30 00 00 00
10800->37129 12 0b 02 00 00 02 00 00 00 [30 00 00 00] [48 00 00 00] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 [30 00 00 00] [48 00 00 00] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 48 00 00 00 48 00 00 00
10800->37129 12 0b 02 00 00 02 00 00 00 48 00 00 00 60 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 48 00 00 00 60 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 60 00 00 00 60 00 00 00
10800->37129 12 0b 02 00 00 02 00 00 00 60 00 00 00 78 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 60 00 00 00 78 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
line 5791
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 78 00 00 00 78 00 00 00
10800->37129 12 0b 02 00 00 02 00 00 00 78 00 00 00 90 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 78 00 00 00 90 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
             00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
...
line 5845
10800->37129 05 00 00 00 32 ce 01 00 51 11 02 00
37129->10800 01 01 00 00 32 ce 01 00 51 11 02 00
...
line 6644
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [50 07 00 00 50 07 00 00]
10800->37129 12 0b 02 00 00 02 00 00 00
             [50 07 00 00] [59 07 00 00] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   允许长度不一致
             [50 07 00 00] [58 07 00 00] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [59 07 00 00 58 07 00 00]
10800->37129 12 0b 02 00 00 02 00 00 00 [59 07 00 00] [5c 07 00 00] 00 00 00 00 00 00
                                        [58 07 00 00] [5c 07 00 00] 00 00 00 00 00 00 00 00
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [5c 07 00 00 5c 07 00 00]
line 6679
10800->37129 12 0b 02 00 00 02 00 00 00 [5c 07 00 00] [60 07 00 00] 00 00 00 00 00 00 00 00
                                        [5c 07 00 00] [5f 07 00 00] 00 00 00 00 00 00
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [60 07 00 00 5f 07 00 00]
10800->37129 12 0b 02 00 00 02 00 00 00 [60 07 00 00] [63 07 00 00] 00 00 00 00 00 00
                                        [5f 07 00 00] [62 07 00 00] 00 00 00 00 00 00
...
line 6748
37129->10800 13 01 09 00 00 00 02 00 00 00 00 00 00 00 [76 07 00 00 75 07 00 00]
10800->37129 12 0b 02 00 00 02 00 00 00 [76 07 00 00] [7a 07 00 00] 00 00 00 00 00 00 00 00
                                        [75 07 00 00] [79 07 00 00] 00 00 00 00 00 00 00 00

line 15993 结束


*/

/*

00 CLIENT_T_ACK 00 [00] [00 00] [random id 76 5c 00 00] [time 1b 6c 00 00 06c1bh] (为 06 删去中间部分)
01 HOST_T_ACK   01 [00] [00 00] [random id 76 5c 00 00] [time 01 8d 04 00 48d01h] (为 05 直接返回)
                01 [client count 01] [00 00] [random id 32 ce 01 00] 51 11 02 00
04 INIT_ACK     04
05 HOST_T       05 [00] [00 00] [random id 76 5c 00 00] [time 01 8d 04 00 48d01h]  host 发出
                05 [00] [00 00] [random id 32 ce 01 00] [time 51 11 02 00 06c1bh]  time 为从 th155.exe 启动开始(?)的毫秒数，在两端链接活跃时 05/06 包均隔 1s 发送一个，可能是测试延时之用
06 CLIENT_T     06 [00] [71 00] [random id 76 5c 00 00] [unknown 固定 03 00 00 00 03 00 00 00] [time 1b 6c 00 00] client 发出
07 PUNCH        07 [punch status 告诉需要打洞的对端 00] [4c 12 02 00] [port d3 3a 54074d] [ip c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
                07 [punch status 两个对端打洞       01] [00 00 02 00] 2a 30 [ip c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
                07 [punch status 回应打洞完成       02] [01 00 02 00] 91 09 [ip c0 a8 00 68 192.168.0.104] 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
08 INIT         08 [unknown 固定 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00] [random id 76 5c 00 00]
09 INIT_REQUEST 09 [unknown 固定 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00] [random id 76 5c 00 00] [对战 00 00 00 00 00 00 6b 00 6b 00 00 01] [87 00 00 00 len] [zlib data 78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd b8 f3 78 25 96 82 42 3f 50 eb e7 4c 5e 52 93 e2 6a 86 17 c2 4c 80 ee 00 b8 bb 64 52 80 9a 18 46 06 f2 3e 5d e6 a2 62 57 cd 5c bd 65 8f 14 9c 74 ae 90 1c 38 17 2a c5 3b 13 8f d0 00 9d fa 2f 3d e7 e9 f2 be a5 28 fc 3e 73 65 59 b3 c9 a7 2c 28 df b2 fa 01 83 4c 14 6b]
                                                                                                           decompress 16 0 0 8 5 0 0 0 [101 120 116 114 97] 1 0 0 1 16 0 0 8 11 0 0 0 [97 108 108 111 119 95 119 97 116 99 104] 8 0 0 1 1 16 0 0 8 4 0 0 0 [110 97 109 101] 16 0 0 8 0 0 0 0 16 0 0 8 10 0 0 0 [98 97 116 116 108 101 95 110 117 109] 2 0 0 5 1 0 0 0 16 0 0 8 7 0 0 0 [118 101 114 115 105 111 110] 2 0 0 5 255 170 23 0 16 0 0 8 5 0 0 0 [99 111 108 111 114] 2 0 0 5 10 0 0 0 1 0 0 1
                                                                                                                                        extra 1 0 0 1                                     allow_watch 8 0 0 1 1                                                   name                  []                       battle_num       2 0 0 5 1 0 0 0                                         version   2 0 0 5 255 170 23 0  (th155)                                color 2 0 0 5 10 0 0 0        1 0 0 1
                09 [unknown 固定 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00] [random id 85 d6 01 00] [对战 00 00 00 00 00 00 6b 00 6b 00 00 01] [87 00 00 00 len] [zlib data 78 9c 2d cd 4b 0a 80 30 0c 04 d0 a9 28 16 dd 74 e7 d9 5c 49 94 82 42 3f 50 ab f5 4e 5e d2 a4 b8 9a e1 85 30 06 d0 1d 00 fb e4 44 0a 50 86 61 64 20 e7 62 59 0a e5 6d d7 cc d5 5b f6 40 de 4a e7 0a c9 81 73 a5 9c 9d 5d c2 e5 1b a0 53 ff a5 e7 bc 6d 3a 8f 18 84 e7 77 aa 2c 6b 5b 74 31 09 ca b7 ac 7e 70 9a 13 c6]
                                                                                                           decompress 16 0 0 8 5 0 0 0 [101 120 116 114 97] 1 0 0 1 16 0 0 8 11 0 0 0 [97 108 108 111 119 95 119 97 116 99 104] 8 0 0 1 1 16 0 0 8 4 0 0 0 [110 97 109 101] 16 0 0 8 0 0 0 0 16 0 0 8 10 0 0 0 [98 97 116 116 108 101 95 110 117 109] 2 0 0 5 1 0 0 0 16 0 0 8 7 0 0 0 [118 101 114 115 105 111 110] 2 0 0 5 89 171 23 0  16 0 0 8 5 0 0 0 [99 111 108 111 114] 2 0 0 5 10 0 0 0 1 0 0 1
                                                                                                                                        extra 1 0 0 1                                     allow_watch 8 0 0 1 1                                                   name                  []                       battle_num       2 0 0 5 1 0 0 0                                         version   2 0 0 5 89 171 23 0  (th155_beta)                            color 2 0 0 5 10 0 0 0        1 0 0 1
                09 [unknown 固定 57 09 f6 67 f0 fd 4b d0 b9 9a 74 f8 38 33 81 88 00 00 00] [random id 32 ce 01 00] [观战 00 00 00 00 00 00 71 00 71 00 00 01] [9c 00 00 00 len] [zlib data 78 9c 45 cd 4b 0a 80 30 0c 04 d0 54 2c 8a 6e ba f3 78 12 a5 a0 d0 0f d4 fa 39 93 97 d4 89 88 ae 66 78 49 88 21 aa 6b 22 9a 97 7e e7 3c 4e e8 8a 0c 50 03 ed 91 13 2b 88 40 0b 60 e7 e2 fe 2f 3e 5e c2 03 7b 2b 1d f5 39 6e 90 03 e7 ec 6c 1f 56 5f 10 69 f5 4e 2a e4 66 d3 32 c7 20 7c 9d dd f7 6d 8c 2e 26 41 b9 96 af 37 7b 90 17 e6]
                                                                                                           decompress 16 0 0 8 8 0 0 0 [105 115 95 119 97 116 99 104] 8 0 0 1 0 16 0 0 8 5 0 0 0 [101 120 116 114 97] 1 0 0 1 16 0 0 8 11 0 0 0 [97 108 108 111 119 95 119 97 116 99 104] 8 0 0 1 1 16 0 0 8 4 0 0 0 [110 97 109 101] 16 0 0 8 0 0 0 0 16 0 0 8 10 0 0 0 [98 97 116 116 108 101 95 110 117 109] 2 0 0 5 1 0 0 0 16 0 0 8 7 0 0 0 [118 101 114 115 105 111 110] 2 0 0 5 255 170 23 0 16 0 0 8 5 0 0 0 [99 111 108 111 114] 2 0 0 5 10 0 0 0 1 0 0 1
                                                                                                                                         is_watch 8 0 0 1 0                                        extra 1 0 0 1                                  allow_watch 8 0 0 1 1                                                  name                     []                       battle_num   2 0 0 5 1 0 0 0                                             version 2 0 0 5 255 170 23 0   (th155)                                color            2 0 0 5 10 0 0 0          1 0 0 1
0b INIT_SUCCESS 0b [00 00 00] [random id 76 5c 00 00] 02 00 [self port 81 d1] [self ip c0 a8 01 0b]       [00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 84 00 84 00 00 01] [b4 00 00 00 dec len=180] [zlib data 78 9c 55 8e d1 0a c2 30 0c 45 ef c4 e1 a6 3e f4 0f 43 d6 05 15 ba 56 9a e9 f0 d1 4f de 07 08 a6 65 3e f8 94 e4 e4 e4 12 07 74 3d 80 87 0a 85 34 0c af 0e 68 e0 8c ee 8d de 7c 8a 8d 01 b7 59 99 e3 48 2a 32 ee 80 f6 b3 be ab 78 2e a2 d2 9d b3 c4 99 9e 5a 12 ea 45 6b 0b 9f 42 ca c5 3e da f0 8b 8d 3c 49 e9 b1 b1 83 d5 49 54 f9 f2 87 4f 56 39 84 b4 d0 c2 b3 bf d6 dc f2 cd 17 c5 df 1d 2f]
                                                                                                           decompress 16 0 0 8 9 0 0 0 [117 115 101 95 108 111 98 98 121] 8 0 0 1 0 16 0 0 8 4 0 0 0 [105 99 111 110] 1 0 0 1 16 0 0 8 9 0 0 0 [114 97 110 100 95 115 101 101 100] 2 0 0 5 253 241 129 0 16 0 0 8 12 0 0 0 [105 115 95 112 97 114 101 110 116 95 118 115] 8 0 0 1 1 16 0 0 8 5 0 0 0 [99 111 108 111 114] 2 0 0 5 10 0 0 0 16 0 0 8 4 0 0 0 [110 97 109 101] 16 0 0 8 0 0 0 0 16 0 0 8 7 0 0 0 [109 101 115 115 97 103 101] 16 0 0 8 0 0 0 0 16 0 0 8 11 0 0 0 [97 108 108 111 119 95 119 97 116 99 104] 8 0 0 1 1 1 0 0 1
                                                                                                                                              use_lobby  8 0 0 1 0                                            icon   1 0 0 1                            rand_seed    2 0 0 5 253 241 129 0                                               is_parent_vs    8 0 0 1 1                                            color   2 0 0 5 10 0 0 0                               name                    []                           message                            []                              allow_watch      8 0 0 1 1            1 0 0 1
                0b [00 00 00] [random id c6 73 01 00] 02 00 [d3 3a 54074d]    [c0 a8 00 68 192.168.0.104] [00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 85 00 85 00 00 01] [b4 00 00 00 dec len=180] [zlib data 78 9c 55 8e 41 0e c2 30 0c 04 b7 a8 15 a5 70 c8 33 f8 55 e4 a6 16 20 a5 31 8a 0b 15 07 be c8 9b 70 a2 72 e0 64 7b 3c 5e d9 01 fd 01 c0 43 d9 47 19 c7 57 0f 34 70 46 5b a3 b7 20 a9 31 e0 36 2b 53 9a bc 32 4f 3b a0 3b 7f de 55 3c 15 51 fd 9d 32 a7 c5 3f b5 24 d4 8b ce 16 41 a2 e4 62 0f 36 fc 62 13 cd 5c 7a 6c 6c 6f 75 66 55 ba fc e1 a3 55 8a 51 56 bf d2 12 ae 35 b7 7c f3 05 4e ee 1c 33]
                                                                                                           decompress 16 0 0 8 9 0 0 0 [117 115 101 95 108 111 98 98 121] 8 0 0 1 0 16 0 0 8 4 0 0 0 [105 99 111 110] 1 0 0 1 16 0 0 8 9 0 0 0 [114 97 110 100 95 115 101 101 100] 2 0 0 5 40 206 125 0 16 0 0 8 12 0 0 0  [105 115 95 112 97 114 101 110 116 95 118 115] 8 0 0 1 1 16 0 0 8 5 0 0 0 [99 111 108 111 114] 2 0 0 5 10 0 0 0 16 0 0 8 4 0 0 0 [110 97 109 101] 16 0 0 8 0 0 0 0 16 0 0 8 7 0 0 0 [109 101 115 115 97 103 101] 16 0 0 8 0 0 0 0 16 0 0 8 11 0 0 0 [97 108 108 111 119 95 119 97 116 99 104] 8 0 0 1 1 1 0 0 1
                                                                                                                                              use_lobby  8 0 0 1 0                                            icon   1 0 0 1                            rand_seed    2 0 0 5 40 206 125 0                                                is_parent_vs    8 0 0 1 1                                            color   2 0 0 5 10 0 0 0                               name                    []                           message                            []                              allow_watch      8 0 0 1 1            1 0 0 1
0c INIT_ERROR   0c [unknown 固定 00 00 00 00 00 00 00 00 00 28 00 28 00 00 01] [22 00 00 00 dec len=34] [zlib data 78 9c 13 60 60 e0 60 67 60 60 c8 4d 2d 2e 4e 4c 4f 15 80 72 cb 52 8b 8a 33 f3 f3 18 19 18 18 01 61 20 06 2c]
                                                                                            decompress 16 0 0 8 7 0 0 0 [109 101 115 115 97 103 101] 16 0 0 8 7 0 0 0 [118 101 114 115 105 111 110] 1 0 0 1
                                                                                                                           message                                         version    版本不匹配
                0c [unknown 固定 00 00 00 00 00 00 00 00 00 28 00 28 00 00 01] [1f 00 00 00 dec len=31] [zlib data 78 9c 13 60 60 e0 60 67 60 60 c8 4d 2d 2e 4e 4c 4f 15 00 72 59 80 dc a4 d2 e2 4a 46 06 06 46 00 4a a5 04 e6]
                                                                                            decompress 16 0 0 8 7 0 0 0 [109 101 115 115 97 103 101] 16 0 0 8 4 0 0 0 [98 117 115 121] 1 0 0 1
                                                                                                                           message                                         busy       已经在对战中
                12 [unknown 发生变化 0 0 0 0 0 0 0 0 0 41 0 41 0 0 1] [32 0 0 0] [120 156 19 96 96 224 96 103 96 96 200 77 45 46 78 76 79 21 0 114 89 129 220 162 212 196 148 74 70 6 6 70 0 81 7 5 57]
                                                                                            INIT_ERROR Unknown [0 0 0 0 0 0 0 0 0 41 0 41 0 0 1]message:[], ready:[]                  请求观战但是对战未开始
                 zlib 压缩数据，首先是一个小端序的解压后长度，然后是 0x78 0x9c 引导的 zlib 数据。解压后为一串数据，这串数据以 1 0 0 1 结尾；这串数据中每个数据以 16 0 0 8 开头，然后是小端序的数据名称长度 len ，然后是 len 长的 ascii 字串为数据名，然后是数据值（可能没有值）
0f HOST_QUIT    0f [00 00 00] [random id f5 61 02 00] 00 00 00 00
10 CLIENT_QUIT  10 [f9 fc 0e] [random id 5d 64 00 00] 00 00 00 00
12 HOST_GAME    12                   [06 00] [game input 02 00 02 00 02 00 02 00 02 00] [match id 02 00 00 00] [host frame 7a 03 00 00] [client frame 78 03 00 00]
                                      GAME_INPUT
                12                   [0a 02] 00 00 [match id 02 00 00 00] 00 01 00 00 ca 00 ca 00 [40 02 00 00 len] [zlib data 78 9c 9d 52 6d 0e c2 20 0c 2d 46 8d 3a 4d f6 c7 43 78 0b 6f b2 e0 c6 16 22 0c 03 c3 64 f7 f2 80 b6 73 1f cd 12 8d fa ab ed eb 7b 8f 52 48 01 36 5b 00 a8 a4 55 99 75 85 5a 00 ac 04 02 29 36 96 18 09 3b 63 8e 38 50 0f 66 51 b0 1a 73 41 ba 84 74 32 34 ca 67 35 fa ce e5 44 a1 c4 2b 6d 23 3f 6f dd e9 bc 0e 72 70 da 21 12 8c bc ab b7 46 93 86 3b 8d f6 83 11 d1 2e 95 cd 74 41 b4 d8 d3 68 d0 9b 91 ed 87 41 f9 3d 87 9a 4f 57 e8 b2 d4 79 34 4d cb 55 c9 38 77 ee 8c f3 df 2c 50 30 e3 fd b4 c0 bf f4 d4 0c 37 65 cc af 2f 77 a0 35 c9 fc 5a 79 17 eb a2 df d6 89 7d 87 a0 54 87 3d 8e 2f d1 13 87 cb 3b 6f]
                                      GAME_REPLAY_MATCH
                        game_mode:[2 0 0 5 1 0 0 0], mode:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 1 0 0 0 2 0 0 5 0 0 0 0 1 0 0 1], master_name:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0], reimu:[2 0 0 5 1 0 0 0], marisa:[1 0 0 1], slave_name:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0], marisa:[2 0 0 5 1 0 0 0], reimu:[1 0 0 1], bgm_id:[2 0 0 5 117 0 0 0], player_name:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0], :[2 0 0 5 1 0 0 0], :[1 0 0 1], difficulty:[2 0 0 5 0 0 0 0], slave_color:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 1 0 0 0 2 0 0 5 1 0 0 0 1 0 0 1], master_color:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 1 0 0 0 2 0 0 5 1 0 0 0 1 0 0 1], spell:[64 0 0 8 2 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 0 0 0 0 2 0 0 5 1 0 0 0 2 0 0 5 0 0 0 0 1 0 0 1], background_id:[2 0 0 5 41 0 0 0], seed:[2 0 0 5 180 21 0 0]
                12                   [0b 02] 00 00 [match id 02 00 00 00] [frame id start 5c 07 00 00 1884d] [frame id end 60 07 00 00 1888d] [frame input data 00 00 00 00 00 00 00 00 length=4 (max=24)] [frame id start 5c 07 00 00 1884d] [frame id end 5f 07 00 00 1887d] [frame input data 00 00 00 00 00 00 length=3]
                                      GAME_REPLAY_DATA
                12                   [0c 00] 00 00 [match id 02 00 00 00]
                                      GAME_REPLAY_END
                12                   [04 02] 01 00 [match id 01 00 00 00]
                                      GAME_SELECT
                    12 由 host 发出
13 CLIENT_GAME  13 [client count 00] [06 01] [game input 10 00 00 00 00 00 00 00 10 00] [match id 01 00 00 00] [host frame 91 00 00 00] [client frame 90 00 00 00]
                                      GAME_INPUT
                13 [client count 01] [09 00] 00 00 [match id 02 00 00 00] 00 00 00 00 [frame id 59 07 00 00] [frame id 58 07 00 00]
                                      GAME_REPLAY_REQUEST
                    13 由 client 发出， client count 为 client 在这个 host 下的编号，从 0 开始

连接过程为 08->04->09->0b/0c
如果已经在对战中，则第三端在连接 host 时会和 client 07 打洞两次
match id 和则不一样，华一个 Round 计为一次 Match

每次观战握手成功后 初始 match id 和 frame id 为 0
观战方 13 01 09 00 00 00 00 00 00 00 00 00 00 00
未开始时 host 不返回数据
开始后发送 12 0a ，其中包含 match id ，此时对战开始
观战方继续发送 match id 为此次对战的 13 01 09 ， host 返回观战数据 12 0b

本次对战结束时 host 发送 12 0c 作为 13 01 09 的回应
观战方知晓已经结束，开始发送 13 01 09 00 00 00 00 00 00 00 00 00 00 00
host 会回应 12 0a ，但是其中的 match id 并不是新对局的，所以观战方不为所动
当新的对局开始时，回应的 12 0a 包含新 match id ，观战方开始获取新一次对战的 replay

注意尽管每个对战有 2-3 个 match ， 一次对战中的 match id 是一直不变的，会在下一次对战跳变

*/
