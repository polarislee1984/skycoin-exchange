// Code generated by protoc-gen-go.
// source: pp.transaction.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InjectTxnReq struct {
	CoinType         *string `protobuf:"bytes,10,opt,name=coin_type" json:"coin_type,omitempty"`
	Tx               *string `protobuf:"bytes,20,opt,name=tx" json:"tx,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InjectTxnReq) Reset()                    { *m = InjectTxnReq{} }
func (m *InjectTxnReq) String() string            { return proto.CompactTextString(m) }
func (*InjectTxnReq) ProtoMessage()               {}
func (*InjectTxnReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *InjectTxnReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *InjectTxnReq) GetTx() string {
	if m != nil && m.Tx != nil {
		return *m.Tx
	}
	return ""
}

type InjectTxnRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	Txid             *string `protobuf:"bytes,10,opt,name=txid" json:"txid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InjectTxnRes) Reset()                    { *m = InjectTxnRes{} }
func (m *InjectTxnRes) String() string            { return proto.CompactTextString(m) }
func (*InjectTxnRes) ProtoMessage()               {}
func (*InjectTxnRes) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *InjectTxnRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *InjectTxnRes) GetTxid() string {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return ""
}

type GetTxReq struct {
	CoinType         *string `protobuf:"bytes,10,opt,name=coin_type" json:"coin_type,omitempty"`
	Txid             *string `protobuf:"bytes,20,opt,name=txid" json:"txid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetTxReq) Reset()                    { *m = GetTxReq{} }
func (m *GetTxReq) String() string            { return proto.CompactTextString(m) }
func (*GetTxReq) ProtoMessage()               {}
func (*GetTxReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *GetTxReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *GetTxReq) GetTxid() string {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return ""
}

type Tx struct {
	Btc              *BtcTx `protobuf:"bytes,10,opt,name=btc" json:"btc,omitempty"`
	Sky              *SkyTx `protobuf:"bytes,20,opt,name=sky" json:"sky,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Tx) Reset()                    { *m = Tx{} }
func (m *Tx) String() string            { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()               {}
func (*Tx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *Tx) GetBtc() *BtcTx {
	if m != nil {
		return m.Btc
	}
	return nil
}

func (m *Tx) GetSky() *SkyTx {
	if m != nil {
		return m.Sky
	}
	return nil
}

type GetTxRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	CoinType         *string `protobuf:"bytes,10,opt,name=coin_type" json:"coin_type,omitempty"`
	Tx               *Tx     `protobuf:"bytes,20,opt,name=tx" json:"tx,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetTxRes) Reset()                    { *m = GetTxRes{} }
func (m *GetTxRes) String() string            { return proto.CompactTextString(m) }
func (*GetTxRes) ProtoMessage()               {}
func (*GetTxRes) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *GetTxRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetTxRes) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *GetTxRes) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

type GetRawTxReq struct {
	CoinType         *string `protobuf:"bytes,10,opt,name=coin_type" json:"coin_type,omitempty"`
	Txid             *string `protobuf:"bytes,20,opt,name=txid" json:"txid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetRawTxReq) Reset()                    { *m = GetRawTxReq{} }
func (m *GetRawTxReq) String() string            { return proto.CompactTextString(m) }
func (*GetRawTxReq) ProtoMessage()               {}
func (*GetRawTxReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *GetRawTxReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *GetRawTxReq) GetTxid() string {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return ""
}

type GetRawTxRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	CoinType         *string `protobuf:"bytes,10,opt,name=coin_type" json:"coin_type,omitempty"`
	Rawtx            *string `protobuf:"bytes,20,opt,name=rawtx" json:"rawtx,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetRawTxRes) Reset()                    { *m = GetRawTxRes{} }
func (m *GetRawTxRes) String() string            { return proto.CompactTextString(m) }
func (*GetRawTxRes) ProtoMessage()               {}
func (*GetRawTxRes) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *GetRawTxRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetRawTxRes) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *GetRawTxRes) GetRawtx() string {
	if m != nil && m.Rawtx != nil {
		return *m.Rawtx
	}
	return ""
}

type BtcTx struct {
	Txid             *string    `protobuf:"bytes,10,opt,name=txid" json:"txid,omitempty"`
	Version          *uint32    `protobuf:"varint,11,opt,name=version" json:"version,omitempty"`
	Locktime         *uint32    `protobuf:"varint,12,opt,name=locktime" json:"locktime,omitempty"`
	Vin              []*BtcVin  `protobuf:"bytes,13,rep,name=vin" json:"vin,omitempty"`
	Vout             []*BtcVout `protobuf:"bytes,14,rep,name=vout" json:"vout,omitempty"`
	Blockhash        *string    `protobuf:"bytes,15,opt,name=blockhash" json:"blockhash,omitempty"`
	Confirmations    *uint64    `protobuf:"varint,16,opt,name=confirmations" json:"confirmations,omitempty"`
	Time             *int64     `protobuf:"varint,17,opt,name=time" json:"time,omitempty"`
	Blocktime        *int64     `protobuf:"varint,18,opt,name=blocktime" json:"blocktime,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *BtcTx) Reset()                    { *m = BtcTx{} }
func (m *BtcTx) String() string            { return proto.CompactTextString(m) }
func (*BtcTx) ProtoMessage()               {}
func (*BtcTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *BtcTx) GetTxid() string {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return ""
}

func (m *BtcTx) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *BtcTx) GetLocktime() uint32 {
	if m != nil && m.Locktime != nil {
		return *m.Locktime
	}
	return 0
}

func (m *BtcTx) GetVin() []*BtcVin {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *BtcTx) GetVout() []*BtcVout {
	if m != nil {
		return m.Vout
	}
	return nil
}

func (m *BtcTx) GetBlockhash() string {
	if m != nil && m.Blockhash != nil {
		return *m.Blockhash
	}
	return ""
}

func (m *BtcTx) GetConfirmations() uint64 {
	if m != nil && m.Confirmations != nil {
		return *m.Confirmations
	}
	return 0
}

func (m *BtcTx) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *BtcTx) GetBlocktime() int64 {
	if m != nil && m.Blocktime != nil {
		return *m.Blocktime
	}
	return 0
}

