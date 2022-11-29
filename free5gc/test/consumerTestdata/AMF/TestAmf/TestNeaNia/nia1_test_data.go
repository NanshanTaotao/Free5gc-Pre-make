package TestNeaNia

type NIA1TestSet struct {
	CountI    uint32
	Bearer    byte
	Direction uint32
	Ik        [16]byte
	Length    uint64
	Msg       []byte
	MacI      uint32
}

const (
	NIA1Test1 = "NIA1Test1"
	NIA1Test2 = "NIA1Test2"
	NIA1Test3 = "NIA1Test3"
	NIA1Test4 = "NIA1Test4"
	NIA1Test5 = "NIA1Test5"
	NIA1Test6 = "NIA1Test6"
	NIA1Test7 = "NIA1Test7"
)

var TestNIA1Table = make(map[string]*NIA1TestSet)

func init() {
	TestNIA1Table[NIA1Test1] = &NIA1TestSet{
		CountI:    0x38a6f056,
		Bearer:    0x1f,
		Direction: 0,
		Ik:        [16]byte{0x2b, 0xd6, 0x45, 0x9f, 0x82, 0xc5, 0xb3, 0x00, 0x95, 0x2c, 0x49, 0x10, 0x48, 0x81, 0xff, 0x48},
		Length:    88,
		Msg: []byte{
			0x33, 0x32, 0x34, 0x62, 0x63, 0x39, 0x38, 0x61, 0x37, 0x34, 0x79, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
		MacI: 0x731f1165,
	}

	TestNIA1Table[NIA1Test2] = &NIA1TestSet{
		CountI:    0x36af6144,
		Bearer:    0x18,
		Direction: 1,
		Ik:        [16]byte{0x7e, 0x5e, 0x94, 0x43, 0x1e, 0x11, 0xd7, 0x38, 0x28, 0xd7, 0x39, 0xcc, 0x6c, 0xed, 0x45, 0x73},
		Length:    254,
		Msg: []byte{
			0xb3, 0xd3, 0xc9, 0x17, 0x0a, 0x4e, 0x16, 0x32, 0xf6, 0x0f, 0x86, 0x10, 0x13, 0xd2, 0x2d, 0x84,
			0xb7, 0x26, 0xb6, 0xa2, 0x78, 0xd8, 0x02, 0xd1, 0xee, 0xaf, 0x13, 0x21, 0xba, 0x59, 0x29, 0xdc,
		},
		MacI: 0xe3259f6f,
	}

	TestNIA1Table[NIA1Test3] = &NIA1TestSet{
		CountI:    0xc7590ea9,
		Bearer:    0x17,
		Direction: 0,
		Ik:        [16]byte{0xd3, 0x41, 0x9b, 0xe8, 0x21, 0x08, 0x7a, 0xcd, 0x02, 0x12, 0x3a, 0x92, 0x48, 0x03, 0x33, 0x59},
		Length:    511,
		Msg: []byte{
			0xbb, 0xb0, 0x57, 0x03, 0x88, 0x09, 0x49, 0x6b, 0xcf, 0xf8, 0x6d, 0x6f, 0xbc, 0x8c, 0xe5, 0xb1,
			0x35, 0xa0, 0x6b, 0x16, 0x60, 0x54, 0xf2, 0xd5, 0x65, 0xbe, 0x8a, 0xce, 0x75, 0xdc, 0x85, 0x1e,
			0x0b, 0xcd, 0xd8, 0xf0, 0x71, 0x41, 0xc4, 0x95, 0x87, 0x2f, 0xb5, 0xd8, 0xc0, 0xc6, 0x6a, 0x8b,
			0x6d, 0xa5, 0x56, 0x66, 0x3e, 0x4e, 0x46, 0x12, 0x05, 0xd8, 0x45, 0x80, 0xbe, 0xe5, 0xbc, 0x7e,
		},
		MacI: 0x9a16c77d,
	}

	TestNIA1Table[NIA1Test4] = &NIA1TestSet{
		CountI:    0x36af6144,
		Bearer:    0x0f,
		Direction: 1,
		Ik:        [16]byte{0x83, 0xfd, 0x23, 0xa2, 0x44, 0xa7, 0x4c, 0xf3, 0x58, 0xda, 0x30, 0x19, 0xf1, 0x72, 0x26, 0x35},
		Length:    768,
		Msg: []byte{
			0x35, 0xc6, 0x87, 0x16, 0x63, 0x3c, 0x66, 0xfb, 0x75, 0x0c, 0x26, 0x68, 0x65, 0xd5, 0x3c, 0x11,
			0xea, 0x05, 0xb1, 0xe9, 0xfa, 0x49, 0xc8, 0x39, 0x8d, 0x48, 0xe1, 0xef, 0xa5, 0x90, 0x9d, 0x39,
			0x47, 0x90, 0x28, 0x37, 0xf5, 0xae, 0x96, 0xd5, 0xa0, 0x5b, 0xc8, 0xd6, 0x1c, 0xa8, 0xdb, 0xef,
			0x1b, 0x13, 0xa4, 0xb4, 0xab, 0xfe, 0x4f, 0xb1, 0x00, 0x60, 0x45, 0xb6, 0x74, 0xbb, 0x54, 0x72,
			0x93, 0x04, 0xc3, 0x82, 0xbe, 0x53, 0xa5, 0xaf, 0x05, 0x55, 0x61, 0x76, 0xf6, 0xea, 0xa2, 0xef,
			0x1d, 0x05, 0xe4, 0xb0, 0x83, 0x18, 0x1e, 0xe6, 0x74, 0xcd, 0xa5, 0xa4, 0x85, 0xf7, 0x4d, 0x7a,
		},
		MacI: 0xbba74492,
	}

	TestNIA1Table[NIA1Test5] = &NIA1TestSet{
		CountI:    0x36af6144,
		Bearer:    0x18,
		Direction: 0,
		Ik:        [16]byte{0x68, 0x32, 0xa6, 0x5c, 0xff, 0x44, 0x73, 0x62, 0x1e, 0xbd, 0xd4, 0xba, 0x26, 0xa9, 0x21, 0xfe},
		Length:    383,
		Msg: []byte{
			0xd3, 0xc5, 0x38, 0x39, 0x62, 0x68, 0x20, 0x71, 0x77, 0x65, 0x66, 0x76, 0x20, 0x32, 0x38, 0x37,
			0x63, 0x62, 0x40, 0x98, 0x1b, 0xa6, 0x82, 0x4c, 0x1b, 0xfb, 0x1a, 0xb4, 0x85, 0x47, 0x20, 0x29,
			0xb7, 0x1d, 0x80, 0x8c, 0xe3, 0x3e, 0x2c, 0xc3, 0xc0, 0xb5, 0xfc, 0x1f, 0x3d, 0xe8, 0xa6, 0xdc,
		},
		MacI: 0x4145e4b0,
	}

	TestNIA1Table[NIA1Test6] = &NIA1TestSet{
		CountI:    0x7827fab2,
		Bearer:    0x05,
		Direction: 1,
		Ik:        [16]byte{0x5d, 0x0a, 0x80, 0xd8, 0x13, 0x4a, 0xe1, 0x96, 0x77, 0x82, 0x4b, 0x67, 0x1e, 0x83, 0x8a, 0xf4},
		Length:    2558,
		Msg: []byte{
			0x70, 0xde, 0xdf, 0x2d, 0xc4, 0x2c, 0x5c, 0xbd, 0x3a, 0x96, 0xf8, 0xa0, 0xb1, 0x14, 0x18, 0xb3,
			0x60, 0x8d, 0x57, 0x33, 0x60, 0x4a, 0x2c, 0xd3, 0x6a, 0xab, 0xc7, 0x0c, 0xe3, 0x19, 0x3b, 0xb5,
			0x15, 0x3b, 0xe2, 0xd3, 0xc0, 0x6d, 0xfd, 0xb2, 0xd1, 0x6e, 0x9c, 0x35, 0x71, 0x58, 0xbe, 0x6a,
			0x41, 0xd6, 0xb8, 0x61, 0xe4, 0x91, 0xdb, 0x3f, 0xbf, 0xeb, 0x51, 0x8e, 0xfc, 0xf0, 0x48, 0xd7,
			0xd5, 0x89, 0x53, 0x73, 0x0f, 0xf3, 0x0c, 0x9e, 0xc4, 0x70, 0xff, 0xcd, 0x66, 0x3d, 0xc3, 0x42,
			0x01, 0xc3, 0x6a, 0xdd, 0xc0, 0x11, 0x1c, 0x35, 0xb3, 0x8a, 0xfe, 0xe7, 0xcf, 0xdb, 0x58, 0x2e,
			0x37, 0x31, 0xf8, 0xb4, 0xba, 0xa8, 0xd1, 0xa8, 0x9c, 0x06, 0xe8, 0x11, 0x99, 0xa9, 0x71, 0x62,
			0x27, 0xbe, 0x34, 0x4e, 0xfc, 0xb4, 0x36, 0xdd, 0xd0, 0xf0, 0x96, 0xc0, 0x64, 0xc3, 0xb5, 0xe2,
			0xc3, 0x99, 0x99, 0x3f, 0xc7, 0x73, 0x94, 0xf9, 0xe0, 0x97, 0x20, 0xa8, 0x11, 0x85, 0x0e, 0xf2,
			0x3b, 0x2e, 0xe0, 0x5d, 0x9e, 0x61, 0x73, 0x60, 0x9d, 0x86, 0xe1, 0xc0, 0xc1, 0x8e, 0xa5, 0x1a,
			0x01, 0x2a, 0x00, 0xbb, 0x41, 0x3b, 0x9c, 0xb8, 0x18, 0x8a, 0x70, 0x3c, 0xd6, 0xba, 0xe3, 0x1c,
			0xc6, 0x7b, 0x34, 0xb1, 0xb0, 0x00, 0x19, 0xe6, 0xa2, 0xb2, 0xa6, 0x90, 0xf0, 0x26, 0x71, 0xfe,
			0x7c, 0x9e, 0xf8, 0xde, 0xc0, 0x09, 0x4e, 0x53, 0x37, 0x63, 0x47, 0x8d, 0x58, 0xd2, 0xc5, 0xf5,
			0xb8, 0x27, 0xa0, 0x14, 0x8c, 0x59, 0x48, 0xa9, 0x69, 0x31, 0xac, 0xf8, 0x4f, 0x46, 0x5a, 0x64,
			0xe6, 0x2c, 0xe7, 0x40, 0x07, 0xe9, 0x91, 0xe3, 0x7e, 0xa8, 0x23, 0xfa, 0x0f, 0xb2, 0x19, 0x23,
			0xb7, 0x99, 0x05, 0xb7, 0x33, 0xb6, 0x31, 0xe6, 0xc7, 0xd6, 0x86, 0x0a, 0x38, 0x31, 0xac, 0x35,
			0x1a, 0x9c, 0x73, 0x0c, 0x52, 0xff, 0x72, 0xd9, 0xd3, 0x08, 0xee, 0xdb, 0xab, 0x21, 0xfd, 0xe1,
			0x43, 0xa0, 0xea, 0x17, 0xe2, 0x3e, 0xdc, 0x1f, 0x74, 0xcb, 0xb3, 0x63, 0x8a, 0x20, 0x33, 0xaa,
			0xa1, 0x54, 0x64, 0xea, 0xa7, 0x33, 0x38, 0x5d, 0xbb, 0xeb, 0x6f, 0xd7, 0x35, 0x09, 0xb8, 0x57,
			0xe6, 0xa4, 0x19, 0xdc, 0xa1, 0xd8, 0x90, 0x7a, 0xf9, 0x77, 0xfb, 0xac, 0x4d, 0xfa, 0x35, 0xec,
		},
		MacI: 0x0fa2b1ee,
	}

	TestNIA1Table[NIA1Test7] = &NIA1TestSet{
		CountI:    0x296f393c,
		Bearer:    0x0b,
		Direction: 1,
		Ik:        [16]byte{0xb3, 0x12, 0x0f, 0xfd, 0xb2, 0xcf, 0x6a, 0xf4, 0xe7, 0x3e, 0xaf, 0x2e, 0xf4, 0xeb, 0xec, 0x69},
		Length:    16448,
		Msg: []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
			0xe0, 0x95, 0x80, 0x45, 0xf3, 0xa0, 0xbb, 0xa4, 0xe3, 0x96, 0x83, 0x46, 0xf0, 0xa3, 0xb8, 0xa7,
			0xc0, 0x2a, 0x01, 0x8a, 0xe6, 0x40, 0x76, 0x52, 0x26, 0xb9, 0x87, 0xc9, 0x13, 0xe6, 0xcb, 0xf0,
			0x83, 0x57, 0x00, 0x16, 0xcf, 0x83, 0xef, 0xbc, 0x61, 0xc0, 0x82, 0x51, 0x3e, 0x21, 0x56, 0x1a,
			0x42, 0x7c, 0x00, 0x9d, 0x28, 0xc2, 0x98, 0xef, 0xac, 0xe7, 0x8e, 0xd6, 0xd5, 0x6c, 0x2d, 0x45,
			0x05, 0xad, 0x03, 0x2e, 0x9c, 0x04, 0xdc, 0x60, 0xe7, 0x3a, 0x81, 0x69, 0x6d, 0xa6, 0x65, 0xc6,
			0xc4, 0x86, 0x03, 0xa5, 0x7b, 0x45, 0xab, 0x33, 0x22, 0x15, 0x85, 0xe6, 0x8e, 0xe3, 0x16, 0x91,
			0x87, 0xfb, 0x02, 0x39, 0x52, 0x86, 0x32, 0xdd, 0x65, 0x6c, 0x80, 0x7e, 0xa3, 0x24, 0x8b, 0x7b,
			0x46, 0xd0, 0x02, 0xb2, 0xb5, 0xc7, 0x45, 0x8e, 0xb8, 0x5b, 0x9c, 0xe9, 0x58, 0x79, 0xe0, 0x34,
			0x08, 0x59, 0x05, 0x5e, 0x3b, 0x0a, 0xbb, 0xc3, 0xea, 0xce, 0x87, 0x19, 0xca, 0xa8, 0x02, 0x65,
			0xc9, 0x72, 0x05, 0xd5, 0xdc, 0x4b, 0xcc, 0x90, 0x2f, 0xe1, 0x83, 0x96, 0x29, 0xed, 0x71, 0x32,
			0x8a, 0x0f, 0x04, 0x49, 0xf5, 0x88, 0x55, 0x7e, 0x68, 0x98, 0x86, 0x0e, 0x04, 0x2a, 0xec, 0xd8,
			0x4b, 0x24, 0x04, 0xc2, 0x12, 0xc9, 0x22, 0x2d, 0xa5, 0xbf, 0x8a, 0x89, 0xef, 0x67, 0x97, 0x87,
			0x0c, 0xf5, 0x07, 0x71, 0xa6, 0x0f, 0x66, 0xa2, 0xee, 0x62, 0x85, 0x36, 0x57, 0xad, 0xdf, 0x04,
			0xcd, 0xde, 0x07, 0xfa, 0x41, 0x4e, 0x11, 0xf1, 0x2b, 0x4d, 0x81, 0xb9, 0xb4, 0xe8, 0xac, 0x53,
			0x8e, 0xa3, 0x06, 0x66, 0x68, 0x8d, 0x88, 0x1f, 0x6c, 0x34, 0x84, 0x21, 0x99, 0x2f, 0x31, 0xb9,
			0x4f, 0x88, 0x06, 0xed, 0x8f, 0xcc, 0xff, 0x4c, 0x91, 0x23, 0xb8, 0x96, 0x42, 0x52, 0x7a, 0xd6,
			0x13, 0xb1, 0x09, 0xbf, 0x75, 0x16, 0x74, 0x85, 0xf1, 0x26, 0x8b, 0xf8, 0x84, 0xb4, 0xcd, 0x23,
			0xd2, 0x9a, 0x09, 0x34, 0x92, 0x57, 0x03, 0xd6, 0x34, 0x09, 0x8f, 0x77, 0x67, 0xf1, 0xbe, 0x74,
			0x91, 0xe7, 0x08, 0xa8, 0xbb, 0x94, 0x9a, 0x38, 0x73, 0x70, 0x8a, 0xef, 0x4a, 0x36, 0x23, 0x9e,
			0x50, 0xcc, 0x08, 0x23, 0x5c, 0xd5, 0xed, 0x6b, 0xbe, 0x57, 0x86, 0x68, 0xa1, 0x7b, 0x58, 0xc1,
			0x17, 0x1d, 0x0b, 0x90, 0xe8, 0x13, 0xa9, 0xe4, 0xf5, 0x8a, 0x89, 0xd7, 0x19, 0xb1, 0x10, 0x42,
			0xd6, 0x36, 0x0b, 0x1b, 0x0f, 0x52, 0xde, 0xb7, 0x30, 0xa5, 0x8d, 0x58, 0xfa, 0xf4, 0x63, 0x15,
			0x95, 0x4b, 0x0a, 0x87, 0x26, 0x91, 0x47, 0x59, 0x77, 0xdc, 0x88, 0xc0, 0xd7, 0x33, 0xfe, 0xff,
			0x54, 0x60, 0x0a, 0x0c, 0xc1, 0xd0, 0x30, 0x0a, 0xaa, 0xeb, 0x94, 0x57, 0x2c, 0x6e, 0x95, 0xb0,
			0x1a, 0xe9, 0x0d, 0xe0, 0x4f, 0x1d, 0xce, 0x47, 0xf8, 0x7e, 0x8f, 0xa7, 0xbe, 0xbf, 0x77, 0xe1,
			0xdb, 0xc2, 0x0d, 0x6b, 0xa8, 0x5c, 0xb9, 0x14, 0x3d, 0x51, 0x8b, 0x28, 0x5d, 0xfa, 0x04, 0xb6,
			0x98, 0xbf, 0x0c, 0xf7, 0x81, 0x9f, 0x20, 0xfa, 0x7a, 0x28, 0x8e, 0xb0, 0x70, 0x3d, 0x99, 0x5c,
			0x59, 0x94, 0x0c, 0x7c, 0x66, 0xde, 0x57, 0xa9, 0xb7, 0x0f, 0x82, 0x37, 0x9b, 0x70, 0xe2, 0x03,
			0x1e, 0x45, 0x0f, 0xcf, 0xd2, 0x18, 0x13, 0x26, 0xfc, 0xd2, 0x8d, 0x88, 0x23, 0xba, 0xaa, 0x80,
			0xdf, 0x6e, 0x0f, 0x44, 0x35, 0x59, 0x64, 0x75, 0x39, 0xfd, 0x89, 0x07, 0xc0, 0xff, 0xd9, 0xd7,
			0x9c, 0x13, 0x0e, 0xd8, 0x1c, 0x9a, 0xfd, 0x9b, 0x7e, 0x84, 0x8c, 0x9f, 0xed, 0x38, 0x44, 0x3d,
			0x5d, 0x38, 0x0e, 0x53, 0xfb, 0xdb, 0x8a, 0xc8, 0xc3, 0xd3, 0xf0, 0x68, 0x76, 0x05, 0x4f, 0x12,
			0x24, 0x61, 0x10, 0x7d, 0xe9, 0x2f, 0xea, 0x09, 0xc6, 0xf6, 0x92, 0x3a, 0x18, 0x8d, 0x53, 0xaf,
			0xe5, 0x4a, 0x10, 0xf6, 0x0e, 0x6e, 0x9d, 0x5a, 0x03, 0xd9, 0x96, 0xb5, 0xfb, 0xc8, 0x20, 0xf8,
			0xa6, 0x37, 0x11, 0x6a, 0x27, 0xad, 0x04, 0xb4, 0x44, 0xa0, 0x93, 0x2d, 0xd6, 0x0f, 0xbd, 0x12,
			0x67, 0x1c, 0x11, 0xe1, 0xc0, 0xec, 0x73, 0xe7, 0x89, 0x87, 0x9f, 0xaa, 0x3d, 0x42, 0xc6, 0x4d,
			0x20, 0xcd, 0x12, 0x52, 0x74, 0x2a, 0x37, 0x68, 0xc2, 0x5a, 0x90, 0x15, 0x85, 0x88, 0x8e, 0xce,
			0xe1, 0xe6, 0x12, 0xd9, 0x93, 0x6b, 0x40, 0x3b, 0x07, 0x75, 0x94, 0x9a, 0x66, 0xcd, 0xfd, 0x99,
			0xa2, 0x9b, 0x13, 0x45, 0xba, 0xa8, 0xd9, 0xd5, 0x40, 0x0c, 0x91, 0x02, 0x4b, 0x0a, 0x60, 0x73,
			0x63, 0xb0, 0x13, 0xce, 0x5d, 0xe9, 0xae, 0x86, 0x9d, 0x3b, 0x8d, 0x95, 0xb0, 0x57, 0x0b, 0x3c,
			0x2d, 0x39, 0x14, 0x22, 0xd3, 0x24, 0x50, 0xcb, 0xcf, 0xae, 0x96, 0x65, 0x22, 0x86, 0xe9, 0x6d,
			0xec, 0x12, 0x14, 0xa9, 0x34, 0x65, 0x27, 0x98, 0x0a, 0x81, 0x92, 0xea, 0xc1, 0xc3, 0x9a, 0x3a,
			0xaf, 0x6f, 0x15, 0x35, 0x1d, 0xa6, 0xbe, 0x76, 0x4d, 0xf8, 0x97, 0x72, 0xec, 0x04, 0x07, 0xd0,
			0x6e, 0x44, 0x15, 0xbe, 0xfa, 0xe7, 0xc9, 0x25, 0x80, 0xdf, 0x9b, 0xf5, 0x07, 0x49, 0x7c, 0x8f,
			0x29, 0x95, 0x16, 0x0d, 0x4e, 0x21, 0x8d, 0xaa, 0xcb, 0x02, 0x94, 0x4a, 0xbf, 0x83, 0x34, 0x0c,
			0xe8, 0xbe, 0x16, 0x86, 0xa9, 0x60, 0xfa, 0xf9, 0x0e, 0x2d, 0x90, 0xc5, 0x5c, 0xc6, 0x47, 0x5b,
			0xab, 0xc3, 0x17, 0x1a, 0x80, 0xa3, 0x63, 0x17, 0x49, 0x54, 0x95, 0x5d, 0x71, 0x01, 0xda, 0xb1,
			0x6a, 0xe8, 0x17, 0x91, 0x67, 0xe2, 0x14, 0x44, 0xb4, 0x43, 0xa9, 0xea, 0xaa, 0x7c, 0x91, 0xde,
			0x36, 0xd1, 0x18, 0xc3, 0x9d, 0x38, 0x9f, 0x8d, 0xd4, 0x46, 0x9a, 0x84, 0x6c, 0x9a, 0x26, 0x2b,
			0xf7, 0xfa, 0x18, 0x48, 0x7a, 0x79, 0xe8, 0xde, 0x11, 0x69, 0x9e, 0x0b, 0x8f, 0xdf, 0x55, 0x7c,
			0xb4, 0x87, 0x19, 0xd4, 0x53, 0xba, 0x71, 0x30, 0x56, 0x10, 0x9b, 0x93, 0xa2, 0x18, 0xc8, 0x96,
			0x75, 0xac, 0x19, 0x5f, 0xb4, 0xfb, 0x06, 0x63, 0x9b, 0x37, 0x97, 0x14, 0x49, 0x55, 0xb3, 0xc9,
			0x32, 0x7d, 0x1a, 0xec, 0x00, 0x3d, 0x42, 0xec, 0xd0, 0xea, 0x98, 0xab, 0xf1, 0x9f, 0xfb, 0x4a,
			0xf3, 0x56, 0x1a, 0x67, 0xe7, 0x7c, 0x35, 0xbf, 0x15, 0xc5, 0x9c, 0x24, 0x12, 0xda, 0x88, 0x1d,
			0xb0, 0x2b, 0x1b, 0xfb, 0xce, 0xbf, 0xac, 0x51, 0x52, 0xbc, 0x99, 0xbc, 0x3f, 0x1d, 0x15, 0xf7,
			0x71, 0x00, 0x1b, 0x70, 0x29, 0xfe, 0xdb, 0x02, 0x8f, 0x8b, 0x85, 0x2b, 0xc4, 0x40, 0x7e, 0xb8,
			0x3f, 0x89, 0x1c, 0x9c, 0xa7, 0x33, 0x25, 0x4f, 0xdd, 0x1e, 0x9e, 0xdb, 0x56, 0x91, 0x9c, 0xe9,
			0xfe, 0xa2, 0x1c, 0x17, 0x40, 0x72, 0x52, 0x1c, 0x18, 0x31, 0x9a, 0x54, 0xb5, 0xd4, 0xef, 0xbe,
			0xbd, 0xdf, 0x1d, 0x8b, 0x69, 0xb1, 0xcb, 0xf2, 0x5f, 0x48, 0x9f, 0xcc, 0x98, 0x13, 0x72, 0x54,
			0x7c, 0xf4, 0x1d, 0x00, 0x8e, 0xf0, 0xbc, 0xa1, 0x92, 0x6f, 0x93, 0x4b, 0x73, 0x5e, 0x09, 0x0b,
			0x3b, 0x25, 0x1e, 0xb3, 0x3a, 0x36, 0xf8, 0x2e, 0xd9, 0xb2, 0x9c, 0xf4, 0xcb, 0x94, 0x41, 0x88,
			0xfa, 0x0e, 0x1e, 0x38, 0xdd, 0x77, 0x8f, 0x7d, 0x1c, 0x9d, 0x98, 0x7b, 0x28, 0xd1, 0x32, 0xdf,
			0xb9, 0x73, 0x1f, 0xa4, 0xf4, 0xb4, 0x16, 0x93, 0x5b, 0xe4, 0x9d, 0xe3, 0x05, 0x16, 0xaf, 0x35,
			0x78, 0x58, 0x1f, 0x2f, 0x13, 0xf5, 0x61, 0xc0, 0x66, 0x33, 0x61, 0x94, 0x1e, 0xab, 0x24, 0x9a,
			0x4b, 0xc1, 0x23, 0xf8, 0xd1, 0x5c, 0xd7, 0x11, 0xa9, 0x56, 0xa1, 0xbf, 0x20, 0xfe, 0x6e, 0xb7,
			0x8a, 0xea, 0x23, 0x73, 0x36, 0x1d, 0xa0, 0x42, 0x6c, 0x79, 0xa5, 0x30, 0xc3, 0xbb, 0x1d, 0xe0,
			0xc9, 0x97, 0x22, 0xef, 0x1f, 0xde, 0x39, 0xac, 0x2b, 0x00, 0xa0, 0xa8, 0xee, 0x7c, 0x80, 0x0a,
			0x08, 0xbc, 0x22, 0x64, 0xf8, 0x9f, 0x4e, 0xff, 0xe6, 0x27, 0xac, 0x2f, 0x05, 0x31, 0xfb, 0x55,
			0x4f, 0x6d, 0x21, 0xd7, 0x4c, 0x59, 0x0a, 0x70, 0xad, 0xfa, 0xa3, 0x90, 0xbd, 0xfb, 0xb3, 0xd6,
			0x8e, 0x46, 0x21, 0x5c, 0xab, 0x18, 0x7d, 0x23, 0x68, 0xd5, 0xa7, 0x1f, 0x5e, 0xbe, 0xc0, 0x81,
			0xcd, 0x3b, 0x20, 0xc0, 0x82, 0xdb, 0xe4, 0xcd, 0x2f, 0xac, 0xa2, 0x87, 0x73, 0x79, 0x5d, 0x6b,
			0x0c, 0x10, 0x20, 0x4b, 0x65, 0x9a, 0x93, 0x9e, 0xf2, 0x9b, 0xbe, 0x10, 0x88, 0x24, 0x36, 0x24,
			0x42, 0x99, 0x27, 0xa7, 0xeb, 0x57, 0x6d, 0xd3, 0xa0, 0x0e, 0xa5, 0xe0, 0x1a, 0xf5, 0xd4, 0x75,
			0x83, 0xb2, 0x27, 0x2c, 0x0c, 0x16, 0x1a, 0x80, 0x65, 0x21, 0xa1, 0x6f, 0xf9, 0xb0, 0xa7, 0x22,
			0xc0, 0xcf, 0x26, 0xb0, 0x25, 0xd5, 0x83, 0x6e, 0x22, 0x58, 0xa4, 0xf7, 0xd4, 0x77, 0x3a, 0xc8,
			0x01, 0xe4, 0x26, 0x3b, 0xc2, 0x94, 0xf4, 0x3d, 0xef, 0x7f, 0xa8, 0x70, 0x3f, 0x3a, 0x41, 0x97,
			0x46, 0x35, 0x25, 0x88, 0x76, 0x52, 0xb0, 0xb2, 0xa4, 0xa2, 0xa7, 0xcf, 0x87, 0xf0, 0x09, 0x14,
			0x87, 0x1e, 0x25, 0x03, 0x91, 0x13, 0xc7, 0xe1, 0x61, 0x8d, 0xa3, 0x40, 0x64, 0xb5, 0x7a, 0x43,
			0xc4, 0x63, 0x24, 0x9f, 0xb8, 0xd0, 0x5e, 0x0f, 0x26, 0xf4, 0xa6, 0xd8, 0x49, 0x72, 0xe7, 0xa9,
			0x05, 0x48, 0x24, 0x14, 0x5f, 0x91, 0x29, 0x5c, 0xdb, 0xe3, 0x9a, 0x6f, 0x92, 0x0f, 0xac, 0xc6,
			0x59, 0x71, 0x2b, 0x46, 0xa5, 0x4b, 0xa2, 0x95, 0xbb, 0xe6, 0xa9, 0x01, 0x54, 0xe9, 0x1b, 0x33,
			0x98, 0x5a, 0x2b, 0xcd, 0x42, 0x0a, 0xd5, 0xc6, 0x7e, 0xc9, 0xad, 0x8e, 0xb7, 0xac, 0x68, 0x64,
			0xdb, 0x27, 0x2a, 0x51, 0x6b, 0xc9, 0x4c, 0x28, 0x39, 0xb0, 0xa8, 0x16, 0x9a, 0x6b, 0xf5, 0x8e,
			0x1a, 0x0c, 0x2a, 0xda, 0x8c, 0x88, 0x3b, 0x7b, 0xf4, 0x97, 0xa4, 0x91, 0x71, 0x26, 0x8e, 0xd1,
			0x5d, 0xdd, 0x29, 0x69, 0x38, 0x4e, 0x7f, 0xf4, 0xbf, 0x4a, 0xab, 0x2e, 0xc9, 0xec, 0xc6, 0x52,
			0x9c, 0xf6, 0x29, 0xe2, 0xdf, 0x0f, 0x08, 0xa7, 0x7a, 0x65, 0xaf, 0xa1, 0x2a, 0xa9, 0xb5, 0x05,
			0xdf, 0x8b, 0x28, 0x7e, 0xf6, 0xcc, 0x91, 0x49, 0x3d, 0x1c, 0xaa, 0x39, 0x07, 0x6e, 0x28, 0xef,
			0x1e, 0xa0, 0x28, 0xf5, 0x11, 0x8d, 0xe6, 0x1a, 0xe0, 0x2b, 0xb6, 0xae, 0xfc, 0x33, 0x43, 0xa0,
			0x50, 0x29, 0x2f, 0x19, 0x9f, 0x40, 0x18, 0x57, 0xb2, 0xbe, 0xad, 0x5e, 0x6e, 0xe2, 0xa1, 0xf1,
			0x91, 0x02, 0x2f, 0x92, 0x78, 0x01, 0x6f, 0x04, 0x77, 0x91, 0xa9, 0xd1, 0x8d, 0xa7, 0xd2, 0xa6,
			0xd2, 0x7f, 0x2e, 0x0e, 0x51, 0xc2, 0xf6, 0xea, 0x30, 0xe8, 0xac, 0x49, 0xa0, 0x60, 0x4f, 0x4c,
			0x13, 0x54, 0x2e, 0x85, 0xb6, 0x83, 0x81, 0xb9, 0xfd, 0xcf, 0xa0, 0xce, 0x4b, 0x2d, 0x34, 0x13,
			0x54, 0x85, 0x2d, 0x36, 0x02, 0x45, 0xc5, 0x36, 0xb6, 0x12, 0xaf, 0x71, 0xf3, 0xe7, 0x7c, 0x90,
			0x95, 0xae, 0x2d, 0xbd, 0xe5, 0x04, 0xb2, 0x65, 0x73, 0x3d, 0xab, 0xfe, 0x10, 0xa2, 0x0f, 0xc7,
			0xd6, 0xd3, 0x2c, 0x21, 0xcc, 0xc7, 0x2b, 0x8b, 0x34, 0x44, 0xae, 0x66, 0x3d, 0x65, 0x92, 0x2d,
			0x17, 0xf8, 0x2c, 0xaa, 0x2b, 0x86, 0x5c, 0xd8, 0x89, 0x13, 0xd2, 0x91, 0xa6, 0x58, 0x99, 0x02,
			0x6e, 0xa1, 0x32, 0x84, 0x39, 0x72, 0x3c, 0x19, 0x8c, 0x36, 0xb0, 0xc3, 0xc8, 0xd0, 0x85, 0xbf,
			0xaf, 0x8a, 0x32, 0x0f, 0xde, 0x33, 0x4b, 0x4a, 0x49, 0x19, 0xb4, 0x4c, 0x2b, 0x95, 0xf6, 0xe8,
			0xec, 0xf7, 0x33, 0x93, 0xf7, 0xf0, 0xd2, 0xa4, 0x0e, 0x60, 0xb1, 0xd4, 0x06, 0x52, 0x6b, 0x02,
			0x2d, 0xdc, 0x33, 0x18, 0x10, 0xb1, 0xa5, 0xf7, 0xc3, 0x47, 0xbd, 0x53, 0xed, 0x1f, 0x10, 0x5d,
			0x6a, 0x0d, 0x30, 0xab, 0xa4, 0x77, 0xe1, 0x78, 0x88, 0x9a, 0xb2, 0xec, 0x55, 0xd5, 0x58, 0xde,
			0xab, 0x26, 0x30, 0x20, 0x43, 0x36, 0x96, 0x2b, 0x4d, 0xb5, 0xb6, 0x63, 0xb6, 0x90, 0x2b, 0x89,
			0xe8, 0x5b, 0x31, 0xbc, 0x6a, 0xf5, 0x0f, 0xc5, 0x0a, 0xcc, 0xb3, 0xfb, 0x9b, 0x57, 0xb6, 0x63,
			0x29, 0x70, 0x31, 0x37, 0x8d, 0xb4, 0x78, 0x96, 0xd7, 0xfb, 0xaf, 0x6c, 0x60, 0x0a, 0xdd, 0x2c,
			0x67, 0xf9, 0x36, 0xdb, 0x03, 0x79, 0x86, 0xdb, 0x85, 0x6e, 0xb4, 0x9c, 0xf2, 0xdb, 0x3f, 0x7d,
			0xa6, 0xd2, 0x36, 0x50, 0xe4, 0x38, 0xf1, 0x88, 0x40, 0x41, 0xb0, 0x13, 0x11, 0x9e, 0x4c, 0x2a,
			0xe5, 0xaf, 0x37, 0xcc, 0xcd, 0xfb, 0x68, 0x66, 0x07, 0x38, 0xb5, 0x8b, 0x3c, 0x59, 0xd1, 0xc0,
			0x24, 0x84, 0x37, 0x47, 0x2a, 0xba, 0x1f, 0x35, 0xca, 0x1f, 0xb9, 0x0c, 0xd7, 0x14, 0xaa, 0x9f,
			0x63, 0x55, 0x34, 0xf4, 0x9e, 0x7c, 0x5b, 0xba, 0x81, 0xc2, 0xb6, 0xb3, 0x6f, 0xde, 0xe2, 0x1c,
			0xa2, 0x7e, 0x34, 0x7f, 0x79, 0x3d, 0x2c, 0xe9, 0x44, 0xed, 0xb2, 0x3c, 0x8c, 0x9b, 0x91, 0x4b,
			0xe1, 0x03, 0x35, 0xe3, 0x50, 0xfe, 0xb5, 0x07, 0x03, 0x94, 0xb7, 0xa4, 0xa1, 0x5c, 0x0c, 0xa1,
			0x20, 0x28, 0x35, 0x68, 0xb7, 0xbf, 0xc2, 0x54, 0xfe, 0x83, 0x8b, 0x13, 0x7a, 0x21, 0x47, 0xce,
			0x7c, 0x11, 0x3a, 0x3a, 0x4d, 0x65, 0x49, 0x9d, 0x9e, 0x86, 0xb8, 0x7d, 0xbc, 0xc7, 0xf0, 0x3b,
			0xbd, 0x3a, 0x3a, 0xb1, 0xaa, 0x24, 0x3e, 0xce, 0x5b, 0xa9, 0xbc, 0xf2, 0x5f, 0x82, 0x83, 0x6c,
			0xfe, 0x47, 0x3b, 0x2d, 0x83, 0xe7, 0xa7, 0x20, 0x1c, 0xd0, 0xb9, 0x6a, 0x72, 0x45, 0x1e, 0x86,
			0x3f, 0x6c, 0x3b, 0xa6, 0x64, 0xa6, 0xd0, 0x73, 0xd1, 0xf7, 0xb5, 0xed, 0x99, 0x08, 0x65, 0xd9,
			0x78, 0xbd, 0x38, 0x15, 0xd0, 0x60, 0x94, 0xfc, 0x9a, 0x2a, 0xba, 0x52, 0x21, 0xc2, 0x2d, 0x5a,
			0xb9, 0x96, 0x38, 0x9e, 0x37, 0x21, 0xe3, 0xaf, 0x5f, 0x05, 0xbe, 0xdd, 0xc2, 0x87, 0x5e, 0x0d,
			0xfa, 0xeb, 0x39, 0x02, 0x1e, 0xe2, 0x7a, 0x41, 0x18, 0x7c, 0xbb, 0x45, 0xef, 0x40, 0xc3, 0xe7,
			0x3b, 0xc0, 0x39, 0x89, 0xf9, 0xa3, 0x0d, 0x12, 0xc5, 0x4b, 0xa7, 0xd2, 0x14, 0x1d, 0xa8, 0xa8,
			0x75, 0x49, 0x3e, 0x65, 0x77, 0x6e, 0xf3, 0x5f, 0x97, 0xde, 0xbc, 0x22, 0x86, 0xcc, 0x4a, 0xf9,
			0xb4, 0x62, 0x3e, 0xee, 0x90, 0x2f, 0x84, 0x0c, 0x52, 0xf1, 0xb8, 0xad, 0x65, 0x89, 0x39, 0xae,
			0xf7, 0x1f, 0x3f, 0x72, 0xb9, 0xec, 0x1d, 0xe2, 0x15, 0x88, 0xbd, 0x35, 0x48, 0x4e, 0xa4, 0x44,
			0x36, 0x34, 0x3f, 0xf9, 0x5e, 0xad, 0x6a, 0xb1, 0xd8, 0xaf, 0xb1, 0xb2, 0xa3, 0x03, 0xdf, 0x1b,
			0x71, 0xe5, 0x3c, 0x4a, 0xea, 0x6b, 0x2e, 0x3e, 0x93, 0x72, 0xbe, 0x0d, 0x1b, 0xc9, 0x97, 0x98,
			0xb0, 0xce, 0x3c, 0xc1, 0x0d, 0x2a, 0x59, 0x6d, 0x56, 0x5d, 0xba, 0x82, 0xf8, 0x8c, 0xe4, 0xcf,
			0xf3, 0xb3, 0x3d, 0x5d, 0x24, 0xe9, 0xc0, 0x83, 0x11, 0x24, 0xbf, 0x1a, 0xd5, 0x4b, 0x79, 0x25,
			0x32, 0x98, 0x3d, 0xd6, 0xc3, 0xa8, 0xb7, 0xd0,
		},
		MacI: 0xabf3e651,
	}
}
