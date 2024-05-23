package packed

import gbres "ghostbb.io/gb/os/gb_res"

func init() {
	if err := gbres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