type BtcVin struct {
	Coinbase         *string       `protobuf:"bytes,10,opt,name=coinbase" json:"coinbase,omitempty"`
	Txid             *string       `protobuf:"bytes,11,opt,name=txid" json:"txid,omitempty"`
	Vout             *uint32       `protobuf:"varint,12,opt,name=vout" json:"vout,omitempty"`
	ScriptSig        *BtcScriptSig `protobuf:"bytes,13,opt,name=scriptSig" json:"scriptSig,omitempty"`
	Sequence         *uint32       `protobuf:"varint,14,opt,name=sequence" json:"sequence,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *BtcVin) Reset()                    { *m = BtcVin{} }
func (m *BtcVin) String() string            { return proto.CompactTextString(m) }
func (*BtcVin) ProtoMessage()               {}
func (*BtcVin) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *BtcVin) GetCoinbase() string {
	if m != nil && m.Coinbase != nil {
		return *m.Coinbase
	}
	return ""
}

func (m *BtcVin) GetTxid() string {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return ""
}

func (m *BtcVin) GetVout() uint32 {
	if m != nil && m.Vout != nil {
		return *m.Vout
	}
	return 0
}

func (m *BtcVin) GetScriptSig() *BtcScriptSig {
	if m != nil {
		return m.ScriptSig
	}
	return nil
}

func (m *BtcVin) GetSequence() uint32 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return 0
}

type BtcScriptSig struct {
	Asm              *string `protobuf:"bytes,10,opt,name=asm" json:"asm,omitempty"`
	Hex              *string `protobuf:"bytes,11,opt,name=hex" json:"hex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BtcScriptSig) Reset()                    { *m = BtcScriptSig{} }
func (m *BtcScriptSig) String() string            { return proto.CompactTextString(m) }
func (*BtcScriptSig) ProtoMessage()               {}
func (*BtcScriptSig) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *BtcScriptSig) GetAsm() string {
	if m != nil && m.Asm != nil {
		return *m.Asm
	}
	return ""
}

func (m *BtcScriptSig) GetHex() string {
	if m != nil && m.Hex != nil {
		return *m.Hex
	}
	return ""
}

