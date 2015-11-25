package gobcy

import (
	"bytes"
	"encoding/json"
)

//GenAssetKeychain generates a public/private key pair, alongside
//an associated OAPAddress for use in the Asset API.
func (api *API) GenAssetKeychain() (pair AddrKeychain, err error) {
	u, err := api.buildURL("/oap/addrs")
	resp, err := postResponse(u, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&pair)
	return
}

//IssueAsset issues new assets onto an Open Asset Address,
//using a private key associated with a funded address
//on the underlying blockchain.
func (api *API) IssueAsset(issue OAPIssue) (tx OAPTX, err error) {
	u, err := api.buildURL("/oap/issue")
	if err != nil {
		return
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&issue); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&tx)
	return
}

//TransferAsset transfers previously issued assets onto a new
//Open Asset Address, based on the assetid and OAPIssue.
func (api *API) TransferAsset(issue OAPIssue, assetID string) (tx OAPTX, err error) {
	u, err := api.buildURL("/oap/" + assetID + "/transfer")
	if err != nil {
		return
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&issue); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&tx)
	return
}

//ListAssetTXs lists the transaction hashes associated
//with the given assetID.
func (api *API) ListAssetTXs(assetID string) (txs []string, err error) {
	u, err := api.buildURL("/oap/" + assetID + "/txs")
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&txs)
	return
}

//GetAssetTX returns a OAPTX associated with the given
//assetID and transaction hash.
func (api *API) GetAssetTX(assetID, hash string) (tx OAPTX, err error) {
	u, err := api.buildURL("/oap/" + assetID + "/txs/" + hash)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&tx)
	return
}

//GetAssetAddr returns an Addr associated with the given
//assetID and oapAddr. Note that while it returns an Address,
//anything that would have represented "satoshis" now represents
//"amount of asset."
func (api *API) GetAssetAddr(assetID, oapAddr string) (addr Addr, err error) {
	u, err := api.buildURL("/oap/" + assetID + "/addrs/" + oapAddr)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&addr)
	return
}
