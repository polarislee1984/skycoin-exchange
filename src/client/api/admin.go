package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/skycoin/skycoin-exchange/src/client/account"
	"github.com/skycoin/skycoin-exchange/src/pp"
	"github.com/skycoin/skycoin-exchange/src/sknet"
	"github.com/skycoin/skycoin/src/cipher"
)

// UpdateCredit admin update credit of specific account
// mode: PUT
// url: /api/v1/admin/credit?dst_pubkey=[:dst_pubkey]&coin_type=[:coin_type]&credit=[:credit]
// params:
//      dst_pubkey: the account pubkey, whose credit will be updated.
//      coin_type: skycoin or bitcoin, the coin you want to credit.
//      credit: credit that will be updated.
func UpdateCredit(se Servicer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var rlt *pp.EmptyRes
		for {
			// get dst_pubkey
			dstPk := r.FormValue("dst_pubkey")
			if dstPk == "" {
				err := errors.New("dst pubkey is empty")
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(err)
				break
			}

			// validate the dst_pubkey
			if _, err := cipher.PubKeyFromHex(dstPk); err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(errors.New("invalid dst pubkey"))
				break
			}

			// get coin type.
			cp := r.FormValue("coin_type")
			if cp == "" {
				err := errors.New("coin_type is empty")
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(err)
				break
			}

			// get credit
			cd := r.FormValue("credit")
			if cd == "" {
				err := errors.New("credit is empty")
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(err)
				break
			}
			credit, err := strconv.ParseUint(cd, 10, 64)
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(err)
				break
			}

			a := account.GetActive()
			r := pp.UpdateCreditReq{
				Pubkey:    pp.PtrString(a.Pubkey),
				CoinType:  pp.PtrString(cp),
				Credit:    pp.PtrUint64(credit),
				DstPubkey: pp.PtrString(dstPk),
			}

			req, err := makeEncryptReq(&r, se.GetServKey().Hex(), a.Seckey)
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_WrongRequest)
				break
			}
			rsp, err := sknet.Get(se.GetServAddr(), "/admin/update/credit", req)
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_ServerError)
				break
			}
			res, err := decodeRsp(rsp.Body, se.GetServKey().Hex(), a.Seckey, &pp.UpdateCreditRes{})
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_ServerError)
				break
			}
			sendJSON(w, res)
			return
		}
		sendJSON(w, rlt)
	}
}