type BtcVout struct {
	Value            *string                `protobuf:"bytes,10,opt,name=value" json:"value,omitempty"`
	N                *uint32                `protobuf:"varint,11,opt,name=n" json:"n,omitempty"`
	ScriptPubkey     *BtcScriptPubKeyResult `protobuf:"bytes,12,opt,name=scriptPubkey" json:"scriptPubkey,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *BtcVout) Reset()                    { *m = BtcVout{} }
func (m *BtcVout) String() string            { return proto.CompactTextString(m) }
func (*BtcVout) ProtoMessage()               {}
func (*BtcVout) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *BtcVout) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *BtcVout) GetN() uint32 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *BtcVout) GetScriptPubkey() *BtcScriptPubKeyResult {
	if m != nil {
		return m.ScriptPubkey
	}
	return nil
}

type BtcScriptPubKeyResult struct {
	Asm              *string  `protobuf:"bytes,10,opt,name=asm" json:"asm,omitempty"`
	Hex              *string  `protobuf:"bytes,11,opt,name=hex" json:"hex,omitempty"`
	ReqSigs          *int32   `protobuf:"varint,12,opt,name=reqSigs" json:"reqSigs,omitempty"`
	Type             *string  `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	Addresses        []string `protobuf:"bytes,14,rep,name=addresses" json:"addresses,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BtcScriptPubKeyResult) Reset()                    { *m = BtcScriptPubKeyResult{} }
func (m *BtcScriptPubKeyResult) String() string            { return proto.CompactTextString(m) }
func (*BtcScriptPubKeyResult) ProtoMessage()               {}
func (*BtcScriptPubKeyResult) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *BtcScriptPubKeyResult) GetAsm() string {
	if m != nil && m.Asm != nil {
		return *m.Asm
	}
	return ""
}

func (m *BtcScriptPubKeyResult) GetHex() string {
	if m != nil && m.Hex != nil {
		return *m.Hex
	}
	return ""
}

func (m *BtcScriptPubKeyResult) GetReqSigs() int32 {
	if m != nil && m.ReqSigs != nil {
		return *m.ReqSigs
	}
	return 0
}

func (m *BtcScriptPubKeyResult) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *BtcScriptPubKeyResult) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type SkyTx struct {
	Length           *uint32        `protobuf:"varint,10,opt,name=length" json:"length,omitempty"`
	Type             *int32         `protobuf:"varint,11,opt,name=type" json:"type,omitempty"`
	Hash             *string        `protobuf:"bytes,12,opt,name=hash" json:"hash,omitempty"`
	InnerHash        *string        `protobuf:"bytes,13,opt,name=inner_hash" json:"inner_hash,omitempty"`
	Sigs             []string       `protobuf:"bytes,14,rep,name=sigs" json:"sigs,omitempty"`
	Inputs           []string       `protobuf:"bytes,15,rep,name=inputs" json:"inputs,omitempty"`
	Outputs          []*SkyTxOutput `protobuf:"bytes,16,rep,name=outputs" json:"outputs,omitempty"`
	Unknow           *bool          `protobuf:"varint,18,opt,name=unknow" json:"unknow,omitempty"`
	Confirmed        *bool          `protobuf:"varint,19,opt,name=confirmed" json:"confirmed,omitempty"`
	Height           *uint64        `protobuf:"varint,20,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SkyTx) Reset()                    { *m = SkyTx{} }
func (m *SkyTx) String() string            { return proto.CompactTextString(m) }
func (*SkyTx) ProtoMessage()               {}
func (*SkyTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *SkyTx) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *SkyTx) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *SkyTx) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *SkyTx) GetInnerHash() string {
	if m != nil && m.InnerHash != nil {
		return *m.InnerHash
	}
	return ""
}

func (m *SkyTx) GetSigs() []string {
	if m != nil {
		return m.Sigs
	}
	return nil
}

func (m *SkyTx) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *SkyTx) GetOutputs() []*SkyTxOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *SkyTx) GetUnknow() bool {
	if m != nil && m.Unknow != nil {
		return *m.Unknow
	}
	return false
}

func (m *SkyTx) GetConfirmed() bool {
	if m != nil && m.Confirmed != nil {
		return *m.Confirmed
	}
	return false
}

func (m *SkyTx) GetHeight() uint64 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type SkyTxOutput struct {
	Hash             *string `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
	Address          *string `protobuf:"bytes,11,opt,name=address" json:"address,omitempty"`
	Coins            *string `protobuf:"bytes,12,opt,name=coins" json:"coins,omitempty"`
	Hours            *uint64 `protobuf:"varint,13,opt,name=hours" json:"hours,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SkyTxOutput) Reset()                    { *m = SkyTxOutput{} }
func (m *SkyTxOutput) String() string            { return proto.CompactTextString(m) }
func (*SkyTxOutput) ProtoMessage()               {}
func (*SkyTxOutput) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *SkyTxOutput) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *SkyTxOutput) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *SkyTxOutput) GetCoins() string {
	if m != nil && m.Coins != nil {
		return *m.Coins
	}
	return ""
}

func (m *SkyTxOutput) GetHours() uint64 {
	if m != nil && m.Hours != nil {
		return *m.Hours
	}
	return 0
}

