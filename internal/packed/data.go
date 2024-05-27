package packed

import gbres "ghostbb.io/gb/os/gb_res"

func init() {
	if err := gbres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
