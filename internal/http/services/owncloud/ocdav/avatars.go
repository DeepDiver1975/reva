// Copyright 2018-2024 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package ocdav

import (
	"encoding/hex"
	"net/http"

	"github.com/cs3org/reva/pkg/appctx"
	"github.com/cs3org/reva/pkg/rhttp/router"
)

// AvatarsHandler handles avatar requests.
type AvatarsHandler struct {
}

func (h *AvatarsHandler) init(c *Config) error {
	return nil
}

// Handler handles requests.
func (h *AvatarsHandler) Handler(s *svc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := appctx.GetLogger(ctx)

		if r.Method == http.MethodOptions {
			// no need for the user, and we need to be able
			// to answer preflight checks, which have no auth headers
			r.URL.Path = "/" // always use / ... we just want the options answered so phoenix doesnt hiccup
			s.handleOptions(w, r)
			return
		}

		_, r.URL.Path = router.ShiftPath(r.URL.Path)
		if r.Method == http.MethodGet && r.URL.Path == "/128.png" {
			// TODO load avatar url from user context?
			const img = "89504E470D0A1A0A0000000D4948445200000080000000800806000000C33E61CB00000006624B474400FF00FF00FFA0BDA793000000097048597300000B1300000B1301009A9C180000000774494D4507E3061B080516D3ECF61E000008F24944415478DAED9D7D8C1D5515C07FDB76774B775BB7454AA54BBB2D5DDD765B6BD34A140B464CB07EA0113518016B4848FC438A1A448D9A18FF40316942524D544C900F49A17C2882120A28604D506C915AB160B7B2A14275B7BB606BCB76779F7FDC79F4CE79F7BD7DEFED7B3377E69E5FF2B233DBED9B7B3EEECCB977CE3D171445511445511445098B9680645D0BAC8C7EAE020A0E5D0C027B80DDC033EA1ED96521B001D80A3C1F19BB9ECF007003F0CEE83B15CFB90C781A189986D1CB7D8E007F06AE5035FBC599C0359181AA35E6716014188A3EA3D1EFAAFDFFAF025F06DEA2EA4F97EB81935318EB047037F0396035300FE8043A8039D1A723FADD3CA01FB80AB817989CE2BB4F0237AA1992E703C00B150CB313D812057DD36555D4DB7756B8DE41E0236A9664B8A982216E897A72B3980BDC5CE1CE70AB9AA779744541984BF1DF03BA136C4B77F4F871B5E519E074355763590E8C9519A62D4DB15DDDC07E47BBC681156AB6C6D0071C7328F93A60A607ED9B017CDED1BEA35140A94C83259122ED67EE316093876DDD28E61F26A3B69EAD66AC9F61D1AB463D1F7BCF075E126D1E5233D6C74EC7E4CBEA0CB47B317048B4FD6135676D5C2E14F83A705686DA3FD771F7D229E41A823E19507D2A83729CEF90A34FCD3B35F70BA5DD906159AE14B2FC5ACD5B99F384C20E016D19966726B04FC874819AB93C434259EFCD814C2B1C2319C5C14542513FCF916C5B856C17ABB94BF915F1A9D43CCDA2AD20FEDAFA5135779CD9A287FC2D8732EE12322E52B39FE28742391B722863BF90F17635BBA115386C296630C7B2DA492CFFC16423A5CA0C0F94B214938A55E4DE9CC73945E691EEAB6C6F1C605D140314F96D8E1DE009EBB82D923D78EE14CFC63C67DA9E2D64DDA1E687D7882751E49D717452E80DE692DFC99F723C26646E0F390638579C3F1280033CEE888182758035E27C57000EF09438EF0BD9017AC5F940000EB0479CF784EC004BACE362E66FDE1916E7DD213BC07CEBF8BF8104BE72B4B330640768B58E8F0734FA39661D7785EA002DE2FA2703790448676F0DD901EC123593013D02267CB90BCF48591105E110A13051A12304E500E3BEDC0A136666858E105410683B407B20778116605699BB41700E30621DCF09E80E709A757C22640778D93A9E1B501C603BFB70C80EF092753C0B3FD6FB27815DC6E65F213B80CCFFEB0DC0F8B27CCC3F43768003E27C6D000E20339E5F08D9019E9B423979E43C71BE97C0B1B3639E0A40DE3F089983E72FC4EBEAE41DBBDED1F36937C687B4703B55BA050F72E59B488F18EA3EAE0E509A07B826C70E2083DC87D5014C143C669DAFCFB103D8B28D3B82E020E9225EEA3DCF2B839EB4E41C414BCABEC19E4022635BC67D3E346886278AF99138BF3487C6DF2CCE7FA2FD3EEE8876EF78368732CA6251AD6AF6D2D180BDA54B9E6AEC2E25BE25CD633EF53C5FD86E1DCF06DE9D2307D8487C09FC1DDADF4B5981C98E29F692277224DB1F2DB926D0BD04CAF2AC784E2ECB814CB236D05E3573792E10CABA270732FD46C874A19AB9320396B286C9F664C9424C1188A23C2FFA38FCF20D3B185C80D9222EAB7C0C7893757EA7F6EFA9E9A174E3C7AC22B797D3E0AF4AEE168AFB520665F8AA90E101356BF57489DEB39F6C958D6FA77467D337AB59ABA705784828F033196AFF15A2ED8F12D6DAC786B086D22D57B2B07A688EA3DDEBD59CF5F103A1C86D1968F336D1E69FAA19EB6701A6744C5666079789B61ED367FFF4F99650EA11FC5C42D64A3CB3A9007C57CDD7189E168AFDBE876DBC91FCE734A4463F66F3485BC11FF4A87D978AB68D11C632B744B99AD2DD44CFF1A05DEB89BFC62E00D7AAB99AC30EA1E8D7800E8F82BE02709F9AA9799C46E9DE820748A7E2F65B8997BA2F06A81D6AA6E6D289C9A9B7153F98F070EB3D8E9E3F4AFCCD9FD244563B0C3044325BB17DC271ED02F02E354BF2C1D70987219AB9A6E0DA32C6FFA49A231DFACA18647B13AE7553996B6D5333A4CB324CA125DB2813C0CA065EA3D731D42B00B7A9FAFDC136CCFF68ECEE638BA2EF94A38F3655BB1FC8F705CDD87CF23E718D6FAADAFD19168E0AE3346338D625AE314C7CB58F921232FBA6995BCFDD21AEF551557FBAB4736AA38924B26F36503AF9A3A95E29F26002C33F89CC58BE4BCD900E1FA2741E3E89A8BC8D78E2C704F03E3547B2F43AC6E4572778FD2D8EEBF7A859926101F04A0AB77E89DCF5FC1029EF0016024B89EFBE5D00FE413AAF83DB319341765B4E92EF4297A97215A519C2C749E60D603916112FFD52DC14F2323557633803F3EEFD49C73377043F52C2CE1141617149DB4398323767AA19AB6739F005E09798248F51DC6FE00EE357DD80D3817F9769EBAB517CF040143C6AB018B10CB818F80EA61ED0781905CACF0EFC4CBBEAC45434A9468613983AC15F073691AF8A6815E9C1E4CF8F44069FAC5261C5D2EAABF07BE6AD0593A3F05C0D724D46BA18C2AC77E8CE93C1DB804F03B746B7F4420D9F21E07EE02BC0BA0CCADE0F5C8399391CA851F641E076E072329864DA1605463B6A10780CB38E6E2F701D8D7D97EF13E702D7037F8F460B6355EAE741E06D789E7FB004933675A04AA186819B31397C6B896FA516029D98E4D64B22BD1DA9426703C08F7DEA20B380B7535A0ACDF59C3B0CEC06BE019CA531B0933E4C8EE100A51948AE6078252916FADA8829803C51A191AF005FC32CA298AFF6ADA9632D8E628017A77874EE8E3A6162F402BFA8D0A8039852E8FD6AC786B10E938CF27205BD6F4F628EE18B94AED22D7E0E621226B40C7AF368053E4EE9CA287B7E6173332EDC8149B4745DF477C087D53689B3391A4DB86C720B0DAEA774D07191A3C0F96A87D4D952663839D8A85BCE2EC7977F9B6C54EC0A851EE0670E3BED9EEEDCC1FB1D51E73AD5B7B75CE888D12E99CE17CAE95B5D04E93F17519A2B5917B2FAC53ED56D6678840614A9FEACF8924DAAD7CCB0BEDA3B77A569C47788F3DFAB5E33C37E71BEB81E07586E1D1F45F7BACF1AF67ECC67D4E30036AFAB03648A494CB26A91B28B6567D5F0A573D07570596176E40045C3774ED7011670EA3DBFE23F2DC2E8EDF538408B389EA77ACD2C6DF5C40007556FB9E1AFD5F472175762B66D9B2D6EFF05F19332E7D4F877AE7F6FF66327EF8FB53F015BB50F288AA2288AA2288A62F83FEC37068C6750398B0000000049454E44AE426082"
			decoded, err := hex.DecodeString(img)
			if err != nil {
				log.Error().Err(err).Msg("error decoding string")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set(HeaderContentType, "image/png")
			if _, err := w.Write(decoded); err != nil {
				log.Error().Err(err).Msg("error writing data response")
			}
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})
}