func init() {
	proto.RegisterType((*InjectTxnReq)(nil), "pp.InjectTxnReq")
	proto.RegisterType((*InjectTxnRes)(nil), "pp.InjectTxnRes")
	proto.RegisterType((*GetTxReq)(nil), "pp.GetTxReq")
	proto.RegisterType((*Tx)(nil), "pp.Tx")
	proto.RegisterType((*GetTxRes)(nil), "pp.GetTxRes")
	proto.RegisterType((*GetRawTxReq)(nil), "pp.GetRawTxReq")
	proto.RegisterType((*GetRawTxRes)(nil), "pp.GetRawTxRes")
	proto.RegisterType((*BtcTx)(nil), "pp.BtcTx")
	proto.RegisterType((*BtcVin)(nil), "pp.BtcVin")
	proto.RegisterType((*BtcScriptSig)(nil), "pp.BtcScriptSig")
	proto.RegisterType((*BtcVout)(nil), "pp.BtcVout")
	proto.RegisterType((*BtcScriptPubKeyResult)(nil), "pp.BtcScriptPubKeyResult")
	proto.RegisterType((*SkyTx)(nil), "pp.SkyTx")
	proto.RegisterType((*SkyTxOutput)(nil), "pp.SkyTxOutput")
}

func init() { proto.RegisterFile("pp.transaction.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x3e, 0x9a, 0xd6, 0x63, 0xbb, 0x49, 0x97, 0x16, 0xdc, 0x9e, 0x22, 0x73, 0xa9, 0x84,
	0x08, 0x52, 0xc4, 0x81, 0x33, 0x17, 0x04, 0x15, 0x02, 0xb5, 0x11, 0xd7, 0xca, 0x71, 0x96, 0x78,
	0x49, 0xbc, 0x76, 0xbc, 0xeb, 0xc4, 0xf9, 0x65, 0x48, 0xfc, 0x3a, 0x66, 0xc7, 0x5e, 0xab, 0x41,
	0x2d, 0x82, 0x9b, 0xf7, 0xed, 0xec, 0x7b, 0x6f, 0x66, 0x9e, 0xe1, 0x3c, 0xcf, 0x27, 0xba, 0x88,
	0xa4, 0x8a, 0x62, 0x2d, 0x32, 0x39, 0xc9, 0x8b, 0x4c, 0x67, 0xac, 0x9b, 0xe7, 0x57, 0x43, 0xbc,
	0x89, 0xb3, 0x34, 0xb5, 0x60, 0xf8, 0x1a, 0xbc, 0x8f, 0xf2, 0x07, 0x8f, 0xf5, 0xac, 0x92, 0xb7,
	0x7c, 0xc3, 0xce, 0xc0, 0x89, 0x33, 0x21, 0xef, 0xf5, 0x3e, 0xe7, 0x01, 0x8c, 0x3b, 0xd7, 0x0e,
	0x03, 0xe8, 0xea, 0x2a, 0x38, 0x37, 0xdf, 0xe1, 0xbb, 0x83, 0x72, 0xc5, 0xae, 0x60, 0x50, 0x70,
	0x55, 0xae, 0x75, 0xd0, 0x19, 0x77, 0xaf, 0xdd, 0x29, 0x4c, 0x50, 0xe0, 0x96, 0x10, 0xe6, 0x41,
	0x5f, 0x57, 0x62, 0x51, 0xb3, 0x84, 0xaf, 0xe0, 0xe4, 0x03, 0xc7, 0x67, 0x4f, 0x88, 0xd8, 0xe2,
	0x5a, 0xe6, 0x2d, 0x74, 0x67, 0x15, 0x7b, 0x0e, 0xbd, 0xb9, 0x8e, 0xa9, 0xc0, 0x9d, 0x3a, 0x86,
	0xf9, 0xbd, 0x8e, 0x6b, 0x5c, 0xad, 0xf6, 0x54, 0xda, 0xe0, 0x77, 0xab, 0xfd, 0xac, 0x0a, 0x3f,
	0xb7, 0x12, 0x7f, 0x37, 0xf6, 0x88, 0x3c, 0x6b, 0x7b, 0x74, 0xa7, 0x03, 0x53, 0x8a, 0x74, 0x13,
	0x70, 0x91, 0xee, 0x36, 0xda, 0xfd, 0xa3, 0xe9, 0x9b, 0x87, 0xf5, 0xff, 0xed, 0xc0, 0x87, 0xa3,
	0x22, 0xda, 0xb5, 0x83, 0xfe, 0xd9, 0x81, 0xa3, 0xba, 0xdb, 0x83, 0x31, 0xb2, 0x21, 0x1c, 0x6f,
	0x79, 0xa1, 0x70, 0xab, 0x81, 0x8b, 0x80, 0xcf, 0x46, 0x70, 0xb2, 0xce, 0xe2, 0x95, 0x16, 0x29,
	0x0f, 0x3c, 0x42, 0x5e, 0x40, 0x6f, 0x2b, 0x64, 0xe0, 0x8f, 0x7b, 0x56, 0x15, 0x89, 0xbe, 0x09,
	0xc9, 0x2e, 0xa1, 0xbf, 0xcd, 0x4a, 0x1d, 0x9c, 0xd2, 0x8d, 0x6b, 0x6f, 0x10, 0x32, 0x86, 0xe6,
	0x86, 0x26, 0x89, 0x54, 0x12, 0x0c, 0x49, 0xe9, 0x02, 0xfc, 0x38, 0x93, 0xdf, 0x45, 0x91, 0x46,
	0x26, 0x44, 0x2a, 0x18, 0x21, 0xdc, 0x27, 0x3b, 0x46, 0xeb, 0x0c, 0x4f, 0xbd, 0xf6, 0x1d, 0x41,
	0xcc, 0x40, 0x61, 0x0a, 0x83, 0x46, 0x0f, 0xad, 0x99, 0x2e, 0xe7, 0x91, 0xfa, 0x73, 0x60, 0xae,
	0x3d, 0x91, 0x9f, 0xda, 0xf6, 0x4b, 0x70, 0x54, 0x5c, 0x88, 0x5c, 0xdf, 0x89, 0x25, 0x9a, 0x37,
	0x9b, 0x18, 0x35, 0x16, 0xef, 0x2c, 0x6e, 0x28, 0x15, 0xdf, 0x94, 0x5c, 0xc6, 0x1c, 0xdb, 0xc0,
	0x67, 0xe1, 0x35, 0x78, 0x07, 0x15, 0x2e, 0xf4, 0x22, 0x95, 0x36, 0x7a, 0x78, 0x48, 0x78, 0x55,
	0xcb, 0x85, 0x33, 0x38, 0xb6, 0xed, 0xe2, 0xb0, 0xb7, 0xd1, 0xba, 0xb4, 0xb6, 0x1c, 0xe8, 0xd8,
	0x71, 0xbe, 0x01, 0xaf, 0x76, 0xf1, 0xb5, 0x9c, 0xaf, 0xf8, 0x9e, 0xbc, 0xb9, 0xd3, 0xcb, 0x03,
	0x23, 0x78, 0x75, 0xc3, 0xf7, 0xf5, 0x2a, 0xc3, 0x18, 0x2e, 0x1e, 0xbd, 0x78, 0xda, 0x88, 0xd9,
	0x61, 0xc1, 0x37, 0x68, 0x56, 0x11, 0xfd, 0x11, 0x8d, 0xc5, 0x24, 0xc1, 0xa7, 0x6b, 0x9c, 0x69,
	0xb4, 0x58, 0x60, 0x76, 0x14, 0x57, 0xb4, 0x2b, 0x27, 0xfc, 0x85, 0x69, 0xa0, 0x8c, 0xb3, 0x53,
	0x18, 0xac, 0xb9, 0x5c, 0xea, 0x84, 0x88, 0xfd, 0xf6, 0xa9, 0x6b, 0x89, 0x68, 0x83, 0x5e, 0x13,
	0x6a, 0x10, 0x52, 0xf2, 0xe2, 0x9e, 0x30, 0xdf, 0xce, 0x5c, 0x19, 0x61, 0xe2, 0x35, 0x6c, 0x42,
	0xe6, 0xa5, 0x56, 0xb8, 0x73, 0x73, 0x1e, 0xc3, 0x31, 0x8e, 0x87, 0x80, 0x11, 0x85, 0x64, 0xd8,
	0xfe, 0x5d, 0x5f, 0x08, 0x37, 0x2f, 0x4a, 0xb9, 0x92, 0xd9, 0x8e, 0xb6, 0x7d, 0x52, 0x27, 0x99,
	0x52, 0xc2, 0x17, 0xc1, 0x33, 0x82, 0xb0, 0x24, 0xe1, 0x62, 0x99, 0x68, 0x8a, 0x72, 0x3f, 0xfc,
	0x04, 0xee, 0x43, 0x06, 0xeb, 0xb1, 0xcd, 0x73, 0xd3, 0x6c, 0x33, 0x1c, 0x5c, 0x8d, 0x09, 0x8d,
	0x6a, 0x7a, 0xc0, 0x63, 0x92, 0x95, 0x85, 0x22, 0xfb, 0xfd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x13, 0x7d, 0x20, 0xd3, 0xda, 0x04, 0x00, 0x00,
}
